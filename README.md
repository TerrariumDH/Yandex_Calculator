# Yandex_Calculator
Этот проект представляет собой веб-сервис, который принимает арифметические выражения через HTTP-запросы и возвращает результат их вычисления.

Установка

1. Установите Go с Офиц. Сайта => https://golang.org/dl/

2. Скопируйте репозиторий:
git clone https://github.com/TerrariumDH/Yandex_Calculator.git
cd Yandex_Calculator

3. Запустите проект с помощью команды:
go run ./cmd/main.go

Сервис будет доступен по адресу: http://localhost:8080/api/v1/calculate

Использование

Эндпоинт
POST /api/v1/calculate

Шаблон Запроса

curl -i -X POST http://localhost:8080/api/v1/calculate/ \
     -H "Content-Type: application/json" \
     -d '{"expression":"EXPRESSION"}'

Тело запроса
Пример:

{
  "expression": "2+2*2"
}

Корректный запрос:
Введя данный запрос:

curl --location 'http://127.0.0.1:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
    "expression":"2+2*2/(2+2)"
}'
вы получите ответ:

{"result":3}
с кодом [200].

Запрос с методом не POST:
Введя данный запрос:

curl --location --request GET 'http://127.0.0.1:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
    "expression":"2+2"
}'
вы получите ответ:

{"error": "only POST method allowed"}
с кодом [405].

Запрос с неправильным телом:
Введя данный запрос:

curl --location 'http://127.0.0.1:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
    "expression":"2+2
}'
вы получите ответ:

{"error": "Bad request"}
с кодом [400].

Запрос с делением на 0(ноль)
Введя данный запрос:

curl --location 'http://127.0.0.1:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
    "expression":"2+2*2/0"
}'
вы получите ответ:

{"error":"Expression is not valid. Division by zero"}
с кодом [422].

Запрос с не закрытой скобкой
Введя данный запрос:

curl --location 'http://127.0.0.1:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
    "expression":"2+(9+7"
}'
вы получите ответ:

{"error": "Expression is not valid. Number of brackets doesn't match"}
с кодом [422].

Запрос с выражением с буквами
Введя данный запрос:

curl --location 'http://127.0.0.1:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
    "expression":"2+(9+x)"
}'
вы получите ответ:

{"error": "Expression is not valid. Only numbers and ( ) + - * / allowed"}
с кодом [422].

Запрос с выражением c лишними знаками действия
Введя данный запрос:

curl --location 'http://127.0.0.1:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
    "expression":"2++2"
}'
вы получите ответ:

{"error": "Expression is not valid. Not enough values"}
с кодом [422].

В случае иной ошибки на стороне сервера будет получен ответ:

{"error":"Internal server error"}
с кодом [500].
