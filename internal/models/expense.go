package models

import (
 "sync"
 "time"
)

// Expense — структура одной траты
type Expense struct {
 ID          int
 Amount      float64
 Description string
 Date        time.Time
}

// Store — хранилище трат в памяти
type Store struct {
 mu       sync.RWMutex
 expenses []Expense
 nextID   int
}

// NewStore создаёт новое хранилище
func NewStore() *Store {
 return &Store{
  expenses: make([]Expense, 0),
  nextID:   1,
 }
}

// Add добавляет новую трату
func (s *Store) Add(amount float64, description string, date time.Time) {
 s.mu.Lock()
 defer s.mu.Unlock()

 s.expenses = append(s.expenses, Expense{
  ID:          s.nextID,
  Amount:      amount,
  Description: description,
  Date:        date,
 })
 s.nextID++
}

// GetAll возвращает все траты
func (s *Store) GetAll() []Expense {
 s.mu.RLock()
 defer s.mu.RUnlock()

 result := make([]Expense, len(s.expenses))
 copy(result, s.expenses)
 return result
}