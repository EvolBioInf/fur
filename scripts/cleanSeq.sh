./cleanSeq ../data/cleanSeqIn.fasta > tmp.fasta
DIFF=$(diff tmp.fasta ../data/cleanSeqOut.fasta)
if [ "$DIFF" == "" ] 
then
    printf "Test(cleanSeq) - 1\tpass\n"
else
    printf "Test(cleanSeq) - 1\tfail\n"
    echo ${DIFF}
fi

./cleanSeq -m 2 -l 2 ../data/test1.fasta > tmp.fasta
DIFF=$(diff tmp.fasta ../data/cleanSeqOut2.fasta)
if [ "$DIFF" == "" ] 
then
    printf "Test(cleanSeq) - 2\tpass\n"
else
    printf "Test(cleanSeq) - 2\tfail\n"
    echo ${DIFF}
fi

./cleanSeq -m 2 -l 2 ../data/test2.fasta > tmp.fasta
DIFF=$(diff tmp.fasta ../data/cleanSeqOut2.fasta)
if [ "$DIFF" == "" ] 
then
    printf "Test(cleanSeq) - 3\tpass\n"
else
    printf "Test(cleanSeq) - 3\tfail\n"
    echo ${DIFF}
fi

./cleanSeq -m 2 -l 2 ../data/test3.fasta > tmp.fasta
DIFF=$(diff tmp.fasta ../data/cleanSeqOut2.fasta)
if [ "$DIFF" == "" ] 
then
    printf "Test(cleanSeq) - 4\tpass\n"
else
    printf "Test(cleanSeq) - 4\tfail\n"
    echo ${DIFF}
fi

./cleanSeq -m 2 -l 2 ../data/test4.fasta > tmp.fasta
DIFF=$(diff tmp.fasta ../data/cleanSeqOut2.fasta)
if [ "$DIFF" == "" ] 
then
    printf "Test(cleanSeq) - 5\tpass\n"
else
    printf "Test(cleanSeq) - 5\tfail\n"
    echo ${DIFF}
fi

rm -r tmp.fasta
