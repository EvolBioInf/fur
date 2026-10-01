./fur -d test.db &> r1.txt
./fur -d test.db -q 0.5 &> r2.txt
./fur -d test.db -q 0.5 -w 150 &> r3.txt
./fur -d test.db -q 0.5 -w 150 -t 8 &> r4.txt
./fur -d test.db -M &> r5.txt
./fur -d masked.db &> r6.txt
./fur -d masked.db -M &> r7.txt
./fur -d testPartial.db -f 0.8 &> r8.txt
./fur -d test.db -m -W 4 &> r9.txt
