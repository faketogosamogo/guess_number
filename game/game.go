package game

import (
	"../models"
	"github.com/google/uuid"
	"strings"
)
const(
	New = iota
	Load
)

type Game struct{
	UserName string
	CountOfDigits int
	CountOfTurns  int

	UserTurns []string
	BotTurns  []string
	ThoughtNumber string
	CountOfSeconds int
	IsFinished bool
}
func(g *Game) GetCountOfDigits() int{
	return g.CountOfDigits
}
func NewGame(userName string, countOfDigits int) *Game{
	game:= 	Game{UserName:userName,
		CountOfDigits:countOfDigits,
		CountOfTurns:0,
		BotTurns:make([]string,0),
		UserTurns:make([]string,0),
		ThoughtNumber:CreateNumber(countOfDigits),
		IsFinished:false,
	}
	return &game
}

func (g *Game) AddTime(countOfSeconds int){
	g.CountOfSeconds+=countOfSeconds
}

func (g *Game) LoadUnfinishedGame(userName string,db * models.DB) error{
	game, err:= db.GetUnfinishedGame(userName)
	if err!=nil{
		return err
	}
	userTurns:=strings.Split(game.UserTurns, ";")
	botTurns:=strings.Split(game.BotTurns, ";")


	g.UserName= game.UserName
	g.CountOfDigits=game.CountOfDigits
	g.CountOfTurns = game.CountOfTurns
	g.ThoughtNumber=game.ThoughtNumber
	g.IsFinished= false
	g.UserTurns=userTurns
	g.BotTurns=botTurns
	g.CountOfSeconds = game.CountOfSeconds
	return nil
}

func(g *Game) Save(db *models.DB) error{
	if g.IsFinished{
		game:= models.FinishedGame{Id:uuid.New().String(),
									UserName:g.UserName,
									CountOfDigits:g.CountOfDigits,
									CountOfTurns:g.CountOfTurns,
									ThoughtNumber:g.ThoughtNumber,
									CountOfSeconds:g.CountOfSeconds}
		err:= db.AddFinishedGame(game)
		return err
	}else{
		botTurnsStr := ""
		userTurnsStr := ""
		for _, el:= range g.UserTurns{
			userTurnsStr+=el+";"
		}
		for _, el:= range g.BotTurns{
			botTurnsStr+=el+";"
		}
		game:= models.UnfinishedGame{
			UserName:g.UserName,
			CountOfDigits:g.CountOfDigits,
			CountOfTurns:g.CountOfTurns,
			UserTurns:userTurnsStr,
			BotTurns:botTurnsStr,
			ThoughtNumber:g.ThoughtNumber,
			CountOfSeconds:g.CountOfSeconds,
		}
		err:= db.AddUnfinishedGame(game)
		return err
	}
}

