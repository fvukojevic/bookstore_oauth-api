package app

import (
	"github.com/fvukojevic/bookstore_oauth-api/src/clients/cassandra"
	"github.com/fvukojevic/bookstore_oauth-api/src/domain/access_token"
	"github.com/fvukojevic/bookstore_oauth-api/src/http"
	"github.com/fvukojevic/bookstore_oauth-api/src/repository/db"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	session, dbErr := cassandra.GetSession()
	if dbErr != nil {
		panic(dbErr)
	}
	session.Close()

	tokenHandler := http.NewHandler(access_token.NewService(db.NewRepository()))

	router.GET("/oauth/access_token/:access_token_id", tokenHandler.GetById)
	router.POST("/oauth/access_token/", tokenHandler.Create)

	router.Run(":8080")
}
