package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	// 検索結果を取ってきて，検索したファイル名が検索結果に含まれているか．
	// そして含まれていたらその類似度は？
	searchResultFiles, err := ioutil.ReadDir("../search_result_" + os.Args[1])
	if err != nil {
		panic(err)
	}
	count7 := 0
	count5 := 0
	count2 := 0
	all := 0
	other := 0

	for _, file := range searchResultFiles {
		input, err := os.Open("../search_result_" + os.Args[1] + "/" + file.Name())
		if err != nil {
			panic(err)
		}
		defer input.Close()

		scanner := bufio.NewScanner(input)

		scanner.Scan()

		result := strings.Split(scanner.Text(), ",")

		var flag = false
		for scanner.Scan() {
			tmp := strings.Split(scanner.Text(), ",")
			if tmp[0] == result[0] {
				flag = true
				flo, err := strconv.ParseFloat(tmp[1], 64)
				if err != nil {
					panic(err)
				}
				if flo >= 0.75 {
					count7++
				}
				if flo >= 0.5 {
					count5++
				}
				if flo >= 0.25 {
					count2++
				}
				all++
			}
		}
		if flag == false {
			other++
		}
		input.Close()
	}
	fmt.Printf("%d, %d, %d, %d, %d\n", other, all, count7, count5, count2)
}
