package settings

import (
	"flag"
	"fmt"
	"strconv"
)

type Argument struct {
	Filename  string
	Birthmark string
	Rows      string
	Host      string
	Port      string
	Modes     []string
	Length    int
	InputDir  string
	OutputDir string
}

type modes []string

func (s *modes) String() string {
	return fmt.Sprintf("%v", multipleModes)
}

func (s *modes) Set(v string) error {
	*s = append(*s, v)
	return nil
}

var multipleModes modes

func SearchInit() Argument {
	// オプション取得
	var f = flag.String("f", "", "filename for search")
	var b = flag.String("b", "2gram", "birthmark for search")
	var r = flag.String("r", "10", "rows for search")
	var h = flag.String("h", "localhost", "host of search engine")
	var p = flag.String("p", "8983", "port of search engine")
	var l = flag.String("l", "0", "threshold length of birthmark")
	var i = flag.String("i", "search_result", "seach result dir")
	var o = flag.String("o", "compare_result", "compare result dir")
	//var m = flag.String("m", "search", "mode of script, modes{search, compare}")
	flag.Var(&multipleModes, "m", "mode of script, modes{search, compare}")
	flag.Parse()
	length, err := strconv.Atoi(*l)
	if err != nil {
		panic(err)
	}
	return Argument{*f, *b, *r, *h, *p, multipleModes, length, *i, *o}
}
