package main

import (
	"flag"
	"log"

	"github.com/noskovm/crypto-storage/internal/config"
	"github.com/noskovm/crypto-storage/internal/httpserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "../configs/server-config.yml", "path to config file")
}

func main() {
	flag.Parse()

	config := config.MustLoad(configPath)
	log.Printf("Прочитан конфигурационный файл: %v", config)

	server := httpserver.BuildServer(config)
	log.Fatal(server.Start())

	// router := gin.Default()

	// router.GET("/download/", download.DownloadResorce)
	// router.GET("/upload/", server.getAllTasksHandler)

	// port := flag.String("p", "8081", "port to serve on")
	// directory := flag.String("d", "C:/Users/noskovm/Desktop/storage/", "the directory of static file to host")
	// flag.Parse()

	// http.Handle("/", http.FileServer(http.Dir(*directory)))
	// http.Handle("/upload", http.FileServer(http.Dir(*directory)))

	// log.Printf("Serving %s on HTTP port: %s\n", *directory, *port)
	// log.Fatal(http.ListenAndServe(":"+*port, nil))

}
