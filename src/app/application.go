package app

import (
	"github.com/Hossam-Eldin/go_Oauth-api/src/domain/accesstoken"
	"github.com/Hossam-Eldin/go_Oauth-api/src/http"
	"github.com/Hossam-Eldin/go_Oauth-api/src/repository/db"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

//StartApplication : start the app with this func
func StartApplication() {

	atHandler := http.NewHandler(accesstoken.NewService(db.New()))
	router.GET("/oauth/at/:access_token_id", atHandler.GetByID)
	router.POST("/oauth/create", atHandler.Create)
	router.Run(":3000")
}
