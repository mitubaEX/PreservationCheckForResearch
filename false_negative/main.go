package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func changeStringOrder(l []string) string {
	if l[0] > l[1] {
		return l[1] + "," + l[0]
	}
	return l[0] + "," + l[1]
}

func contains(s []string, e string) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}

func main() {
	// 検索結果を取ってきて，検索したファイル名が検索結果に含まれているか．
	// そして含まれていたらその類似度は？
	searchResultFiles, err := ioutil.ReadDir("../search_result")
	if err != nil {
		panic(err)
	}
	// count7 := 0
	// count5 := 0
	// count2 := 0
	// all := 0
	// other := 0
	var classSlice []string = []string{}
	var resultSlice []string = []string{}

	for _, file := range searchResultFiles {
		if !strings.Contains(file.Name(), os.Args[1]) {
			continue
		}
		input, err := os.Open("../search_result/" + file.Name())
		if err != nil {
			panic(err)
		}
		defer input.Close()

		scanner := bufio.NewScanner(input)

		scanner.Scan()

		// result := strings.Split(scanner.Text(), ",")
		//
		// var flag = false
		for scanner.Scan() {
			tmp := strings.Split(scanner.Text(), ",")
			if len(tmp) <= 1 {
				continue
			}
			changeString := changeStringOrder(tmp)
			if !contains(classSlice, changeString) {
				flo, err := strconv.ParseFloat(tmp[1], 64)
				if err != nil {
					panic(err)
				}
				classSlice = append(classSlice, changeString)
				if flo >= 0.25 {
					resultSlice = append(resultSlice, changeString+","+tmp[1])
				}
			}

			// if tmp[0] == result[0] {
			// 	flag = true
			// 	flo, err := strconv.ParseFloat(tmp[1], 64)
			// 	if err != nil {
			// 		panic(err)
			// 	}
			// 	if flo >= 0.75 {
			// 		count7++
			// 	}
			// 	if flo >= 0.5 {
			// 		count5++
			// 	}
			// 	if flo >= 0.25 {
			// 		count2++
			// 	}
			// 	all++
			// }
		}
		// if flag == false {
		// 	other++
		// }
		input.Close()
	}
	// fmt.Printf("other: %d, all: %d, count7: %d, count5: %d, count2: %d\n", other, all, count7, count5, count2)
	for _, v := range resultSlice {
		fmt.Println(v)
	}
}
