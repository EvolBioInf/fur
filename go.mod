module github.com/evolbioinf/fur

go 1.23.0

toolchain go1.23.11

require (
	github.com/evolbioinf/clio v0.0.0-20240827074707-cb9ff755a85b
	github.com/evolbioinf/esa v0.0.0-20240208112648-445905ef2b6d
	github.com/evolbioinf/fasta v0.0.0-20230419094527-219cc47d94b2
	github.com/evolbioinf/sus v0.0.0-20230320163303-b6d16dd4ec1f
)

require (
	github.com/ivantsers/chr v0.0.0-20251113132902-d017496cc4c3 // indirect
	github.com/ivantsers/fastautils v0.0.0-20241118142913-f5f2f9b175e5 // indirect
)

replace github.com/evolbioinf/fur/util => ../util
