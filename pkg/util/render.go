package util

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//RenderError render error
func RenderError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, fmt.Sprintf("%+v", err))
}

//RenderSuccess render success
func RenderSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}
