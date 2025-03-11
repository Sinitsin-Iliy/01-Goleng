# 01-Goleng
# cmd
# main.go                 - точка входа в приложение
# pkg
# database
# db.go                   - работа с базой данных
# models.go               - модели данных
# routes
# routes.go               - обработчики маршрутов
# Dockerfile              - файл конфигурации Docker
# docker-compose.yml      - файл для управления контейнерами Docker


# Установка зависимостей
go mod init example.com/zakonm
go get github.com/gorilla/mux
go get github.com/go-sql-driver/mysql


# Сборка и запуск через Docker
# Соберем и запустим наше приложение с помощью Docker Compose:

docker-compose up --build
# Теперь ваше приложение доступно по адресу http://localhost:8000/.


# Cоздал простое веб-приложение на языке Go с использованием базы данных MySQL через Docker. Приложение включает обработку маршрутов, взаимодействие с базой данных и простейшую аутентификацию пользователей.