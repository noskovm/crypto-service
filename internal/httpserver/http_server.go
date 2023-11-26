package httpserver

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/noskovm/crypto-storage/internal/config"
	"github.com/noskovm/crypto-storage/internal/handlers/download"
	"github.com/noskovm/crypto-storage/internal/handlers/upload"
	"github.com/sirupsen/logrus"
)

type HTTPServer struct {
	hhtpServer *http.Server
	config     *config.HTTPServerConfig
	logger     *logrus.Logger
	router     *gin.Engine
}

// возвращает экземпляр *http.Server с заполненными полями конфигов
func BuildServer(cfg *config.HTTPServerConfig) *HTTPServer {
	s := &HTTPServer{
		hhtpServer: &http.Server{
			Addr: config.ParseAdress(cfg),
		},
		config: cfg,
		logger: logrus.New(),
	}

	return s
}

func (s *HTTPServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()

	s.logger.Info("Запуск сервера ...")
	return s.hhtpServer.ListenAndServe()
}

func (s *HTTPServer) Shutdown(ctx context.Context) error {
	return s.hhtpServer.Shutdown(ctx)
}

func (s *HTTPServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

func (s *HTTPServer) configureRouter() {
	router := gin.Default()

	router.StaticFS("/fs", gin.Dir("C:/Users/noskovm/Desktop/storage", true))
	router.GET("/download:fileId", download.DownloadHandler)
	router.POST("/upload", upload.UploadHandler)

	s.hhtpServer.Handler = router
}
