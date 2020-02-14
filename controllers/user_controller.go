package controllers

import (
	"../models"
	"database/sql"
	"github.com/google/uuid"
	"log"
	"net/http"
	"time"
)
var DB *models.DB


func RegisterController(w http.ResponseWriter, r *http.Request){
	name:=r.Header.Get("name")
	password:= r.Header.Get("password")

	if len(name)<5 || len(password)< 5{
		http.Error(w, "Длины логина и пароля не должны быть меньше 5", 422)
		return
	}
	_, err := DB.GetUser(name)

	if err==nil{
		http.Error(w, "Ошибка регистрации возможно данное имя занято", 422)
	}else{
		err = DB.AddUser(models.User{name, password})
		if err!=nil{
			log.Println(err)
		}
	}

}

func AuthController(w http.ResponseWriter, r *http.Request){
	name:=r.Header.Get("name")
	password:= r.Header.Get("password")

	if len(name)<5 || len(password)< 5{
		http.Error(w, "Длины логина и пароля не должны быть меньше 5", 422)
		return
	}
	user, err := DB.GetUser(name)

	if err!=nil{
		if err== sql.ErrNoRows{
			http.Error(w, "Ошибка авторизации!", 422)
			return
		}
		http.Error(w, "Ошибка получения пользователя", 500)
		return
	}
	if user.Password!=password{
		http.Error(w, "Ошибка авторизации!", 422)
		return
	}
	token := uuid.New().String()
	tokens[token]= name

	cookie := http.Cookie{
		Name:    "token",
		Value: token,
		Expires: time.Now().AddDate(1, 0, 0),
	}
	http.SetCookie(w, &cookie)
}