CF  = -O3 -Wall -Wshadow -pedantic -std=gnu99 -g
CC  = gcc $(CF)
OBJ = eprintf.o
libcommon.a: $(OBJ)
	ar -cvq $@ $^
	ranlib $@
.PHONY: clean
clean:
	rm -f *.o libcommon.a
