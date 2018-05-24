#!/bin/bash -v
mkdir DNR
for i in ./jar/*.jar ;do java -jar ../donquixote/target/donquixote-3.0.1.jar --processor apiblinder --destination des1 --summary report1.txt $i && jar cvf DNR/`basename "$i"` des1/* && rm dest1;done
