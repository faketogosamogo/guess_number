package main

import (
	"./game"
	"./models"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"time"
)
import "./controllers"
func main(){
	rand.Seed(time.Now().UTC().UnixNano())
	//TODO: доделать html странички и сделать авторизацию

	gameManager, err := game.NewGameManager()
	if err!=nil{
		log.Fatal(err.Error())
	}
	controllers.DB , err= models.NewDB(models.DriverName, models.ConnectionString)
	if err!=nil{
		log.Fatal("Ошибка подключения к БД")
	}
	controllers.GameManager = gameManager


	r := mux.NewRouter()
	r.HandleFunc("/", controllers.HomeController)
	r.HandleFunc("/register", controllers.RegisterController)
	r.HandleFunc("/auth", controllers.AuthController)

	rAuth:= r.PathPrefix("/user").Subrouter()
	rAuth.HandleFunc("/finished_games", controllers.UserFinishedGamesController)
	rAuth.HandleFunc("/gameWS", controllers.GameController)
	rAuth.HandleFunc("/game", controllers.GamePageController)
	rAuth.HandleFunc("/scoreboard", controllers.ScoreboardController)
	rAuth.Use(controllers.AuthMiddleware)

	http.Handle("/",r)

	http.ListenAndServe("localhost:8080", nil)

}
