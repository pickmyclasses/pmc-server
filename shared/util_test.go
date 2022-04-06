package shared

import (
	"reflect"
	"testing"
)

func TestConvertTimestamp(t *testing.T) {
	tests := []struct {
		timestamp string
		translate float32
	}{
		{"09:00am", 9},
		{"10:00am", 10},
		{"00:00am", 0},
		{"11:59am", 11.59},
		{"12:00pm", 12},
		{"11:59pm", 23.59},
		{"03:00pm", 15},
		{"04:32pm", 16.32},
		{"12:50pm", 12.50},
	}

	for _, tt := range tests {
		if actual, _ := ConvertTimestamp(tt.timestamp); actual != tt.translate {
			t.Errorf("TestConvertTimestamp(%s), should be %f, got %f instead.\n",
				tt.timestamp, tt.translate, actual)
		}
	}
}

func TestToFixed(t *testing.T) {
	tests := []struct {
		origin    float64
		translate float64
	}{
		{0.333333, 0.33},
		{11223.323, 11223.32},
		{78.9999, 79.00},
		{1000000.999999999, 1000001.00},
		{78.563, 78.56},
		{78.588, 78.59},
	}

	for _, tt := range tests {
		if actual := ToFixed(tt.origin, 2); actual != tt.translate {
			t.Errorf("TestToFixed(%f, %d), should be %f, got %f instead.\n",
				tt.origin, 2, tt.translate, actual)
		}
	}

	longerFloats := []struct {
		origin    float64
		translate float64
	}{
		{12900000000.1200000000000000000, 12900000000.120000},
		{12900.0, 12900.000000},
		{12900.012, 12900.012000},
	}

	for _, tt := range longerFloats {
		if actual := ToFixed(tt.origin, 6); actual != tt.translate {
			t.Errorf("TestToFixed(%f, %d), should be %f, got %f instead.\n",
				tt.origin, 2, tt.translate, actual)
		}
	}
}

func TestRemoveTopStruct(t *testing.T) {
	tests := []struct {
		origin    map[string]string
		translate map[string]string
	}{
		{map[string]string{"error": "unknown error"}, map[string]string{"error": "unknown error"}},
	}

	for _, tt := range tests {
		if actual := RemoveTopStruct(tt.origin); reflect.DeepEqual(tt.translate, actual) {
			t.Errorf("TestToFixed(%+v), should be %+v, got %+v instead.\n",
				tt.origin, tt.translate, actual)
		}
	}
}

func TestIntersection(t *testing.T) {
	tests := []struct {
		s1           []int64
		s2           []int64
		intersection []int64
	}{
		{[]int64{2, 4, 5, 7}, []int64{2, 4, 5, 7}, []int64{2, 4, 5, 7}},
		{[]int64{1, 3, 5, 7}, []int64{2, 4, 6, 8}, []int64{}},
		{[]int64{4, 6, 5, 4}, []int64{4, 6}, []int64{4, 6}},
	}

	for _, tt := range tests {
		actual := Intersection(tt.s1, tt.s2)
		for i, ttt := range tt.intersection {
			if actual[i] != ttt {
				t.Errorf("TestIntersection(%+v, %+v), should be %+v, got %+v instead \n",
					tt.s1, tt.s2, tt.intersection, actual)
			}
		}
	}
}

func TestContains(t *testing.T) {
	tests := []struct {
		s   []string
		e   string
		res bool
	}{
		{[]string{"test", "test1", "test2"}, "test1", true},
		{[]string{"test", "test1", "test2"}, "test3", false},
		{[]string{}, "test3", false},
	}

	for _, tt := range tests {
		if actual := Contains(tt.s, tt.e); actual != tt.res {
			t.Errorf("TestContains(%+v, %s), should be %v, got %v instead \n",
				tt.s, tt.e, tt.res, actual)
		}
	}
}

func TestConvertSliceToDateString(t *testing.T) {
	tests := []struct {
		s []int
		d string
	}{
		{[]int{1}, "Mo"},
		{[]int{1, 2, 3}, "MoTuWe"},
		{[]int{1, 3, 4, 5}, "MoWeThFr"},
	}

	for _, tt := range tests {
		if actual := ConvertSliceToDateString(tt.s); actual != tt.d {
			t.Errorf("TestConvertSliceToDateString(%+v), should be %s, got %s instead \n",
				tt.s, tt.d, actual)
		}
	}
}

func TestParseString(t *testing.T) {
	tests := []struct {
		origin string
		letters string
		number string
	}{
		{"CS4400", "CS", "4400"},
		{"HE EDU310", "HE EDU", "310"},
		{"Math2210", "Math", "2210"},
		{"CS5500", "CS", "5500"},
	}

	for _, tt := range tests {
		if l, n := ParseString(tt.origin, false); l != tt.letters || n != tt.number {
			t.Errorf("TestParseString(%+v), received letters %s and numbers %s \n",
				tt.origin, l, n)
		}
	}
}

func TestSeparateLettersAndNums(t *testing.T) {
	tests := []struct {
		origin string
		converted string
	} {
		{"CS4400", "CS 4400"},
		{"CS 4000", "CS 4000"},
		{"Math 2210", "Math 2210"},
		{"Math2210", "Math 2210"},
		{"EP SN5000", "EP SN 5000"},
		{"ESL1", "ESL 1"},
	}

	for _, tt := range tests {
		if actual := SeparateLettersAndNums(tt.origin); actual != tt.converted {
			t.Errorf("TestSeparateLettersAndNums(%+v), should be %s, got %s instead \n",
				tt.origin, tt.converted, actual)
		}
	}
}
