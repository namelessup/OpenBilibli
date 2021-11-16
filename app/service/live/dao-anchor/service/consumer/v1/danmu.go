package v1

import (
	"context"
	"encoding/json"

	"github.com/namelessup/bilibili/app/service/live/dao-anchor/dao"

	"github.com/namelessup/bilibili/app/service/live/dao-anchor/model"
	"github.com/namelessup/bilibili/library/log"
	"github.com/namelessup/bilibili/library/queue/databus"
)

//消费弹幕消息的业务逻辑

//DanmuCountStatistics  弹幕数统计
func (s *ConsumerService) DanmuCountStatistics(ctx context.Context, msg *databus.Message) (err error) {
	defer msg.Commit()
	//@todo 等弹幕加了msg_id 再换
	msgInfo := new(model.DMSendMsgContent)
	if err = json.Unmarshal(msg.Value, &msgInfo); err != nil {
		log.Error("DanmuCountStatistics_error:msg=%v;err=%v;msgInfo=%v", msg, err, msgInfo)
		return
	}
	log.Info("DanmuCountStatistics:msgInfo=%v", string(msg.Value))
	roomId := msgInfo.RoomId
	redisKeyInfo := s.dao.SRoomRecordCurrent(dao.DANMU_NUM, roomId)
	s.dao.Incr(ctx, redisKeyInfo.RedisKey, 1, redisKeyInfo.TimeOut)
	return
}
