package upper

import (
	"context"
	"time"

	"github.com/namelessup/bilibili/app/interface/main/app-feed/conf"
	article "github.com/namelessup/bilibili/app/interface/openplatform/article/model"
	feed "github.com/namelessup/bilibili/app/service/main/feed/model"
	feedrpc "github.com/namelessup/bilibili/app/service/main/feed/rpc/client"
	"github.com/namelessup/bilibili/library/cache/redis"
	"github.com/namelessup/bilibili/library/ecode"
	"github.com/namelessup/bilibili/library/net/metadata"

	"github.com/pkg/errors"
)

// Dao is feed dao.
type Dao struct {
	// rpc
	feedRPC *feedrpc.Service
	// redis
	redis     *redis.Pool
	expireRds int32
}

// New new a archive dao.
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		// rpc
		feedRPC: feedrpc.New(c.FeedRPC),
		// redis init
		redis:     redis.NewPool(c.Redis.Upper.Config),
		expireRds: int32(time.Duration(c.Redis.Upper.ExpireUpper) / time.Second),
	}
	return
}

// Ping check redis connection
func (d *Dao) Ping(c context.Context) (err error) {
	conn := d.redis.Get(c)
	_, err = conn.Do("SET", "PING", "PONG")
	conn.Close()
	return
}

func (d *Dao) Feed(c context.Context, mid int64, pn, ps int) (fs []*feed.Feed, err error) {
	ip := metadata.String(c, metadata.RemoteIP)
	arg := &feed.ArgFeed{Mid: mid, Pn: pn, Ps: ps, RealIP: ip}
	if fs, err = d.feedRPC.AppFeed(c, arg); err != nil {
		if err == ecode.NothingFound {
			err = nil
			return
		}
		err = errors.Wrapf(err, "%v", arg)
	}
	return
}

func (d *Dao) ArchiveFeed(c context.Context, mid int64, pn, ps int) (fs []*feed.Feed, err error) {
	ip := metadata.String(c, metadata.RemoteIP)
	arg := &feed.ArgFeed{Mid: mid, Pn: pn, Ps: ps, RealIP: ip}
	if fs, err = d.feedRPC.ArchiveFeed(c, arg); err != nil {
		if err == ecode.NothingFound {
			err = nil
			return
		}
		err = errors.Wrapf(err, "%v", arg)
	}
	return
}

func (d *Dao) BangumiFeed(c context.Context, mid int64, pn, ps int) (fs []*feed.Feed, err error) {
	ip := metadata.String(c, metadata.RemoteIP)
	arg := &feed.ArgFeed{Mid: mid, Pn: pn, Ps: ps, RealIP: ip}
	if fs, err = d.feedRPC.BangumiFeed(c, arg); err != nil {
		if err == ecode.NothingFound {
			err = nil
			return
		}
		err = errors.Wrapf(err, "%v", arg)
	}
	return
}

func (d *Dao) Recent(c context.Context, mid, aid int64) (fs []*feed.Feed, err error) {
	ip := metadata.String(c, metadata.RemoteIP)
	arg := &feed.ArgFold{Mid: mid, Aid: aid, RealIP: ip}
	if fs, err = d.feedRPC.Fold(c, arg); err != nil {
		if err == ecode.NothingFound {
			err = nil
			return
		}
		err = errors.Wrapf(err, "%v", arg)
	}
	return
}

func (d *Dao) AppUnreadCount(c context.Context, mid int64, withoutBangumi bool) (unread int, err error) {
	ip := metadata.String(c, metadata.RemoteIP)
	arg := &feed.ArgUnreadCount{Mid: mid, WithoutBangumi: withoutBangumi, RealIP: ip}
	if unread, err = d.feedRPC.AppUnreadCount(c, arg); err != nil {
		err = errors.Wrapf(err, "%v", arg)
	}
	return
}

func (d *Dao) ArticleFeed(c context.Context, mid int64, pn, ps int) (fs []*article.Meta, err error) {
	ip := metadata.String(c, metadata.RemoteIP)
	arg := &feed.ArgFeed{Mid: mid, Pn: pn, Ps: ps, RealIP: ip}
	if fs, err = d.feedRPC.ArticleFeed(c, arg); err != nil {
		if err == ecode.NothingFound {
			err = nil
			return
		}
		err = errors.Wrapf(err, "%v", arg)
	}
	return
}

func (d *Dao) ArticleUnreadCount(c context.Context, mid int64) (unread int, err error) {
	ip := metadata.String(c, metadata.RemoteIP)
	arg := &feed.ArgMid{Mid: mid, RealIP: ip}
	if unread, err = d.feedRPC.ArticleUnreadCount(c, arg); err != nil {
		err = errors.Wrapf(err, "%v", arg)
	}
	return
}
