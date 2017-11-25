# ToolForResearch

## clone project

```
git clone --recursive https://github.com/mitubaEX/PreservationCheckForResearch.git
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
  -m string
    	mode of script, modes{search, compare} (default "search")
  -p string
    	port of search engine (default "8983")
  -r string
    	rows for search (default "10")
```


## Dependencies

- glide:0.12.3
- donquixote
- sandomark:V3.4
