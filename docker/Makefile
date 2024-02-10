path = /home/haubold/Programming/Go/src/github.com/evolbioinf
all: bin/cleanSeq bin/cres bin/cutSeq bin/fur bin/makeFurDb bin/midRoot bin/nj bin/plotLine bin/plotTree bin/stan furDoc.pdf
	sudo docker build -t fox .
bin/cleanSeq: $(path)/fur/bin/cleanSeq
	cp $(path)/fur/bin/cleanSeq bin
bin/cres: $(path)/biobox/bin/cres
	cp $(path)/biobox/bin/cres bin
bin/cutSeq: $(path)/biobox/bin/cutSeq
	cp $(path)/biobox/bin/cutSeq bin
bin/fur: $(path)/fur/bin/fur
	cp $(path)/fur/bin/fur bin
bin/makeFurDb: $(path)/fur/bin/makeFurDb
	cp $(path)/fur/bin/makeFurDb bin
bin/midRoot: $(path)/biobox/bin/midRoot
	cp $(path)/biobox/bin/midRoot bin
bin/nj: $(path)/biobox/bin/nj
	cp $(path)/biobox/bin/nj bin
bin/plotLine: $(path)/biobox/bin/plotLine
	cp $(path)/biobox/bin/plotLine bin
bin/plotTree: $(path)/biobox/bin/plotTree
	cp $(path)/biobox/bin/plotTree bin
bin/stan: $(path)/stan/bin/stan
	cp $(path)/stan/bin/stan bin
furDoc.pdf: $(path)/fur/doc/furDoc.pdf
	cd ../doc && make && cd - && cp ../doc/furDoc.pdf .

