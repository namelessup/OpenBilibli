package version

import (
	"context"
	"github.com/namelessup/bilibili/app/interface/main/creative/model/version"
	"github.com/namelessup/bilibili/library/ecode"
)

// Versions fn
func (s *Service) versionMap(c context.Context) (versions map[string][]*version.Version, err error) {
	if s.VersionCache == nil {
		err = ecode.NothingFound
		return
	}
	versions = make(map[string][]*version.Version)
	for _, v := range s.VersionCache {
		vs := &version.Version{
			ID:       v.ID,
			Ty:       v.Ty,
			Title:    v.Title,
			Content:  v.Content,
			Link:     v.Link,
			Ctime:    v.Ctime,
			Dateline: v.Dateline,
		}
		versions[vs.Ty] = append(versions[vs.Ty], vs)
	}
	return
}
