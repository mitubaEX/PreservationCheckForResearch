rm compare_result/comp_result.csv
java -jar -Xmx20480m ./pochi/pochi-runner/target/pochi-runner-1.0-SNAPSHOT.jar ./compare.js ./search_result/a.csv ./search_result/b.csv > ./compare_result/comp_result.csv
