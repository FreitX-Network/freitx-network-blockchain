// Copyright (c) FreitX Network
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package testaddress

import (
	"github.com/freitx-project/freitx-network-blockchain/config"
	"github.com/freitx-project/freitx-network-blockchain/onexaddress"
)

const (
	pubkeyA        = "2c9ccbeb9ee91271f7e5c2103753be9c9edff847e1a51227df6a6b0765f31a4b424e84027b44a663950f013a88b8fd8cdc53b1eda1d4b73f9d9dc12546c8c87d68ff1435a0f8a006"
	prikeyA        = "b5affb30846a00ef5aa39b57f913d70cd8cf6badd587239863cb67feacf6b9f30c34e800"
	pubkeyB        = "881504d84a0659e14dcba59f24a98e71cda55b139615342668840c64678f1514941bbd053c7492fb9b719e6050cfa972efa491b79e11a1713824dda5f638fc0d9fa1b68be3c0f905"
	prikeyB        = "b89c1ec0fb5b192c8bb8f6fcf9a871e4a67ef462f40d2b8ff426da1d1eaedd9696dc9d00"
	pubkeyC        = "252fc7bc9a993b68dd7b13a00213c9cf4befe80da49940c52220f93c7147771ba2d783045cf0fbf2a86b32a62848befb96c0f38c0487a5ccc806ff28bb06d9faf803b93dda107003"
	prikeyC        = "3e05de562a27fb6e25ac23ff8bcaa1ada0c253fa8ff7c6d15308f65d06b6990f64ee9601"
	pubkeyD        = "29aa28cc21c3ee3cc658d3a322997ceb8d5d352f45d052192d3ab57cd196d3375af558067f5a2cfe5fc65d5249cc07f991bab683468382a3acaa4c8b7af35156b46aeda00620f307"
	prikeyD        = "d4b7b441382751d9a1955152b46a69f3c9f9559c6205757af928f5181ff207060d0dab00"
	pubkeyE        = "64dc2d5f445a78b884527252a3dba1f72f52251c97ec213dda99868882024d4d1442f100c8f1f833d0c687871a959ee97665dea24de1a627cce6c970d9db5859da9e4295bb602e04"
	prikeyE        = "53a827f7c5b4b4040b22ae9b12fcaa234e8362fa022480f50b8643981806ed67c7f77a00"
	pubkeyF        = "ed3540a4a01ee813cf133273ade824523d27bcb3165a16035c61051f25ecdb1429e8600173d4038d50f09cf951cf962be839b8871eac85e1a5bf6533fc6ddfb546869e8a0cd3c501"
	prikeyF        = "6bdd24a9b0a2a4225a286e31633af0f24a2fdbc81e8c18a11885b7689d0fa0c87cd04101"
	pubkeyG        = "390d6ef794d27d4ddab1a56fc879f20c0f855caf661d3e0ffb4d32fb5251b16efbf8ca0192977784b986cf01b9c917352df8912e026a7abcc5d41456af5da8ab0ff16ce6ff3efc02"
	prikeyG        = "864186b54785dbe31b2c115ea8f66d28307af4b6d24886f6179f8b0907e0e1d8e6863b01"
	pubkeyProducer = "5e24c6c19eb50a6da14d0d2841ee8b7f8e31771f31413466526f7a726f70d8a619421f066d2033c413cbaeb710de4056061c1ed728274cfaaa69a91436ec3fe2135d0e40e7fbae03"
	prikeyProducer = "c0c08b8f4887c62a07d07388575e03d650cdcf1cf8050ba09881ea49ecb9746a4854ee01"
)

// Addrinfo contains the address information
var Addrinfo map[string]*OnexAddress.Address

func init() {
	Addrinfo = make(map[string]*OnexAddress.Address)

	chainID := config.Default.Chain.ID
	Addrinfo["producer"] = ConstructAddress(chainID, pubkeyProducer, prikeyProducer)
	Addrinfo["alfa"] = ConstructAddress(chainID, pubkeyA, prikeyA)
	Addrinfo["bravo"] = ConstructAddress(chainID, pubkeyB, prikeyB)
	Addrinfo["charlie"] = ConstructAddress(chainID, pubkeyC, prikeyC)
	Addrinfo["delta"] = ConstructAddress(chainID, pubkeyD, prikeyD)
	Addrinfo["echo"] = ConstructAddress(chainID, pubkeyE, prikeyE)
	Addrinfo["foxtrot"] = ConstructAddress(chainID, pubkeyF, prikeyF)
	Addrinfo["galilei"] = ConstructAddress(chainID, pubkeyG, prikeyG)
}
