package err

import (
	"errors"
	"github.com/gin-gonic/gin"
	"pmc_server/shared"
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
			if errors.Is(err, shared.ContentNotFoundErr{}) {
				parsedErr = shared.ContentNotFoundErr{}
			}
			if errors.Is(err, shared.ParamIncompatibleErr{}) {
				parsedErr = shared.ParamInsufficientErr{}
			}
			if errors.Is(err, shared.ParamInsufficientErr{}) {
				parsedErr = shared.ParamInsufficientErr{}
			}
			if errors.Is(err, shared.UserInfoNotFoundErr{}) {
				parsedErr = shared.UserInfoNotFoundErr{}
			}
			if errors.Is(err, shared.MalformedIDErr{}) {
				parsedErr = shared.MalformedIDErr{}
			}
			if errors.Is(err, shared.ResourceConflictErr{}) {
				parsedErr = shared.ResourceConflictErr{}
			}
			if errors.Is(err, shared.InfoUnmatchedErr{}) {
				parsedErr = shared.InfoUnmatchedErr{}
			}
			if errors.Is(err, shared.InternalErr{}) {
				parsedErr = shared.InternalErr{}
			}

			c.IndentedJSON(parsedErr.Code(), gin.H{
				shared.ERROR: parsedErr.Error(),
			})
			c.Abort()
			return
		}
	}
}
