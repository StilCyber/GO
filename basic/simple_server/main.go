package main

import (
	crypto "crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/go-chi/chi/v5"
)

const (
	baseUrl         = "localhost:8080"
	home            = "/"
	getWomanPostfix = "/woman/{id}"
	createWoman     = "/woman/create"
	updateWoman     = "/woman/update"
)

type WomanInfo struct {
	Name                 string `json:"name"`
	Venereal_disease     int64  `json:"venereal_disease"`
	Personality_disorder int64  `json:"Personality_disorder"`
}

type Woman struct {
	Id         int64      `json:"id"`
	Info       *WomanInfo `json:"info"`
	Created_at time.Time  `json:"created_at"`
	Updated_at time.Time  `json:"updated_at"`
}

type SyncWomen struct {
	elements map[int64]*Woman
	mutex    sync.Mutex
}

var women = &SyncWomen{elements: make(map[int64]*Woman)}

func main() {
	r := chi.NewRouter()

	r.Get(getWomanPostfix, getWomanHandler)
	r.Get(home, getMainPageHandler)

	r.Post(createWoman, postWomanHandler)
	r.Put(updateWoman, putWomanHandler)

	r.Delete(getWomanPostfix, deleteWomanHandler)

	fmt.Println("Server started on port", baseUrl)
	err := http.ListenAndServe(baseUrl, r)
	if err != nil {
		log.Fatal(err)
	}
}

func getWomanHandler(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "id")

	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		http.Error(w, "Incorrect ID", http.StatusBadRequest)
		return
	}

	women.mutex.Lock()
	defer women.mutex.Unlock()

	woman, ok := women.elements[id]
	if !ok {
		errorText := "Woman with ID " + userID + " not found"
		http.Error(w, errorText, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(woman); err != nil {
		http.Error(w, "Failed to encode woman data", http.StatusInternalServerError)
	}
}

func postWomanHandler(w http.ResponseWriter, r *http.Request) {
	womanInfo := &WomanInfo{}
	if err := json.NewDecoder(r.Body).Decode(womanInfo); err != nil {
		http.Error(w, "Failed to decode woman data", http.StatusBadRequest)
	}

	id := newCryptoRand()
	timeNow := time.Now()

	woman := &Woman{
		Id:         id,
		Info:       womanInfo,
		Created_at: timeNow,
		Updated_at: timeNow,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(woman); err != nil {
		http.Error(w, "Failed to encode woman data", http.StatusInternalServerError)
		return
	}

	women.mutex.Lock()
	defer women.mutex.Unlock()

	women.elements[woman.Id] = woman
}

func getMainPageHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Main page"))
}

func putWomanHandler(w http.ResponseWriter, r *http.Request) {
	womanUpdate := &Woman{}
	if err := json.NewDecoder(r.Body).Decode(womanUpdate); err != nil {
		http.Error(w, "Failed to decode woman data", http.StatusBadRequest)
	}

	woman, ok := women.elements[womanUpdate.Id]
	if !ok {
		http.Error(w, "Woman not found", http.StatusBadRequest)
		return
	}

	timeNow := time.Now()

	women.mutex.Lock()
	defer women.mutex.Unlock()

	woman = &Woman{
		Id:         womanUpdate.Id,
		Info:       womanUpdate.Info,
		Created_at: woman.Created_at,
		Updated_at: timeNow,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(woman); err != nil {
		http.Error(w, "Failed to encode woman data", http.StatusInternalServerError)
		return
	}

	women.elements[woman.Id] = woman
}

func deleteWomanHandler(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "id")

	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		http.Error(w, "Incorrect ID", http.StatusBadRequest)
		return
	}

	women.mutex.Lock()
	defer women.mutex.Unlock()

	_, ok := women.elements[id]
	if !ok {
		errorText := "Woman with ID " + userID + " not found"
		http.Error(w, errorText, http.StatusInternalServerError)
		return
	}

	message := "Woman with id" + userID + " deleted"

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(message); err != nil {
		http.Error(w, "Failed to encode woman data", http.StatusInternalServerError)
	}

	delete(women.elements, id)
}

func newCryptoRand() int64 {
	safeNum, err := crypto.Int(crypto.Reader, big.NewInt(100234))
	if err != nil {
		panic(err)
	}
	return safeNum.Int64()
}
