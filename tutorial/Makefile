exe = tut
nw = $(shell which noweb)

all : $(exe).sh

$(exe).sh: tutorial.org
	if [ "$(nw)" != "" ]; then\
		bash ../scripts/org2nw tutorial.org | notangle -R$(exe) > $(exe).sh;\
	fi
test: $(exe).sh
	bash $(exe).sh
clean:
	rm $(exe).sh

