#!/bin/bash
# for i in ../birthmark_server/data/search_birthmark/*2gram* ; do python3 ./row_search_get_only_last_result.py $i 15 2gram $1; done
# for i in ../birthmark_server/data/search_birthmark/*3gram* ; do python3 ./row_search_get_only_last_result.py $i 24 3gram $1; done
# for i in ../birthmark_server/data/search_birthmark/*4gram* ; do python3 ./row_search_get_only_last_result.py $i 31 4gram $1; done
# for i in ../birthmark_server/data/search_birthmark/*5gram* ; do python3 ./row_search_get_only_last_result.py $i 37 5gram $1; done
# for i in ../birthmark_server/data/search_birthmark/*6gram* ; do python3 ./row_search_get_only_last_result.py $i 41 6gram $1; done
# for i in ../birthmark_server/data/search_birthmark/*uc* ; do python3 ./row_search_get_only_last_result.py $i 2 uc $1; done

# paste <(for i in 2gram 3gram 4gram 5gram 6gram uc;do echo $i ;done) <(for j in 15 24 31 37 41 2; do echo $j ;done) | while read birth threshold ; do for i in birthmark_server/data/birthmark/*"$birth".csv; do for ethre in 0.0 0.1 0.2 0.3 0.4 0.5 0.6 0.7 0.8 0.9 ;do python3 row_search.py $i $threshold "$birth" ;done; done > "$birth-$threshold.csv";done

# paste <(for i in 2gram 3gram 4gram 5gram 6gram uc;do echo $i ;done) <(for j in 15 24 31 37 41 2; do echo $j ;done) | while read birth threshold ; do
#
# done

for i in ../birthmark_server/data/search_birthmark/*2gram*.csv
do
    python3 row_search_get_only_last_result.py $i 15 2gram "$1"
done > "2gram-$1.csv"

for i in ../birthmark_server/data/search_birthmark/*3gram*.csv
do
    python3 row_search_get_only_last_result.py $i 24 3gram "$1"
done > "3gram-$1.csv"

for i in ../birthmark_server/data/search_birthmark/*4gram*.csv
do
    python3 row_search_get_only_last_result.py $i 31 4gram "$1"
done > "4gram-$1.csv"

for i in ../birthmark_server/data/search_birthmark/*5gram*.csv
do
    python3 row_search_get_only_last_result.py $i 37 5gram "$1"
done > "5gram-$1.csv"

for i in ../birthmark_server/data/search_birthmark/*6gram*.csv
do
    python3 row_search_get_only_last_result.py $i 41 6gram "$1"
done > "6gram-$1.csv"

for i in ../birthmark_server/data/search_birthmark/*uc*.csv
do
    python3 row_search_get_only_last_result.py $i 2 uc  "$1"
done > "uc-$1.csv"
