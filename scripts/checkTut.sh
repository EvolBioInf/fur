cd data
update_blastdb.pl --decompress Betacoronavirus
update_blastdb.pl --decompress taxdb
export BLASTDB=$(pwd)
cd ..
checkPrim -v query=data/p1.fa \
                    -v db=data/Betacoronavirus \
                    -v taxids=2697049
checkPrim -v query=data/p1.fa \
                    -v db=data/Betacoronavirus \
                    -v negativeTaxids=2697049
