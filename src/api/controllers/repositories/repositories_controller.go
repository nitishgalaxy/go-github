package repositories

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nitishgalaxy/go-github/src/api/models/repositories"
	"github.com/nitishgalaxy/go-github/src/api/services"
	"github.com/nitishgalaxy/go-github/src/api/utils/errors"
)

func CreateRepo(c *gin.Context) {
	var request repositories.CreateRepoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		apiErr := errors.NewBadRequestError("invalid json body")
		c.JSON(apiErr.GetStatus(), apiErr)
		return
	}

	result, err := services.RepoService.CreateRepo(request)
	if err != nil {
		c.JSON(err.GetStatus(), err)
		return
	}

	c.JSON(http.StatusCreated, result)
}
