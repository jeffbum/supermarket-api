package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestGetProduce(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(get)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `[{"produceCode":"L6M9-5P3N-Y5QR-LHEL","name":"Apple","unitPrice":1.23},{"produceCode":"YR7K-B6KU-SRUK-MND6","name":"Orange","unitPrice":2.45}]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestGetProduceById(t *testing.T) {

	req, err := http.NewRequest("GET", "api/v1/produce", nil)
	if err != nil {
		t.Fatal(err)
	}
	vars := map[string]string{"produceCode": "L6M9-5P3N-Y5QR-LHEL"}
	req = mux.SetURLVars(req, vars)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getOne)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"produceCode":"L6M9-5P3N-Y5QR-LHEL","name":"Apple","unitPrice":1.23}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestGetProduceByIdCasingDifferent(t *testing.T) {

	req, err := http.NewRequest("GET", "api/v1/produce", nil)
	if err != nil {
		t.Fatal(err)
	}
	vars := map[string]string{"produceCode": "L6M9-5P3N-Y5QR-lhel"}
	req = mux.SetURLVars(req, vars)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getOne)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"produceCode":"L6M9-5P3N-Y5QR-LHEL","name":"Apple","unitPrice":1.23}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestGetProduceByIdNotFound(t *testing.T) {
	req, err := http.NewRequest("GET", "/produce", nil)
	if err != nil {
		t.Fatal(err)
	}
	vars := map[string]string{"produceCode": "L6M9-5P3N-Y5QR-LHE"}

	req = mux.SetURLVars(req, vars)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getOne)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status == http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}


func TestDeleteEntry(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/produce", nil)
	if err != nil {
		t.Fatal(err)
	}
	vars := map[string]string{"produceCode": "L6M9-5P3N-Y5QR-LHEL"}

	req = mux.SetURLVars(req, vars)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(delete)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNoContent)
	}
}

func TestCreateEntry(t *testing.T) {

	var jsonStr = []byte(`[{"name":"Oatmeal","unitPrice":2.59, "produceCode":"WQ9B-5P3N-Y5QR-LHEL"}]`)

	req, err := http.NewRequest("POST", "/produce", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(post)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}
	var produce []Produce
	json.NewDecoder(rr.Body).Decode(&produce)

	if produce[0].Name != "Oatmeal" || produce[0].UnitPrice != 2.59 {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), produce)
	}
}