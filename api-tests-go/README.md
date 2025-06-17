# API-тесты на Go (ReqRes)

Этот проект содержит автоматизированные тесты для API [ReqRes](https://reqres.in/) с использованием Go, Resty и Testify.

## Описание
Тесты проверяют GET-запросы к эндпоинту `/users` API ReqRes, обеспечивая корректность статус-кода и структуры ответа. Проект демонстрирует навыки работы с Go для автоматизации API-тестирования.

## Технологии
- Go 1.22+
- Resty v2
- Testify v1

## Структура проекта
api-tests-go/
├── client/
│   └── api_client.go     # Клиент для API-запросов
├── models/
│   └── user.go           # Модель данных пользователя
├── tests/
│   └── get_users_test.go # Тесты GET-запросов
├── go.mod                # Зависимости
└── README.md             # Документация


## Установка
1. Клонируйте репозиторий:
   ```bash
   git clone https://github.com/mono1eet/AQA-Portfolio.git
   cd AQA-Portfolio/api-tests-go
2. Установите Go 1.22+ (скачать с golang.org/dl).
3. Установите зависимости:
   ```bash
   go mod tidy

Запуск тестов
   ```bash
   go test ./tests -v

Пример теста
Тест отправляет GET-запрос к /users и проверяет:

Отсутствие ошибок.
Статус-код 200.
Наличие ключа "data" в JSON-ответе.


Автор
Степан К. (mono1eet)
