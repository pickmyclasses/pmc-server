package migrate

import (
	"pmc_server/init/es"

	"github.com/spf13/viper"
)

func ESCourse() error {
	if err := es.Init(viper.GetString("elastic.url"), viper.GetString("elastic.username"), viper.GetString("elastic.password")); err != nil {
		return err
	}

	// create index for course

	return nil
}
