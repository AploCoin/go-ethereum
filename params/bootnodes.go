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

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	// Ethereum Foundation Go Bootnodes
	//"enode://3f87316c80ee6ec015e21ce7e2cd5b1067d946a52e87ae83aacc03fa050379c1c3c5412fc31a7092e210636848e2d57db28453779aa1ce7da7113bcd80060adf@65.109.62.117:30303",
	//"enode://5c771ff0c5a0400edaeb12dfb9b405b35a3b6fcb1e780cc223eebb050563f67e122fbbbd0d617f89c2163f97062a9bc89b8ba4c03d039b2d09ba7ce8d2c3c979@35.82.33.142:30303",
	//"enode://8a66b37212a72f338e86c63cc05fa7a1faa954bd384a48f3656032239388fa889c9b59b01879078a44776ca0c4562137a929d3a78662cfc3de36af154937f81d@176.9.38.209:30303",
	//"enode://aae79ff5076aecaef5e1748bb50f02bb14703823200f184c7beb9b57b12ebe58416dfc96dd41a459fcd91c7931f9c05323b124be03ee2f5a19febaee0d15678a@65.109.62.118:30303",
	//"enode://873631a6313fe26e309b566fa16b8b12b9e37bf019180e7c18e982290cc3df489664d57480e52a0acb6f8c33322010379bcf11173bcfb018a91f06b475c4bac1@144.76.17.180:30303",
	//"enode://efe4155a02f678d8f85a66e9d0c6aac7e7f68e94735aaca4410f8f1ed5b04c8134e619b98ca2b81d4debae76bb36074c43bfcf278bfad1babb90218784865801@54.176.239.225:30303",
	//"enode://a2cb0517c88b71580ab46595c5dab4aca73d248404b7d1773bb089a1ea5a0cbd88b2c6a315ff3199a62c804b05783c568628933facf351f6e7f88eef527f51ea@54.177.117.107:30303",
	//"enode://8cf6f46d351c319c708e54eb693050ef8e717115c96073d6961e2ded71fc609f3b2f8ba42ec005ab554e4d4147393fe12c1bf7ce88b6353585897e9c758bb3a2@54.156.57.38:30303",
	//"enode://8b05b73b65e6307a09e0c7a92d2b8a83384939f81618783069de894a9ba8bde42ea5a88f52c02d54fc9ffa12b080fb5cf625412a2a63d0f9baa19dae9f7cce39@35.158.233.108:30303",
	//"enode://c6bdbf1d7eb319f6ea54df20a270fb6bc15d31f7b8b9dcaf918f8161277769038df7834d0dc94b250b1e92017b477fd2c19ce920c7a30b1dbb9eea3769108c61@3.38.46.69:30303",
	//"enode://dafb0b584ba443219fea0bd172e9d3457eabd8c913e7aed7ab9627577850355903b85c64611de056433458d89a4e39c5d48fb44f18181d650fedf43808368d70@44.230.118.175:30303",
}

// RopstenBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var RopstenBootnodes = []string{
	"enode://30b7ab30a01c124a6cceca36863ece12c4f5fa68e3ba9b0b51407ccc002eeed3b3102d20a88f1c1d3c3154e2449317b8ef95090e77b312d5cc39354f86d5d606@52.176.7.10:30303",    // US-Azure geth
	"enode://865a63255b3bb68023b6bffd5095118fcc13e79dcf014fe4e47e065c350c7cc72af2e53eff895f11ba1bbb6a2b33271c1116ee870f266618eadfc2e78aa7349c@52.176.100.77:30303",  // US-Azure parity
	"enode://6332792c4a00e3e4ee0926ed89e0d27ef985424d97b6a45bf0f23e51f0dcb5e66b875777506458aea7af6f9e4ffb69f43f3778ee73c81ed9d34c51c4b16b0b0f@52.232.243.152:30303", // Parity
	"enode://94c15d1b9e2fe7ce56e458b9a3b672ef11894ddedd0c6f247e0f1d3487f52b66208fb4aeb8179fce6e3a749ea93ed147c37976d67af557508d199d9594c35f09@192.81.208.223:30303", // @gpip
}

// SepoliaBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Sepolia test network.
var SepoliaBootnodes = []string{
	// geth
	"enode://9246d00bc8fd1742e5ad2428b80fc4dc45d786283e05ef6edbd9002cbc335d40998444732fbe921cb88e1d2c73d1b1de53bae6a2237996e9bfe14f871baf7066@18.168.182.86:30303",
	// besu
	"enode://ec66ddcf1a974950bd4c782789a7e04f8aa7110a72569b6e65fcd51e937e74eed303b1ea734e4d19cfaec9fbff9b6ee65bf31dcb50ba79acce9dd63a6aca61c7@52.14.151.177:30303",
}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{
	"enode://a24ac7c5484ef4ed0c5eb2d36620ba4e4aa13b8c84684e1b4aab0cebea2ae45cb4d375b77eab56516d34bfbd3c1a833fc51296ff084b770b94fb9028c4d25ccf@52.169.42.101:30303", // IE
	"enode://343149e4feefa15d882d9fe4ac7d88f885bd05ebb735e547f12e12080a9fa07c8014ca6fd7f373123488102fe5e34111f8509cf0b7de3f5b44339c9f25e87cb8@52.3.158.184:30303",  // INFURA
	"enode://b6b28890b006743680c52e64e0d16db57f28124885595fa03a562be1d2bf0f3a1da297d56b13da25fb992888fd556d4c1a27b1f39d531bde7de1921c90061cc6@159.89.28.211:30303", // AKASHA
}

// GoerliBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Görli test network.
var GoerliBootnodes = []string{
	// Upstream bootnodes
	"enode://011f758e6552d105183b1761c5e2dea0111bc20fd5f6422bc7f91e0fabbec9a6595caf6239b37feb773dddd3f87240d99d859431891e4a642cf2a0a9e6cbb98a@51.141.78.53:30303",
	"enode://176b9417f511d05b6b2cf3e34b756cf0a7096b3094572a8f6ef4cdcb9d1f9d00683bf0f83347eebdf3b81c3521c2332086d9592802230bf528eaf606a1d9677b@13.93.54.137:30303",
	"enode://46add44b9f13965f7b9875ac6b85f016f341012d84f975377573800a863526f4da19ae2c620ec73d11591fa9510e992ecc03ad0751f53cc02f7c7ed6d55c7291@94.237.54.114:30313",
	"enode://b5948a2d3e9d486c4d75bf32713221c2bd6cf86463302339299bd227dc2e276cd5a1c7ca4f43a0e9122fe9af884efed563bd2a1fd28661f3b5f5ad7bf1de5949@18.218.250.66:30303",

	// Ethereum Foundation bootnode
	"enode://a61215641fb8714a373c80edbfa0ea8878243193f57c96eeb44d0bc019ef295abd4e044fd619bfc4c59731a73fb79afe84e9ab6da0c743ceb479cbb6d263fa91@3.11.147.67:30303",

	// Goerli Initiative bootnodes
	"enode://d4f764a48ec2a8ecf883735776fdefe0a3949eb0ca476bd7bc8d0954a9defe8fea15ae5da7d40b5d2d59ce9524a99daedadf6da6283fca492cc80b53689fb3b3@46.4.99.122:32109",
	"enode://d2b720352e8216c9efc470091aa91ddafc53e222b32780f505c817ceef69e01d5b0b0797b69db254c586f493872352f5a022b4d8479a00fc92ec55f9ad46a27e@88.99.70.182:30303",
}

