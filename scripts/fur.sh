./fur > tmp.out
DIFF=$(diff tmp.out ../data/fur.out)
if [ "$DIFF" == "" ] 
then
    printf "Test(fur)\tpass\n"
else
    printf "Test(fur)\tfail\n"
    echo ${DIFF}
fi

rm tmp.out
