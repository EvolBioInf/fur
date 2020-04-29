all:
	make -C common
	make -C src
	make -C src tutorial
	mkdir -p build
	cp src/fur build
	cp src/checkTut.sh src/furTut.sh scripts
test:
	@make -s -C src test
eco105:
	make -C src eco105
clean:
	make -C common clean
	make -C src    clean
	make -C doc    clean
.PHONY:	doc
doc:	
	make -C doc
