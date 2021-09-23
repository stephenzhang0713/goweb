package controller

import (
	"fmt"
	"goweb/chapter3/model"
	"net/http"
	"strconv"
	"text/template"
)

type UserController struct {
}

func (c UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	uid, _ := strconv.Atoi(query["uid"][0])

	user := model.GetUser(uid)
	fmt.Println(user)

	t, _ := template.ParseFiles("chapter3/view/t3.html")
	userInfo := []string{user.Name, user.Phone}
	t.Execute(w, userInfo)
}
