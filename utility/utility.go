package utility

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Mkdir(directoryName string) error {
	exists := func(filename string) bool {
		_, err := os.Stat(filename)
		return err == nil
	}
	if !exists(directoryName) {
		if err := os.Mkdir(directoryName, 0777); err != nil {
			return err
		}
	}
	return nil
}

// 16進数変換してその結果を連結して返す
func StringToHexString(str string) string {
	var mergedStringSlice = []string{}
	getHexString := func(str string) string {
		var hexString string = ""
		for _, v := range strings.Split(str, " ") {
			atoi, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			hexString += fmt.Sprintf("%x", atoi)
		}
		return hexString
	}
	for _, v := range strings.Split(str, ",") {
		mergedStringSlice = append(mergedStringSlice, getHexString(v))
	}
	return strings.Join(mergedStringSlice, ",")
}

func WriteFile(filename string, fileBody string) {
	// file open
	output, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer output.Close()

	// output
	output.WriteString(fileBody)
}
