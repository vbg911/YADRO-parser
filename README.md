# YADRO-parser

Тестовое задание для стажировки в YADRO.

Требуется написать прототип системы, которая следит за работой компьютерного клуба, обрабатывает события и подсчитывает выручку за день и время занятости каждого стола.


## Запуск

### Обычный запуск

1. Клонировать репозиторий:

```bash
   git clone https://github.com/vbg911/YADRO-parser.git
```

2. Перейти в директроию cmd/parser:

```bash
   cd cmd/parser
```

3. Собрать проект:

```bash
   go build -o task.exe
```

### Запуск используя Docker

1. Собрать образ приложения:

```bash
   docker build -t task-yadro .
```

2. Запустить контейнер:

следуя шаблону

```bash  
   docker run -v {абсолютный путь к тестовому файлу}:/{имя в контейнере} task-yadro ./task.exe {имя в контейнере}
```  

например

```bash
   docker run -v C:\Users\vbg\Desktop\YADRO-parser\test\case1\simple_test.txt:/text.txt task-yadro ./task.exe text.txt
```

## Тестирование
1. Запуск тестов находясь в корне проекта:
```bash
    go test ./... -v
```

[Файл](test/cover.html) отображающий покрытие кода тестами.