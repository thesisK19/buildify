package handler

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/thesisK19/buildify/app/file_mgt/internal/constant"
	"github.com/thesisK19/buildify/app/file_mgt/internal/util"
)

func UploadAvatarHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("UploadAvatarHandler processing...")

	key := "avatar"
	remoteFolder := "avatars"

	err := r.ParseForm()
	if err != nil {
		log.Printf("UploadAvatarHandler | failed to parse form %s", err.Error())
		http.Error(w, fmt.Sprintf("failed to parse form: %s", err.Error()), http.StatusBadRequest)
		return
	}

	// Get the file from the request
	r.Body = http.MaxBytesReader(w, r.Body, constant.SIZE_5MB) // max 5mb
	file, header, err := r.FormFile(key)
	if err != nil {
		log.Printf("UploadAvatarHandler | Exceeded limit 5MB %s", err.Error())
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

	url, err := util.UploadFile(file, remoteFolder+"/"+filename)
	if err != nil {
		log.Printf("UploadAvatarHandler | Fail to UploadFile %s", err.Error())
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

func UploadImageHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("UploadImageHandler processing...")

	key := "image"
	remoteFolder := "images"

	err := r.ParseForm()
	if err != nil {
		log.Printf("UploadImageHandler | failed to parse form %s", err.Error())
		http.Error(w, fmt.Sprintf("failed to parse form: %s", err.Error()), http.StatusBadRequest)
		return
	}

	// Get the file from the request
	r.Body = http.MaxBytesReader(w, r.Body, constant.SIZE_5MB) // max 5mb
	file, header, err := r.FormFile(key)
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

	url, err := util.UploadFile(file, remoteFolder+"/"+filename)
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
