package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"victord/daemon/internal/dto"
	entity "victord/daemon/internal/entity/vector"

	"github.com/gorilla/mux"
)

func (h *Handler) InsertVectorHandler(w http.ResponseWriter, r *http.Request) {
	var req dto.InsertVectorRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Println("Error decoding request:", err)
		http.Error(w, "Invalid insert vector request payload", http.StatusBadRequest)
		return
	}

	urlParams := mux.Vars(r)
	indexNameParam := urlParams["indexName"]

	vecId, err := h.VectorService.InsertVector(&req, indexNameParam)
	if err != nil {
		fmt.Println("Error inserting vector:", err)
		http.Error(w, "Failed to insert vector", http.StatusInternalServerError)
		return
	}

	res := dto.InsertVectorResponse{
		Status:  "success",
		Message: "Vector inserted successfully",
		Results: entity.InsertVectorResult{
			ID:     *vecId,
			Vector: req.Vector,
		},
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

func (h *Handler) DeleteVectorHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	indexName := vars["indexName"]
	vectorID := vars["vectorID"]

	vectId, err := strconv.ParseUint(vectorID, 10, 64)
	if err != nil {
		http.Error(w, "Invalid vector ID", http.StatusBadRequest)
		return
	}

	err = h.VectorService.DeleteVector(vectId, indexName)
	if err != nil {
		http.Error(w, "Failed to delete vector", http.StatusInternalServerError)
		return
	}

	res := dto.DeleteVectorResponse{
		Status:  "success",
		Message: "Vector deleted successfully",
		Results: entity.DeleteVectorResult{
			ID: vectId,
		},
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (h *Handler) SearchVectorHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	indexName := vars["indexName"]

	vectorParam := r.URL.Query().Get("vector")
	if vectorParam == "" {
		http.Error(w, "Missing vector parameter", http.StatusBadRequest)
		return
	}

	vectorChunks := strings.Split(vectorParam, ",")
	var vec []float32

	for _, p := range vectorChunks {
		val, err := strconv.ParseFloat(strings.TrimSpace(p), 32)
		if err != nil {
			http.Error(w, "Invalid vector value: "+p, http.StatusBadRequest)
			return
		}
		v := float32(val)
		vec = append(vec, v)
	}

	fmt.Println("Vector to search:", vec)

	var k int = 5
	if kParam := r.URL.Query().Get("top_k"); kParam != "" {
		if parsedK, err := strconv.Atoi(kParam); err != nil || parsedK <= 0 {
			http.Error(w, "Invalid 'top_k' parameter: must be greater than zero", http.StatusBadRequest)
			return
		} else {
			k = parsedK
		}
	}

	fmt.Println("topK:", k)

	result, err := h.VectorService.SearchVector(vec, indexName, k)
	if err != nil {
		fmt.Println("Error searching vector:", err)
		http.Error(w, "Failed to search vector", http.StatusInternalServerError)
		return
	}

	res := dto.SearchVectorResponse{
		Status:  "success",
		Message: "Vector search completed successfully",
		Results: entity.SearchVectorResult{
			ID:       result.ID,
			Distance: result.Distance,
		},
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
