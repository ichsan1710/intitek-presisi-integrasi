package productcontroller

import (
	"encoding/json"
	"net/http"
	"server-side/entities"
	"server-side/models/productmodel"
	"strconv"
)

func Index(w http.ResponseWriter, r *http.Request) {
	products := productmodel.GetAll()

	// Changing data to JSON format
	jsonData, err := json.Marshal(products)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Adding Content-Type Header
	w.Header().Set("Content-Type", "application/json")

	// Sending JSON response
	w.Write(jsonData)
}

func Add(w http.ResponseWriter, r *http.Request) {
	// Making sure the requested method is only POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode JSON data
	var product entities.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Adding product to database
	lastID, err := productmodel.Create(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Sending JSON response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "Product added successfully", "id": lastID})

}

func Detail(w http.ResponseWriter, r *http.Request) {
	// Getting ID from query
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Getting data from database
	product, err := productmodel.Detail(id)
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	// Sending JSON response
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(product)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func Update(w http.ResponseWriter, r *http.Request) {
	// Making sure the requested method is only PUT
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Getting ID from query
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Decode JSON data
	var updatedProduct entities.Product
	err = json.NewDecoder(r.Body).Decode(&updatedProduct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Updating new data in database
	result, err := productmodel.Update(id, updatedProduct)
	if err != nil {
		http.Error(w, "Product ID does not exist", http.StatusNotFound)
		return
	}

	// Sending JSON Response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]interface{}{"message": "Product has been updated successfully", "data": result})
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	// Getting ID from query
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Deleting data from database
	if err := productmodel.Delete(id); err != nil {
		http.Error(w, "Product ID does not exist", http.StatusNotFound)
		return
	}

	// Sending JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]interface{}{"message": "Product has been deleted succesfully", "id": id})
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}
}
