build:
	docker-compose build ims-staff-api

run:
	docker-compose up ims-staff-api

migrate:
	migrate -path ./schema -database 'postgres://postgres:qwerty@0.0.0.0:5432/postgres?sslmode=disable' up


