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
	router.GET("/ouath/at/:access_token_id", atHandler.GetByID)
	router.Run(":3000")
}
