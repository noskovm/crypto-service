package upload

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//const dst

func UploadHandler(ctx *gin.Context) {
	file, _ := ctx.FormFile("file")
	log.Println(file.Filename)

	// Upload the file to specific dst.
	if err := ctx.SaveUploadedFile(file, fmt.Sprintf("C:/Users/noskovm/Desktop/storage/Browse/%s", file.Filename)); err != nil {
		ctx.JSON(http.StatusOK, fmt.Sprintf(`"error":"%s"`, err))
	}

	ctx.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	ctx.JSON(http.StatusOK, `"value":"uploaded"`)
}
