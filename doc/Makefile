SRC1 =   checkPrim.tex cleanSeq.tex fur.tex fur2prim.tex madis.tex makeFurDb.tex prim2fasta.tex tutorial.tex util.tex
SRC2 =   header.tex intro.tex furDoc.tex

all: furDoc.pdf
furDoc.pdf: furDoc.tex $(SRC1) $(SRC2) version.txt date.txt
	latex furDoc.tex
	bibtex furDoc
	latex furDoc.tex
	latex furDoc.tex
	dvips furDoc -o -q
	ps2pdf -dALLOWPSTRANSPARENCY furDoc.ps
%.tex: ../*/%.org furDoc.tex
	bash ../scripts/org2nw $< | awk -f ../scripts/preWeave.awk | noweave -n -x | sed 's/_test/\\_test/g' > $@
version.txt: ../data/version.txt
	cat ../data/version.txt | tr '-' ' ' | awk '{printf "%s-%s\n", $$1, $$2}' > version.txt
date.txt: ../data/date.txt
	cat ../data/date.txt | tr '_' ' ' > date.txt
clean:
	rm -f furDoc.pdf furDoc.ps furDoc.dvi furDoc.aux furDoc.bbl  furDoc.blg  furDoc.dvi  furDoc.log $(SRC1) version.txt date.txt
