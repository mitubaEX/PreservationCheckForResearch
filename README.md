# ToolForResearch

## clone project

```
git clone --recursive https://github.com/mitubaEX/ToolForResearch.git
```

## Install donquixote

```
cd donquixote
mvn package
```

## Install sandmark

Please put sandmark into sandmark directory.

[sandomark](http://sandmark.cs.arizona.edu/)

## Install pochi

```
cd pochi
mvn package
```

## Install Dependencies

```
glide install
```

## Usage

```
go build
./ToolForResearch -f <filename>
```

all option setting

```
./ToolForResearch -f <filename> -h <host> -p <port> -r <rows> -b <birthmark>
```

### check false positive

```
cd false_positive
go build
./false_positive
```

## Option

```
flag needs an argument: -h
Usage of ./ToolForResearch:
  -b string
        birthmark for search (default "2gram")
  -f string
        filename for search
  -h string
        host of search engine (default "localhost")
  -i string
        seach result dir (default "search_result")
  -l string
        threshold length of birthmark
  -m value
        mode of script, modes{search, compare}
  -o string
        compare result dir (default "compare_result")
  -p string
        port of search engine (default "8983")
  -r string
        rows for search (default "10")
```


## Dependencies

- glide:0.12.3
- donquixote
- sandomark:V3.4
