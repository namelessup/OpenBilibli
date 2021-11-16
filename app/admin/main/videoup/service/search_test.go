package service

import (
	"context"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/namelessup/bilibili/app/admin/main/videoup/model/search"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"testing"
)

func TestService_SearchArchive(t *testing.T) {
	var (
		c = &bm.Context{
			Context: context.Background(),
		}
		ap = &search.ArchiveParams{
			Aids: "35084740,35088598,35088785,35084740,35010110,35088598,35088785,29398125,35089978,35091989,35066022,35093090,35093223,35066022,35089909,35089586,35093908,35093223,35089909,13654260,35095497,35094073,35095209,35095209,13654260,35094073,35097561,35097561,35097561,35099099,35078689,35099099,33764038,35100034,35100034,35101407,35102575,35102934,35103059,35104205,35105495,35081994,35106475,35101407,35103059,35105495,35081994,35102934,35106475,35107286,35106731,35060418,35108599,35108463,35109748,35111706,35110453,35113524,35113604,35111549,35114277,35114010,35114736,35114637,35116637,35116836,35114534,35117091,32831400,35118300,35119337,9286322,35107286,35060418,35106731,35108463,35109748,35111706,35110453,35113524,35113604,35111549,35114736,32831400,35114277,35114637,35118300,35119337,35118873,35052973,35118873,35052973,35122858,35122995,35097561,35097561,35125531,35126909,35126641,35126651,35127290,35126553,35127729,35128312,35127395,35127569,21342393,35127488,35129538,34881916,35127290,35075183,23584443,35130428,23191210,35129543,35132998,35133396,35134689",
		}
	)
	Convey("SearchArchive", t, WithService(func(s *Service) {
		_, err := svr.SearchArchive(c, ap)
		So(err, ShouldBeNil)
	}))
}
