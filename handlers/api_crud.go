package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/aZ4ziL/go_crud/models"
)

type ErrorHandler struct {
	Type    int
	Message string
}

// DataAPIIndex is a function to handler for all Data of database
// and encode JSON type with Content-Type: application/json
func DataAPIIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		if dataID, ok := r.URL.Query()["id"]; ok {
			dataIDInt, _ := strconv.Atoi(dataID[0])
			data, err := models.GetDataByID(uint(dataIDInt))
			if err != nil {
				errorHandler := ErrorHandler{
					Type:    http.StatusNotFound,
					Message: "Error not found",
				}
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(errorHandler)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data)
			return
		}

		datas := models.GetAllData()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(datas)
	}
}

// DataAPIPost is handler with method POST for create new Data
func DataAPIPost(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		fullName := r.PostFormValue("full_name")
		email := r.PostFormValue("email")
		address := r.PostFormValue("address")

		if fullName == "" || email == "" || address == "" {
			http.Error(w, "Please input...", http.StatusBadRequest)
			return
		}

		data := models.Data{
			FullName: fullName,
			Email:    email,
			Address:  address,
		}

		err := models.NewData(&data)
		if err != nil {
			errorHandler := ErrorHandler{
				Type:    http.StatusBadRequest,
				Message: err.Error(),
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(errorHandler)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(data)
		return
	}
}

// DataAPIPut
func DataAPIPut(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		id := r.FormValue("id")
		fullName := r.FormValue("full_name")
		email := r.FormValue("email")
		address := r.FormValue("address")

		idInt, _ := strconv.Atoi(id)

		// Get Data with ID
		data, err := models.GetDataByID(uint(idInt))
		if err != nil {
			errorHandler := ErrorHandler{
				Type:    http.StatusNotFound,
				Message: err.Error(),
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(errorHandler)
			return
		}

		// Update Data
		data.FullName = fullName
		data.Email = email
		data.Address = address
		models.GetDB().Save(&data)

		errorHandler := ErrorHandler{
			Type:    http.StatusCreated,
			Message: "Successfully to update data with ID:" + id,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(errorHandler)
		return
	}
}

// DataAPIDelete
func DataAPIDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		id := r.URL.Query().Get("id")
		idInt, _ := strconv.Atoi(id)

		data, err := models.GetDataByID(uint(idInt))
		if err != nil {
			errorHandler := ErrorHandler{
				Type:    http.StatusNotFound,
				Message: err.Error(),
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(errorHandler)
			return
		}

		err = models.GetDB().Delete(&data).Error
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		errorHandler := ErrorHandler{
			Type:    http.StatusOK,
			Message: "Successfully to delete data by ID: " + id,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(errorHandler)
		return
	}
}
