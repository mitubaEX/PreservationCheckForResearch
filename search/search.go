package search

import (
	"../settings"
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
	"time"
)

// 16進数変換してその結果を連結して返す
func getEncodedString(str string) string{
	var mergedStringSlice = []string{}
	getHexString := func(str string) string{
		var hexString string = ""
		for _, v := range strings.Split(str, " "){
			atoi, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			hexString += fmt.Sprintf("%x", atoi)
		}
		return hexString
	}
	for _, v := range strings.Split(str, ","){
		mergedStringSlice = append(mergedStringSlice, getHexString(v))
	}
	return strings.Join(mergedStringSlice, ",")
}

func splitString(text string) []string{
	split := strings.SplitN(text,",", 4)
	return split
}

func Search(argument settings.Argument) {
	// file read
	input, err := os.Open(argument.Filename)
	if err != nil {
		panic(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		row := splitString(scanner.Text())
		if len(row) < 4{
			continue
		}
		// 16進変換
		encodedString := getEncodedString(row[3])

		start := time.Now()
		// 検索
		response, err := postJson(argument, encodedString)
		if err != nil {
			panic(err)
		}

		// 検索結果をファイルに書き込む
		output, err := os.Create("data/"+ row[0] + argument.Birthmark)
		if err != nil {
			panic(err)
		}
		defer output.Close()
		output.WriteString(strings.Join(row, ","))
		output.WriteString(response)
		end := time.Now();
		output.WriteString(strconv.FormatFloat((end.Sub(start)).Seconds(), 'f', -1, 64) + "sec")
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
