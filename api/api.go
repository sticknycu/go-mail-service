package api

import "github.com/emicklei/go-restful/v3"

type API struct {
}

func NewAPI() *API {
	return &API{}
}

func (api *API) RegisterRoutes(ws *restful.WebService) {
	ws.Path("/api")
	ws.Route(ws.POST("/email").To(api.handleSendEmail))
}
