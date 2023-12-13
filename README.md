Первый опыт по написанию каких-либо сервисов на go, был бы рад получить любой фидбэк по проекту или советы на будущее.  

# О сервисе
Сервис по сокращению ссылок в 10 символьные уникальные токены. Поддерживает два метода http запроса:  
1. Post - сохраняет оригинальный URL в базе и возвращает сокращённый  
2. Get - принимает сокращённый URL и возвращает оригинальный URL, если он есть в базе  

Проект был сделан для получения практического опыта и в нем реализована:  
- Следование REST API дизайну  
- Работа с фремворком gin  
- Работа с базой данных postgresql  
- Работа с окружением godotenv + viper  
- Работа с mocks  
- Частичное unit тестирование  

# Работа с api:  
Post - получить сокращенную ссылку  
Запрос: curl.exe -X POST localhost:8080/api/send_url -H "Content-Type: application/json" -d '{"base_url" : "https://github.com/Xonesent"}'  
Ответ: {"Short_URL": "g1DOFTf3uP"}  

Get - получить обычную ссылку  
Запрос: curl.exe -X GET localhost:8080/api/get_url/g1DOFTf3uP  
Ответ: {"Base_URL":"https://github.com/Xonesent"}  

# Что хотелось бы доделать / исправить  
- Запуск из Docker (docker-compose build app)  
Не очень понимаю, что именно работает не так, но мне вылетают две следующие ошибки:  
1. /pkg/mod/github.com/spf13/afero@v1.10.0/internal/common/adapters.go:16:8: package io/fs is not in GOROOT (/usr/local/go/src/io/fs)  
2. /pkg/mod/github.com/sagikazarmark/slog-shim@v0.1.0/attr.go:10:2: package log/slog is not in GOROOT (/usr/local/go/src/log/slog)  
Самостоятельный запуск работает отлично, но не из Docker. Надеюсь на чью-то поддержку!  

Почта - 1pyankov.d.s@gmail.com  
Телега - @Xonesent  