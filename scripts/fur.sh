wget -q guanine.evolbio.mpg.de/fur/test.tar.gz
tar -xzf test.tar.gz
./makeFurDb -t testTar -n testNei -d furDb 2>/dev/null
./fur -d furDb > tmp.out 2>/dev/null
DIFF=$(diff tmp.out ../data/fur.out)
if [ "$DIFF" == "" ] 
then
    printf "Test(fur)\tpass\n"
else
    printf "Test(fur)\tfail\n"
    echo ${DIFF}
fi

rm -r testTar testNei test.tar.gz furDb tmp.out
