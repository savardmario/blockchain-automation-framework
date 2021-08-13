/*
SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"log"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"drone-reservation/chaincode"
)

func main() {
	assetChaincode, err := contractapi.NewChaincode(&chaincode.SmartContract{})
	if err != nil {
		log.Panicf("Error creating drone reservation chaincode: %v", err)
	}

	if err := assetChaincode.Start(); err != nil {
		log.Panicf("Error starting drone reservation chaincode: %v", err)
	}
}
