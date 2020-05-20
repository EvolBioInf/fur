CC=gcc
export VERSION = $(shell bash ../scripts/version.sh) # The single, trailing blank is essential
export DATE    = $(shell bash ../scripts/date.sh) #    The single, trailing blank is essential
CFLAGS         = -DVERSION=\"$(VERSION)\" -DDATE=\"$(DATE)\" -O3 -Wall -Wshadow -pedantic -std=gnu99 -g -I../common/ -L../common/ #-pg
.SUFFIXES:
.SUFFIXES: .c .o .h

EXECFILE1  = fur 
SRCFILES1  = fur.c interface.c
OBJFILES1  = fur.o interface.o

EXECFILE2  = makeFurDb
SRCFILES2  = makeFurDb.c mfdbI.c
OBJFILES2  = makeFurDb.o mfdbI.o

LIBS      = -lcommon -lm -lbsd -lgsl -lgslcblas


.PHONY : all
all : $(EXECFILE1) $(EXECFILE2) tutorial aux
$(EXECFILE1) : $(OBJFILES1)
	$(CC) $(CFLAGS) -o $(EXECFILE1) $(OBJFILES1) $(LIBS)
fur.c: fur.org
	bash ../scripts/org2nw fur.org | notangle -Rfur.c > fur.c
interface.o: interface.h

$(EXECFILE2) : $(OBJFILES2)
	$(CC) $(CFLAGS) -o $(EXECFILE2) $(OBJFILES2) $(LIBS)
makeFurDb.c: makeFurDb.org
	bash ../scripts/org2nw makeFurDb.org | notangle -RmakeFurDb.c > makeFurDb.c
mfdbI.o: mfdbI.h

tutorial: furTut.sh count.awk checkTut.sh
furTut.sh: tutorial.org
	bash ../scripts/org2nw tutorial.org | notangle -RfurTut.sh > furTut.sh
count.awk: tutorial.org
	bash ../scripts/org2nw tutorial.org | notangle -Rcount.awk > count.awk
checkTut.sh: tutorial.org
	bash ../scripts/org2nw tutorial.org | notangle -RcheckTut.sh > checkTut.sh

aux: fur2prim.awk prim2fasta.awk checkPrim.awk
fur2prim.awk: fur2prim.org
	bash ../scripts/org2nw fur2prim.org | notangle -Rfur2prim.awk > ../build/fur2prim.awk
	chmod +x ../build/fur2prim.awk
prim2fasta.awk: prim2fasta.org
	bash ../scripts/org2nw prim2fasta.org | notangle -Rprim2fasta.awk > ../build/prim2fasta.awk
	chmod +x ../build/prim2fasta.awk
checkPrim.awk: checkPrim.org
	bash ../scripts/org2nw checkPrim.org | notangle -RcheckPrim.awk > ../build/checkPrim.awk
	chmod +x ../build/checkPrim.awk

clean:
	rm -f *.o *~ count.awk checkPrim.awk furTur.sh checkTut.sh rm makeFurDb fur.c makeFurDb.c
test:
	bash ../scripts/fur.sh 
