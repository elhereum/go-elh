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
//"enr:-KO4QM6oq6g3YstgN9jZtEv6Qy9XihDfixRw7t4slnsSuOHYRs53D3KGcTe0uEfkPjqD73n7WrGwjFriWuu_fcx_vUGGAYf6dFh3g2V0aMfGhMHpFjOAgmlkgnY0gmlwhCW7TCWJc2VjcDI1NmsxoQNK26z_KuxybXxBnj4X4S6VIOUjZDjvcMcpORq-66e1kYRzbmFwwIN0Y3CCmd6DdWRwgpne",
//"enr:-KO4QPxTB0KISMOZ9pxuJQksd-wuz0br7GBPmGwAx5e3U76fb3uWIUVA4WLJgoN59xbMMXFWSJlQTKP71fIJupAwkeCGAYf6dt2Pg2V0aMfGhMHpFjOAgmlkgnY0gmlwhLyl1UWJc2VjcDI1NmsxoQJZm7S5-_ViFpXPKeRHlZZYegFruzMpzqJeqHr_BEfvIIRzbmFwwIN0Y3CCmd6DdWRwgpne",
//"enr:-KO4QC9Nn1vgKhUw9daUV1AkQom5Ey2CnEkPxQvMAyL4t14-fyxYRIi4DtNpDWsJlbckkFocmEOxLl2Fom_JbGfj5MKGAYf6dqQWg2V0aMfGhMHpFjOAgmlkgnY0gmlwhCW7g2mJc2VjcDI1NmsxoQJufI7hIwxrEJQD92XI4AZsZ7wjPPWw-p4xCSLI1IpPrYRzbmFwwIN0Y3CCmd6DdWRwgpne",
//"enr:-J24QHYZUEzTXdfHdxRwrd3UWqG7x0DvZ8gTHbbyEQklQIuDIzKQjGxOxGE5e3SKaYoh3JqmkI2Pdf6_1zWyRY7uAA6GAYgEDb3jg2V0aMfGhMHpFjOAgmlkgnY0gmlwhDPDKMmJc2VjcDI1NmsxoQLfSjOW_UdgMfGOtXERRN_khuac5Ai-w_OqW1aJSzlTh4N0Y3CCmd6DdWRwgpne",
}
var AgaMainnetBootnodes = []string{
"enode://fd72e1ce0310362c9f0df0dfee0b1c27de8b907ac3a79a3d8020d9c2da7a9f0cce5d2dd29e3fff1e7a6adb119b6ef50a30d4f87720adcdd48be511ba844e2a3a@31.220.84.74:27390",
"enode://75ada21d94664ca5737166e17c6d8ba82f18330f293449fbc0c6ebd22dd2cfa3cde7eca5971b9a93680a056927fb671749a4b70131665218e6231ec834bdd999@188.165.213.69:39390",
"enode://55e7f3666f6c697f47cb7be968d954edc26565ba59b73f4d79a918c3c218f9b23aff728f6b956e0f43c9f632a7568488248fd17138cdbd4b75b57b1342fe2b53@193.108.113.23:39390",
"enode://7fd9cc4b1455ba0c7faaa2ee603421b52e0abc78c814b897ccea0fba5aaa41da47c32fcd1f406a577130c29ce31f000689ffd809400863ff5412e427495d6a42@107.174.115.73:56134",
"enode://a30b4753bb20bb9045520f8b0911267ceaa7e32ad63411f09d241e89c4542edc055270fe8d7d66f7d245352587e598687e3934f457ad79ee1ac6fcefe2c761d7@107.174.115.67:37670",
"enode://61b87a31fc6e3f6860bbe9342f3b663898ff55e2f3a4b8a142e01a7a54f579ccc53ecd6d4692940f40fd8b8234f5f1048d102c9214e20898fbcbfaf888e7226f@60.247.148.7:38450",
"enode://5281f63d2ecee35953c78d247951e6b26c5a6153b2d7a73eedc967fc8b483601a5373c2fdfa8fe451acd59df58dc442ab2ceb0bcb86971554dffc98ee27b4cc2@79.157.40.102:50946",
"enode://ccaf083d82ce66637b4f1603230805af00edf6b708191529a3ebc77ffbb3fa2310358ef2a3b51cd74d219748ab0912aa09bc05ce8bd8280531d25278c0931d72@50.21.227.209:58564",
"enode://a2f853b93444197f5f0e71219fb66f14ef531be9d8c3046a7e47d3512aa494689d8c9ea27bfbfa5bd6238f8b6cd03b8e6a61a2f73c90334d4336b625be7efff5@96.126.96.29:55262",
"enode://8a3d6dca942b9ae839129451e29577eee6cf40751045b64956f0b41c0f54150004bb752eaf58e16d5b63dd8541e1e1188f99aa909f6d8e152272089a76807380@181.45.56.131:33468",
"enode://ee9373f0ff08121f91145d49618d3dd288a6dfdec8c531d97e15357fad8220d51eaa480af7725d66b4f4f2df1e566f08760662b633551d19fd5c144220c1c551@191.83.169.44:43668",
"enode://0f11fd43fd2d62b24c1e47c340eadb21feddbdae6256a1fc1f62e13d98dff63c7daf7c6c350b5f774f691877e877e0e96e903e00afd3e23e2da336795ca96451@186.122.96.97:45462",
"enode://dce0329372035a837334407c6274dcd6544c2573d5f13324f5ec5182b193e2984855df5b7099721593c6088d2e3299ff2944b63cb340a631f8bf85a484c6bd6d@172.245.226.72:50230",
"enode://3aa8852457a6c10725105cc250325a97f014e491a627b1f96fc21ce062e790c1a611dc81053250c9baf1ed1c7e05b17d0394e802c1f6c54b57d42a3c947c1e5f@173.82.248.116:39390",
"enode://aba317e02d4bd36bac4114144688bdd15df90e3f8fdc4b9df6d1234793575bb99a0c99333c07c857dda72701119925d67ea026a7fe0b12dfa800e2ce459cf750@31.42.10.1:39180",
"enode://d9c4e4a35bc7be8434dd4f749f7345bf529dcd88b327311e5478cfa9152e0ef5f97990dd087174a3a7c82a4d957aa64c8cced04984dc6abc8d88df83f5aa6230@198.148.123.132:48982",
"enode://9f782d0913693d6462ec3f402e46092764c869c8fac3d8499166ca3d89995c1afeb7a465524b8b426896ad804a2740cd9ab68568ed0c19908f498c839b0bfa49@51.83.212.194:49048",
"enode://f92811fd09231fa91148a780e027e9b3adf5ca5e10b8072f8ca633cfb12f10d529e4c2094e044892ec30f4430b4454177e6e13e01c6595e94206bce50229f1f1@107.174.156.114:39390",
"enode://df0c64419667b867394cf8762592f44142e0b020d6248f86cd141669331c07dca77a94d20e57f9b8489616f1267887e501ae9ea5851f2640a9d6d5f7b37079d3@73.132.67.157:36342",
"enode://f2f109274e34e0556e3a9138ab2a03d22d397a3a2f27d87f381814013e79b8e9bbc778f7c40f0f4f12fed6540e1ef166f76c76cef06dbbaecb437bbad7bc55af@109.108.71.107:33033",
"enode://f86e8c28508f864f59b3c92308940cff5e37fa7a40f311bd63728eb4102bfa0ebd49462587a43c3d80ad4e99253322952f5934f7d8b7e2c8ab95823f34c0f82f@89.33.193.195:39390",
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
