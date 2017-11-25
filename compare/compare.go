package compare

import (
	"../settings"
	"../utility"
	"bufio"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

// ファイルの一行目が検索バースマーク，三行目以降が検索結果
// 一行目と三行目のファイルを作成
// 比較を実行してファイルとしてcompare_resultに吐き出す
func Compare(argument settings.Argument) {
	outputDirectoryName := "./compare_result"
	// ディレクトリ作成
	utility.Mkdir(outputDirectoryName)

	// search_resultからファイルのリストを取って来る
	files, err := ioutil.ReadDir("./search_result")
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		input, err := os.Open("./search_result/" + file.Name())
		if err != nil {
			panic(err)
		}
		defer input.Close()

		var lineCount int = 0
		var searchResult = []string{}
		var outputFilename string
		scanner := bufio.NewScanner(input)

		// 検索バースマークと比較バースマークを取得する
		for scanner.Scan() {
			if lineCount == 0 {
				searchBirthmark := scanner.Text()
				outputFilename = strings.Split(searchBirthmark, ",")[0]
				utility.WriteFile("a.csv", searchBirthmark)
			} else if lineCount >= 2 {
				searchResult = append(searchResult, scanner.Text())
			}
			lineCount += 1
		}

		// 例外処理
		if len(searchResult) <= 0 {
			continue
		}
		utility.WriteFile("b.csv", strings.Join(searchResult[:len(searchResult)-1], "\n"))

		// compare
		v, err := exec.Command("java", "-jar",
			"./pochi/pochi-runner/target/pochi-runner-1.0-SNAPSHOT.jar",
			"compare.js", "a.csv", "b.csv").Output()
		if err != nil {
			panic(err)
		}
		utility.WriteFile(outputDirectoryName+"/"+outputFilename, string(v))
	}
}
