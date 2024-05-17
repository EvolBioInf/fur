date = $(shell git log | grep Date | head -n 1 | sed -r 's/Date: +[A-Z][a-z]+ ([A-Z][a-z]+) ([0-9]+) [^ ]+ ([0-9]+) .+/\2_\1_\3/')
version = $(shell git describe)

exe = cleanSeq
vf = -X github.com/evolbioinf/fur/util.version=$(version)
df = -X github.com/evolbioinf/fur/util.date=$(date)

all: $(exe)

$(exe):
	go build -ldflags "$(vf) $(df)" $(exe).go
tangle: $(exe).go
$(exe).go: $(exe).org
	bash ../scripts/org2nw $(exe).org | notangle -R$(exe).go > $(exe).go
build: tangle
	go build -ldflags "$(vf) $(df)" $(exe).go
clean:
	rm $(exe).go $(exe)
test:
	@echo "The program $(exe) is not tested."
