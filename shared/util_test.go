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
		{"12:00pm", 24},
		{"11:59pm", 23.59},
		{"03:00pm", 15},
		{"04:32pm", 16.32},
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
