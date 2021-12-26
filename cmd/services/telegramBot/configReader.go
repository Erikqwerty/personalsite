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

	File := strings.Replace(wr.String(), " ", "", -1)

	config := make(map[string]string)

	slFile := breakString(File, ';')
	for _, arg := range slFile {
		slArg := breakString(arg+" ", '=')
		config[slArg[0]] = slArg[1]
	}

	return config, nil
}
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
