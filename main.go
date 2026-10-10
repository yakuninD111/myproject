package main

import (
    "fmt"
    "log"
    "net/http"
)

// homeHandler — обработчик главной страницы
func homeHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/plain; charset=utf-8")
    fmt.Fprintln(w, "Добро пожаловать! Это главная страница моего Go-сервера.")
}

// aboutHandler — обработчик страницы "О проекте"
func aboutHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/plain; charset=utf-8")
    fmt.Fprintln(w, "О проекте: учебный HTTP-сервер на стандартной библиотеке net/http.")
}

// pingHandler — обработчик /ping
func pingHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/plain; charset=utf-8")
    w.WriteHeader(http.StatusOK)
    fmt.Fprintln(w, "pong")
}

func main() {
    // Регистрируем обработчики в стандартном маршрутизаторе (ServeMux)
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/about", aboutHandler)
    http.HandleFunc("/ping", pingHandler)

    addr := ":8080"
    fmt.Printf("Сервер запущен на http://localhost%s\n", addr)

    // Запускаем сервер; log.Fatal выведет ошибку и завершит программу, если порт занят
    log.Fatal(http.ListenAndServe(addr, nil))
}