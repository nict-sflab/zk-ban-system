module github.com/akakou/zk-ban-system/dump

go 1.24.3

replace github.com/akakou/zk-ban-system => ../

replace github.com/akakou/zk-ban => ../../zk-ban

replace github.com/akakou/gnark-precomputes => ../../gnark-precomputes

replace github.com/akakou/gnark-serializable => ../../gnark-serializable

require (
	github.com/akakou/snark-utils v0.0.2
	github.com/akakou/zk-ban v0.0.0-00010101000000-000000000000
	github.com/akakou/zk-ban-system v0.0.0-00010101000000-000000000000
	github.com/consensys/gnark v0.13.0
)

require (
	github.com/akakou/gnark-precomputes v0.0.0-00010101000000-000000000000 // indirect
	github.com/akakou/gnark-serializable v0.0.0-00010101000000-000000000000 // indirect
	github.com/bits-and-blooms/bitset v1.22.0 // indirect
	github.com/blang/semver/v4 v4.0.0 // indirect
	github.com/consensys/gnark-crypto v0.18.0 // indirect
	github.com/fxamacker/cbor/v2 v2.8.0 // indirect
	github.com/google/pprof v0.0.0-20250607225305-033d6d78b36a // indirect
	github.com/iden3/go-iden3-crypto v0.0.17 // indirect
	github.com/ingonyama-zk/icicle-gnark/v3 v3.2.2 // indirect
	github.com/liyue201/gnark-circomlib v0.0.0-20241024021655-892bf7c71a20 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/ronanh/intcomp v1.1.1 // indirect
	github.com/rs/zerolog v1.34.0 // indirect
	github.com/x448/float16 v0.8.4 // indirect
	golang.org/x/crypto v0.39.0 // indirect
	golang.org/x/sync v0.15.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
)
