TAG     := $(shell git describe | sed 's/-.*//')
DATE    := $(shell git log -1 --format=%ai $(TAG) | sed 's/-/ /g' | awk '{printf "\\\\newdate{tagDate}{%s}{%s}{%s}\n", $$3, $$2, $$1}')
VERSION := $(shell git describe | sed 's/^[vV]//; s/-.*//')
NAME    = fur

all: $(NAME).tex text.tex header.tex ref.bib
	printf "%s\n" $(VERSION) > version.tex
	printf "%s\n" $(DATE)    > date.tex
	latex $(NAME)
	bibtex $(NAME)
	latex $(NAME)
	latex $(NAME)
	dvipdf $(NAME)
text.tex: ../src/$(NAME).org
	bash ../scripts/org2nw ../src/$(NAME).org | noweave -n -x > text.tex
clean:
	rm -f text.tex $(NAME).pdf
