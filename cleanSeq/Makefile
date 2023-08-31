version = $(shell git describe)
date = $(shell git log | grep Date | head -n 1 | sed -r 's/Date: +[A-Z][a-z]+ ([A-Z][a-z]+) ([0-9]+) [^ ]+ ([0-9]+) .+/\2_\1_\3/')

exe = cleanSeq
all: $(exe)
$(exe): $(exe).org
	bash ../scripts/org2nw $(exe).org | notangle -R$(exe).go > $(exe).go
	go build -ldflags "-X github.com/evolbioinf/fur/util.version=$(version) -X github.com/evolbioinf/fur/util.date=$(date)" $(exe).go
clean:
	rm $(exe)
test:
	@echo "The program $(exe) is not tested."
