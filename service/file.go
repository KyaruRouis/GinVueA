package service

import (
	"GinVueA/define"
	"GinVueA/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
	"path/filepath"
	"strconv"
	"time"
)

// UploadFile 上传图片文件
func UploadFile(c *gin.Context) {
	file, err := c.FormFile("fileResource")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "上传失败",
		})
		return
	}

	// 获取文件名称
	fmt.Println(file.Filename)
	// 文件大小
	fmt.Println(file.Size)

	// 获取文件的后缀名
	fileExt := path.Ext(file.Filename)
	fmt.Println(fileExt)
	// 允许上传文件后缀
	allowExt := "jpg,gif,png,bmp,jpeg,JPG"
	// 检查文件上传的后缀
	if !checkFileExt(fileExt, allowExt) {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "上传文件文件格式不正确，文件后缀值允许为:" + allowExt + "的文件",
		})
		return
	}

	// 根据当前时间戳生成新的文件名
	fileNameInt := time.Now().Unix()
	fileNameStr := strconv.FormatInt(fileNameInt, 10)
	// 新的文件名
	fileName := fileNameStr + fileExt
	// 保存上传文件
	filePath := filepath.Join(define.StaticResource, "/", fileName)
	err2 := c.SaveUploadedFile(file, filePath)
	if err2 != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err2,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":     200,
		"msg":      "上传文件成功",
		"fileName": fileName,
	})

}

// 检查文件格式是否合法
func checkFileExt(fileExt string, typeString string) bool {
	// 允许上传文件后缀
	exts := utils.Split(typeString, ",")
	// 是否验证通过
	isValid := false
	for _, v := range exts {
		// 对比文件后缀
		if utils.Equal(fileExt, "."+v) {
			isValid = true
			break
		}
	}
	return isValid
}
