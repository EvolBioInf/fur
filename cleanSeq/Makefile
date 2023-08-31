version = $(shell cat ../data/version.txt)
date = $(shell cat ../data/date.txt)

exe = cleanSeq
all: $(exe)
$(exe): $(exe).org
	bash ../scripts/org2nw $(exe).org | notangle -R$(exe).go > $(exe).go
	go build -ldflags "-X github.com/evolbioinf/fur/util.version=$(version) -X github.com/evolbioinf/fur/util.date=$(date)" $(exe).go
clean:
	rm $(exe)
test:
	@echo "The program $(exe) is not tested."
