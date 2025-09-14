name: Go CI

on:
push:
branches: [ main ]
pull_request:
branches: [ main ]

jobs:
- name: Install golangci-lint
uses: golangci/golangci-lint-action@v4
with:
version: v1.59.1
lint:
runs-on: ubuntu-latest
steps:
- name: Checkout code
uses: actions/checkout@v3

- name: Set up Go
uses: actions/setup-go@v5
with:
go-version: '1.22'

- name: Установка зависимостей
run: go mod download

- name: Линт
run: make lint

# -------------------------------
# Юнит и интеграционные тесты
# -------------------------------
test:
runs-on: ubuntu-latest
needs: lint   # тесты запускаются только после успешного линта

services:
postgres:
image: postgres:15
ports:
- 5432:5432
env:
POSTGRES_USER: postgres
POSTGRES_PASSWORD: postgres
POSTGRES_DB: testdb
options: >-
--health-cmd "pg_isready -U postgres"
--health-interval 10s
--health-timeout 5s
--health-retries 5

env:
TEST_DATABASE_DSN: postgres://postgres:postgres@localhost:5432/testdb?sslmode=disable

steps:
- name: Checkout code
uses: actions/checkout@v3

- name: Set up Go
uses: actions/setup-go@v5
with:
go-version: '1.22'

- name: Установка зависимостей
run: go mod download

- name: Запуск тестов
run: make test
