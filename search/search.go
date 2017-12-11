package search

import (
	"../settings"
	"../utility"
	"bufio"
	"os"
	"strconv"
	"strings"
	"time"
)

func splitString(text string) []string {
	split := strings.SplitN(text, ",", 4)
	return split
}

func Search(argument settings.Argument) {
	// file read
	input, err := os.Open(argument.Filename)
	if err != nil {
		panic(err)
	}
	defer input.Close()

	outputDirectoryName := "./search_result"
	utility.Mkdir(outputDirectoryName)

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		row := splitString(scanner.Text())
		if len(row) < 4 {
			continue
		}
		// 16進変換
		encodedString := utility.StringToHexString(row[3])

		start := time.Now()
		// 検索
		response, err := postJson(argument, encodedString)
		if err != nil {
			panic(err)
		}

		// 検索結果をファイルに書き込む
		output, err := os.Create(outputDirectoryName + "/" + row[0] + argument.Birthmark)
		if err != nil {
			panic(err)
		}
		defer output.Close()

		// output
		output.WriteString(strings.Join(row, ",") + "\n")
		//output.WriteString(utility.ScoreNormalization(strings.Replace(
		//	strings.Replace(response, "quot;", "", -1), "\"", "", -1)) + "\n")
		output.WriteString(strings.Replace(
			strings.Replace(response, "quot;", "", -1), "\"", "", -1) + "\n")
		output.WriteString(strconv.FormatFloat((time.Now().Sub(start)).Seconds(), 'f', -1, 64) + "sec")
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
