package main

import (
	"bufio"
	"bytes"
	"os"
	"strings"
)

func readerConfig(dir string) (map[string]string, error) {
	f, err := os.Open(dir)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Чтение файла с ридером
	wr := bytes.Buffer{}
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		wr.WriteString(sc.Text())
	}

	// Удаление пробелова из конфигурационного файла для упращения алгоритма.
	File := strings.Replace(wr.String(), " ", "", -1)

	// Разделяем перечень параметров в конфигурационном файле через ;
	// Получаем срез строк в формате параметр=значение
	slFile := breakString(File, ';')

	// Инициализируем карту в которую будем складывать map[параметр]значение.
	config := make(map[string]string)
	// Цикл для переборки среза строк slFile (в формате параметр=значение) и создания
	// карты. Разделяем параметр и значение на 2 строки записываем в срез для того чтобы
	// создать карту map[параметр]значение. arg+" " - костыль алгоритм не совершенен.
	for _, arg := range slFile {
		slArg := breakString(arg+" ", '=')
		config[slArg[0]] = slArg[1]
	}

	return config, nil
}

// Функция для создания среза из строки. Разделяет строку (текст|char|текст).
// Новые элементы среза ([0]текст(|char|[1]текст... и т.д.).
func breakString(str string, char rune) []string {
	slice := []string{}
	for len(str) > 0 {
		for i, ch := range str {
			if ch == char || len(str) == i+1 {
				slice = append(slice, str[:i])
				str = str[i+1:]
				break
			}
		}
	}
	return slice
}
