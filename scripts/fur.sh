echo -n "# Downloading data..."
curl guanine.evolbio.mpg.de/fur/test.tar.gz -o test.tar.gz
tar -xzf test.tar.gz
echo "done."
echo -n "# Making fur database..."
./makeFurDb -t testTar -n testNei -d furDb 2>/dev/null
echo "done."
echo -n "# Running fur..."
./fur -d furDb > tmp.out 2>/dev/null
echo "done."
DIFF=$(diff tmp.out ../data/fur.out)
if [ "$DIFF" == "" ] 
then
    printf "Test(fur)\tpass\n"
else
    printf "Test(fur)\tfail\n"
    echo ${DIFF}
fi

rm -r testTar testNei test.tar.gz furDb tmp.out
