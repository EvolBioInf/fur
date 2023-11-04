for s in 1 0.5 0.25 0; do
    for a in 1 2 5 10 20 50 100; do
	echo -n $a ' '
	ranseq -l ${a}000000 |
	    /usr/bin/time -f "%U %E %M" ./stream -i -s $s  2>&1 |
	    tr -d '\n'
	echo " i $s"
    done
    for a in 1 2 5 10 20 50 100; do
	echo -n $a ' '
	ranseq -l ${a}000000 |
	    /usr/bin/time -f "%U %E %M" ./stream -s $s  2>&1 |
	    tr -d '\n'
	echo " s $s"
    done
done
