# for i in ../birthmark_server/data/birthmark/*2gram* ; do python3 row_search.py $i 15 2gram ; done
# for i in ../birthmark_server/data/birthmark/*3gram* ; do python3 row_search.py $i 24 3gram ; done
# for i in ../birthmark_server/data/birthmark/*4gram* ; do python3 row_search.py $i 31 4gram ; done
# for i in ../birthmark_server/data/birthmark/*5gram* ; do python3 row_search.py $i 37 5gram ; done
# for i in ../birthmark_server/data/birthmark/*6gram* ; do python3 row_search.py $i 41 6gram ; done
for i in ../birthmark_server/data/birthmark/*uc* ; do python3 row_search.py $i 2 uc ; done
