package routes

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"

	"example.com/zakonm/pkg/database"
)

// SetupRoutes настраивает маршруты для API
func SetupRoutes(router *mux.Router) {
	router.HandleFunc("/api/v1/users", CreateUserHandler).Methods("POST")
	router.HandleFunc("/api/v1/users/{id:[0-9]+}", GetUserByIDHandler).Methods("GET")
}

// CreateUserHandler обрабатывает POST-запросы для создания новых пользователей
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user database.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = database.CreateUser(user.Username, user.Email, user.Password)
	if err != nil {
		http.Error(w, errors.Wrap(err, "Не удалось создать пользователя").Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Пользователь успешно создан"})
}

// GetUserByIDHandler обрабатывает GET-запросы для получения информации о пользователе по идентификатору
func GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, ok := params["id"]
	if !ok {
		http.Error(w, "Идентификатор пользователя не указан", http.StatusBadRequest)
		return
	}

	userID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		http.Error(w, "Неверный формат идентификатора пользователя", http.StatusBadRequest)
		return
	}

	user, err := database.GetUserByID(uint32(userID))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Пользователь не найден", http.StatusNotFound)
			return
		}
		http.Error(w, errors.Wrap(err, "Ошибка при получении информации о пользователе").Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
