FROM golang:1.22

RUN go version

COPY ./ ./

# install psql
RUN apt-get update
RUN apt-get -y install postgresql-client

# make wait-for-postgres.sh executable
RUN chmod +x wait-for-postgres.sh

# build go app
RUN go mod download
RUN go build -o ims-staff-api ./cmd/main.go

# Команда для запуска приложения
CMD ["./ims-staff-app"]
