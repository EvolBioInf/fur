packs = util
progs = cleanSeq fur madis makeFurDb stream
all:
	test -d bin || mkdir bin
	for pack in $(packs); do \
		make -C $$pack; \
	done
	for prog in $(progs); do \
		make -C $$prog; \
		cp $$prog/$$prog bin; \
	done
tangle:
	test -d bin || mkdir bin
	for pack in $(packs); do \
		make tangle -C $$pack; \
	done
	for prog in $(progs); do \
		make tangle -C $$prog; \
	done
test:
	test -d bin || mkdir bin
	for pack in $(packs); do \
		make -C $$pack; \
	done
	for prog in $(progs) $(packs); do \
		make test -C $$prog; \
	done
.PHONY: weave test docker
docker:
	make -C docker
clean:
	for prog in $(progs) $(packs) doc; do \
		make clean -C $$prog; \
	done
	rm -f bin/*
