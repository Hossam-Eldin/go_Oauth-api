//package http:  package this like controller for mvc patteren
package http

import (
	"net/http"

	"github.com/Hossam-Eldin/go_Oauth-api/src/domain/accesstoken"
	"github.com/gin-gonic/gin"
)

//AccessTokenHandler : interface method
type AccessTokenHandler interface {
	GetByID(*gin.Context)
}

type accesstokenHandler struct {
	service accesstoken.Service
}

//NewHandler :
func NewHandler(service accesstoken.Service) AccessTokenHandler {
	return &accesstokenHandler{
		service: service,
	}
}

func (h *accesstokenHandler) GetByID(c *gin.Context) {
	accessToken, err := h.service.GetByID(c.Param("access_token_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return

	}
	c.JSON(http.StatusOK, accessToken)
}
