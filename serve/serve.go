package serve

import (
	"github.com/gin-gonic/gin"
)

type Configuration struct {
	ListenAddr  string `yaml:"listenAddr"`
	NotifyUrl   string `yaml:"notifyUrl"`
	ReturnUrl   string `yaml:"returnUrl"`
	LocalUrl    string `yaml:"localUrl"`
	ClientId    string `yaml:"clientId"`
	ClientKey   string `yaml:"clientKey"`
	ThreePriKey string `yaml:"threePriKey"`
}

func NewConfiguration() *Configuration {
	return &Configuration{}
}

type Serve struct {
	cfg *Configuration
}

func NewServer(cfg *Configuration) *Serve {
	return &Serve{
		cfg: cfg,
	}
}

func (s *Serve) Run() error {
	UrlMap = make(map[uint8]string)
	UrlMap[NotifyUrl] = s.cfg.NotifyUrl
	UrlMap[ReturnUrl] = s.cfg.ReturnUrl
	UrlMap[LocalUrl] = s.cfg.LocalUrl
	UrlMap[ClientId] = s.cfg.ClientId
	UrlMap[ClientKey] = s.cfg.ClientKey
	UrlMap[ThreePriKey] = s.cfg.ThreePriKey

	r := NewRouter()
	r.Run(s.cfg.ListenAddr)
	return nil
}

var UrlMap map[uint8]string

const (
	NotifyUrl uint8 = iota + 1
	ReturnUrl
	LocalUrl
	ClientId
	ClientKey
	ThreePriKey
)

func GetConf(url uint8) string {
	return UrlMap[url]
}

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	v1 := router.Group("v1")
	gp := v1.Group("order")
	acc := new(OrderController)
	gp.POST("create", acc.Create)
	gp.POST("notify", acc.DispatchReturn)
	return router
}
