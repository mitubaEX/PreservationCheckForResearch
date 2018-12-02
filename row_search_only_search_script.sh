for i in ../birthmark_server/data/search_birthmark/*2gram*.csv
do
    python3 row_search_only_search.py $i 15 2gram "$1"
done > "2gram-$1-only-search.csv"

for i in ../birthmark_server/data/search_birthmark/*3gram*.csv
do
    python3 row_search_only_search.py $i 24 3gram "$1"
done > "3gram-$-only-search1.csv"

for i in ../birthmark_server/data/search_birthmark/*4gram*.csv
do
    python3 row_search_only_search.py $i 31 4gram "$1"
done > "4gram-$1-only-search.csv"

for i in ../birthmark_server/data/search_birthmark/*5gram*.csv
do
    python3 row_search_only_search.py $i 37 5gram "$1"
done > "5gram-$1-only-search.csv"

for i in ../birthmark_server/data/search_birthmark/*6gram*.csv
do
    python3 row_search_only_search.py $i 41 6gram "$1"
done > "6gram-$1-only-search.csv"

for i in ../birthmark_server/data/search_birthmark/*uc*.csv
do
    python3 row_search_only_search.py $i 2 uc  "$1"
done > "uc-$1-only-search.csv"

