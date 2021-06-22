TAG     := $(shell git describe | sed 's/-.*//')
DATE    := $(shell git log -1 --format=%ai $(TAG) | sed 's/-/ /g' | awk '{printf "\\\\newdate{tagDate}{%s}{%s}{%s}\n", $$3, $$2, $$1}')
VERSION := $(shell git describe | sed 's/^[vV]//; s/-.*//')
NAME   = fur

all: makeFurDb.tex $(NAME).tex furIn.tex header.tex ref.bib cdf.tex senSpec.tex fur2prim.tex checkPrim.tex prim2fasta.tex tutorial.tex cleanSeq.tex
	printf "%s\n" $(VERSION) > version.tex
	printf "%s\n" $(DATE)    > date.tex
	latex $(NAME)
	bibtex $(NAME)
	latex $(NAME)
	latex $(NAME)
	dvips -o -q $(NAME)
	ps2pdf $(NAME).ps
makeFurDb.tex: ../src/makeFurDb.org
	bash ../scripts/org2nw ../src/makeFurDb.org | noweave -n -x > makeFurDb.tex
senSpec.tex: ../src/senSpec.org
	bash ../scripts/org2nw ../src/senSpec.org | noweave -n -x > senSpec.tex
furIn.tex: ../src/$(NAME).org 
	bash ../scripts/org2nw ../src/$(NAME).org | noweave -n -x > furIn.tex
cdf.tex: fig.gp cdf.dat
	gnuplot fig.gp
checkPrim.tex: ../src/checkPrim.org
	bash ../scripts/org2nw ../src/checkPrim.org | noweave -n -x > checkPrim.tex
fur2prim.tex: ../src/fur2prim.org
	bash ../scripts/org2nw ../src/fur2prim.org | noweave -n -x > fur2prim.tex
prim2fasta.tex: ../src/prim2fasta.org
	bash ../scripts/org2nw ../src/prim2fasta.org | noweave -n -x > prim2fasta.tex
tutorial.tex: ../src/tutorial.org
	bash ../scripts/org2nw ../src/tutorial.org | noweave -n -x > tutorial.tex
cleanSeq.tex: ../src/cleanSeq.org
	bash ../scripts/org2nw ../src/cleanSeq.org | noweave -n -x > cleanSeq.tex
clean:
	rm -f text.tex $(NAME).pdf $(NAME).ps makeFurDb.tex furIn.tex cleanSeq.tex senSpec.tex checkPrim.tex fur2prim.tex prim2fasta.tex tutorial.tex
