package database

import (
	"database/sql"
	"errors"
	"time"
)

// User представляет модель пользователя
type User struct {
	ID        uint32    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CreateUser создает нового пользователя в базе данных
func CreateUser(username, email, password string) error {
	stmt, err := DB.Prepare("INSERT INTO users (username, email, password, created_at, updated_at) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	hashedPassword, err := hashPassword(password)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(username, email, hashedPassword, time.Now(), time.Now())
	if err != nil {
		return err
	}

	return nil
}

// GetUserByID получает информацию о пользователе по идентификатору
func GetUserByID(id uint32) (*User, error) {
	user := &User{}
	row := DB.QueryRow("SELECT id, username, email, created_at, updated_at FROM users WHERE id = ?", id)
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, errors.New("Пользователь не найден")
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}

// hashPassword хэширует пароль перед сохранением в базу данных
func hashPassword(password string) (string, error) {
	return "", nil // Здесь должен быть реальный алгоритм хэширования пароля
}
