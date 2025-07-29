module github.com/evolbioinf/fur

go 1.22.5

toolchain go1.23.11

require (
	github.com/evolbioinf/clio v0.0.0-20240827074707-cb9ff755a85b
	github.com/evolbioinf/esa v0.0.0-20230428092833-66d4eac05d77
	github.com/evolbioinf/fasta v0.0.0-20230419094527-219cc47d94b2
	github.com/evolbioinf/sus v0.0.0-20230123102713-cc3fd6887965
)

replace github.com/evolbioinf/fur/util => ../util
