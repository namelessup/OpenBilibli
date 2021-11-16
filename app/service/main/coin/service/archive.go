package service

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	coin "github.com/namelessup/bilibili/app/service/main/coin/model"
	"github.com/namelessup/bilibili/library/log"
	xtime "github.com/namelessup/bilibili/library/time"
	"github.com/namelessup/bilibili/library/xstr"
)

// sendArcMsg send msg to databus
func (s *Service) sendArcMsg(c context.Context, mid, aid, tp, upmid, multiply, pubTime int64, ip string, typeid int32, ts int64) {
	msg := &coin.DataBus{
		Mid:      mid,
		Avid:     aid,
		AvType:   int8(tp),
		UpID:     upmid,
		Multiply: multiply,
		Time:     xtime.Time(pubTime),
		IP:       ip,
		TypeID:   int16(typeid),
		Ctime:    ts,
		MsgID:    fmt.Sprintf("%d%06d%02d", time.Now().Unix(), rand.Intn(_maxRandVal), mid%_maxEXP),
	}
	ids, err := s.coinDao.TagIds(c, msg.Avid)
	if err != nil {
		log.Error("s.coinDao.TagIds error(%v)", err)
		return
	}
	msg.Tags = xstr.JoinInts(ids)
	if err = s.coinDao.PubBigData(c, msg.Avid, msg); err == nil {
		log.Info("mabd m(%d)a(%d)", msg.Mid, msg.Avid)
	}
}
