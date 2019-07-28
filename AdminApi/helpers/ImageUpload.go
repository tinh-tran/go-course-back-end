package helpers

import (
	"io"
	"net/http"
	"os"
)

func FileUpload(r *http.Request) (string, error) {
	//this function returns the filename(to save in database) of the saved file or an error if it occurs
	r.ParseMultipartForm(32 << 20)

	//ParseMultipartForm parses a request body as multipart/form-data

	var file_name string
	var errors error

	file, handler, err := r.FormFile("file") //retrieve the file from form data
	defer file.Close()                       //close the file when we finish

	if err != nil {
		errors = err

	}
	//this is path which  we want to store the file
	f, err := os.OpenFile("./images/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {

		errors = err
	}
	file_name = handler.Filename
	defer f.Close()
	io.Copy(f, file)
	//here we save our file to our path

	return file_name, errors

}
