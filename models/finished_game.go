package models

import (
	"fmt"
	"log"
)

type FinishedGame struct {
	Id            string
	UserName string
	CountOfDigits int
	CountOfTurns  int
	ThoughtNumber string
	CountOfSeconds int
}


func(db *DB) AddFinishedGame(game FinishedGame) error{
	_, err := db.Exec("insert into finished_games values(?,?,?,?,?, ?)", game.Id, game.UserName,game.CountOfDigits, game.CountOfTurns, game.ThoughtNumber, game.CountOfSeconds)
	if err!=nil{
		log.Println(err.Error())
	}
	return err
}
func (db *DB) GetUserFinishedGames(userName string)([]FinishedGame, error){
	games := make([]FinishedGame,0)
	rows, err := db.Query("select * from finished_games where username=?", userName)
	if err != nil {
		log.Println(err)
		return games, err
	}
	defer rows.Close()
	for rows.Next(){
		game := FinishedGame{}
		err := rows.Scan(&game.Id, &game.UserName, &game.CountOfDigits, &game.CountOfTurns, &game.ThoughtNumber, &game.CountOfSeconds)
		if err != nil{
			fmt.Println(err)
			continue
		}
		games = append(games, game)
	}
	return games, nil
}
func (db *DB) GetFinishedGame(id string)(FinishedGame, error) {
	game := FinishedGame{}
	row := db.QueryRow("select * from finished_games where id=?", id)
	err := row.Scan(&game.Id,&game.UserName, &game.CountOfDigits, &game.CountOfTurns, &game.ThoughtNumber, &game.CountOfSeconds)
	if err != nil {
		log.Println(err.Error())
	}
	return game, err
}
func (db *DB) GetScoreboard(countOfDigits int)([]FinishedGame, error){
	games := make([]FinishedGame,0)
	rows, err := db.Query("select * from finished_games where countofdigits=? order by countofseconds asc limit 10", countOfDigits)
	if err != nil {
		log.Println(err)
		return games, err
	}
	defer rows.Close()
	for rows.Next(){
		game := FinishedGame{}
		err := rows.Scan(&game.Id, &game.UserName, &game.CountOfDigits, &game.CountOfTurns, &game.ThoughtNumber, &game.CountOfSeconds)
		if err != nil{
			fmt.Println(err)
			continue
		}
		games = append(games, game)
	}
	return games, nil
}