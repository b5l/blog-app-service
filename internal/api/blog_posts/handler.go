package blogPosts

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BlogPostsHandler(c *gin.Context) {
	fmt.Println("worked")
	c.JSON(http.StatusOK, "run")
}
