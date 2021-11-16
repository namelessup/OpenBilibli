package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/namelessup/bilibili/app/admin/main/videoup/model/utils"
	"github.com/namelessup/bilibili/library/sync/errgroup"
	"github.com/namelessup/bilibili/library/xstr"
	"time"

	"github.com/namelessup/bilibili/app/admin/main/videoup/model/archive"
	accApi "github.com/namelessup/bilibili/app/service/main/account/api"
	"github.com/namelessup/bilibili/library/log"
)

// weightConf 所有激活的配置项
func (s *Service) weightConf(c context.Context) (wconfs map[int8]map[int64]*archive.WCItem, err error) {
	items, err := s.arc.WeightConf(c)
	if err != nil {
		log.Error("WeightConf error(%v)", err)
		return
	}
	wconfs = map[int8]map[int64]*archive.WCItem{}
	for _, item := range items {
		conf, ok := wconfs[item.Radio]
		if !ok {
			conf = make(map[int64]*archive.WCItem)
		}
		conf[item.CID] = item
		wconfs[item.Radio] = conf
	}
	return
}

// 同一类型的同id配置只能有一个
func (s *Service) checkConflict(c context.Context, mcases map[int64]*archive.WCItem) (err error) {
	wcf, err := s.weightConf(c)
	if err != nil {
		log.Error("s.weightConf error(%v)", err)
		return
	}
	s.twConCache = wcf
	for _, item := range mcases {
		if cfgs, ok := s.twConCache[item.Radio]; ok {
			if cfg, ok := cfgs[item.CID]; ok {
				return fmt.Errorf("config(%d) conflict with (id=%d,desc=%s)", item.CID, cfg.ID, cfg.Desc)
			}
		}
	}
	return
}

// AddWeightConf 配置权重
func (s *Service) AddWeightConf(c context.Context, cfg *archive.WeightConf, uid int64, uname string) (err error) {
	//0. 一级分区转化为二级分区
	if cfg.Radio == archive.WConfType {
		cfg.Ids = s.tarnsType(c, cfg.Ids)
	}

	//1. 解析表单提交的配置
	mcases, istaskid, err := archive.ParseWeightConf(cfg, uid, uname)
	if err != nil {
		log.Error("archive.ParseWeightConf(%v, %v) error(%v)", cfg, uname, err)
		return
	}
	//2. 检查是否互相冲突
	if err = s.checkConflict(c, mcases); err != nil {
		log.Error("s.checkConflict error(%v)", err)
		return
	}

	//3. 更新任务的权重配置
	if istaskid {
		if err = s.setWeightConf(c, cfg.Ids, mcases); err != nil {
			log.Error("s.setWeightConf error(%v)", err)
			return
		}
	}

	//4. 将配置存储到数据库表
	if err = s.arc.InWeightConf(c, mcases); err != nil {
		log.Error("s.arc.InWeightConf(%v) error(%v)", mcases, err)
		return
	}

	return
}

// DelWeightConf 删除配置项
func (s *Service) DelWeightConf(c context.Context, id int64) (err error) {
	if _, err = s.arc.DelWeightConf(c, id); err != nil {
		log.Error("s.DelWeightConf(%d) error(%v)", id, err)
	}
	return
}

// ListWeightConf 列出配置
func (s *Service) ListWeightConf(c context.Context, v *archive.Confs) (cfg []*archive.WCItem, err error) {
	if cfg, err = s.arc.ListWeightConf(c, v); err != nil {
		log.Error("s.ListWeightConf(%+v) error(%v)", v, err)
		return
	}
	// 根据taskid补充filename和title
	switch v.Radio {
	case archive.WConfMid: // 补充昵称和粉丝数
		s.singleIDtoName(c, cfg, func(c context.Context, cid int64) (res []interface{}, err error) {
			stat, err := s.profile(c, cid)
			if err != nil {
				log.Error("s.profile(%d) error(%v)", cid, err)
				return
			}
			res = []interface{}{stat.Profile.Name, stat.Follower}
			return
		}, false, "CID", "Creator", "Fans")
	case archive.WConfTaskID: //补充title和filename
		s.mulIDtoName(c, cfg, s.arc.LWConfigHelp, "CID", "FileName", "Title", "Vid")
	case archive.WConfType: //补充分区名称
		s.singleIDtoName(c, cfg, func(c context.Context, cid int64) (res []interface{}, err error) {
			if item, ok := s.typeCache[int16(cid)]; ok {
				res = []interface{}{item.Name}
			}
			return
		}, false, "CID", "TypeName")
	case archive.WConfUpFrom: //补充投稿来源
		s.singleIDtoName(c, cfg, func(c context.Context, cid int64) (res []interface{}, err error) {
			res = []interface{}{archive.UpFrom(int8(cid))}
			return
		}, false, "CID", "UpFrom")
	default:
	}

	return
}

// ListWeightLogs 权重变更日志
func (s *Service) ListWeightLogs(c context.Context, taskid int64, page int) (cfg []*archive.TaskWeightLog, items int64, err error) {
	cfg, err = s.arc.WeightLog(c, taskid)
	if err != nil {
		log.Info("s.ListWeightLogs(%d) error(%v)", taskid, err)
		return
	}
	cfg, items, err = s.weightLogHelp(c, taskid, cfg, page)
	return
}

// MaxWeight 当前的权重最大值, 供用户配置做参考
func (s *Service) MaxWeight(c context.Context) (max int64, err error) {
	return s.arc.GetMaxWeight(c)
}

