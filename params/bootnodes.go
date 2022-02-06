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

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	"enode://89e878dc69f09fda0878f4f69e7b379ae39f1d5ee24c6381e80d2cace9bb7db16ade620852813134c25c83808c1732f67fd2c7d33ade531a929b1e895269af53@178.128.108.236:32668",
	"enode://2b9ea88df4b75fd1d6e202a798887be4c0481ba495abeea0e5fdfbf3576c4f89f4e7678adf5f5485c78e6e3ab2cf6b79d6e9c9df52259180233538958d077179@159.223.66.37:32668",
	"enode://f67b5a26c429d7dd1e42e0d6e637e1303e1131e56a5edf3563fa1e7d49a2c102e00bf74dda14e5c21dd470494507eee67c23c48961bde354144056410f06fa27@149.28.144.79:32668",
	"enode://a02542f319b3aa9d21dc38c2a5cad0f5a15d1cddca0fd474672d3297cce6296cf2f3360723eba9d9de085463db9378114bc55d8283461aa72a326bec5045b013@149.28.151.228:32668",
	"enode://79232e5f601b31c00300550315ebed0f1c31e384010b42a84e6eddf5bf8b62daafcb5b21fcd04a6c8df4b59d019d50c4558958c1d800fa5183fe229f45481b18@45.76.147.205:32668",
	"enode://f6057385c20a071e7fbe784d1187916eb5e05a81e344a6459d95d821f41de77ba6a75e9579aba094dec10c69e5efb804585ed2ed035bd8b5ec878574cc4e3684@45.77.241.232:32668",
	"enode://2f7b3065880a1f56f2cd5d355707bd14153cf312c607253d3fc62980d5f5766818ac60958c435ecd93b60b63c956580891d73c4d984abef85da7a69b0a7a0a2b@139.180.208.140:32668",
	"enode://00271c24d7516312670eacb9ec2e521a0755db559510702fddcce6417f646b8429ede262bed9a0c01a45f3973d96d480560e66b4e511523f11dfeff939c5011d@149.28.150.244:32668",
	"enode://1d56970ad4ae2bd8efcd560e43f59f10aa6d65c5489bcaa9748c702148b3e97f5a2e4e2cc7c443f0e5809999340c93ceba628bcf9434e04ff9aeb1a9ad5bf1f6@45.77.43.226:32668",
	"enode://6ea0bb52fe1b6afb6c50c4ba4ece0dd4f1f96af17963b5f1d3a80d6ae115b6f5c211651e9a7eaa7c8b863e51075ca8281857e96ca6cc5ef1ba98d19f14a595ff@139.180.156.86:32668",
	"enode://7ab047b5645add9bb558b9eadad06e15787be056c32fb87f0995bf6402d3a3673c45b0932356f3466b110b853e7ed6c545b14e19840a51888a337358d19be71a@45.77.32.219:32668",
}
var TestnetBootnodes = []string{}

// KnownDNSNetwork returns the address of a public DNS-based node list for the given
// genesis hash and protocol. See https://github.com/ethereum/discv4-dns-lists for more
// information.
func KnownDNSNetwork(genesis common.Hash, protocol string) string {
	return ""
}
