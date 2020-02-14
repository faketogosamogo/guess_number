package controllers

import (
	"net/http"

)

var tokens = make(map[string]string)

func AuthMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//token := r.Header.Get("token")
		token, err := r.Cookie("token")
		if err!=nil{
			http.Error(w, "Не авторизованы!", 422)
			return
		}
		_, ok:= tokens[token.Value]
		if !ok{
			http.Error(w, "Неверный токен!", 422)
			return
		}
		next.ServeHTTP(w, r)
	})
}
