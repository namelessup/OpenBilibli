package v1

import (
	"context"
	"encoding/json"

	"github.com/namelessup/bilibili/app/service/live/dao-anchor/dao"
	"github.com/namelessup/bilibili/app/service/live/dao-anchor/model"
	"github.com/namelessup/bilibili/library/log"
	"github.com/namelessup/bilibili/library/queue/databus"
)

func (s *ConsumerService) GuardBuyStatistics(ctx context.Context, msg *databus.Message) (err error) {
	defer msg.Commit()

	msgInfo := new(model.MessageValue)
	if err = json.Unmarshal(msg.Value, &msgInfo); err != nil {
		log.Error("GuardBuyStatistics_unmarshal_message_error:msg=%v;err=%v;msgInfo=%v", msg, err, msgInfo)
		return
	}
	log.Info("GuardBuyStatistics:msgInfo=%v", string(msg.Value))
	msgContent := new(model.GuardBuyMessageContent)
	if err = json.Unmarshal([]byte(msgInfo.MsgContent), &msgContent); err != nil {
		log.Error("GuardBuyStatistics_unmarshal_msgcontent_error:msg=%v;err=%v;msgInfo=%v", msg, err, msgInfo)
		return
	}

	roomId := msgContent.RoomId
	num := msgContent.Num
	coin := msgContent.Coin

	goldNumKeyInfo := s.dao.SRoomRecordCurrent(dao.GIFT_GOLD_NUM, roomId)
	s.dao.Incr(ctx, goldNumKeyInfo.RedisKey, num, goldNumKeyInfo.TimeOut)

	goldAmountKeyInfo := s.dao.SRoomRecordCurrent(dao.GIFT_GOLD_AMOUNT, roomId)
	s.dao.Incr(ctx, goldAmountKeyInfo.RedisKey, coin, goldAmountKeyInfo.TimeOut)

	return
}
