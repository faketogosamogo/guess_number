package controllers

import (
	"html/template"
	"net/http"
	"strconv"
)

func UserFinishedGamesController(w http.ResponseWriter, r *http.Request) {
	token, _ := r.Cookie("token")
	games, err:= DB.GetUserFinishedGames(tokens[token.Value])
	if err!=nil{
		http.Error(w, "Ошибка загрузки игр", 500)
		return
	}
	tmpl, _ := template.ParseFiles("templates/user_games.html")
	tmpl.Execute(w, games)
}
func ScoreboardController(w http.ResponseWriter, r* http.Request){
	countOfDigitsCookie, err := r.Cookie("countOfDigits")
	if err!=nil{
		http.Error(w, "Ошибка получения количества цифр!", 422)
		return
	}

	countOfDigits , err := strconv.Atoi(countOfDigitsCookie.Value)
	if err!=nil{
		http.Error(w, "Ошибка преобразования количества цифр!", 422)
		return
	}
	games, err:= DB.GetScoreboard(countOfDigits)
	if err!=nil{
		http.Error(w, "Ошибка получения игр!", 500)
		return
	}
	tmpl, _ := template.ParseFiles("templates/user_games.html")
	tmpl.Execute(w, games)
}