package service

import (
	"context"
	"time"

	"github.com/namelessup/bilibili/app/admin/main/videoup/model/archive"
	accgrpc "github.com/namelessup/bilibili/app/service/main/account/api"
	"github.com/namelessup/bilibili/app/service/main/up/dao/global"
	"github.com/namelessup/bilibili/app/service/main/up/model"
	"github.com/namelessup/bilibili/library/log"
	"github.com/namelessup/bilibili/library/net/metadata"
	"github.com/namelessup/bilibili/library/queue/databus/report"
)

const (
	//ActionAdd add
	ActionAdd = "add"
	//ActionEdit edit
	ActionEdit = "edit"
	//ActionDelete delete
	ActionDelete = "delete"
)

//UpSpecialLogInfo special log
type UpSpecialLogInfo struct {
	Up     *model.UpSpecialWithName
	UpOld  *model.UpSpecialWithName
	UID    int64
	UName  string
	CTime  time.Time
	Action string
}

// send to log special up edit
func (s *Service) sendUpSpecialLog(c context.Context, opInfo *UpSpecialLogInfo) (err error) {
	logData := &report.ManagerInfo{
		Uname:    opInfo.UName,
		UID:      opInfo.UID,
		Business: archive.LogClientUp,
		Type:     int(opInfo.Up.GroupID),
		Oid:      opInfo.Up.Mid,
		Action:   opInfo.Action,
		Ctime:    opInfo.CTime,
		Content: map[string]interface{}{
			"up": opInfo.Up,
		},
	}
	s.fillGroupInfo(c, opInfo.Up)
	if opInfo.Action == ActionEdit {
		s.fillGroupInfo(c, opInfo.UpOld)
		logData.Content["old_up"] = opInfo.UpOld
	}
	report.Manager(logData)
	log.Info("sendUpSpecialLog logData(%+v) opInfo(%+v)", logData, opInfo)
	return
}

func (s *Service) fillGroupInfo(c context.Context, up *model.UpSpecialWithName) {
	if up == nil || up.GroupName != "" {
		return
	}
	var (
		err        error
		infosReply *accgrpc.InfoReply
		group      = s.getGroupCache(up.GroupID)
	)
	if group != nil {
		up.GroupName = group.Name
		up.GroupTag = group.Tag
	}
	if up.Mid == 0 {
		return
	}
	if infosReply, err = global.GetAccClient().Info3(c, &accgrpc.MidReq{Mid: up.Mid, RealIp: metadata.String(c, metadata.RemoteIP)}); err != nil {
		return
	}
	if infosReply == nil || infosReply.Info == nil {
		return
	}
	up.UName = infosReply.Info.Name
}
