packs = util
progs = checkPrim cleanSeq fur fur2prim madis makeFurDb prim2fasta
all:
	test -d bin || mkdir bin
	for pack in $(packs); do \
		make -C $$pack; \
	done
	for prog in $(progs); do \
		make -C $$prog; \
		cp $$prog/$$prog bin; \
	done
test:
	test -d bin || mkdir bin
	for pack in $(packs); do \
		make -C $$pack; \
	done
	for prog in $(progs) $(packs); do \
		make test -C $$prog; \
	done
.PHONY: doc test docker
doc:
	make -C doc
docker:
	make -C docker
clean:
	for prog in $(progs) $(packs) doc; do \
		make clean -C $$prog; \
	done
