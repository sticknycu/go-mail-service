package api

import (
	"encoding/json"
	"github.com/emicklei/go-restful/v3"
	log "github.com/sirupsen/logrus"
	"go-mail-service/domain"
	"io/ioutil"
	"net/http"
)

func (api *API) sendEmail(req *restful.Request, resp *restful.Response) {
	body := req.Request.Body
	if body == nil {
		log.Printf("[ERROR] Couldn't read request body")
		resp.WriteError(http.StatusInternalServerError, restful.NewError(http.StatusInternalServerError, "nil body"))
		return
	}

	defer body.Close()

	var err error
	if err != nil {
		log.Printf("[ERROR] Couldn't read request body")
		resp.WriteError(http.StatusInternalServerError, restful.NewError(http.StatusInternalServerError, err.Error()))
		return
	}

	data, err := ioutil.ReadAll(body)
	if err != nil {
		log.Printf("[ERROR] Couldn't read request body")
		resp.WriteError(http.StatusInternalServerError, restful.NewError(http.StatusInternalServerError, err.Error()))
		return
	}

	var emailTemplate domain.EmailTemplate
	unmarshalErr := json.Unmarshal(data, &emailTemplate)
	if unmarshalErr != nil {
		log.Printf("[ERROR] Failed to unmarshal EmailTemplate, err=%v", unmarshalErr)
		resp.WriteError(http.StatusInternalServerError, restful.NewError(http.StatusInternalServerError, err.Error()))
		return
	}

	log.Infof("Sending email to the user %v", emailTemplate.ToMail)
}
