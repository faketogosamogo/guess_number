package controllers

import (
	"github.com/gorilla/websocket"
	"log"
	"../game"
	"time"
)
var GameManager *game.GameManager

func writeMessageToPlayer(conn *websocket.Conn, message string)error{
	err := conn.WriteMessage(websocket.TextMessage, []byte(message))
	if err!=nil {
		log.Println("ошибка отправления сообщения игроку message: ", message)
	}
	return err
}

func startNewGame(conn *websocket.Conn, userName string, countOfDigits int) {
	err := GameManager.StartNewGame(userName, countOfDigits)
	if err != nil {
		log.Println("ошибка создания игры username: ", userName)
		err = writeMessageToPlayer(conn, "Ошибка создания игры!")
		if err != nil {
			return
		}
	}
	reader(conn, userName, countOfDigits)
}
func loadGame(conn *websocket.Conn, userName string){
	err:= GameManager.LoadUnfinishedGame(userName)
	if err!=nil {
		log.Println("ошибка загрузки игры username: ", userName)
		err = writeMessageToPlayer(conn, "Не найденно незаконченных игр!")
		if err != nil {
			return
		}
	}
	countOfDigits, err := GameManager.GetCountOfDigits(userName)
	if err!=nil{
		return
	}
	reader(conn, userName, countOfDigits)
}

func reader(conn *websocket.Conn, userName string, countOfDigits int){
	start := time.Now();
	for {
		_, number, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
		if len(string(number)) != countOfDigits{
			err:= writeMessageToPlayer(conn, "Неверная длина числа!")
			if err!=nil{
				break
			}
			continue
		}
		result, isFinish, err := GameManager.MakeTurn(userName, string(number))
		if err!=nil{
			err:= writeMessageToPlayer(conn, "Ошибка совершения хода!!")
			if err!=nil{
				break
			}
		}
		if isFinish{
			writeMessageToPlayer(conn, "Вы выиграли!")

			break
		}else{
			err:= writeMessageToPlayer(conn, result)
			if err!=nil{
				break
			}
		}
	}
	finish := time.Now();
	err:= GameManager.AddSeconds(userName, int(finish.Sub(start).Seconds()))
	if err!=nil{
		log.Println(err)

	}
	err= GameManager.DeleteUnfinishedGame(userName)
	if err!=nil{
		log.Println(err)
	}
	err	= GameManager.SaveGame(userName)
		if err!=nil{
			log.Println(err)
		}
	GameManager.DeleteGame(userName)
	log.Println("отключился!")
}
