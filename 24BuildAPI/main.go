package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// model for Music
type Music struct {
	MusicId      int     `json:"id"`
	MusicName    string  `json:"name"`
	LenthOfMusic float64 `json:"length"`
	Artist       *Artist `json:"artist"`
}

type Artist struct {
	FullName string `json:"fullname"`
	Link     string `json:"link"`
}

// init Music
var music []Music

// display on browser
func serverPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>welcome <br/> This is Music REST API.</h1>"))
}

// get all musics
func getAllMusics(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all Musics")

	w.Header().Set("Contant-type", "application/json")
	json.NewEncoder(w).Encode(music)
}

// get one musics
func getOneMusic(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one Musics")

	w.Header().Set("Contant-type", "application/json")

	param := mux.Vars(r) // Get id of music

	// Loop through all music and mache give id
	for _, item := range music {

		// convert id string to int
		conId, _ := strconv.Atoi(param["id"])

		// check the id is same as give id
		if item.MusicId == conId {
			json.NewEncoder(w).Encode(item)
			return
		}

	}

	json.NewEncoder(w).Encode(&Music{})
}

// add music
func addMusic(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Add music")

	w.Header().Set("Contant-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	var m Music
	_ = json.NewDecoder(r.Body).Decode(&m)

	rand.Seed(time.Now().UnixNano())
	m.MusicId = rand.Intn(100)
	music = append(music, m)
	json.NewEncoder(w).Encode(music)

}

// edit music
func editMusic(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Edit music")

	w.Header().Set("Contant-Type", "application/json")

	param := mux.Vars(r)

	for index, m := range music {
		// convert id string to int
		conId, _ := strconv.Atoi(param["id"])

		// check the id is same as give id
		if m.MusicId == conId {
			music = append(music[:index], music[index+1:]...)

			var m Music
			_ = json.NewDecoder(r.Body).Decode(&m)
			m.MusicId = conId
			music = append(music, m)
			json.NewEncoder(w).Encode(m)
			return
		}
	}
}

// delete music
func deleteMusic(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete music")

	w.Header().Set("Contant-Type", "application/json")

	param := mux.Vars(r)

	for index, m := range music {
		conId, _ := strconv.Atoi(param["id"])

		// check the id is same as give id
		if m.MusicId == conId {
			music = append(music[:index], music[index+1:]...)
			break
		}
	}
}
func main() {
	fmt.Println("Server is running at port 5000...")

	r := mux.NewRouter()

	// mock data
	music = append(music, Music{MusicId: 1, MusicName: "All Too Well", LenthOfMusic: 5.29, Artist: &Artist{FullName: "Taylor swift", Link: "alltoowell.com"}})
	music = append(music, Music{MusicId: 2, MusicName: "Bad Decision", LenthOfMusic: 2.52, Artist: &Artist{FullName: "bb, jk, snoop dogg", Link: "baddecision.com"}})

	// Routing
	// Create sample handler to returns 404
	r.HandleFunc("/", serverPage).Methods("GET")
	r.HandleFunc("/api/musics", getAllMusics).Methods("GET")
	r.HandleFunc("/api/musics/{id}", getOneMusic).Methods("GET")
	r.HandleFunc("/api/musics", addMusic).Methods("POST")
	r.HandleFunc("/api/musics/{id}", editMusic).Methods("PUT")
	r.HandleFunc("/api/musics/{id}", deleteMusic).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":5000", r))
}
