package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/my.blogs/config"
	"github.com/my.blogs/global"
	"github.com/my.blogs/model"
	"github.com/my.blogs/model/system"
	"github.com/my.blogs/router"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	initialize()
	//quit()
}

func getGormConfig(gormConfig *config.GormConfig) *gorm.Config {
	return &gorm.Config{
		DryRun: gormConfig.DryRun,
	}
}

func initialize() {
	// 读取yaml配置
	file, err := os.ReadFile("server/config.yaml")
	if err != nil {
		panic(err)
	}
	c := new(config.Server)
	err = yaml.Unmarshal(file, c)
	fmt.Println(c)
	if err != nil {
		panic(err)
	}
	// 初始化mysql连接
	db, err := gorm.Open(mysql.Open(c.Mysql.GetDSN()), getGormConfig(&c.GormConfig))
	if err != nil {
		panic(err)
	}
	// 建表
	err = db.AutoMigrate(&model.User{}, &model.Comment{}, &model.Post{}, &system.ErrorLog{})
	if err != nil {
		panic(err)
	}
	// 初始化gin
	global.GORM_DB = db
	global.GLOBAL_CONFIG = c
	global.GLOBAL_ROUTERS = router.Routers()

}

func quit() {
	srv := &http.Server{
		Addr:    ":8080",
		Handler: global.GLOBAL_ROUTERS.Handler(),
	}
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	// kill (no params) by default sends syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Println("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
