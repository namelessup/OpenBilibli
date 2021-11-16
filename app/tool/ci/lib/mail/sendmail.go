package sendmail

import (
	"github.com/namelessup/bilibili/app/tool/saga/model"
	"github.com/namelessup/bilibili/app/tool/saga/service/mail"
	"github.com/namelessup/bilibili/library/log"
)

/*
mail model:

【Saga 提醒】+mailTitle
Saga 事件通知
执行状态 : 成功
Pipeline信息: http://gitlab.bilibili.co/platform/github.com/namelessup/bilibili/pipelines/1551
来源分支 : ci-commit/test
修改说明 :
额外信息: 你真是棒棒的，合并成功了
*/
func SendMail(sendTo []string, url string, data string, sourceBranch string, extraData string, pipeStatus string) {
	var (
		mAddress model.Mail
		mDada    model.MailData
	)
	for _, to := range sendTo {
		singleMail := &model.MailAddress{Address: to}
		mAddress.ToAddress = append(mAddress.ToAddress, singleMail)
	}
	if pipeStatus == "failed" {
		mAddress.Subject = "【Saga 提醒】Pipeline 执行失败 "
		mDada.PipeStatus = "失败"
	} else if pipeStatus == "success" {
		mAddress.Subject = "【Saga 提醒】Pipeline 执行成功 "
		mDada.PipeStatus = "成功"
	} else {
		mAddress.Subject = "【Saga 提醒】 " + pipeStatus
	}
	mDada.Info = extraData
	mDada.Description = data
	mDada.SourceBranch = sourceBranch
	mDada.URL = url
	mDada.PipelineStatus = pipeStatus
	if err := mail.SendMail3("saga@bilibili.com", "SAGA", "Lexgm8AAQrF7CcNA", &mAddress, &mDada); err != nil {
		log.Error("Error(%v)", err)
	}
}
