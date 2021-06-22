./cleanSeq ../data/cleanSeqIn.fasta > tmp.fasta
DIFF=$(diff tmp.fasta ../data/cleanSeqOut.fasta)
if [ "$DIFF" == "" ] 
then
    printf "Test(cleanSeq)\tpass\n"
else
    printf "Test(cleanSeq)\tfail\n"
    echo ${DIFF}
fi

rm -r tmp.fasta
