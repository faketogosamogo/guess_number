package models

import "log"

type UnfinishedGame struct {
	UserName string
	CountOfDigits int
	CountOfTurns  int

	UserTurns string
	BotTurns  string
	ThoughtNumber string
	CountOfSeconds int
}

func(db *DB) AddUnfinishedGame(game UnfinishedGame) error{
	_, err := db.Exec("insert into unfinished_games values(?,?,?,?,?,?,?)", game.UserName, game.CountOfDigits, game.CountOfTurns, game.UserTurns, game.BotTurns, game.ThoughtNumber, game.CountOfSeconds)
	if err!=nil{
		log.Println(err.Error())
	}
	return err
}
func(db *DB)DeleteUnfinishedGame(userName string)error{
	_, err:= db.Exec("delete from unfinished_games where UserName=?", userName)
	return err
}
func (db *DB) GetUnfinishedGame(userName string)(UnfinishedGame, error) {
	game := UnfinishedGame{}
	row := db.QueryRow("select * from unfinished_games where username=?", userName)
	err := row.Scan(&game.UserName, &game.CountOfDigits, &game.CountOfTurns,&game.UserTurns, &game.BotTurns, &game.ThoughtNumber, &game.CountOfSeconds)
	if err != nil {
		log.Println(err.Error())
	}
	return game, err
}
