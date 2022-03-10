package shared

import (
	"encoding/json"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"
	"unicode"

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

// Paginate takes the page number and page size to paginate the data
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

// HandlePagination checks for the input parameters for pagination and returns the paginated data
func HandlePagination(c *gin.Context, defaultVal string) (int, int, error) {
	pn := c.DefaultQuery("pn", "0")
	pSize := c.DefaultQuery("psize", defaultVal)
	pnInt, err := strconv.Atoi(pn)
	if err != nil {
		return 0, 0, ParamIncompatibleErr{}
	}

	pSizeInt, err := strconv.Atoi(pSize)

	if err != nil {
		return 0, 0, ParamIncompatibleErr{}
	}

	if pnInt < 0 || pSizeInt < 0 {
		return 0, 0, ParamIncompatibleErr{}
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

// round up a floats
func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

// ToFixed fixed a float number to the given precision
func ToFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

// ConvertTimestamp converts timestamp string such as 6:00am to 6, and 8:45pm to 20.45
func ConvertTimestamp(timestamp string) int {
	offerTime := []rune(strings.ToLower(timestamp))
	res := 0
	hours := 0
	minutes := 0
	plus := 0
	appeared := false
	for i, t := range offerTime {
		if t == '-' || t == ' ' {
			continue
		}
		if t == ':' {
			appeared = true
			continue
		}
		if unicode.IsDigit(t) {
			if appeared {
				if offerTime[i-1] == ':' {
					first, _ := strconv.Atoi(string(t))
					minutes += first / 10
				} else {
					second, _ := strconv.Atoi(string(t))
					minutes += second / 100
				}
			} else {
				if i == 0 {
					first, _ := strconv.Atoi(string(t))
					minutes += first * 10
				} else {
					second, _ := strconv.Atoi(string(t))
					minutes += second
				}
			}
		}

		if t == 'a' {
			continue
		}
		if t == 'p' {
			plus += 12
		}
	}
	res = hours + minutes + plus

	return res
}
