package handler

import (
	"encoding/json"
	"fmt"
	"learn_service/service"
	"net/http"
	"reflect"
)

type TestingHandler struct {
	TestingService service.TestingService
}

func (h *TestingHandler) HealthCheck(w http.ResponseWriter, req *http.Request) {
	queryParams := req.URL.Query()
	name := queryParams.Get("name")

	// Check the type of the variable
	fmt.Printf("The type of 'name' is: %s\n", reflect.TypeOf(name))

	response, err := h.TestingService.HealthCheck()

	if err != nil {
		WriteResponse(w, http.StatusInternalServerError, err.Error())
	}

	if name == "mithun" {
		// Call WriteResponse to set status code to 400
		WriteResponse(w, http.StatusBadRequest, response)
	} else {
		WriteResponse(w, http.StatusOK, response)
	}
}

// responseHandler returns an HTTP handler function that responds with the specified status code and body
func WriteResponse(w http.ResponseWriter, code int, body interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "")
	w.Header().Set("Access-Control-Allow-Methods", "")
	w.Header().Set("Access-Control-Allow-Headers", "")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(body)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
