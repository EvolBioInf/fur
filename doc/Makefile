SRC1 =   checkPrim.tex cleanSeq.tex fur.tex fur2prim.tex madis.tex makeFurDb.tex prim2fasta.tex tutorial.tex util.tex
SRC2 =   header.tex intro.tex furDoc.tex

all: furDoc.pdf
furDoc.pdf: furDoc.tex $(SRC1) $(SRC2)
	git log | grep Date | head -n 1 | awk '{print $$3, $$4, $$6}' > date.tex
	latex furDoc.tex
	bibtex furDoc
	latex furDoc.tex
	latex furDoc.tex
	dvips furDoc -o -q
	ps2pdf -dALLOWPSTRANSPARENCY furDoc.ps
%.tex: ../*/%.org furDoc.tex
	bash ../scripts/org2nw $< | awk -f ../scripts/preWeave.awk | noweave -n -x | sed 's/_test/\\_test/g' > $@
clean:
	rm -f furDoc.pdf furDoc.ps furDoc.dvi furDoc.aux furDoc.bbl  furDoc.blg  furDoc.dvi  furDoc.log $(SRC1)
