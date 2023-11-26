package download

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2/log"
)

func DownloadHandler(ctx *gin.Context) {
	fileId := string(ctx.Query("fileId"))
	log.Info(fmt.Sprintf("C:/Users/noskovm/Desktop/storage/Browse/%s", fileId))
	data, _ := os.ReadFile(fmt.Sprintf("C:/Users/noskovm/Desktop/storage/Browse/%s", fileId))
	_, _ = ctx.Writer.Write(data)
}
