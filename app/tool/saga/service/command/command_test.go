package command

import (
	"context"
	"encoding/json"
	"flag"
	"path/filepath"
	"testing"

	"github.com/namelessup/bilibili/app/tool/saga/conf"
	"github.com/namelessup/bilibili/app/tool/saga/dao"
	"github.com/namelessup/bilibili/app/tool/saga/model"
	"github.com/namelessup/bilibili/app/tool/saga/service/gitlab"

	"github.com/smartystreets/goconvey/convey"
)

var (
	c                     *Command
	gitlabHookCommentTest = []byte(`{
	"object_kind":"note",
	"event_type":"note",
	"user":{
		"name":"changhengyuan",
		"username":"changhengyuan",
		"avatar_url":"https://www.gravatar.com/avatar/d3218d34473c6fb4d18a770f14e59a89?s=80\u0026d=identicon"
	},
	"project_id":35,
	"project":{
		"id":35,
		"name":"test-saga",
		"description":"",
		"web_url":"http://gitlab.bilibili.co/changhengyuan/test-saga",
		"avatar_url":null,
		"git_ssh_url":"git@gitlab.bilibili.co:changhengyuan/test-saga.git",
		"git_http_url":"http://gitlab.bilibili.co/changhengyuan/test-saga.git",
		"namespace":"changhengyuan",
		"visibility_level":20,
		"path_with_namespace":"changhengyuan/test-saga",
		"default_branch":"master",
		"ci_config_path":null,
		"homepage":"http://gitlab.bilibili.co/changhengyuan/test-saga",
		"url":"git@gitlab.bilibili.co:changhengyuan/test-saga.git",
		"ssh_url":"git@gitlab.bilibili.co:changhengyuan/test-saga.git",
		"http_url":"http://gitlab.bilibili.co/changhengyuan/test-saga.git"},
		"object_attributes":{
			"id":3040,
			"note":"test",
			"noteable_type":"MergeRequest",
			"author_id":15,
			"created_at":"2018-09-26 06:55:13 UTC",
			"updated_at":"2018-09-26 06:55:13 UTC",
			"project_id":35,
			"attachment":null,
			"line_code":null,
			"commit_id":"",
			"noteable_id":390,
			"system":false,
			"st_diff":null,
			"updated_by_id":null,
			"type":null,
			"position":null,
			"original_position":null,
			"resolved_at":null,
			"resolved_by_id":null,
			"discussion_id":"450c34e4c0f9e925bdc6a24c2ae4920d7a394ebc",
			"change_position":null,
			"resolved_by_push":null,
			"url":"http://gitlab.bilibili.co/changhengyuan/test-saga/merge_requests/52#note_3040"
		},
		"repository":{
			"name":"test-saga",
			"url":"git@gitlab.bilibili.co:changhengyuan/test-saga.git",
			"description":"",
			"homepage":"http://gitlab.bilibili.co/changhengyuan/test-saga"
		},
		"merge_request":{
			"assignee_id":null,
			"author_id":15,
			"created_at":"2018-09-26 06:41:55 UTC",
			"description":"",
			"head_pipeline_id":4510,
			"id":390,
			"iid":52,
			"last_edited_at":null,
			"last_edited_by_id":null,
			"merge_commit_sha":null,
			"merge_error":null,
			"merge_params":{
				"force_remove_source_branch":"0"
			},
			"merge_status":"cannot_be_merged",
			"merge_user_id":null,
			"merge_when_pipeline_succeeds":false,
			"milestone_id":null,
			"source_branch":"test-branch",
			"source_project_id":35,
			"state":"opened",
			"target_branch":"master",
			"target_project_id":35,
			"time_estimate":0,
			"title":"Test branch",
			"updated_at":"2018-09-26 06:54:33 UTC",
			"updated_by_id":null,
			"url":"http://gitlab.bilibili.co/changhengyuan/test-saga/merge_requests/52",
			"source":{
				"id":35,
				"name":"test-saga",
				"description":"",
				"web_url":"http://gitlab.bilibili.co/changhengyuan/test-saga",
				"avatar_url":null,
				"git_ssh_url":"git@gitlab.bilibili.co:changhengyuan/test-saga.git",
				"git_http_url":"http://gitlab.bilibili.co/changhengyuan/test-saga.git",
				"namespace":"changhengyuan",
				"visibility_level":20,
				"path_with_namespace":"changhengyuan/test-saga",
				"default_branch":"master",
				"ci_config_path":null,
				"homepage":"http://gitlab.bilibili.co/changhengyuan/test-saga",
				"url":"git@gitlab.bilibili.co:changhengyuan/test-saga.git",
				"ssh_url":"git@gitlab.bilibili.co:changhengyuan/test-saga.git",
				"http_url":"http://gitlab.bilibili.co/changhengyuan/test-saga.git"
			},
			"target":{
				"id":35,
				"name":"test-saga",
				"description":"",
				"web_url":"http://gitlab.bilibili.co/changhengyuan/test-saga",
				"avatar_url":null,
				"git_ssh_url":"git@gitlab.bilibili.co:changhengyuan/test-saga.git",
				"git_http_url":"http://gitlab.bilibili.co/changhengyuan/test-saga.git",
				"namespace":"changhengyuan",
				"visibility_level":20,
				"path_with_namespace":"changhengyuan/test-saga",
				"default_branch":"master",
				"ci_config_path":null,
				"homepage":"http://gitlab.bilibili.co/changhengyuan/test-saga",
				"url":"git@gitlab.bilibili.co:changhengyuan/test-saga.git",
				"ssh_url":"git@gitlab.bilibili.co:changhengyuan/test-saga.git",
				"http_url":"http://gitlab.bilibili.co/changhengyuan/test-saga.git"
			},
			"last_commit":{
				"id":"51e9c3ba2ceac496dbaf55f0db564ab6a15e20eb",
				"message":"add CONTRIBUTORS.md\n",
				"timestamp":"2018-09-17T18:02:13+08:00",
				"url":"http://gitlab.bilibili.co/changhengyuan/test-saga/commit/51e9c3ba2ceac496dbaf55f0db564ab6a15e20eb",
				"author":{
					"name":"哔哩哔哩",
					"email":"bilibili@bilibilideMac-mini.local"
				}
			},
			"work_in_progress":false,
			"total_time_spent":0,
			"human_total_time_spent":null,
			"human_time_estimate":null}}`)
)

