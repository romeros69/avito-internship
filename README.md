# Тестовое задание а Авито на бэкенд

#### Выполнил: Логинов Роман Алексеевич

### Как запустить (из корня проекта):
    docker-compose up --build

### SWAGGER UI:
http://localhost:9000/swagger/index.html#

### Postman collection:
https://github.com/romeros69/avito-internship/blob/main/assets/endpoints_postman.json

### Database diagram:
https://github.com/romeros69/avito-internship/blob/main/assets/Diagram.png

### Описание реализованных задач:
* Получение баланса пользователя
* Пополнение баланса пользователя
* Резервирование денег с основного счета
* Принятие выполнения услуги, списываение денег с резерва
* Разрезервирование - отмена услуги, возврат денег с резерва на основной счёт
* **Доп задание 1** Получение отчета по выручке с генерацией csv файла
* **Доп задание 2** Получение истории тразакций баланса пользователя 
* Swagger документация к API
* Тесты

### Что было использовано в проекте:
* [gin](https://github.com/gin-gonic/gin) - Web framework
* [pgx](https://github.com/jackc/pgx) - PostgreSQL driver and toolkit for Go
* [uuid](https://github.com/google/uuid) - UUID
* [swag](https://github.com/swaggo/swag) - Swagger
* [gomock](https://github.com/golang/mock) - Mocking framework
* [Docker](https://www.docker.com/) - Docker

### Ограничения и заметки
* Тип id у сущностей - UUID
* Предпологается, что изначально в базе данных есть информация о доступных сервисах
В данной версии бд имеется 3 сервиса:
  1) UUID: 0ba5b953-9df7-4170-80bf-50d3d8e1111d; Tittle: "cleaning"
  2) UUID: 0ba5b953-9df7-4170-80bf-50d3d8e2222d; Tittle: "repair"
  3) UUID: 0ba5b953-9df7-4170-80bf-50d3d8e3333d; Tittle: "massage"

    ***Эта информация нужна для проведения тестирования запросов*** 
* При пополнении баланса в теле запроса добавлено поле source - номер карты, 
иначе говоря источник пополнения. Это поле было добавлено, что бы можно было отслеживать
историю тразакций баланса.
* В методе запроса генерации отчета по выручке в ответе приходит ссылка на скачивание *scv* файла
    
