CC=gcc
export VERSION = $(shell bash ../scripts/version.sh) # The single, trailing blank is essential
export DATE    = $(shell bash ../scripts/date.sh) #    The single, trailing blank is essential
CFLAGS         = -DVERSION=\"$(VERSION)\" -DDATE=\"$(DATE)\" -O3 -Wall -Wshadow -pedantic -std=gnu99 -g -I../common/ -L../common/ #-pg

SRCFILES  = fur.c interface.c
OBJFILES  = fur.o interface.o
LIBS      = -lcommon -lm -lbsd -lgsl -lblas
EXECFILE  = fur

.PHONY : all
all : $(EXECFILE)
$(EXECFILE) : $(OBJFILES)
	$(CC) $(CFLAGS) -o $(EXECFILE) $(OBJFILES) $(LIBS)
interface.o: interface.h

fur.c: fur.org
	bash ../scripts/org2nw fur.org | notangle -Rfur.c > fur.c

clean:
	rm -f *.o *~
test:
	bash ../scripts/fur.sh 
eco105:
	wget guanine.evolbio.mpg.de/fur/eco105.tar.gz
	tar -xzf eco105.tar.gz
	fur -t Focus -n Neighborhood
	rm -r Focus Neighborhood eco105.tar.gz
