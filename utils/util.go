package utils

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"
	"unicode"

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

func GetLetterInfo(str string) (letter, number string) {
	var l, n []rune
	for index, r := range []rune(str) {
		if r >= 'A' && r <= 'Z' {
			if index != 0 && str[index-1] >= '0' && str[index-1] <= '9' {
				break
			}
			if index != len(str) && str[index+1] >= 'a' && str[index+1] <= 'z' {
				continue
			}
			l = append(l, r)
		}
		if r >= '0' && r <= '9' {
			n = append(n, r)
		}
		if r >= 'a' && r <= 'z' {
			continue
		}
	}
	return string(l), string(n)
}

// ParseString separate letters and nums in a string
func ParseString(s string, ignoreSpace bool) (letters, numbers string) {
	var l, n []rune
	for _, r := range s {
		switch {
		case r >= 'A' && r <= 'Z':
			l = append(l, r)
		case r >= 'a' && r <= 'z':
			l = append(l, r)
		case unicode.IsSpace(r):
			{
				if !ignoreSpace {
					l = append(l, r)
				}
			}
		case r >= '0' && r <= '9':
			n = append(n, r)
		}
	}
	return string(l), string(n)
}

// ParseDate parses a date time info to integers
func ParseDate(dates string) []int {
	var daysInt []int
	if dates == "" || len(dates) == 0 {
		return daysInt
	}
	dates = strings.ToLower(dates)
	dateMap := GenerateDateMap()

	var curStr string
	for _, r := range dates {
		curStr += string(r)
		if val, ok := dateMap[curStr]; ok {
			daysInt = append(daysInt, val)
			curStr = ""
		}
	}

	return daysInt
}

// ParseTime parses given time string to a start and finish time
func ParseTime(t string) (start, end string) {
	if t == "" {
		return
	}
	t = strings.ToLower(t)
	dateMap := GenerateDateMap()
	timeStr := string(t[0])
	for i := 1; i < len(t); i++ {
		if i != 0 && t[i-1] == 'm' && t[i] == 'm' {
			break
		}
		timeStr += string(t[i])
		if _, ok := dateMap[timeStr]; ok {
			timeStr = ""
		}
	}
	timeStrSplit := strings.Split(timeStr, "-")
	start, end = timeStrSplit[0], timeStrSplit[1]
	return
}

// GenerateDateMap generates a map to represent mapping between date and number
func GenerateDateMap() map[string]int {
	dateMap := make(map[string]int)
	dateMap["mo"] = 1
	dateMap["tu"] = 2
	dateMap["we"] = 3
	dateMap["th"] = 4
	dateMap["fr"] = 5
	dateMap["st"] = 6
	dateMap["su"] = 7
	return dateMap
}

// GetJson gets the json object from a response
func GetJson(url string, target interface{}) error {
	client := &http.Client{Timeout: 10 * time.Second}
	r, err := client.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
