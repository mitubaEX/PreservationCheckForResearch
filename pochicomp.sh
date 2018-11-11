rm compare_result/compare_result.csv
java -jar $HOME/pochi/pochi-runner/target/pochi-runner-1.0-SNAPSHOT.jar $HOME/compare.js ./search_result/a.csv ./search_result/b.csv > ./compare_result/comp_result.csv
