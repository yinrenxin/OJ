package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"inherited/internal/services"
	"log"
	"net/http"
	"time"
)

const (
	OJ_DATA          = "/home/judge/data"
	OJ_ZIP_TEMP_DATA = OJ_DATA + "/tempzip"
	SESSION_USER_KEY = "dongzhiwei"
)

func QueryAllProblem(ctx *gin.Context) {
	problem := new(services.Problem)
	pro_set, err := problem.QueryAllProblem()
	if err != nil {
		fmt.Printf("[api.QueryAllProblem], err: %v", err)

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": pro_set,
	})

	return

}

// FileUpload 上传测试数据zip
func FileUpload(ctx *gin.Context) {

	file, err := ctx.FormFile("file")
	key := time.Now().Format("2006-01-02 15:04:05")
	zipDir := OJ_ZIP_TEMP_DATA + key + file.Filename
	fileName := key+file.Filename
	if err != nil {
		log.Println("[api.FileUpload]", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "文件上传失败, name因用file"})
		return
	}

	// 保存文件
	//dir := "/Users/apple/test/" + key + file.Filename
	err = ctx.SaveUploadedFile(file, zipDir)

	if err != nil {
		log.Println(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":      "上传成功",
		"fileName": fileName,
	})
}

func AddProblem(ctx *gin.Context) {
	//data := models.Problem{}
	//err := ctx.ShouldBindJSON(&data)

}
