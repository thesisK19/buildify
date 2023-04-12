package handler

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/thesisK19/buildify/app/file-management/internal/constant"
	"github.com/thesisK19/buildify/app/file-management/internal/util"
)

func UploadImageHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to parse form: %s", err.Error()), http.StatusBadRequest)
		return
	}

	// Get the file from the request
	r.Body = http.MaxBytesReader(w, r.Body, 10*1024*1024) // Max 10 Mb
	file, header, err := r.FormFile("image")
	if err != nil {
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
