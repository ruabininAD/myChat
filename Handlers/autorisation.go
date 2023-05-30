package Handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func Authorization(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Получение данных из формы
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Обработка данных, например, проверка правильности авторизации

		// Вывод данных, полученных из формы
		fmt.Println("Username:", username)
		fmt.Println("Password:", password)

		// Перенаправление пользователя на другую страницу или выполнение другой логики
		http.Redirect(w, r, "/chat", http.StatusFound)
	} else {
		// Обработка GET запроса на страницу авторизации
		tmpl, err := template.ParseFiles("template/autorisation.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
	}

}

func Chat(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

	} else {
		// Обработка GET запроса на страницу авторизации
		tmpl, err := template.ParseFiles("template/chat.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
	}

}
