# Yandex_Calculator
Калькулятор для Яндекс Лицея
Запуск:
git clone https://github.com/TerrariumDH/Yandex_Calculator
cd Yandex_Calculator
go run main/main.go
Приложение поддерживает post запросы с json формата {"expression": "ваше выражение"}
Пример работы:
status code 200
curl --location http://localhost:8080/api/v1/calculate --header 'Content-Type: application/json' --data '{"expression": "2+2*2"}'
status code 422
curl --location http://localhost:8080/api/v1/calculate --header 'Content-Type: application/json' --data '{"expression": "2+2*2/0"}'
Спасибо Максу за ридми, я не умею их делать просто
