for a in $(seq 100); do
    stan -T t -N n -t 1 -l 1000000 -o
    ./makeFurDb -s 0 -t t -n n -d bla.db -o
    fur -d bla.db 2>&1 |
	grep _1
done
