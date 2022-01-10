package main

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/gorilla/mux"
)

type Produce struct {
	ProduceCode string  `json:"produceCode"`
	Name       string `json:"name"`
	UnitPrice  float64  `json:"unitPrice"`
}

var produceCollection = []Produce {
	{"L6M9-5P3N-Y5QR-LHEL","Apple",1.23},
	{"YR7K-B6KU-SRUK-MND6","Orange",2.45},
}

func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
    newProduceBytes, _ := json.Marshal(produceCollection)
    w.Write(newProduceBytes)
}

func post(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var produce []Produce
    json.NewDecoder(r.Body).Decode(&produce)
     if len(produce) == 0 {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte(`{"error": "body should be an array of produce"}`))
        return
     }
     for _, p := range produce {
        if (p.Name == "") {
            w.WriteHeader(http.StatusBadRequest)
            w.Write([]byte(`{"error": "missing name field"}`))
            return
        }
        if (p.UnitPrice == 0) {
            w.WriteHeader(http.StatusBadRequest)
            w.Write([]byte(`{"error": "missing unit price field"}`))
            return
        }
        re, _ := regexp.Compile("^[a-zA-Z0-9][a-zA-Z0-9-]*$")
        if (p.ProduceCode == "" || !re.MatchString(p.ProduceCode) || len(p.ProduceCode) != 19) {
            w.WriteHeader(http.StatusBadRequest)
            w.Write([]byte(`{"error": "error with the produceCode field"}`))
            return
        }
        go createProduce(p)
     }
     w.WriteHeader(http.StatusCreated)
        newProduceBytes, _ := json.Marshal(produce)
        w.Write(newProduceBytes)
    
}

func createProduce(produce Produce) {
        newProduce := Produce{produce.ProduceCode, produce.Name, produce.UnitPrice}
        produceCollection = append(produceCollection, newProduce)
        
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
            w.WriteHeader(http.StatusBadRequest)
            w.Write([]byte(`{"message": "need a produce code"}`))
            return
        }
    }
    for _, p := range produceCollection {
        if (strings.EqualFold(p.ProduceCode, produceCode)) {
            produce = p
        }
    }
    if (produce.Name == "") {
        w.WriteHeader(http.StatusNotFound)
        w.Write([]byte(`{"message": "no produce with the produceCode provided"}`))
        return
    }
    newProduceBytes, _ := json.Marshal(produce)
    w.WriteHeader(http.StatusOK)
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
