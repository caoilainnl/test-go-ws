module example.com/test-go-ws

go 1.24.4

require github.com/ccxt/ccxt/go/v4 v4.4.92
require github.com/ccxt/ccxt/go/v4/ws v0.0.0-00010101000000-000000000000

replace github.com/ccxt/ccxt/go/v4 => ../ccxt/go/v4
replace github.com/ccxt/ccxt/go/v4/ws => ../ccxt/go/v4/ws

require (
	github.com/bits-and-blooms/bitset v1.13.0 // indirect
	github.com/consensys/bavard v0.1.13 // indirect
	github.com/consensys/gnark-crypto v0.12.1 // indirect
	github.com/crate-crypto/go-ipa v0.0.0-20240223125850-b1e8a79f509c // indirect
	github.com/crate-crypto/go-kzg-4844 v1.0.0 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.0.1 // indirect
	github.com/ethereum/c-kzg-4844 v1.0.0 // indirect
	github.com/ethereum/go-ethereum v1.14.13 // indirect
	github.com/ethereum/go-verkle v0.1.1-0.20240829091221-dffa7562dbe9 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/holiman/uint256 v1.3.1 // indirect
	github.com/mitchellh/mapstructure v1.4.1 // indirect
	github.com/mmcloughlin/addchain v0.4.0 // indirect
	github.com/supranational/blst v0.3.13 // indirect
	github.com/vmihailenco/msgpack/v5 v5.4.1 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	golang.org/x/crypto v0.39.0 // indirect
	golang.org/x/net v0.41.0 // indirect
	golang.org/x/sync v0.15.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	rsc.io/tmplfunc v0.0.3 // indirect
)
