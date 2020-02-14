package game

import (
	"../models"
	"errors"
	"math/rand"
)


var ErrGameNotFound  = errors.New("Игра для этого игрока не найдена!")

func CreateNumber(n int)string{
	alph := []rune("1234567890")
	number:= make([]rune,n)

	for i:=range number{
		number[i]+=alph[rand.Intn(len(alph))]
	}
	return string(number)
}

type GameManager struct{
	games map[string]*Game
	db *models.DB
}

func NewGameManager()(*GameManager, error){
	db, err:= models.NewDB(models.DriverName, models.ConnectionString)
	if err!=nil{
		return nil, err
	}
	return &GameManager{games:make(map[string]*Game), db:db}, nil
}

func (g *GameManager) CheckGameByUserName(userName string) bool{
	_, ok := g.games[userName]
	return ok
}

func (g* GameManager) StartNewGame(userName string, countOfDigits int) error{
	if g.CheckGameByUserName(userName){
		return errors.New("Игра для этого игрока уже начата!")
	}else{
		game:= NewGame(userName, countOfDigits)
		g.games[userName] = game
		return nil
	}
}

func (g* GameManager) LoadUnfinishedGame(userName string) error{
	game := &Game{}
	err:= game.LoadUnfinishedGame(userName, g.db)
	if err!=nil{
		return  err
	}
	g.games[userName]= game
	return nil
}

func (g *GameManager) MakeTurn(userName, number string) (string, bool, error){
	game, ok := g.games[userName]
	if ok{
		res, isFinish, err := game.MakeTurn(number)
		return res, isFinish , err
	}else{
		return "", false, ErrGameNotFound
	}
}

func (g* GameManager) SaveGame(userName string) error{
	game, ok := g.games[userName]
	if ok{
		game.Save(g.db)
		return  nil
	}else{
		return ErrGameNotFound
	}
}
func (g* GameManager) AddSeconds(userName string, count int) (error){
	game, ok := g.games[userName]
	if ok{
		game.AddTime(count)
		return nil
	}else{
		return ErrGameNotFound
	}
}
func (g* GameManager) GetGame(userName string) (*Game, error){
	game, ok := g.games[userName]
	if ok{
		return game, nil
	}else{
		return game, ErrGameNotFound
	}
}

func (g* GameManager) DeleteGame(userName string){
	delete(g.games, userName)
}

func(g *GameManager) DeleteUnfinishedGame(userName string)error{
	return g.db.DeleteUnfinishedGame(userName)
}

func (g *GameManager) GetCountOfDigits(userName string)(int, error){
	game, ok := g.games[userName]
	if ok{
		return game.GetCountOfDigits(), nil
	}else{
		return 0, ErrGameNotFound
	}
}