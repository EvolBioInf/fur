module github.com/evolbioinf/fur

go 1.23.0

toolchain go1.23.11

require (
	github.com/evolbioinf/clio v0.0.0-20250730155633-f17ebc5319c4
	github.com/evolbioinf/esa v0.0.0-20260221181549-4ec39e983a94
	github.com/evolbioinf/fasta v0.0.0-20251121105511-f74cf90e08b9
	github.com/evolbioinf/sus v0.0.0-20230320163303-b6d16dd4ec1f
	github.com/ivantsers/chr v0.0.0-20260305132247-73e9872f7573
)

require github.com/ivantsers/fastautils v0.0.0-20260305131012-2ab126fbf7b5 // indirect

replace github.com/evolbioinf/fur/util => ../util
