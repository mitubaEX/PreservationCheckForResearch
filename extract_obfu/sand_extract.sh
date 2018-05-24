 mkdir DNR DR IOP MLI IR
for i in ./jar/*.jar ; do java -cp sandmark.jar sandmark.smash.SandmarkCLI -O -A "Insert Opaque Predicates" -i $i -o ./IOP/`basename "$i"`-IOP.jar ;done
for i in ./jar/*.jar ; do java -cp sandmark.jar sandmark.smash.SandmarkCLI -O -A "Irreducibility" -i $i -o ./IR/`basename "$i"`-IR.jar ;done
for i in ./jar/*.jar ; do java -cp sandmark.jar sandmark.smash.SandmarkCLI -O -A "Duplicate Registers" -i $i -o ./DR/`basename "$i"`-DR.jar ;done
for i in ./jar/*.jar ; do java -cp sandmark.jar sandmark.smash.SandmarkCLI -O -A "Merge Local Integers" -i $i -o ./MLI/`basename "$i"`-MLI.jar ;done
