package service

import (
	runtime "github.com/banzaicloud/logrus-runtime-formatter"
	"github.com/emicklei/go-restful/v3"
	log "github.com/sirupsen/logrus"
	"go-mail-service/api"
	"net/http"
	"os"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) StartWebService() {
	formatter := runtime.Formatter{ChildFormatter: &log.TextFormatter{
		FullTimestamp:          true,
		DisableLevelTruncation: true,
	}}

	formatter.Line = true
	log.SetFormatter(&formatter)
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)

	ws := new(restful.WebService)
	restful.Add(ws)

	apiManager := api.NewAPI()
	apiManager.RegisterRoutes(ws)

	log.Printf("Started mail service on port 9094")
	log.Fatal(http.ListenAndServe(":9094", nil))
}
