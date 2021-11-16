package dao

import (
	"context"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"github.com/namelessup/bilibili/app/job/live/gift/internal/model"
	"github.com/namelessup/bilibili/library/log"
	"strconv"
)

var (
	_getBag       = "SELECT id,gift_num FROM user_gift_%s WHERE uid = ? AND gift_id = ? AND expireat = ? LIMIT 1"
	_getBagByID   = "SELECT id,gift_num FROM user_gift_%s WHERE id = ?"
	_updateBagNum = "UPDATE user_gift_%s SET gift_num = gift_num + ? WHERE id = ?"
	_insertBag    = "INSERT INTO user_gift_%s (uid,gift_id,gift_num,expireat) VALUES (?,?,?,?)"
)

// GetBag GetBag
func (d *Dao) GetBag(ctx context.Context, uid, giftID, expireAt int64) (res *model.BagInfo, err error) {
	log.Info("GetBag,uid:%d,giftID:%d,expireAt:%d", uid, giftID, expireAt)
	row := d.db.QueryRow(ctx, fmt.Sprintf(_getBag, getPostFix(uid)), uid, giftID, expireAt)
	res = &model.BagInfo{}
	if err = row.Scan(&res.ID, &res.GiftNum); err != nil {
		if err == sql.ErrNoRows {
			err = nil
			return
		}
		log.Error("GetBag row.Scan error(%v)", err)
	}
	return
}

// UpdateBagNum UpdateBagNum
func (d *Dao) UpdateBagNum(ctx context.Context, uid, id, num int64) (affected int64, err error) {
	log.Info("UpdateBagNum,uid:%d,id:%d,num:%d", uid, id, num)
	res, err := d.db.Exec(ctx, fmt.Sprintf(_updateBagNum, getPostFix(uid)), num, id)
	if err != nil {
		log.Error("UpdateBagNum error(%v)", err)
		return
	}
	return res.RowsAffected()
}

// AddBag AddBag
func (d *Dao) AddBag(ctx context.Context, uid, giftID, giftNum, expireAt int64) (affected int64, err error) {
	log.Info("AddBag,uid:%d,giftID:%d,giftNum:%d,expireAt:%d", uid, giftID, giftNum, expireAt)
	res, err := d.db.Exec(ctx, fmt.Sprintf(_insertBag, getPostFix(uid)), uid, giftID, giftNum, expireAt)
	if err != nil {
		log.Error("AddBag error(%v)", err)
		return
	}
	return res.LastInsertId()
}

// GetBagByID GetBagByID
func (d *Dao) GetBagByID(ctx context.Context, uid, id int64) (res *model.BagInfo, err error) {
	log.Info("GetBagByID,uid:%d,id:%d", uid, id)
	row := d.db.QueryRow(ctx, fmt.Sprintf(_getBagByID, getPostFix(uid)), id)
	res = &model.BagInfo{}
	if err = row.Scan(&res.ID, &res.GiftNum); err != nil {
		if err == sql.ErrNoRows {
			err = nil
			return
		}
		log.Error("GetBagByID row.Scan error(%v)", err)
	}
	return
}

func getPostFix(uid int64) string {
	uidStr := strconv.Itoa(int(uid))
	h := md5.New()
	h.Write([]byte(uidStr))
	md5Str := hex.EncodeToString(h.Sum(nil))
	return md5Str[0:1]
}
