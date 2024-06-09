package ORPXI

import "goCmd/structs"

var commands = []structs.Command{
	{"whois", "Информация о домене"},
	{"pingview", "Показывает пинг"},
	{"traceroute", "Трассировка маршрута"},
	{"extractzip", "Распаковывает архивы .zip"},
	{"signout", "Пользователь выходит из ORPXI"},
	{"newshablon", "Создает новый шаблон комманд для выполнения"},
	{"shablon", "Выполняет определенный шаблон комманд"},
	{"newuser", "Новый пользователь для ORPXI"},
	{"promptSet", "Изменяет ORPXI"},
	{"systemgocmd", "Вывод информации о ORPXI"},
	{"rename", "Переименовывает файл"},
	{"remove", "Удаляет файл"},
	{"read", "Выводит на экран содержимое файла"},
	{"write", "Записывает данные в файл"},
	{"create", "Создает новый файл"},
	{"exit", "Выход из программы"},
	{"orpxi", "Запускает ещё одну ORPXI"},
	{"wifiutils", "Запускает утилиту для работы с WiFi"},
	{"clean", "Очистка экрана"},
	{"matrixmul", "Умножение больших матриц"},
	{"primes", "Поиск больших простых чисел"},
	{"picalc", "Вычисление числа π"},
	{"fileio", "Тест на интенсивную работу с файлами"},
	{"cd", "Смена текущего каталога"},
	{"edit", "Редактирует файл"},
	{"ls", "Выводит содержимое каталога"},
	{"scanport", "Сканирование портов"},
	{"dnslookup", "DNS-запросы"},
	{"ipinfo", "Информация об IP-адресе"},
	{"geoip", "Геолокация IP-адреса"},
}

var commandHistory []string