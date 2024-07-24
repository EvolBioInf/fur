mkdir tutorial
cd tutorial
stan -s 3
ls targets/ | wc -l
ls neighbors/ | wc -l
phylonium targets/* neighbors/* |
    nj |
    midRoot |
    plotTree
cres targets/*
cres neighbors/*
makeFurDb -t targets -n neighbors -d test.db
fur -d test.db/
cutSeq -r 4492-5500 test.db/r.fasta
fur -d test.db -u > unique1.fasta
cres unique1.fasta
fur -d test.db -U > unique2.fasta
cres unique2.fasta
stan -o -l 5000000
/usr/bin/time -f "user: %Us; elapsed: %es; memory: %Mkb" \
                makeFurDb -t targets/ -n neighbors/ \
                -d test.db -o
/usr/bin/time -f "user: %Us; elapsed: %es; memory: %Mkb" \
                makeFurDb -t targets/ -n neighbors/ \
                -d test.db -o -T 1
/usr/bin/time -f "user: %Us; elapsed: %es; memory: %Mkb" \
                fur -d test.db/
stan -o -l 5000000 -n 20
/usr/bin/time -f "user: %Us; elapsed: %es; memory: %Mkb" \
                makeFurDb -t targets/ -n neighbors/ \
                -d test.db/ -o
rm -r test.db targets neighbors
wget guanine.evolbio.mpg.de/fur/eco105.tar.gz
tar -xvzf eco105.tar.gz
makeFurDb -t targets -n neighbors -d eco105.db
fur -d eco105.db/ > eco105.fasta
for a in $(seq 60 120)
do
    echo -n $a " "
    fur -w $a -q 0.4 -d eco105.db/ 2>&1 |
          grep "Subtraction_2" |
          awk '{print $3}'
done > yield.dat
plotLine yield.dat
fur -m -d eco105.db/ > eco105_2.fasta
fur -e 1e-2 -d eco105.db/ > eco105_3.fasta
