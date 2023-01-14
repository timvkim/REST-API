.PHONY:
.SILENT:
.DEFAULT_GOAL := all

build:
	go build -o ./.bin/app ./cmd/app/main.go

run: build
	./.bin/app

migrate:
	migrate -database postgres://admin:admin123@localhost:5432/course_db?sslmode=disable -path migrations up

drop-tables:
	migrate -database postgres://admin:admin123@localhost:5432/course_db?sslmode=disable -path migrations down