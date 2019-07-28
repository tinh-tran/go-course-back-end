package services

import (
	. "ImgService/common"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type VideoPdfService struct {
}

func (VpService *VideoPdfService) UploadVideo(c *gin.Context) {
	//-----------
	// Read file
	//-----------

	// Source
	file, err := c.FormFile("File")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}
	src, err := file.Open()
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}
	defer src.Close()
	filepath := GetFilePathVideo()
	// Destination
	dst, err := os.Create(filepath + file.Filename)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}
	c.Writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully with fields ", file.Filename))
}
func (VpService *VideoPdfService) UploadPdf(c *gin.Context) {
	//-----------
	// Read file
	//-----------

	// Source
	file, err := c.FormFile("File")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}
	src, err := file.Open()
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}
	defer src.Close()
	filepath := GetFilePathPdf()
	// Destination
	dst, err := os.Create(filepath + file.Filename)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}
	c.Writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully with fields ", file.Filename))
}
