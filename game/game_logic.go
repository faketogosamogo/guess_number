package game

import (
	"errors"
	"strings"
)

func(g *Game) MakeTurn(number string)(string, bool, error){
	if len(number)!=g.CountOfDigits{
		return "", false, errors.New("Не совпадают количество цифр")
	}
	res:=make([]rune,len(number))

	for i, _ := range g.ThoughtNumber{
		if g.ThoughtNumber[i]==number[i]{
			res = append(res, 'В')
		}else if strings.Contains(g.ThoughtNumber, string(number[i])){
			res = append(res, 'К')
		}else{
			res = append(res, rune(number[i]))
		}
	}
	isFinished := true
	for i, el := range number{
		if el != rune(g.ThoughtNumber[i]){
			isFinished = false
		}
	}
	g.UserTurns = append(g.UserTurns, number)
	g.BotTurns = append(g.BotTurns, string(res))
	g.CountOfTurns++
	g.IsFinished = isFinished
	return string(res), g.IsFinished, nil
}