func init() {
	dir, _ := filepath.Abs("../../cmd/saga-test.toml")
	flag.Set("conf", dir)
	conf.Init()
	c = New(&dao.Dao{}, &gitlab.Gitlab{})
}

func TestCommandNew(t *testing.T) {
	convey.Convey("New", t, func(ctx convey.C) {
		ctx.Convey("When everything goes positive", func(ctx convey.C) {
			ctx.Convey("Then c should not be nil.", func(ctx convey.C) {
				ctx.So(c, convey.ShouldNotBeNil)
			})
		})
	})
}

func TestCommandExec(t *testing.T) {
	convey.Convey("Exec", t, func(ctx convey.C) {
		var (
			ct    = context.Background()
			cmd   = "+1"
			event = &model.HookComment{}
			repo  = &model.Repo{}
			c     = &Command{}
		)
		_ = json.Unmarshal(gitlabHookCommentTest, event)
		ctx.Convey("When everything goes positive", func(ctx convey.C) {
			err := c.Exec(ct, cmd, event, repo)
			ctx.Convey("Then err should be nil.", func(ctx convey.C) {
				ctx.So(err, convey.ShouldBeNil)
			})
		})
	})
}

func TestCommandRegister(t *testing.T) {
	convey.Convey("register", t, func(ctx convey.C) {
		var (
			cmd = "test_cmd"
			f   cmdFunc
		)
		ctx.Convey("When everything goes positive", func(ctx convey.C) {
			c.register(cmd, f)
			ctx.Convey("No return values", func(ctx convey.C) {
				cmd, ok := c.cmds["test_cmd"]
				ctx.So(ok, convey.ShouldEqual, true)
				ctx.So(cmd, convey.ShouldEqual, f)
			})
		})
	})
}
