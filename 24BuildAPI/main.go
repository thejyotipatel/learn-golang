package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

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

// middleware
func (m *Music) IsEmpty() bool {
	return m.MusicName == "" && m.LenthOfMusic == 0.0
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

	param := mux.Vars(r)

	for _, item := range music {
		conId, _ := strconv.Atoi(param["id"])
		if item.MusicId == conId {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&Music{})
}

func main() {
	fmt.Println("Server is running at port 5000...")

	r := mux.NewRouter()

	// mock data
	music = append(music, Music{MusicId: 1, MusicName: "All Too Well", LenthOfMusic: 5.29, Artist: &Artist{FullName: "Taylor swift", Link: "alltoowell.com"}})
	music = append(music, Music{MusicId: 2, MusicName: "Bad Decision", LenthOfMusic: 2.52, Artist: &Artist{FullName: "bb, jk, snoop dogg", Link: "baddecision.com"}})

	// Routing
	r.HandleFunc("/", serverPage).Methods("GET")
	r.HandleFunc("/api/musics", getAllMusics).Methods("GET")
	r.HandleFunc("/api/musics/{id}", getOneMusic).Methods("GET")
	// r.HandleFunc("/api/musics", addMusic).Methods("POST")
	// r.HandleFunc("/api/musics/{id}", editMusic).Methods("PUT")
	// r.HandleFunc("/api/musics/{id}", deleteMusic).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":5000", r))
}
