package polo

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReturnPolo(c *gin.Context) {
	log.Println("In polo controller")
	c.String(http.StatusOK, "polo")
}
