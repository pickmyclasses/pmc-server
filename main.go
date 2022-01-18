package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"pmc_server/config"
	"pmc_server/init/logger"
	"pmc_server/init/postgres"
	"pmc_server/init/redis"
	libs "pmc_server/libs/snowflake"
	"pmc_server/routes"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

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

	// init redis
	if err = redis.Init(); err != nil {
		fmt.Printf("Init redis failed %s", err)
		return
	}
	defer redis.Close()

	// init router
	r := routes.SetUp(viper.GetString("app.mode"))

	// init snowflake
	if err := libs.Init(viper.GetString("snowflake.start_time"), viper.GetInt64("snowflake.machine_id")); err != nil {
		fmt.Printf("Init snowflake failed %v", err)
	}

	// start logic
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", viper.GetInt("app.port")),
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
