# TODO-go
---

*todo-go* - todo, написанный с применением монолитной архитектуры, построенный на фреймворке [gin](https://github.com/gin-gonic/gin)

# Цели и задачи
---
Основная цель проекта — практика разработки REST API, работы с базой данных и построения серверных приложений на GO.

## Описание
---

Сервис предоставляет REST API для выполнения CRUD-операций над задачами (создание, получение, обновление и удаление).

Проект реализован в виде монолита с разделением на слои:
- HTTP-обработчики
- сервисный слой
- слой работы с данными (repository)

## Стек технологий
---

- Go
- [Gin](https://github.com/gin-gonic/gin)
- PostgreSQL
- Docker / Docker Compose
- [Minimock](https://github.com/gojuno/minimock)
- [Testify](https://github.com/stretchr/testify)
- [Goose](https://github.com/pressly/goose)

## Архитектура
---

Проект построен с использованием монолитной архитектуры
и классического layered-подхода:

```
cmd/            - точка входа в приложение
internal/
  app/          - DI контейнер 
  handlers/     - HTTP-обработчики
  service/      - бизнес-логика
  repo/         - работа с БД
  models/       - модели
  router/       - настройка маршрутов
```

# Запуск проекта
---
### Локальный запуск

```bash
go mod tidy
go run cmd/todo/main.go
```

---

### Docker

```bash
docker compose up -d --build
```

Хост и порт настраивается в переменной окружения.