var KilnBootnodes = []string{
	"enode://c354db99124f0faf677ff0e75c3cbbd568b2febc186af664e0c51ac435609badedc67a18a63adb64dacc1780a28dcefebfc29b83fd1a3f4aa3c0eb161364cf94@164.92.130.5:30303",
	"enode://d41af1662434cad0a88fe3c7c92375ec5719f4516ab6d8cb9695e0e2e815382c767038e72c224e04040885157da47422f756c040a9072676c6e35c5b1a383cce@138.68.66.103:30303",
	"enode://91a745c3fb069f6b99cad10b75c463d527711b106b622756e9ef9f12d2631b6cb885f831d1c8731b9bc7177cae5e1ea1f1be087f86d7d30b590a91f22bc041b0@165.232.180.230:30303",
	"enode://b74bd2e8a9f0c53f0c93bcce80818f2f19439fd807af5c7fbc3efb10130c6ee08be8f3aaec7dc0a057ad7b2a809c8f34dc62431e9b6954b07a6548cc59867884@164.92.140.200:30303",
}

var V5Bootnodes = []string{
	// Teku team's bootnode
	"enr:-KG4QOtcP9X1FbIMOe17QNMKqDxCpm14jcX5tiOE4_TyMrFqbmhPZHK_ZPG2Gxb1GE2xdtodOfx9-cgvNtxnRyHEmC0ghGV0aDKQ9aX9QgAAAAD__________4JpZIJ2NIJpcIQDE8KdiXNlY3AyNTZrMaEDhpehBDbZjM_L9ek699Y7vhUJ-eAdMyQW_Fil522Y0fODdGNwgiMog3VkcIIjKA",
	"enr:-KG4QDyytgmE4f7AnvW-ZaUOIi9i79qX4JwjRAiXBZCU65wOfBu-3Nb5I7b_Rmg3KCOcZM_C3y5pg7EBU5XGrcLTduQEhGV0aDKQ9aX9QgAAAAD__________4JpZIJ2NIJpcIQ2_DUbiXNlY3AyNTZrMaEDKnz_-ps3UUOfHWVYaskI5kWYO_vtYMGYCQRAR3gHDouDdGNwgiMog3VkcIIjKA",
	// Prylab team's bootnodes
	"enr:-Ku4QImhMc1z8yCiNJ1TyUxdcfNucje3BGwEHzodEZUan8PherEo4sF7pPHPSIB1NNuSg5fZy7qFsjmUKs2ea1Whi0EBh2F0dG5ldHOIAAAAAAAAAACEZXRoMpD1pf1CAAAAAP__________gmlkgnY0gmlwhBLf22SJc2VjcDI1NmsxoQOVphkDqal4QzPMksc5wnpuC3gvSC8AfbFOnZY_On34wIN1ZHCCIyg",
	"enr:-Ku4QP2xDnEtUXIjzJ_DhlCRN9SN99RYQPJL92TMlSv7U5C1YnYLjwOQHgZIUXw6c-BvRg2Yc2QsZxxoS_pPRVe0yK8Bh2F0dG5ldHOIAAAAAAAAAACEZXRoMpD1pf1CAAAAAP__________gmlkgnY0gmlwhBLf22SJc2VjcDI1NmsxoQMeFF5GrS7UZpAH2Ly84aLK-TyvH-dRo0JM1i8yygH50YN1ZHCCJxA",
	"enr:-Ku4QPp9z1W4tAO8Ber_NQierYaOStqhDqQdOPY3bB3jDgkjcbk6YrEnVYIiCBbTxuar3CzS528d2iE7TdJsrL-dEKoBh2F0dG5ldHOIAAAAAAAAAACEZXRoMpD1pf1CAAAAAP__________gmlkgnY0gmlwhBLf22SJc2VjcDI1NmsxoQMw5fqqkw2hHC4F5HZZDPsNmPdB1Gi8JPQK7pRc9XHh-oN1ZHCCKvg",
	// Lighthouse team's bootnodes
	"enr:-IS4QLkKqDMy_ExrpOEWa59NiClemOnor-krjp4qoeZwIw2QduPC-q7Kz4u1IOWf3DDbdxqQIgC4fejavBOuUPy-HE4BgmlkgnY0gmlwhCLzAHqJc2VjcDI1NmsxoQLQSJfEAHZApkm5edTCZ_4qps_1k_ub2CxHFxi-gr2JMIN1ZHCCIyg",
	"enr:-IS4QDAyibHCzYZmIYZCjXwU9BqpotWmv2BsFlIq1V31BwDDMJPFEbox1ijT5c2Ou3kvieOKejxuaCqIcjxBjJ_3j_cBgmlkgnY0gmlwhAMaHiCJc2VjcDI1NmsxoQJIdpj_foZ02MXz4It8xKD7yUHTBx7lVFn3oeRP21KRV4N1ZHCCIyg",
	// EF bootnodes
	"enr:-Ku4QHqVeJ8PPICcWk1vSn_XcSkjOkNiTg6Fmii5j6vUQgvzMc9L1goFnLKgXqBJspJjIsB91LTOleFmyWWrFVATGngBh2F0dG5ldHOIAAAAAAAAAACEZXRoMpC1MD8qAAAAAP__________gmlkgnY0gmlwhAMRHkWJc2VjcDI1NmsxoQKLVXFOhp2uX6jeT0DvvDpPcU8FWMjQdR4wMuORMhpX24N1ZHCCIyg",
	"enr:-Ku4QG-2_Md3sZIAUebGYT6g0SMskIml77l6yR-M_JXc-UdNHCmHQeOiMLbylPejyJsdAPsTHJyjJB2sYGDLe0dn8uYBh2F0dG5ldHOIAAAAAAAAAACEZXRoMpC1MD8qAAAAAP__________gmlkgnY0gmlwhBLY-NyJc2VjcDI1NmsxoQORcM6e19T1T9gi7jxEZjk_sjVLGFscUNqAY9obgZaxbIN1ZHCCIyg",
	"enr:-Ku4QPn5eVhcoF1opaFEvg1b6JNFD2rqVkHQ8HApOKK61OIcIXD127bKWgAtbwI7pnxx6cDyk_nI88TrZKQaGMZj0q0Bh2F0dG5ldHOIAAAAAAAAAACEZXRoMpC1MD8qAAAAAP__________gmlkgnY0gmlwhDayLMaJc2VjcDI1NmsxoQK2sBOLGcUb4AwuYzFuAVCaNHA-dy24UuEKkeFNgCVCsIN1ZHCCIyg",
	"enr:-Ku4QEWzdnVtXc2Q0ZVigfCGggOVB2Vc1ZCPEc6j21NIFLODSJbvNaef1g4PxhPwl_3kax86YPheFUSLXPRs98vvYsoBh2F0dG5ldHOIAAAAAAAAAACEZXRoMpC1MD8qAAAAAP__________gmlkgnY0gmlwhDZBrP2Jc2VjcDI1NmsxoQM6jr8Rb1ktLEsVcKAPa08wCsKUmvoQ8khiOl_SLozf9IN1ZHCCIyg",
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
	case SepoliaGenesisHash:
		net = "sepolia"
	default:
		return ""
	}
	return dnsPrefix + protocol + "." + net + ".ethdisco.net"
}

