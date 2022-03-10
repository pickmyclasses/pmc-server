package shared

import "testing"

func TestConvertTimestamp(t *testing.T) {
	tests := []struct {
		timestamp string
		translate float32
	}{
		{"09:00am", 9},
		{"10:00am", 10},
		{"00:00am", 0},
		{"11:59pm", 23.59},
		{"03:00pm", 15},
		{"04:32pm", 16.32},
	}

	for _, tt := range tests {
		if actual := ConvertTimestamp(tt.timestamp); actual != tt.translate {
			t.Errorf("TestConvertTimestamp(%s), should be %f, got %f instead.\n",
				tt.timestamp, tt.translate, actual)
		}
	}
}


