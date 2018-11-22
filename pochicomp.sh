rm compare_result/comp_result.csv
export JAVA_HOME=`/usr/libexec/java_home -v "11"`
java -jar ./pochi_new/pochi-cli/target/pochi-cli-1.0-SNAPSHOT.jar ./compare_new.js ./search_result/a.csv ./search_result/b.csv > ./compare_result/comp_result.csv
