cd ../data
update_blastdb.pl --decompress Betacoronavirus
update_blastdb.pl --decompress taxdb
export BLASTDB=$(pwd)
cd ..
./build/checkPrim -v query=data/p1.fa \
                    -v db=data/Betacoronavirus \
                    -v taxid=2697049
./build/checkPrim -v query=data/p1.fa \
                    -v db=data/Betacoronavirus \
                    -v negativeTaxid=2697049
