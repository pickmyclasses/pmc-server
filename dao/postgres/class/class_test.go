package dao

import (
	"pmc_server/init/postgres"
	"testing"
)

func TestGetClassListByComponent(t *testing.T) {
	if err := postgres.Init(); err != nil {
		t.Errorf("Init database failed %s", err)
	}

	tests := []struct {
		components []string
		count      int
	}{
		{[]string{"IVC"}, 0},
	}

	for _, tt := range tests {
		if actual, err := GetClassListByComponent(tt.components); len(*actual) != tt.count || err != nil {
			t.Errorf("TestGetClassListByComponent(%+v), should be %d, got %d instead \n",
				tt.components, tt.count, len(*actual))
		}
	}
}
