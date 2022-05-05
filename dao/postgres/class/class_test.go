package dao

import (
	"pmc_server/model"
	"testing"
)

func TestGetClassListByComponent(t *testing.T) {
	tests := []struct {
		components []string
		count      int
	}{
		{[]string{"IVC"}, 0},
	}

	for _, tt := range tests {
		if actual, err := GetClassListByComponent(tt.components, 0); len(*actual) != tt.count || err != nil {
			t.Errorf("TestGetClassListByComponent(%+v), should be %d, got %d instead \n",
				tt.components, tt.count, len(*actual))
		}
	}
}

func TestIntersection(t *testing.T) {
	t1 := &model.Class{
		Base:     model.Base{},
		CourseID: 123,
	}
	t2 := &model.Class{
		Base:     model.Base{},
		CourseID: 234,
	}

	t3 := &model.Class{
		Base:     model.Base{},
		CourseID: 345,
	}

	tests := []struct {
		s1 []model.Class
		s2 []model.Class
		e  []model.Class
	}{
		{
			[]model.Class{*t1, *t2},
			[]model.Class{*t1, *t3},
			[]model.Class{*t1},
		},
	}
	for _, tt := range tests {
		t.Error(tt.s2)
	}
}
