package err

import (
	"pmc_server/shared"

	"github.com/gin-gonic/gin"
)

func JsonErrReporter() gin.HandlerFunc {
	return jsonErrReporter(gin.ErrorTypeAny)
}

// jsonErrReporter is the middleware that reports
// this middleware basically does all the jobs that intercepts errors and pass out as JSON
func jsonErrReporter(errType gin.ErrorType) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		detectedErr := c.Errors.ByType(errType)

		if len(detectedErr) > 0 {
			err := detectedErr[0].Err
			var parsedErr shared.AppErr
			switch err.(type) {
			case *shared.ContentNotFoundErr:
				parsedErr = err.(*shared.ContentNotFoundErr)
			case *shared.UserInfoNotFoundErr:
				parsedErr = err.(*shared.UserInfoNotFoundErr)
			case *shared.InternalErr:
				parsedErr = err.(*shared.InternalErr)
			case *shared.ParamInsufficientErr:
				parsedErr = err.(*shared.ParamInsufficientErr)
			case *shared.ParamIncompatibleErr:
				parsedErr = err.(*shared.ParamIncompatibleErr)
			case *shared.MalformedIDErr:
				parsedErr = err.(*shared.MalformedIDErr)
			default:
				parsedErr = err.(*shared.InternalErr)
			}
			c.IndentedJSON(parsedErr.Code(), parsedErr)
			c.Abort()
			return
		}
	}
}

