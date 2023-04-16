package handler

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/thesisK19/buildify/app/file-mgt/internal/constant"
	"github.com/thesisK19/buildify/app/file-mgt/internal/util"
)

func UploadImageHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("UploadImageHandler processing...")
	err := r.ParseForm()
	if err != nil {
		log.Printf("UploadImageHandler | Failed to parse form %s", err.Error())
		http.Error(w, fmt.Sprintf("Failed to parse form: %s", err.Error()), http.StatusBadRequest)
		return
	}

	// Get the file from the request
	r.Body = http.MaxBytesReader(w, r.Body, 5*1024*1024) // Max 10 Mb
	file, header, err := r.FormFile("image")
	if err != nil {
		log.Printf("UploadImageHandler | Exceeded limit 5MB %s", err.Error())
		Send(w, http.StatusBadRequest, UploadImageResponse{
			Code:    constant.Code_INVALID_ARGUMENT,
			Message: err.Error(),
		})
		return
	}
	defer file.Close()

	ext := filepath.Ext(header.Filename)
	name := strings.TrimSuffix(header.Filename, ext)
	filename := fmt.Sprintf("%s_%s%s", name, util.GenUUID(5), ext)

	url, err := util.UploadFile(file, "images/"+filename)
	if err != nil {
		log.Printf("UploadImageHandler | Fail to UploadFile %s", err.Error())
		Send(w, http.StatusInternalServerError, UploadImageResponse{
			Code:    constant.Code_INTERNAL,
			Message: err.Error(),
		})
		return
	}

	// Send a success response
	Send(w, http.StatusOK, UploadImageResponse{
		Code:    constant.Code_OK,
		Message: constant.Code_OK.String(),
		Url:     *url,
	})
}
