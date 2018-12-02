time for i in ../birthmark_server/data/search_birthmark/*2gram*.csv
do
    curl 'http://localhost:'+ port +'/solr/2gram/select?q=*:*&fl=output,score,place,barthmark,data&rows=1000000&sort=score%20desc&wt=csv' > /dev/null
done

time for i in ../birthmark_server/data/search_birthmark/*3gram*.csv
do
    curl 'http://localhost:'+ port +'/solr/3gram/select?q=*:*&fl=output,score,place,barthmark,data&rows=1000000&sort=score%20desc&wt=csv' > /dev/null
done
time for i in ../birthmark_server/data/search_birthmark/*4gram*.csv
do
    curl 'http://localhost:'+ port +'/solr/4gram/select?q=*:*&fl=output,score,place,barthmark,data&rows=1000000&sort=score%20desc&wt=csv' > /dev/null
done
time for i in ../birthmark_server/data/search_birthmark/*5gram*.csv
do
    curl 'http://localhost:'+ port +'/solr/5gram/select?q=*:*&fl=output,score,place,barthmark,data&rows=1000000&sort=score%20desc&wt=csv' > /dev/null
done
time for i in ../birthmark_server/data/search_birthmark/*6gram*.csv
do
    curl 'http://localhost:'+ port +'/solr/6gram/select?q=*:*&fl=output,score,place,barthmark,data&rows=1000000&sort=score%20desc&wt=csv' > /dev/null
done
time for i in ../birthmark_server/data/search_birthmark/*uc*.csv
do
    curl 'http://localhost:'+ port +'/solr/uc/select?q=*:*&fl=output,score,place,barthmark,data&rows=1000000&sort=score%20desc&wt=csv' > /dev/null
done
