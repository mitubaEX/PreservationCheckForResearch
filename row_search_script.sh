for i in ../birthmark_server/data/search_birthmark/*2gram* ; do python3 row_search.py $i 15 2gram ; done
for i in ../birthmark_server/data/search_birthmark/*3gram* ; do python3 row_search.py $i 24 3gram ; done
for i in ../birthmark_server/data/search_birthmark/*4gram* ; do python3 row_search.py $i 31 4gram ; done
for i in ../birthmark_server/data/search_birthmark/*5gram* ; do python3 row_search.py $i 37 5gram ; done
for i in ../birthmark_server/data/search_birthmark/*6gram* ; do python3 row_search.py $i 41 6gram ; done
for i in ../birthmark_server/data/search_birthmark/*uc* ; do python3 row_search.py $i 2 uc ; done
