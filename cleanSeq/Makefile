date = $(shell bash ../scripts/getDate.sh)
version = $(shell bash ../scripts/getVersion.sh)

exe = cleanSeq
vf = -X github.com/evolbioinf/fur/util.version=$(version)
df = -X github.com/evolbioinf/fur/util.date=$(date)

all: $(exe)

$(exe):
	go build -ldflags "$(vf) $(df)" $(exe).go
tangle: $(exe).go
$(exe).go: $(exe).org
	bash ../scripts/org2nw $(exe).org | notangle -R$(exe).go > $(exe).go
clean:
	rm $(exe).go $(exe)
test:
	@echo "The program $(exe) is not tested."
