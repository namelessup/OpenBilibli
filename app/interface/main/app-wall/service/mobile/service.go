package mobile

import (
	"bufio"
	"io"
	"os"
	"strings"
	"time"

	"github.com/namelessup/bilibili/app/interface/main/app-wall/conf"
	mobileDao "github.com/namelessup/bilibili/app/interface/main/app-wall/dao/mobile"
	"github.com/namelessup/bilibili/app/interface/main/app-wall/model"
	"github.com/namelessup/bilibili/app/interface/main/app-wall/model/mobile"
	"github.com/namelessup/bilibili/library/log"
	"github.com/namelessup/bilibili/library/stat/prom"
)

type Service struct {
	c             *conf.Config
	dao           *mobileDao.Dao
	tick          time.Duration
	mobileIpCache []*mobile.MobileIP
	ipPath        string
	// prom
	pHit  *prom.Prom
	pMiss *prom.Prom
}

func New(c *conf.Config) (s *Service) {
	s = &Service{
		c:             c,
		dao:           mobileDao.New(c),
		tick:          time.Duration(c.Tick),
		mobileIpCache: []*mobile.MobileIP{},
		ipPath:        c.IPLimit.MobileIPFile,
		// prom
		pHit:  prom.CacheHit,
		pMiss: prom.CacheMiss,
	}
	s.loadIP()
	return
}

func (s *Service) loadIP() {
	var (
		ip   *mobile.MobileIP
		file *os.File
		line string
		err  error
		ips  []*mobile.MobileIP
	)
	if file, err = os.Open(s.ipPath); err != nil {
		log.Error("mobileIPFile is null")
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		if line, err = reader.ReadString('\n'); err != nil {
			if err == io.EOF {
				err = nil
				break
			}
			continue
		}
		lines := strings.Fields(line)
		if len(lines) < 3 {
			continue
		}
		ip = &mobile.MobileIP{
			IPStartUint: model.InetAtoN(lines[1]),
			IPEndUint:   model.InetAtoN(lines[2]),
		}
		ips = append(ips, ip)
	}
	s.mobileIpCache = ips
	log.Info("loadMobileIPCache success")
}