var GAploContractAddress = common.HexToAddress("0x0000000000000000000000000000000000001234")
var GAploMineSelector = [4]byte{47, 220, 80, 94}
var GAploRewardCoef = big.NewInt(66)
var AploContractAddress = common.HexToAddress("0x0000000000000000000000000000000000001235")

const GAPLO = "608060405234801561001057600080fd5b50600436106101165760003560e01c8063609be203116100a2578063a457c2d711610071578063a457c2d71461030b578063a9059cbb1461033b578063dd62ed3e1461036b578063e31305ae1461039b578063ffd58ff5146103ce57610116565b8063609be2031461028157806370a082311461029f5780637f05b9ef146102cf57806395d89b41146102ed57610116565b80632fdc505e116100e95780632fdc505e146101b7578063313ce567146101d357806339509351146101f1578063410085df146102215780635f7619a41461025157610116565b806306fdde031461011b578063095ea7b31461013957806318160ddd1461016957806323b872dd14610187575b600080fd5b6101236103ec565b6040516101309190611d93565b60405180910390f35b610153600480360381019061014e91906116c2565b61047e565b6040516101609190611d34565b60405180910390f35b61017161049c565b60405161017e9190611f55565b60405180910390f35b6101a1600480360381019061019c9190611673565b6104a6565b6040516101ae9190611d34565b60405180910390f35b6101d160048036038101906101cc91906116fe565b6105f7565b005b6101db610932565b6040516101e89190611fb5565b60405180910390f35b61020b600480360381019061020691906116c2565b61093b565b6040516102189190611d34565b60405180910390f35b61023b600480360381019061023691906116c2565b6109e7565b6040516102489190611d34565b60405180910390f35b61026b60048036038101906102669190611727565b610a78565b6040516102789190611d34565b60405180910390f35b610289610a99565b6040516102969190611f55565b60405180910390f35b6102b960048036038101906102b4919061160e565b610abd565b6040516102c69190611f55565b60405180910390f35b6102d7610b05565b6040516102e49190611f55565b60405180910390f35b6102f5610b0e565b6040516103029190611d93565b60405180910390f35b610325600480360381019061032091906116c2565b610ba0565b6040516103329190611d34565b60405180910390f35b610355600480360381019061035091906116c2565b610c94565b6040516103629190611d34565b60405180910390f35b61038560048036038101906103809190611637565b610d75565b6040516103929190611f55565b60405180910390f35b6103b560048036038101906103b0919061160e565b610dfc565b6040516103c59493929190611f70565b60405180910390f35b6103d6610e2c565b6040516103e39190611d78565b60405180910390f35b6060600380546103fb90612134565b80601f016020809104026020016040519081016040528092919081815260200182805461042790612134565b80156104745780601f1061044957610100808354040283529160200191610474565b820191906000526020600020905b81548152906001019060200180831161045757829003601f168201915b5050505050905090565b600061049261048b610e50565b8484610e58565b6001905092915050565b6000600254905090565b6000806104b1610e50565b90506000600560008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206040518060800160405290816000820154815260200160018201548152602001600282015481526020016003820154815250509050600a816040015111610570576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161056790611e55565b60405180910390fd5b600061057c8784610d75565b9050848110156105c1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105b890611e35565b60405180910390fd5b6105de836105cd610e50565b87846105d99190612042565b610e58565b6105e9878787611023565b600193505050509392505050565b6000600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206040518060800160405290816000820154815260200160018201548152602001600282015481526020016003820154815250509050600081602001511415610710577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81602001818152505080600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082015181600001556020820151816001015560408201518160020155606082015181600301559050505b60003383836020015184606001518560400151604051602001610737959493929190611cd5565b60405160208183030381529060405280519060200120905060008160001c90508260200151811061079d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161079490611ef5565b60405180910390fd5b438360000181815250506001836040018181516107ba9190611fec565b915081815250508083606001818152505060008360000151436107dd9190612042565b905060058110156107fe57808460200151901b84602001818152505061086c565b60058111801561083257507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff846020015114155b1561086b577e8000000000000000000000000000000000000000000000000000000000000060018560200151901c178460200181815250505b5b83600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082015181600001556020820151816001015560408201518160020155606082015181600301559050503373ffffffffffffffffffffffffffffffffffffffff167f634a203b5db8721108a45f719c962a5efff15dfebb1005764878de1795d47716868660600151604051610923929190611d4f565b60405180910390a25050505050565b60006012905090565b60006109dd610948610e50565b848460016000610956610e50565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546109d89190611fec565b610e58565b6001905092915050565b6000806109f2610e50565b9050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610a63576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a5a90611e95565b60405180910390fd5b610a6d84846112a2565b600191505092915050565b600080610a83610e50565b9050610a8f81846113f6565b6001915050919050565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6402540be40081565b606060048054610b1d90612134565b80601f0160208091040260200160405190810160405280929190818152602001828054610b4990612134565b8015610b965780601f10610b6b57610100808354040283529160200191610b96565b820191906000526020600020905b815481529060010190602001808311610b7957829003601f168201915b5050505050905090565b60008060016000610baf610e50565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905082811015610c6c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c6390611f15565b60405180910390fd5b610c89610c77610e50565b858584610c849190612042565b610e58565b600191505092915050565b600080610c9f610e50565b90506000600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206040518060800160405290816000820154815260200160018201548152602001600282015481526020016003820154815250509050600a816040015111610d5e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d5590611e55565b60405180910390fd5b610d69828686611023565b60019250505092915050565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b60056020528060005260406000206000915090508060000154908060010154908060020154908060030154905084565b7f2fdc505e386e491ea84055436bf3dd7ff759f946cca6cf3d132998e18196a4e481565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415610ec8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ebf90611ed5565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610f38576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f2f90611df5565b60405180910390fd5b80600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925836040516110169190611f55565b60405180910390a3505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415611093576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161108a90611eb5565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611103576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110fa90611db5565b60405180910390fd5b61110e8383836115ca565b60008060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905081811015611194576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161118b90611e15565b60405180910390fd5b81816111a09190612042565b6000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546112309190611fec565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516112949190611f55565b60405180910390a350505050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611312576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161130990611f35565b60405180910390fd5b61131e600083836115ca565b80600260008282546113309190611fec565b92505081905550806000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546113859190611fec565b925050819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516113ea9190611f55565b60405180910390a35050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611466576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161145d90611e75565b60405180910390fd5b611472826000836115ca565b60008060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050818110156114f8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114ef90611dd5565b60405180910390fd5b81816115049190612042565b6000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555081600260008282546115589190612042565b92505081905550600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516115bd9190611f55565b60405180910390a3505050565b505050565b6000813590506115de8161221a565b92915050565b6000813590506115f381612231565b92915050565b60008135905061160881612248565b92915050565b60006020828403121561162057600080fd5b600061162e848285016115cf565b91505092915050565b6000806040838503121561164a57600080fd5b6000611658858286016115cf565b9250506020611669858286016115cf565b9150509250929050565b60008060006060848603121561168857600080fd5b6000611696868287016115cf565b93505060206116a7868287016115cf565b92505060406116b8868287016115f9565b9150509250925092565b600080604083850312156116d557600080fd5b60006116e3858286016115cf565b92505060206116f4858286016115f9565b9150509250929050565b60006020828403121561171057600080fd5b600061171e848285016115e4565b91505092915050565b60006020828403121561173957600080fd5b6000611747848285016115f9565b91505092915050565b61176161175c82612076565b612166565b82525050565b61177081612088565b82525050565b61177f81612094565b82525050565b61179661179182612094565b612178565b82525050565b6117a58161209e565b82525050565b60006117b682611fd0565b6117c08185611fdb565b93506117d0818560208601612101565b6117d9816121fc565b840191505092915050565b60006117f1602383611fdb565b91507f45524332303a207472616e7366657220746f20746865207a65726f206164647260008301527f65737300000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000611857602283611fdb565b91507f45524332303a206275726e20616d6f756e7420657863656564732062616c616e60008301527f63650000000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b60006118bd602283611fdb565b91507f45524332303a20617070726f766520746f20746865207a65726f20616464726560008301527f73730000000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000611923602683611fdb565b91507f45524332303a207472616e7366657220616d6f756e742065786365656473206260008301527f616c616e636500000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000611989602883611fdb565b91507f45524332303a207472616e7366657220616d6f756e742065786365656473206160008301527f6c6c6f77616e63650000000000000000000000000000000000000000000000006020830152604082019050919050565b60006119ef602483611fdb565b91507f546f74616c206d696e65642073686f756c64206265206869676865722074686160008301527f6e203130000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000611a55602183611fdb565b91507f45524332303a206275726e2066726f6d20746865207a65726f2061646472657360008301527f73000000000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000611abb601683611fdb565b91507f4f6e6c7920726f6f742063616e2074616b6520666565000000000000000000006000830152602082019050919050565b6000611afb602583611fdb565b91507f45524332303a207472616e736665722066726f6d20746865207a65726f20616460008301527f64726573730000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000611b61602483611fdb565b91507f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460008301527f72657373000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000611bc7601583611fdb565b91507f496e76616c69642070726f6f66206f6620776f726b00000000000000000000006000830152602082019050919050565b6000611c07602583611fdb565b91507f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760008301527f207a65726f0000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000611c6d601f83611fdb565b91507f45524332303a206d696e7420746f20746865207a65726f2061646472657373006000830152602082019050919050565b611ca9816120ea565b82525050565b611cc0611cbb826120ea565b612194565b82525050565b611ccf816120f4565b82525050565b6000611ce18288611750565b601482019150611cf18287611785565b602082019150611d018286611caf565b602082019150611d118285611caf565b602082019150611d218284611caf565b6020820191508190509695505050505050565b6000602082019050611d496000830184611767565b92915050565b6000604082019050611d646000830185611776565b611d716020830184611ca0565b9392505050565b6000602082019050611d8d600083018461179c565b92915050565b60006020820190508181036000830152611dad81846117ab565b905092915050565b60006020820190508181036000830152611dce816117e4565b9050919050565b60006020820190508181036000830152611dee8161184a565b9050919050565b60006020820190508181036000830152611e0e816118b0565b9050919050565b60006020820190508181036000830152611e2e81611916565b9050919050565b60006020820190508181036000830152611e4e8161197c565b9050919050565b60006020820190508181036000830152611e6e816119e2565b9050919050565b60006020820190508181036000830152611e8e81611a48565b9050919050565b60006020820190508181036000830152611eae81611aae565b9050919050565b60006020820190508181036000830152611ece81611aee565b9050919050565b60006020820190508181036000830152611eee81611b54565b9050919050565b60006020820190508181036000830152611f0e81611bba565b9050919050565b60006020820190508181036000830152611f2e81611bfa565b9050919050565b60006020820190508181036000830152611f4e81611c60565b9050919050565b6000602082019050611f6a6000830184611ca0565b92915050565b6000608082019050611f856000830187611ca0565b611f926020830186611ca0565b611f9f6040830185611ca0565b611fac6060830184611ca0565b95945050505050565b6000602082019050611fca6000830184611cc6565b92915050565b600081519050919050565b600082825260208201905092915050565b6000611ff7826120ea565b9150612002836120ea565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156120375761203661219e565b5b828201905092915050565b600061204d826120ea565b9150612058836120ea565b92508282101561206b5761206a61219e565b5b828203905092915050565b6000612081826120ca565b9050919050565b60008115159050919050565b6000819050919050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b60005b8381101561211f578082015181840152602081019050612104565b8381111561212e576000848401525b50505050565b6000600282049050600182168061214c57607f821691505b602082108114156121605761215f6121cd565b5b50919050565b600061217182612182565b9050919050565b6000819050919050565b600061218d8261220d565b9050919050565b6000819050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000601f19601f8301169050919050565b60008160601b9050919050565b61222381612076565b811461222e57600080fd5b50565b61223a81612094565b811461224557600080fd5b50565b612251816120ea565b811461225c57600080fd5b5056fea2646970667358221220c579e1fd1107697ab0fa7525b50b39cfd1f352ea734f274d5db0c60f0104eb8a64736f6c63430008000033"
