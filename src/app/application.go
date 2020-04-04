package app

import (
	"github.com/fvukojevic/bookstore_oauth-api/src/http"
	"github.com/fvukojevic/bookstore_oauth-api/src/repository/db"
	"github.com/fvukojevic/bookstore_oauth-api/src/repository/rest"
	"github.com/fvukojevic/bookstore_oauth-api/src/services"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	tokenHandler := http.NewHandler(services.NewService(db.NewRepository(), rest.NewRepository()))

	router.GET("/oauth/access_token/:access_token_id", tokenHandler.GetById)
	router.POST("/oauth/access_token/", tokenHandler.Create)

	router.Run(":8081")
}