// 补充查询日志需要的信息
func (s *Service) weightLogHelp(c context.Context, taskid int64, twl []*archive.TaskWeightLog, page int) (retwl []*archive.TaskWeightLog, items int64, err error) {
	var (
		mid       int64
		ps        *accApi.ProfileStatReply
		name      string
		fans      int
		upspecial []int8
	)

	task, err := s.arc.TaskByID(c, taskid)
	if task == nil {
		log.Error("s.arc.TaskByID(%d) err(%v)", taskid, err)
		return
	}
	// 1.获取用户信息
	items = int64(len(twl))
	if items == 0 {
		return
	}
	mid = twl[items-1].Mid
	if mid == 0 {
		log.Error("weightLogHelp taskid=%d miss mid", taskid)
		goto EMPTYMID
	}

	ps, err = s.profile(c, mid)
	if err != nil || ps == nil {
		log.Error("can't get Account.Card(%d) err(%v)", mid, err)
		err = nil
	} else {
		name = ps.Profile.Name
		fans = int(ps.Follower)
	}

	upspecial = s.getSpecial(mid)

EMPTYMID:
	retwl = []*archive.TaskWeightLog{}
	// 反转日志顺序
	var count int
	for i := len(twl) - 1 - (page-1)*20; i >= 0 && count <= 20; func() { i--; count++ }() {
		twl[i].Creator = name
		twl[i].UpSpecial = upspecial
		twl[i].Fans = int64(fans)
		twl[i].Wait = twl[i].Uptime.TimeValue().Sub(task.CTime.TimeValue()).Seconds() * 1000.0
		if !task.PTime.TimeValue().IsZero() {
			twl[i].Ptime = string(task.PTime)
		}

		if len(twl[i].CfItems) > 0 {
			for _, item := range twl[i].CfItems {
				var desc = item.Desc
				var v = new(archive.WCItem)
				if err = json.Unmarshal([]byte(item.Desc), v); err == nil { //兼容新旧日志格式
					desc = v.Desc
				}
				err = nil
				twl[i].Desc += fmt.Sprintf("%s:%s; ", archive.CfWeightDesc(item.Radio), desc)
			}
		}
		retwl = append(retwl, twl[i])
	}
	return
}

func (s *Service) getSpecial(mid int64) []int8 {
	upspecial := []int8{}
	for k, v := range s.upperCache {
		if _, ok := v[mid]; ok {
			upspecial = append(upspecial, k)
		}
	}
	return upspecial
}

func (s *Service) setWeightConf(c context.Context, ids string, mcases map[int64]*archive.WCItem) (err error) {
	var mitems map[int64]*archive.TaskPriority
	arrid, _ := xstr.SplitInts(ids)
	// 1. 获取旧的配置
	if mitems, err = s.getTWCache(c, arrid); err != nil {
		log.Error("s.getTWCache(%v) error(%v)", arrid, err)
		return
	}
	if len(mitems) == 0 {
		return fmt.Errorf("没有找到任务(%v)", arrid)
	}
	// 2. 将新加配置加入
	for _, item := range mitems {
		if one, ok := mcases[item.TaskID]; ok && one != nil {
			one.Mtime = utils.NewFormatTime(time.Now())
			item.CfItems = append(item.CfItems, one)
		}
	}

	// 3. 更新配置到redis和数据库
	var wg errgroup.Group
	wg.Go(func() (err error) {
		return s.task.SetWeightRedis(c, mitems)
	})
	for _, item := range mitems {
		var descb []byte
		v := item
		wg.Go(func() (err error) {
			if descb, err = json.Marshal(v.CfItems); err != nil {
				return
			}
			_, err = s.arc.UpCwAfterAdd(c, v.TaskID, string(descb))
			return
		})
	}
	err = wg.Wait()

	return
}

// ShowWeightVC 展示权重配置值
func (s *Service) ShowWeightVC(c context.Context) (wvc *archive.WeightVC, err error) {
	if wvc, err = s.arc.WeightVC(c); err != nil {
		log.Error("s.arc.WeightVC error(%v)", err)
		return
	}

	if wvc == nil {
		wvc = archive.WLVConf
		s.arc.InWeightVC(c, wvc, "默认权重配置")
	}
	return
}

// SetWeightVC 设置权重配置
func (s *Service) SetWeightVC(c context.Context, wvc *archive.WeightVC) (err error) {
	_, err = s.arc.SetWeightVC(c, wvc, "设置权重配置")
	return
}

// 处理一级分区转二级分区
func (s *Service) tarnsType(c context.Context, ids string) (res string) {
	var (
		mapNoNeed = make(map[int16]struct{}) // 不需要的一级分区
		mapNeed   = make(map[int16]struct{}) // 需要的一级分区
		idres     = []int64{}
	)

	arrid, _ := xstr.SplitInts(ids)
	for _, id := range arrid {
		if mod, ok := s.typeCache[int16(id)]; ok {
			if mod.PID == 0 {
				mapNeed[int16(id)] = struct{}{}
			} else {
				mapNoNeed[mod.PID] = struct{}{}
				idres = append(idres, id)
			}
		}
	}
	for k := range mapNoNeed {
		delete(mapNeed, k)
	}
	for k := range mapNeed {
		idres = append(idres, s.typeCache2[k]...)
	}

	return xstr.JoinInts(idres)
}
