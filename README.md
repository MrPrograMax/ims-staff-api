# Запуск
Сначала создайте файл .env в корне и добавьте пароль:
```yaml
DB_PASSWORD:qwerty
```
## База данных
```cmd
docker run --name=ims-staff-db -e POSTGRES_PASSWORD='qwerty' -p 5432:5432 --rm postgres
```
## Миграция таблиц
```cmd
make migrate
```
Или
```cmd
migrate -path ./schema -database 'postgres://postgres:qwerty@0.0.0.0:5432/postgres?sslmode=disable' up
```
## Запуск приложения
```cmd
go run cmd/main.go
```
