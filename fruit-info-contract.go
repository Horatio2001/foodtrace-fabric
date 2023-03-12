/*
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// FruitInfoContract contract for managing CRUD for FruitInfo
type FruitInfoContract struct {
	contractapi.Contract
}

// FruitInfoExists returns true when asset with given ID exists in world state
func (c *FruitInfoContract) FruitInfoExists(ctx contractapi.TransactionContextInterface, fruitInfoID string) (bool, error) {
	data, err := ctx.GetStub().GetState(fruitInfoID)

	if err != nil {
		return false, err
	}

	return data != nil, nil
}

// CreateFruitInfo creates a new instance of FruitInfo
func (c *FruitInfoContract) CreateFruitInfo(ctx contractapi.TransactionContextInterface, fruitInfoID string, processID string, collectID string) error {
	exists, err := c.FruitInfoExists(ctx, fruitInfoID)
	if err != nil {
		return fmt.Errorf("could not read from world state. %s", err)
	} else if exists {
		return fmt.Errorf("the asset %s already exists", fruitInfoID)
	}

	fruitInfo := new(FruitInfo)
	fruitInfo.ProcessInstanceID = processID
	fruitInfo.CollectID = collectID

	bytes, _ := json.Marshal(fruitInfo)

	return ctx.GetStub().PutState(fruitInfoID, bytes)
}

// SetCollectInfo retrieves an instance of FruitInfo from the world state and set collection information for it
func (c *FruitInfoContract) SetCollectInfo(ctx contractapi.TransactionContextInterface, fruitInfoID string, args []string) error {
	if len(args) != 11 {
		return fmt.Errorf("collect args num should be 11, but now is %d", len(args))
	}
	exists, err := c.FruitInfoExists(ctx, fruitInfoID)
	if err != nil {
		return fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return fmt.Errorf("the asset %s does not exist", fruitInfoID)
	}
	var collectInfo CollectInfo
	collectInfo.CollectPlaceLongitude = args[0]
	collectInfo.CollectPlaceLatitude = args[1]
	collectInfo.CollectPlaceAltitude = args[2]
	collectInfo.CollectPlaceSoilType = args[3]
	collectInfo.CollectPlaceEcologyType = args[4]
	collectInfo.CollectMaterialType = args[5]
	collectInfo.CollectPeople = args[6]
	collectInfo.CollectUnit = args[7]
	collectInfo.CollectTime = args[8]
	collectInfo.SpeciesName = args[9]
	collectInfo.Image = args[10]
	fruitInfo := new(FruitInfo)
	fruitInfo, _ = c.ReadFruitInfo(ctx, fruitInfoID)
	fruitInfo.CollectInfo = collectInfo
	bytes, _ := json.Marshal(fruitInfo)
	return ctx.GetStub().PutState(fruitInfoID, bytes)
}

// SetSpeciesInfo retrieves an instance of FruitInfo from the world state and set species information for it
func (c *FruitInfoContract) SetSpeciesInfo(ctx contractapi.TransactionContextInterface, fruitInfoID string, args []string) error {
	if len(args) != 10 {
		return fmt.Errorf("speciesInfo args num should be 10, but now is %d", len(args))
	}
	exists, err := c.FruitInfoExists(ctx, fruitInfoID)
	if err != nil {
		return fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return fmt.Errorf("the asset %s does not exist", fruitInfoID)
	}
	var speciesInfo SpeciesInfo
	speciesInfo.Type = args[0]
	speciesInfo.Name = args[1]
	speciesInfo.GermplasmName = args[2]
	speciesInfo.GermplasmNameEn = args[3]
	speciesInfo.SectionName = args[4]
	speciesInfo.GenericName = args[5]
	speciesInfo.ScientificName = args[6]
	speciesInfo.ResourceType = args[7]
	speciesInfo.CollectMethod = args[8]
	speciesInfo.GermplasmSource = args[9]
	fruitInfo := new(FruitInfo)
	fruitInfo, _ = c.ReadFruitInfo(ctx, fruitInfoID)
	fruitInfo.SpeciesInfo = speciesInfo
	bytes, _ := json.Marshal(fruitInfo)
	return ctx.GetStub().PutState(fruitInfoID, bytes)
}

// SetSourceInfo retrieves an instance of FruitInfo from the world state and set source information for it
func (c *FruitInfoContract) SetSourceInfo(ctx contractapi.TransactionContextInterface, fruitInfoID string, args []string) error {
	if len(args) != 6 {
		return fmt.Errorf("speciesInfo args num should be 6, but now is %d", len(args))
	}
	exists, err := c.FruitInfoExists(ctx, fruitInfoID)
	if err != nil {
		return fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return fmt.Errorf("the asset %s does not exist", fruitInfoID)
	}
	var sourceInfo SourceInfo
	sourceInfo.SourceCountry = args[0]
	sourceInfo.SourceProvince = args[1]
	sourceInfo.Source = args[2]
	sourceInfo.SourceOrg = args[3]
	sourceInfo.OriginCountry = args[4]
	sourceInfo.OriginPlace = args[5]
	fruitInfo := new(FruitInfo)
	fruitInfo, _ = c.ReadFruitInfo(ctx, fruitInfoID)
	fruitInfo.SourceInfo = sourceInfo
	bytes, _ := json.Marshal(fruitInfo)
	return ctx.GetStub().PutState(fruitInfoID, bytes)
}

// AddTransportInfo retrieves an instance of FruitInfo from the world state and add transport information for it
func (c *FruitInfoContract) AddTransportInfo(ctx contractapi.TransactionContextInterface, fruitInfoID string, args []string) error {
	if len(args) != 8 {
		return fmt.Errorf("speciesInfo args num should be 8, but now is %d", len(args))
	}
	exists, err := c.FruitInfoExists(ctx, fruitInfoID)
	if err != nil {
		return fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return fmt.Errorf("the asset %s does not exist", fruitInfoID)
	}
	var transportInfo TransportInfo
	transportInfo.TransportDepartureTime = args[0]
	transportInfo.TransportArrivalTime = args[1]
	transportInfo.TransportMission = args[2]
	transportInfo.TransportDeparturePlace = args[3]
	transportInfo.TransportDestination = args[4]
	transportInfo.TransportMethod = args[5]
	transportInfo.TransportDepartmentName = args[6]
	transportInfo.TransporterName = args[7]
	fruitInfo := new(FruitInfo)
	fruitInfo, _ = c.ReadFruitInfo(ctx, fruitInfoID)
	fruitInfo.TransportInfo = append(fruitInfo.TransportInfo, transportInfo)
	bytes, _ := json.Marshal(fruitInfo)
	return ctx.GetStub().PutState(fruitInfoID, bytes)
}

// UpdateProcessID retrieves an instance of FruitInfo from the world state and update its processID
func (c *FruitInfoContract) UpdateProcessID(ctx contractapi.TransactionContextInterface, fruitInfoID string, processID string) error {
	exists, err := c.FruitInfoExists(ctx, fruitInfoID)
	if err != nil {
		return fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return fmt.Errorf("the asset %s does not exist", fruitInfoID)
	}
	fruitInfo := new(FruitInfo)
	fruitInfo, _ = c.ReadFruitInfo(ctx, fruitInfoID)
	fruitInfo.ProcessInstanceID = processID
	bytes, _ := json.Marshal(fruitInfo)
	return ctx.GetStub().PutState(fruitInfoID, bytes)
}

// UpdateCollectID retrieves an instance of FruitInfo from the world state and update its collectID
func (c *FruitInfoContract) UpdateCollectID(ctx contractapi.TransactionContextInterface, fruitInfoID string, collectID string) error {
	exists, err := c.FruitInfoExists(ctx, fruitInfoID)
	if err != nil {
		return fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return fmt.Errorf("the asset %s does not exist", fruitInfoID)
	}
	fruitInfo := new(FruitInfo)
	fruitInfo, _ = c.ReadFruitInfo(ctx, fruitInfoID)
	fruitInfo.CollectID = collectID
	bytes, _ := json.Marshal(fruitInfo)
	return ctx.GetStub().PutState(fruitInfoID, bytes)
}

// ReadFruitInfo retrieves an instance of FruitInfo from the world state
func (c *FruitInfoContract) ReadFruitInfo(ctx contractapi.TransactionContextInterface, fruitInfoID string) (*FruitInfo, error) {
	exists, err := c.FruitInfoExists(ctx, fruitInfoID)
	if err != nil {
		return nil, fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return nil, fmt.Errorf("the asset %s does not exist", fruitInfoID)
	}

	bytes, _ := ctx.GetStub().GetState(fruitInfoID)

	fruitInfo := new(FruitInfo)

	err = json.Unmarshal(bytes, fruitInfo)

	if err != nil {
		return nil, fmt.Errorf("could not unmarshal world state data to type FruitInfo")
	}

	return fruitInfo, nil
}

// ReadFruitInfoByRange retrieves an instance of FruitInfo from the world state by range
func (c *FruitInfoContract) ReadFruitInfoByRange(ctx contractapi.TransactionContextInterface, fruitInfoIDStart string, fruitInfoIDEnd string) ([]*FruitInfo, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange(fruitInfoIDStart, fruitInfoIDEnd)
	if err != nil {
		return nil, err
	}
	defer func(resultsIterator shim.StateQueryIteratorInterface) {
		err := resultsIterator.Close()
		if err != nil {
			return
		}
	}(resultsIterator)

	var fruitInfos []*FruitInfo
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, fmt.Errorf("ite.next dont exists, %s", err)
		}
		var fruitInfo FruitInfo
		err = json.Unmarshal(queryResponse.Value, &fruitInfo)
		if err != nil {
			return nil, fmt.Errorf("unmarshl error, %s", err)
		}
		fruitInfos = append(fruitInfos, &fruitInfo)
	}
	return fruitInfos, nil
}

// ReadHistory retrieves an instance of FruitInfo from the world state and return its alter history
func (c *FruitInfoContract) ReadHistory(ctx contractapi.TransactionContextInterface, fruitInfoID string) ([]*FruitInfo, error) {
	historyIterator, err := ctx.GetStub().GetHistoryForKey(fruitInfoID)
	if err != nil {
		return nil, err
	}
	defer func(historyIterator shim.HistoryQueryIteratorInterface) {
		err := historyIterator.Close()
		if err != nil {
			return
		}
	}(historyIterator)
	var fruitInfos []*FruitInfo
	for historyIterator.HasNext() {
		queryResponse, err := historyIterator.Next()
		if err != nil {
			return nil, fmt.Errorf("ite.next dont exists, %s", err)
		}
		var fruitInfo FruitInfo
		err = json.Unmarshal(queryResponse.Value, &fruitInfo)
		if err != nil {
			return nil, fmt.Errorf("unmarshl error, %s", err)
		}
		fruitInfos = append(fruitInfos, &fruitInfo)
	}
	return fruitInfos, nil
}

// DeleteFruitInfo deletes an instance of FruitInfo from the world state
func (c *FruitInfoContract) DeleteFruitInfo(ctx contractapi.TransactionContextInterface, fruitInfoID string) error {
	exists, err := c.FruitInfoExists(ctx, fruitInfoID)
	if err != nil {
		return fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return fmt.Errorf("the asset %s does not exist", fruitInfoID)
	}

	return ctx.GetStub().DelState(fruitInfoID)
}
