package shared

import (
	"encoding/json"
	"errors"
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

// GetLetterInfo gets the letters and number separation from the given string
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

func ConvertOfferDate(offerDates string) []int {
	mapping := make(map[string]int)
	mapping["mo"] = 1
	mapping["tu"] = 2
	mapping["we"] = 3
	mapping["th"] = 4
	mapping["fr"] = 5

	offerLower := []rune(strings.ToLower(offerDates))
	res := make([]int, 0)
	curStr := ""
	for _, s := range offerLower {
		if s == '-' || s == ' ' {
			continue
		}
		curStr = curStr + string(s)
		if num, found := mapping[curStr]; found {
			res = append(res, num)
			curStr = ""
		}
	}
	return res
}

func ConvertSliceToDateString(dates []int) (res string) {
	mapping := make(map[int]string)
	mapping[1] = "Mo"
	mapping[2] = "Tu"
	mapping[3] = "We"
	mapping[4] = "Th"
	mapping[5] = "Fr"

	for _, date := range dates {
		res += mapping[date]
	}
	return
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
func ConvertTimestamp(timestamp string) (float32, error) {
	offerTime := []rune(strings.ToLower(timestamp))
	var res float32
	hour := 0
	var mins float32
	if unicode.IsDigit(offerTime[0]) {
		first, _ := strconv.Atoi(string(offerTime[0]))
		hour += first * 10
	}
	if unicode.IsDigit(offerTime[1]) {
		second, _ := strconv.Atoi(string(offerTime[1]))
		hour += second
	}
	if offerTime[2] != ':' {
		return 0, errors.New("unable to convert time stamp")
	}
	if unicode.IsDigit(offerTime[3]) {
		first, _ := strconv.Atoi(string(offerTime[3]))
		mins += float32(first) * 0.1
	}
	if unicode.IsDigit(offerTime[4]) {
		first, _ := strconv.Atoi(string(offerTime[4]))
		mins += float32(first) * 0.01
	}
	if unicode.IsLetter(offerTime[5]) {
		if offerTime[5] == 'p' && !(offerTime[0] == '1' && offerTime[1] == '2') {
			res += 12
		}
	}
	res += float32(hour) + mins
	return res, nil
}

// Intersection finds intersection of two slices
func Intersection(s1, s2 []int64) (inter []int64) {
	hash := make(map[interface{}]bool)
	for _, e := range s1 {
		hash[e] = true
	}

	for _, e := range s2 {
		if hash[e] {
			inter = append(inter, e)
		}
	}
	return
}

// Contains checks if an element is in the slice
func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
