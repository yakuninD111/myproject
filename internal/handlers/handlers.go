package handlers

import (
 "myserver/internal/models"
 "html/template"
 "net/http"
 "strconv"
 "time"
)

// Handler хранит зависимости
type Handler struct {
 store     *models.Store
 templates *template.Template
}

// NewHandler создаёт обработчик
func NewHandler(store *models.Store, templates *template.Template) *Handler {
 return &Handler{
  store:     store,
  templates: templates,
 }
}

// IndexHandler — главная страница
func (h *Handler) IndexHandler(w http.ResponseWriter, r *http.Request) {
 if r.URL.Path != "/" {
  http.NotFound(w, r)
  return
 }

 expenses := h.store.GetAll()

 var total float64
 for _, e := range expenses {
  total += e.Amount
 }

 data := map[string]interface{}{
  "Title":    "Список трат",
  "Expenses": expenses,
  "Total":    total,
 }

 w.Header().Set("Content-Type", "text/html; charset=utf-8")
 h.templates.ExecuteTemplate(w, "layout", data)
}

// AddPageHandler — страница формы (GET)
func (h *Handler) AddPageHandler(w http.ResponseWriter, r *http.Request) {
 data := map[string]interface{}{
  "Title": "Добавить трату",
 }

 w.Header().Set("Content-Type", "text/html; charset=utf-8")
 h.templates.ExecuteTemplate(w, "layout", data)
}

// AddHandler — обработка формы (POST)
func (h *Handler) AddHandler(w http.ResponseWriter, r *http.Request) {
 if r.Method != http.MethodPost {
  http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
  return
 }

 err := r.ParseForm()
 if err != nil {
  http.Error(w, "Ошибка формы", http.StatusBadRequest)
  return
 }

 amountStr := r.FormValue("amount")
 description := r.FormValue("description")
 dateStr := r.FormValue("date")

 amount, err := strconv.ParseFloat(amountStr, 64)
 if err != nil {
  http.Error(w, "Неверная сумма", http.StatusBadRequest)
  return
 }

 date, err := time.Parse("2006-01-02", dateStr)
 if err != nil {
  http.Error(w, "Неверная дата", http.StatusBadRequest)
  return
 }

 h.store.Add(amount, description, date)

 http.Redirect(w, r, "/", http.StatusSeeOther)
}