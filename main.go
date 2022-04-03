package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"pmc_server/init/aura"
	"pmc_server/init/es"
	"syscall"
	"time"

	"pmc_server/config"
	"pmc_server/init/logger"
	"pmc_server/init/postgres"
	libs "pmc_server/libs/snowflake"
	"pmc_server/routes"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// @title PickMyClasses API Guide
// @version 1.0
// @description General API Description for PickMyClasses
// @termsOfService http://swagger.io/terms/
// @contact.name Kaijie Fu
// @contact.url http://www.swagger.io/support
// @contact.email fio827601499@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost
// @BasePath /user
func main() {
	var err error
	// load config
	if err = config.Init(); err != nil {
		fmt.Printf("Init config failed %s", err)
		return
	}

	// init logger
	if err = logger.Init(viper.GetString("log.mode")); err != nil {
		fmt.Printf("Init logger failed %s", err)
		return
	}
	defer func(l *zap.Logger) {
		err := l.Sync()
		if err != nil {
			fmt.Printf("Close logger failed %s", err)
		}
	}(zap.L())

	// init database
	if err = postgres.Init(); err != nil {
		fmt.Printf("Init database failed %s", err)
		return
	}

	//// init redis
	//if err = redis.Init(); err != nil {
	//	fmt.Printf("Init redis failed %s", err)
	//	return
	//}
	//defer redis.Close()

	// init elastic search
	if err = es.Init(viper.GetString("elastic.url"), viper.GetString("elastic.username"), viper.GetString("elastic.password")); err != nil {
		fmt.Printf("Init res failed %+v", err)
		return
	}

	// init neo4j
	if err = aura.Init(viper.GetString("auradb.uri"), viper.GetString("auradb.username"), viper.GetString("auradb.password")); err != nil {
		fmt.Printf("Init neo4j failed %+v", err)
	}

	// init router
	r := routes.SetUp(viper.GetString("app.mode"))

	// init snowflake
	if err := libs.Init(viper.GetString("snowflake.start_time"), viper.GetInt64("snowflake.machine_id")); err != nil {
		fmt.Printf("Init snowflake failed %v", err)
	}

	// start logic
	var port string
	if viper.GetString("app.mode") == "dev" {
		port = "3000"
	} else {
		port = os.Getenv("PORT")
	}
	if port == "" {
		zap.L().Fatal("port must be set")
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Fatal("listen: %s\n", zap.Error(err))
		}
	}()

	// quit with channel signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	zap.L().Info("Shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server shutdown:", zap.Error(err))
	}

	zap.L().Info("Server exiting")
}
