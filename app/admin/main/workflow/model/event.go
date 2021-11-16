package model

import (
	"net/url"
	"strings"

	xtime "github.com/namelessup/bilibili/library/time"
)

// consts for workflow event
const (
	// EventAdminReply 管理员回复
	EventAdminReply = 1
	// EventAdminNote 管理员回复并记录
	EventAdminNote = 2
	// EventUserReply 用户回复
	EventUserReply = 3
	// EventSystem 系统回复
	EventSystem = 4
)

// EventSlice is a Event slice struct
type EventSlice []*Event

// Event model is the model for challenge changes
type Event struct {
	Eid         int64      `json:"eid" gorm:"column:id"`
	Cid         int64      `json:"cid" gorm:"column:cid"`
	AdminID     int64      `json:"adminid" gorm:"column:adminid"`
	Content     string     `json:"content" gorm:"column:content"`
	Attachments string     `json:"attachments" gorm:"column:attachments"`
	Event       int8       `json:"event" gorm:"column:event"`
	CTime       xtime.Time `json:"ctime" gorm:"column:ctime"`
	MTime       xtime.Time `json:"mtime" gorm:"column:mtime"`
	Admin       string     `json:"admin" gorm:"-"`
}

// TableName is used to identify table name in gorm
func (Event) TableName() string {
	return "workflow_event"
}

// FixAttachments will fix attachments url as user friendly
// ignore https case
// FIXME: this should be removed after attachment url is be normed
func (e *Event) FixAttachments() {
	if len(e.Attachments) <= 0 {
		return
	}
	sep := ";"
	atts := strings.Split(e.Attachments, sep)
	fixed := make([]string, 0, len(atts))
	for _, a := range atts {
		u, err := url.Parse(a)
		if err != nil {
			continue
		}
		u.Scheme = "http"
		fixed = append(fixed, u.String())
	}
	e.Attachments = strings.Join(fixed, sep)
}
