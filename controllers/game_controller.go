package controllers

import (
	"../game"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
)


var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func GameController(w http.ResponseWriter,  r* http.Request){
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Client Connected")
	errStr:=""
	typeGameCookie, err:= r.Cookie("typeGame")
	if err!=nil{
		errStr+="ошибка получения типа игры!;"
	}
	tokenCookie, err:= r.Cookie("token")
	if err!=nil{
		errStr+="ошибка токена!;"
	}
	countOfDigitsCookie, err:= r.Cookie("countOfDigits")
	if err!=nil{
		errStr+="ошибка получения количества цифр!;"
	}
	if len(errStr) !=0 {
		writeMessageToPlayer(conn, errStr)
		log.Println(errStr)
		return
	}

	userName := tokens[tokenCookie.Value]

	countOfDigits, err := strconv.Atoi(countOfDigitsCookie.Value)
	if err!=nil{
		errStr+="Ошибка обработки числа цифр;"
	}else if countOfDigits>10 || countOfDigits <1{
		errStr+="Неверное количество цифр в числе от 1 до 10;"
	}
	typeGame, err := strconv.Atoi(typeGameCookie.Value)
	if err!=nil{
		errStr+="Ошибка типа игры;"
	}else if  typeGame<0 || typeGame>1{
		errStr+="Неверный тип игры;"
	}
	if len(errStr) !=0 {
		writeMessageToPlayer(conn, errStr)
		return
	}
	if typeGame==game.New{
		go startNewGame(conn, userName, countOfDigits)
	}else if typeGame== game.Load{
		go loadGame(conn, userName)
	}
}
func GamePageController(w http.ResponseWriter, r* http.Request){
	http.ServeFile(w,r, "templates/game.html")
}