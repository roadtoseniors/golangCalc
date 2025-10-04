
#  Cервис подсчета арифметических выражений

Этот проект представляет собой сервис для вычисления арифметических выражений, доступный через HTTP-запросы. Вы можете отправить выражение, такое как `2+2*2`, и получить результат в формате JSON

---

## Функционал

✅ Поддержка сложных арифметических выражений: `+`, `-`, `*`, `/`, `()`.  
✅ Ответы в формате JSON.  
✅ Обработка ошибок и корректные HTTP-коды.  
  
---

## Установка и запуск

### 1️⃣ Клонируйте репозиторий:

```bash
git clone https://github.com/roadtoseniors/golangCalc.git
```

### 2️⃣ Запустите сервер:

```bash
go run ./cmd/main.go
```

**По умолчанию сервер будет запущен на**: `http://localhost:8080`.

---

## Использование

### Эндпоинт: `/api/v1/calculate`

### Успешный пример:

#### Запрос:

```bash
curl --location --request POST 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+2*2"
}'
```

#### Ответ:

```json
{
  "result": 6
}
```

---

### Пример с некорректным выражением:

#### Запрос:

```bash
curl --location --request POST 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+2*a"
}'
```

#### Ответ:

```json
{
  "error": "Expression is not valid"
}
```

**Код ответа**: `422 Unprocessable Entity`.

---

### Пример с ошибкой сервера:

Если произойдет внутренняя ошибка, вы получите:

#### Ответ:

```json
{
  "error": "Internal server error"
}
```

**Код ответа**: `500 Internal Server Error`.

---

## Пример с использованием Postman

1. Откройте Postman.
2. Создайте новый запрос.
3. Выберите метод **POST**.
4. Введите URL: `http://localhost:8080/api/v1/calculate`.
5. Перейдите во вкладку **Body**, выберите **raw** и формат **JSON**.
6. Введите тело запроса:
   ```json
   {
     "expression": "2+2*2"
   }
   ```
7. Нажмите **Send**.

Вы получите результат в формате JSON!

---

## Структура проекта

```
.
├── cmd
│   └── main.go
├── go.mod
├── go.sum
├── internal
│   ├── calc
│   │   └── calc.go
│   └── handlers
│       └── handlers.go
└── README.md

5 directories, 6 files
```

---

