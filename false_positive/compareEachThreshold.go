package main

import (
	"../utility"
	"bufio"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func CompareFromSearchResult(threshold string) {
	utility.Mkdir("../each_compare_result")
	utility.Mkdir("../each_compare_result/" + threshold)
	searchResultFiles, err := ioutil.ReadDir("../search_result/" + threshold)
	if err != nil {
		panic(err)
	}
	for _, file := range searchResultFiles {
		if file.IsDir() {
			continue
		}

		input, err := os.Open("../search_result/" + threshold + "/" + file.Name())
		if err != nil {
			panic(err)
		}
		defer input.Close()
		scanner := bufio.NewScanner(input)

		var lineCount int = 0
		var searchResult = []string{}
		var outputFilename string

		// 検索バースマークと比較バースマークを取得する
		for scanner.Scan() {
			if lineCount == 0 {
				searchBirthmark := scanner.Text()
				outputFilename = strings.Split(searchBirthmark, ",")[0]
				utility.WriteFile("a.csv", searchBirthmark)
			} else if lineCount >= 1 {
				searchResult = append(searchResult, scanner.Text())
			}
			lineCount += 1
		}

		// 例外処理
		if len(searchResult) <= 0 {
			continue
		}
		utility.WriteFile("b.csv", strings.Replace(
			strings.Join(searchResult[:len(searchResult)-1], "\n"), "\\", "", -1))

		// compare
		v, err := exec.Command("java", "-Xms20480m", "-Xmx20480m", "-jar",
			"../pochi/pochi-runner/target/pochi-runner-1.0-SNAPSHOT.jar",
			"../compare.js", "a.csv", "b.csv").Output()
		if err != nil {
			panic(err)
		}
		utility.WriteFile("../each_compare_result"+"/"+threshold+"/"+outputFilename, string(v))
	}
}
