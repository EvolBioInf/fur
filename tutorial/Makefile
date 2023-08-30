exe = tut
all : $(exe).sh

$(exe).sh: tutorial.org
	bash ../scripts/org2nw tutorial.org | notangle -R$(exe) > $(exe).sh
test: $(exe).sh
	bash $(exe).sh
clean:
	rm $(exe).sh

