package controllers

import "net/http"

func HomeController(w http.ResponseWriter, r* http.Request){
	http.ServeFile(w,r, "templates/home.html")
}
