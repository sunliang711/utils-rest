// Package server
// @Project:      nft-studio-backend
// @File:          http.go
// @Author:        eagle
// @Create:        2021/08/09 15:23:31
// @Description:
package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"nft-studio-backend/database"
	"nft-studio-backend/handlers"
	"nft-studio-backend/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type HttpServer struct {
	mode             string
	addr             string
	tls              bool
	certFile         string
	keyFile          string
	corsAllowOrigins []string
	corsAllowMethods []string
	corsAllowHeaders []string
	corsMaxAge       time.Duration
	srv              *http.Server
}

func New(mode string, addr string, tls bool, certFile string, keyFile string, corsAllowOrigins []string, corsAllowMethods []string, corsAllowHeaders []string, corsMaxAge time.Duration) (*HttpServer, error) {
	if mode != "debug" && mode != "release" {
		return nil, errors.New("invalid mode(expect debug or release)")
	}
	server := &HttpServer{
		mode:     mode,
		addr:     addr,
		tls:      tls,
		certFile: certFile,
		keyFile:  keyFile,

		corsAllowOrigins: corsAllowOrigins,
		corsAllowMethods: corsAllowMethods,
		corsAllowHeaders: corsAllowHeaders,
		corsMaxAge:       corsMaxAge,
	}
	return server, nil
}

func (server *HttpServer) Init() {
}

func (server *HttpServer) Start() {
	// setup db
	tables := []interface{}{}
	database.InitDB(tables)

	utils.InitFileUploader()

	// MUST SetMode first
	switch server.mode {
	case "debug":
		gin.SetMode(gin.DebugMode)
	case "release":
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	corsCfg := cors.Config{
		AllowOrigins: server.corsAllowOrigins,
		AllowMethods: server.corsAllowMethods,
		AllowHeaders: server.corsAllowHeaders,
		MaxAge:       server.corsMaxAge,
	}
	logrus.Infof(utils.CorsConfigStringify(&corsCfg))

	router.Use(cors.New(corsCfg))

	// Put normal handlers below
	router.GET("/api/v1/health", handlers.Health)
	// router.GET("/api/v1/PATH", handlers.XXX)
	v1Group := router.Group("/api/v1")

	v1User := v1Group.Group("/user")
	{
		// upload
		v1User.POST("/upload", handlers.UploadFile)
	}

	// Put need-auth handlers below
	// router.GET("/api/v1/PATH", middleware.Auth)
	// router.POST("/api/v1/PATH", middleware.Auth)

	logrus.Infof("Start server on %v, tls enabled: %v", server.addr, server.tls)
	server.srv = &http.Server{
		Addr:    server.addr,
		Handler: router,
	}

	go func() {
		if server.tls {
			if err := server.srv.ListenAndServeTLS(server.certFile, server.keyFile); err != nil && err != http.ErrServerClosed {
				logrus.Fatalf("listen: %s", err)
			}
		} else {
			if err := server.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				logrus.Fatalf("listen: %s", err)
			}
		}
	}()
}

func (server *HttpServer) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.srv.Shutdown(ctx); err != nil {
		log.Fatal("server Shutdown:", err)
	}
	logrus.Infof("server stoped")

	database.Release()
}

func StartHttpServerWithConfig() *HttpServer {
	mode := viper.GetString("server.gin_mode")
	addr := viper.GetString("server.addr")
	tls := viper.GetBool("tls.enable")
	certFile := viper.GetString("tls.cert_file")
	keyFile := viper.GetString("tls.key_file")

	corsAllowOrigins := viper.GetStringSlice("cors.allow_origins")
	corsAllowMethods := viper.GetStringSlice("cors.allow_methods")
	corsAllowHeaders := viper.GetStringSlice("cors.allow_headers")
	corsMaxAge, err := time.ParseDuration(viper.GetString("cors.max_age"))
	if err != nil {
		logrus.Fatalf("parse cors.max_age error: %v", err)
	}

	srv, err := New(mode, addr, tls, certFile, keyFile, corsAllowOrigins, corsAllowMethods, corsAllowHeaders, corsMaxAge)
	if err != nil {
		logrus.Fatalf("new http server error: %v", err)
	}
	srv.Start()

	return srv
}
