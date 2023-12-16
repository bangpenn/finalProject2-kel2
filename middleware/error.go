package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/novita/finalproject2-revisi/pkg/errs"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		err := c.Errors.Last()
		if err != nil {
			switch e := err.Err.(type) {
			case *errs.CustomError:
				c.AbortWithStatusJSON(e.Code, gin.H{
					"error": e.Message,
				})
				return
			default:
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
				return
			}
		}
	}
}
