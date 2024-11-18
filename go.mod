module github.com/evolbioinf/fur

go 1.23.0

require (
	github.com/evolbioinf/clio v0.0.0-20240827074707-cb9ff755a85b
	github.com/evolbioinf/esa v0.0.0-20240208112648-445905ef2b6d
	github.com/evolbioinf/fasta v0.0.0-20230419094527-219cc47d94b2
	github.com/evolbioinf/sus v0.0.0-20230320163303-b6d16dd4ec1f
	github.com/ivantsers/chr v0.0.0-20241118144551-5f2b65c51950
)

require github.com/ivantsers/fastautils v0.0.0-20241118142913-f5f2f9b175e5 // indirect

replace github.com/evolbioinf/fur/util => ../util
