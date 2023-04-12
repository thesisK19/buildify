package handler

// func UploadHandler(w http.ResponseWriter, r *http.Request) {
// 	err := r.ParseForm()
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("Failed to parse form: %s", err.Error()), http.StatusBadRequest)
// 		return
// 	}

// 	// Get the file from the request
// 	file, header, err := r.FormFile("file")
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	defer file.Close()

// 	filename := fmt.Sprintf("%s_%s%s", header.Filename, util.GenUUID(5), filepath.Ext(header.Filename))

// 	util.UploadFile(file, filename)

// 	// Send a success response
// 	Send(w, Response{
// 		Code:    "OK",
// 		Message: "OK",
// 		Data:    "hii data",
// 	})
// }
