# architecture-lab-3

## Лабораторна робота №3 з Архітектури ПЗ

### Запуск серверу:

1. Спочатку необхідно створити нову базу данних з назвою "restaurant" і додати користувача, після чого запустити код для ініціалізації БД, що знаходиться в "./db/schema.sql". В межах лабораторної роботи використовується користувач 'vsevolodpavlovskyi' з паролем 'admin'. За необхідності можна змінити дані для підключення до БД як назва БД, користувач, пароль та host у файлі "./server/cmd/server/main.go".

2. Сервер запускається з директорії "./server/" командою 
```shell script
    go run ./cmd/server
```

Можна також викликати команду запуску серверу зі флагом -p, після якого вказати номер порту для запуску серверу, наприклад
```shell script
    go run ./cmd/server -p 1234
```

### Запуск прикладів реалізації заданих сценаріїв:

1. Перейти в директорію "./client/" та ввести наступні команди
```shell script
    npm install
    node examples.js
```

### Лабораторну роботу виконали:

- Василашко Анна ІП-84
- Павловський Всеволод ІП-84
