// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

import "github.com/ethereum/go-ethereum/common"

var MainnetBootnodes = []string{
}
var RopstenBootnodes = []string{
}
var SepoliaBootnodes = []string{
}
var RinkebyBootnodes = []string{
}
var GoerliBootnodes = []string{
}
var KilnBootnodes = []string{
}
var V5Bootnodes = []string{
}
var AgaMainnetBootnodes = []string{
"enode://4adbacff2aec726d7c419e3e17e12e9520e5236438ef70c729391abeeba7b591c9edaf6e380cbbf0857a44b7b9a328402b7719ceb2c1ec0aaf569fd39ec9026d@37.187.76.37:39390",
"enode://599bb4b9fbf5621695cf29e4479596587a016bbb3329cea25ea87aff0447ef20a3ca20a8c60c2325d8a67a08ccfa5b1d50f35cb46d955093991fd08e57677f6e@188.165.213.69:39390",
"enode://6e7c8ee1230c6b109403f765c8e0066c67bc233cf5b0fa9e310922c8d48a4fad59275af292272a6429fedface52eba99db2990cdc7f8e2e0009e51e2ff07694e@37.187.131.105:39390",
}
var AgaTestnetBootnodes = []string{
}

const dnsPrefix = "enrtree://AKA3AM6LPBYEUDMVNU3BSVQJ5AD45Y7YPOHJLEF6W26QOE4VTUDPE@"

// KnownDNSNetwork returns the address of a public DNS-based node list for the given
// genesis hash and protocol. See https://github.com/ethereum/discv4-dns-lists for more
// information.
func KnownDNSNetwork(genesis common.Hash, protocol string) string {
	var net string
	switch genesis {
	case MainnetGenesisHash:
		net = "mainnet"
	case RopstenGenesisHash:
		net = "ropsten"
	case RinkebyGenesisHash:
		net = "rinkeby"
	case GoerliGenesisHash:
		net = "goerli"
	default:
		return ""
	}
	return dnsPrefix + protocol + "." + net + ".ethdisco.net"
}
