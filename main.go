package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type CreateProduce struct {
    Name       string `json:"name"`
    UnitPrice  string `json:"unitPrice"`
}

type Produce struct {
	ProduceCode string  `json:"produceCode"`
	Name       string `json:"name"`
	UnitPrice  string  `json:"unitPrice"`
}

var produceCollection = []Produce {
	{"l6m9-5p3n-y5qr-lhel","Apple","$1.23"},
	{"yr7k-b6ku-sruk-mnd6","Orange","$2.45"},
}

func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
    newProduceBytes, _ := json.Marshal(produceCollection)
    w.Write(newProduceBytes)
}

func post(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var produce CreateProduce
    produceCode := createProduceCode()
    json.NewDecoder(r.Body).Decode(&produce)
    if (produce.Name == "") {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte(`{"error": "missing name field"}`))
        return

    }
    if (produce.UnitPrice == "") {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte(`{"error": "missing unit price field"}`))
        return
    }
    newProduce := Produce{produceCode, produce.Name, produce.UnitPrice}

    w.WriteHeader(http.StatusCreated)
    newProduceBytes, _ := json.Marshal(newProduce)
    w.Write(newProduceBytes)
}

func delete(w http.ResponseWriter, r *http.Request) {
    pathParams := mux.Vars(r)
    w.Header().Set("Content-Type", "application/json")
    produceCode := ""
    if val, ok := pathParams["produceCode"]; ok {
        produceCode = val
        if produceCode == "" {
            w.WriteHeader(http.StatusBadRequest)
            w.Write([]byte(`{"message": "need a produce code"}`))
            return
        }
    }
    for i := len(produceCollection) - 1; i >= 0; i-- {
        produce := produceCollection[i]
        if (produce.ProduceCode == produceCode) {
            produceCollection = append(produceCollection[:i],
                produceCollection[i+1:]...)
            w.WriteHeader(http.StatusNoContent)
            w.Write([]byte(`{"message": "produce was successfully deleted"}`))
            return
        }
    }
    w.WriteHeader(http.StatusNotFound)
    w.Write([]byte(`{"message": "produce was not found with the provided produceCode"}`))
    
}

func getOne(w http.ResponseWriter, r *http.Request) {
    pathParams := mux.Vars(r)
    w.Header().Set("Content-Type", "application/json")
    var produce Produce
    produceCode := ""
    if val, ok := pathParams["produceCode"]; ok {
        produceCode = val
        if produceCode == "" {
            w.WriteHeader(http.StatusInternalServerError)
            w.Write([]byte(`{"message": "need a produce code"}`))
            return
        }
    }
    for _, p := range produceCollection {
        if (p.ProduceCode == produceCode) {
            produce = p
        }
    }
    if (produce.Name == "") {
        w.WriteHeader(http.StatusNotFound)
        w.Write([]byte(`{"message": "no produce with the produceCode provided"}`))
        return
    }
    newProduceBytes, _ := json.Marshal(produce)
    w.Write(newProduceBytes)
}

func main() {
    r := mux.NewRouter()

    api := r.PathPrefix("/api/v1").Subrouter()
    api.HandleFunc("/produce", get).Methods(http.MethodGet)
    api.HandleFunc("/produce", post).Methods(http.MethodPost)
    api.HandleFunc("/produce/{produceCode}", delete).Methods(http.MethodDelete)
    api.HandleFunc("/produce/{produceCode}", getOne).Methods(http.MethodGet)

    log.Fatal(http.ListenAndServe(":8080", r))
}
