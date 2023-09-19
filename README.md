## Запуск сервиса:
1) Скачать текущий репозиторий.
2) Скачать docker. [docker.com/get-started/](https://www.docker.com/get-started/)
3) Собрать все сервисы, описанные в docker-compose. `docker compose build`
4) Запустить все сервисы, описанные в  docker-compose. `docker compose up -d`

## Для работы с API необходимо: 
1) Создать базу данных. `use mongodb`
2) Добавить коллекцию `user_authentication` и создать в ней документ с полем `user_guid`. Например:
   `db.user_authentication.insertOne({user_guid: "4f075160-5099-11ee-be56-0242ac120002"})`

   **P.S** Все данные(логины, пароли, адреса, имена и др.) есть в [config/config.go](https://github.com/RoKZZ-55/authentication-service/blob/main/config/config.go)

## Документация к запросам:
Ссылка на документацию в сервисе [postman doc](https://documenter.getpostman.com/view/14585414/2s9YC8wWaf)