TAG     := $(shell git describe | sed 's/-.*//')
DATE    := $(shell git log -1 --format=%ai $(TAG) | sed 's/-/ /g' | awk '{printf "\\\\newdate{tagDate}{%s}{%s}{%s}\n", $$3, $$2, $$1}')
VERSION := $(shell git describe | sed 's/^[vV]//; s/-.*//')
PROG    := fur
all:
	printf "%s\n" $(VERSION) > version.tex
	printf "%s\n" $(DATE)    > date.tex
	latex  $(PROG)
	# bibtex $(PROG)
	# latex  $(PROG)
	latex  $(PROG)
	dvipdf $(PROG)
clean:	
	rm -f $(PROG).pdf *.log *.bbl *.blg *.dvi
