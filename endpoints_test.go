package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
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

	expected := `[{"produceCode":"A7Fc-VgWD-7m5S-zb9Z","name":"Apple","unitPrice":"$1.23"},{"produceCode":"IgGJ-hHGW-QrWe-sIGY","name":"Orange","unitPrice":"$2.45"}]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

// func TestGetProduceById(t *testing.T) {

// 	req, err := http.NewRequest("GET", "/produce", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	q := req.URL.Query()
// 	q.Add("id", "1")
// 	req.URL.RawQuery = q.Encode()
// 	rr := httptest.NewRecorder()
// 	handler := http.HandlerFunc(getOne)
// 	handler.ServeHTTP(rr, req)
// 	if status := rr.Code; status != http.StatusOK {
// 		t.Errorf("handler returned wrong status code: got %v want %v",
// 			status, http.StatusOK)
// 	}

// 	// Check the response body is what we expect.
// 	expected := `{"id":1,"first_name":"Krish","last_name":"Bhanushali","email_address":"krishsb2405@gmail.com","phone_number":"0987654321"}`
// 	if rr.Body.String() != expected {
// 		t.Errorf("handler returned unexpected body: got %v want %v",
// 			rr.Body.String(), expected)
// 	}
// }