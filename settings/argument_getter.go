package settings

import (
	"flag"
)

type Argument struct {
	Filename  string
	Birthmark string
	Rows      string
	Host      string
	Port      string
	Mode      string
}

func SearchInit() Argument {
	// オプション取得
	var f = flag.String("f", "1234", "filename for search")
	var b = flag.String("b", "1234", "birthmark for search")
	var r = flag.String("r", "1234", "rows for search")
	var h = flag.String("h", "1234", "host of search engine")
	var p = flag.String("p", "1234", "port of search engine")
	var m = flag.String("m", "1234", "mode of script")
	flag.Parse()
	return Argument{*f, *b, *r, *h, *p, *m}
}
