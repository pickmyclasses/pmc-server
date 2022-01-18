package utils

import (
	"errors"
	"strconv"
	"strings"

	. "pmc_server/consts"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RemoveTopStruct removes the struct name in the error message
func RemoveTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}

func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}

		switch {
		case pageSize > 20:
			pageSize = 20
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func HandlePagination(c *gin.Context, defaultVal string) (int, int, error) {
	pn := c.DefaultQuery("pn", "0")
	pSize := c.DefaultQuery("psize", defaultVal)
	pnInt, err := strconv.Atoi(pn)
	if err != nil {
		return 0, 0, err
	}

	pSizeInt, err := strconv.Atoi(pSize)

	if err != nil {
		return 0, 0, err
	}

	if pnInt < 0 || pSizeInt < 0 {
		return 0, 0, errors.New(BAD_PAGE_ERR)
	}

	return pnInt, pSizeInt, nil
}
