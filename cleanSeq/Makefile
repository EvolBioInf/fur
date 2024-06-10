date = $(shell bash ../scripts/getDate.sh)
version = $(shell bash ../scripts/getVersion.sh)

exe = cleanSeq
vf = -X github.com/evolbioinf/fur/util.version=$(version)
df = -X github.com/evolbioinf/fur/util.date=$(date)
nw = $(shell which noweb)

all: $(exe)

$(exe): $(exe).go
	go build -ldflags "$(vf) $(df)" $(exe).go
tangle: $(exe).go
$(exe).go: $(exe).org
	if [ $(nw) != "" ]; then\
		bash ../scripts/org2nw $(exe).org | notangle -R$(exe).go > $(exe).go;\
	fi
clean:
	rm $(exe).go $(exe)
test:
	@echo "The program $(exe) is not tested."
