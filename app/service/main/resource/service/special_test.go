package service

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	pb "github.com/namelessup/bilibili/app/service/main/resource/api/v1"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Special(t *testing.T) {
	Convey("get app banner", t, WithService(func(s *Service) {
		req := pb.RelateRequest{Id: 1, MobiApp: "iphone", Build: 7000}
		tmp, err := s.Relate(context.Background(), &req)
		if err != nil {
			panic(err)
		}
		byte, _ := json.Marshal(tmp)
		fmt.Println(string(byte))
	}))
}
