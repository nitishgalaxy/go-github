package app

import (
	"github.com/nitishgalaxy/go-github/src/api/controllers/polo"
	"github.com/nitishgalaxy/go-github/src/api/controllers/repositories"
)

func mapUrls() {
	router.GET("/marco", polo.ReturnPolo)
	router.POST("/repositories", repositories.CreateRepo)
}
