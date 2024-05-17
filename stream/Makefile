date = $(shell git log | grep Date | head -n 1 | sed -r 's/Date: +[A-Z][a-z]+ ([A-Z][a-z]+) ([0-9]+) [^ ]+ ([0-9]+) .+/\2_\1_\3/')
version = $(shell git describe)

exe = stream
vf = -X github.com/evolbioinf/fur/util.version=$(version)
df = -X github.com/evolbioinf/fur/util.date=$(date)

all: $(exe) time_1.ps time_0.5.ps time_0.25.ps time_0.ps memory.ps
$(exe):
	go build -ldflags "$(vf) $(df)" $(exe).go
tangle: $(exe).go $(exe)_test.go
build: tangle
	go build -ldflags "$(vf) $(df)" $(exe).go
$(exe).go: $(exe).org
	awk -f ../scripts/preTangle.awk $(exe).org | bash ../scripts/org2nw | notangle -R$(exe).go > $(exe).go
$(exe)_test.go: $(exe).org
	awk -f ../scripts/preTangle.awk $(exe).org | bash ../scripts/org2nw | notangle -R$(exe)_test.go | gofmt > $(exe)_test.go
time_1.ps: time_1.dat
	plotLine -L -x "Sequence Length (Mb)" -y "Time (s)" -d "6cm,5cm" -p time_1.ps time_1.dat
time_0.5.ps: time_0.5.dat
	plotLine -L -x "Sequence Length (Mb)" -y "Time (s)" -d "6cm,5cm" -p time_0.5.ps time_0.5.dat
time_0.25.ps: time_0.25.dat
	plotLine -L -x "Sequence Length (Mb)" -y "Time (s)" -d "6cm,5cm" -p time_0.25.ps time_0.25.dat
time_0.ps: time_0.dat
	plotLine -L -x "Sequence Length (Mb)" -y "Time (s)" -d "6cm,5cm" -p time_0.ps time_0.dat
memory.ps: memory.dat
	plotLine -L -x "Sequence Length (Mb)" -y "Memory (GB)" -d "6cm,5cm" -p memory.ps memory.dat
time_1.dat: sim.dat
	awk '$$6==1' sim.dat | awk '{print $$1, $$2, $$5}' >  time_1.dat
time_0.5.dat: sim.dat
	awk '$$6==0.5' sim.dat | awk '{print $$1, $$2, $$5}' >  time_0.5.dat
time_0.25.dat: sim.dat
	awk '$$6==0.25' sim.dat | awk '{print $$1, $$2, $$5}' >  time_0.25.dat
time_0.dat: sim.dat
	awk '$$6==0' sim.dat | awk '{print $$1, $$2, $$5}' >  time_0.dat
memory.dat: sim.dat
	awk '$$6==1' sim.dat | awk '{print $$1, $$4/1e6, $$5}' >  memory.dat
test: $(exe)
	@go test
clean:
	rm -f $(exe) *.go time_*.dat memory.dat

