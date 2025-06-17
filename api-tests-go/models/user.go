package models

// User представляет структуру пользователя из ответа API
type User struct {
    ID        int    `json:"id"`
    Email     string `json:"email"`
    FirstName string `json:"first_name"`
    LastName  string `json:"last_name"`
}