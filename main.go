package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	ISAN     string    `json:"isan"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func getMovies(res_w http.ResponseWriter, req *http.Request) {
	res_w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res_w).Encode(movies) // response that it will be giving want to encode into json
}
func getMovie(res_w http.ResponseWriter, req *http.Request) {
	res_w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	for _, movie := range movies {
		if movie.ID == params["id"] {
			json.NewEncoder(res_w).Encode(movie)
			return
		}
	}
}
func createMovie(res_w http.ResponseWriter, req *http.Request) {
	res_w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(req.Body).Decode(&movie) // decode body of request into movie variable
	movie.ID = strconv.Itoa(rand.Intn(1000))
	movies = append(movies, movie)
	json.NewEncoder(res_w).Encode(movie)
}
func updateMovie(res_w http.ResponseWriter, req *http.Request) {
	res_w.Header().Set("Content-Type", "application/json")
	var updateValues Movie
	_ = json.NewDecoder(req.Body).Decode(&updateValues)
	params := mux.Vars(req)
	for index, movie := range movies {
		if movie.ID == params["id"] {
			movies[index].Title = updateValues.Title
			movies[index].ISAN = updateValues.ISAN
			movies[index].Director = updateValues.Director
			json.NewEncoder(res_w).Encode(movies[index])
			return
		}
	}
}
func deleteMovie(res_w http.ResponseWriter, req *http.Request) {
	res_w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	for index, movie := range movies {
		if movie.ID == params["id"] { // json id matching with golang movie struct's ID
			json.NewEncoder(res_w).Encode(movie)
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
}

var movies []Movie

func main() {
	router := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", ISAN: "432887", Title: "Contact", Director: &Director{FirstName: "Robert", LastName: "Zemeckis"}})
	movies = append(movies, Movie{ID: "2", ISAN: "454250", Title: "Interstellar", Director: &Director{FirstName: "Christopher", LastName: "Nolan"}})

	router.HandleFunc("/movies", getMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	router.HandleFunc("/movies", createMovie).Methods("POST")
	router.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	router.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at post 8000...\n")
	log.Fatal(http.ListenAndServe(":8000", router))
}
