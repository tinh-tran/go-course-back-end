package common

import (
	"fmt"
)

var config = Config{}

func init() {
	fmt.Println("Init service ....")
	config.Read()
}
func GetServicePort() string {
	return config.ServicePort
}
func GetFilePathCat() string {
	return config.FilePathCat
}
func GetFilePathCour() string {
	return config.FilePathCour
}

func GetFilePathVideo() string {
	return config.FilePathVideo
}

func GetFilePathPdf() string {
	return config.FilePathPdf
}
