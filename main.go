/*
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/fabric-contract-api-go/metadata"
)

func main() {
	fruitInfoContract := new(FruitInfoContract)
	fruitInfoContract.Info.Version = "0.0.1"
	fruitInfoContract.Info.Description = "My Smart Contract"
	fruitInfoContract.Info.License = new(metadata.LicenseMetadata)
	fruitInfoContract.Info.License.Name = "Apache-2.0"
	fruitInfoContract.Info.Contact = new(metadata.ContactMetadata)
	fruitInfoContract.Info.Contact.Name = "John Doe"

	chaincode, err := contractapi.NewChaincode(fruitInfoContract)
	chaincode.Info.Title = "realChainCode chaincode"
	chaincode.Info.Version = "0.0.1"

	if err != nil {
		panic("Could not create chaincode from FruitInfoContract." + err.Error())
	}

	err = chaincode.Start()

	if err != nil {
		panic("Failed to start chaincode. " + err.Error())
	}
}
