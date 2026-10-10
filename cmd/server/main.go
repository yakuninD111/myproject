package main

import (
 "myserver/internal/handlers"
 "myserver/internal/models"
 "fmt"
 "html/template"
 "log"
 "net/http"
)

func main() {
 // Создаём хранилище
 store := models.NewStore()

 // Парсим шаблоны
 tmpl, err := template.ParseGlob("web/templates/*.html")
 if err != nil {
  log.Fatalf("Ошибка шаблонов: %v", err)
 }

 // Создаём обработчик
 handler := handlers.NewHandler(store, tmpl)

 // Регистрируем маршруты
 http.HandleFunc("/", handler.IndexHandler)

 http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
  if r.Method == http.MethodGet {
   handler.AddPageHandler(w, r)
  } else if r.Method == http.MethodPost {
   handler.AddHandler(w, r)
  } else {
   http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
  }
 })

 // Статические файлы
 fileServer := http.FileServer(http.Dir("web/static"))
 http.Handle("/static/", http.StripPrefix("/static/", fileServer))

 // Запуск сервера
 addr := ":8080"
 fmt.Printf("Сервер запущен на http://localhost%s\n", addr)
 log.Fatal(http.ListenAndServe(addr, nil))
}