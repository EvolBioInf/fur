packs = fur biobox stan
all:
	for pack in $(packs); do \
		test -d $$pack || git clone git@github.com:EvolBioInf/$$pack; \
		cd $$pack && git pull && make && cd ../; \
	done
	cp fur/bin/cleanSeq fur/bin/fur fur/bin/makeFurDb bin
	cp biobox/bin/cres biobox/bin/cutSeq biobox/bin/midRoot biobox/bin/nj biobox/bin/plotLine biobox/bin/plotTree bin
	cp stan/bin/stan bin
	cd fur && make doc && cp doc/furDoc.pdf ../ && cd ../
	sudo docker build -t fox .
clean:
	rm -rf bin/* $(packs) furDoc.pdf 

