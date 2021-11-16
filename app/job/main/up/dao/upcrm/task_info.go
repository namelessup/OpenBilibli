package upcrm

import (
	"github.com/namelessup/bilibili/app/job/main/up/model/upcrmmodel"
	"github.com/namelessup/bilibili/library/log"
	xtime "github.com/namelessup/bilibili/library/time"
	"time"
)

//StartTask start task
func (d *Dao) StartTask(taskType int, now time.Time) (affectedRow int64, err error) {
	var task = &upcrmmodel.TaskInfo{}
	task.TaskType = int8(taskType)
	task.GenerateDate = now.Format(upcrmmodel.TimeFmtDate)
	task.StartTime = xtime.Time(now.Unix())
	task.TaskState = upcrmmodel.TaskStateStart
	var db = d.crmdb.Model(task).Save(task)
	err = db.Error
	if err != nil {
		log.Error("error start task info, err=%+v", err)
		return
	}

	affectedRow = db.RowsAffected
	return
}

//FinishTask finish task
func (d *Dao) FinishTask(taskType int, now time.Time, state int) (affectedRow int64, err error) {
	var task = &upcrmmodel.TaskInfo{}
	task.TaskType = int8(taskType)
	task.GenerateDate = now.Format(upcrmmodel.TimeFmtDate)
	task.EndTime = xtime.Time(now.Unix())
	task.TaskState = int16(state)
	var db = d.crmdb.Model(task).Where("generate_date=? and task_type=?", task.GenerateDate, taskType).Update(task)
	err = db.Error
	if err != nil {
		log.Error("error end task info, err=%+v", err)
		return
	}

	affectedRow = db.RowsAffected
	return
}
