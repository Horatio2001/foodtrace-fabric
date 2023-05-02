/*
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"realChainCode/util"
	"time"
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
func (c *FruitInfoContract) CreateFruitInfo(ctx contractapi.TransactionContextInterface, fruitInfoID string, args []string) error {
	if len(args) != 29 {
		return fmt.Errorf("collect args num should be 29, but now is %d", len(args))
	}
	exists, err := c.FruitInfoExists(ctx, fruitInfoID)
	if err != nil {
		return fmt.Errorf("Could not read from world state. %s", err)
	} else if exists {
		return fmt.Errorf("The asset %s already exists", fruitInfoID)
	}
	var collectInfo CollectInfo
	fruitInfo := new(FruitInfo)
	fruitInfo.ID = fruitInfoID
	collectInfo.CollectID = args[0]
	collectInfo.Type = args[1]
	collectInfo.Name = args[2]
	collectInfo.GermplasmName = args[3]
	collectInfo.GermplasmNameEn = args[4]
	collectInfo.SectionName = args[5]
	collectInfo.GenericName = args[6]
	collectInfo.ScientificName = args[7]
	collectInfo.ResourceType = args[8]
	collectInfo.CollectMethod = args[9]
	collectInfo.GermplasmSource = args[10]
	collectInfo.SourceCountry = args[11]
	collectInfo.SourceProvince = args[12]
	collectInfo.Source = args[13]
	collectInfo.SourceOrg = args[14]
	collectInfo.OriginCountry = args[15]
	collectInfo.OriginPlace = args[16]
	collectInfo.CollectPlaceLongitude = args[17]
	collectInfo.CollectPlaceLatitude = args[18]
	collectInfo.CollectPlaceAltitude = args[19]
	collectInfo.CollectPlaceSoilType = args[20]
	collectInfo.CollectPlaceEcologyType = args[21]
	collectInfo.CollectMaterialType = args[22]
	collectInfo.CollectPeople = args[23]
	collectInfo.CollectUnit = args[24]
	collectInfo.CollectTime = args[25]
	collectInfo.SpeciesName = args[26]
	collectInfo.Image = args[27]
	collectInfo.CollectRemark = args[28]
	collectInfo.CollectHash = util.GetSHA256String(args)
	//createTime, _ := ctx.GetStub().GetTxTimestamp()
	//fruitInfo.CreateTime = time.Unix(createTime.GetSeconds(), int64(createTime.GetNanos())).Local().Format("2006-01-02 15:04:05")
	//fruitInfo.TxId = ctx.GetStub().GetTxID()
	fruitInfo.CollectInfo = collectInfo
	fruitInfo.Status = "0"
	fruitInfo.IsContradict = "0"
	fruitInfo.IsDeleted = "0"
	fruitInfo.IsLoaded = "0"
	bytes, _ := json.Marshal(fruitInfo)
	return ctx.GetStub().PutState(fruitInfoID, bytes)
}

func (c *FruitInfoContract) RejectCreate(ctx contractapi.TransactionContextInterface, fruitInfoID string) error {
	exists, err := c.FruitInfoExists(ctx, fruitInfoID)
	if err != nil {
		return fmt.Errorf("Could not read from world state. %s", err)
	} else if !exists {
		return fmt.Errorf("The asset %s does not exist", fruitInfoID)
	}
	fruitInfo := new(FruitInfo)
	fruitInfo, _ = c.ReadFruitInfo(ctx, fruitInfoID)
	if fruitInfo.IsLoaded != "0" {
		return fmt.Errorf("could not reject createinfo, because the info is already loaded")
	}
	if fruitInfo.Status != "0" {
		return fmt.Errorf("Could not reject create, because the status is %s, but we except 0. %s", fruitInfo.Status, err)
	}
	createTime, _ := ctx.GetStub().GetTxTimestamp()
	fruitInfo.CreateTime = time.Unix(createTime.GetSeconds(), int64(createTime.GetNanos())).Local().Format("2006-01-02 15:04:05")
	fruitInfo.TxId = ctx.GetStub().GetTxID()

	fruitInfo.IsContradict = "1"
	bytes, _ := json.Marshal(fruitInfo)
	return ctx.GetStub().PutState(fruitInfoID, bytes)
}

func (c *FruitInfoContract) ModifyCreateFruitInfo(ctx contractapi.TransactionContextInterface, fruitInfoID string, args []string) error {
	if len(args) != 29 {
		return fmt.Errorf("collect args num should be 29, but now is %d", len(args))
	}
	exists, err := c.FruitInfoExists(ctx, fruitInfoID)
	if err != nil {
		return fmt.Errorf("Could not read from world state. %s", err)
	} else if !exists {
		return fmt.Errorf("The asset %s does not exist", fruitInfoID)
	}
	fruitInfo := new(FruitInfo)
	fruitInfo, _ = c.ReadFruitInfo(ctx, fruitInfoID)
	if fruitInfo.IsLoaded != "0" {
		return fmt.Errorf("could not modify createInfo, because the info is already loaded")
	}

	var collectInfo CollectInfo
	collectInfo.CollectID = args[0]
	collectInfo.Type = args[1]
	collectInfo.Name = args[2]
	collectInfo.GermplasmName = args[3]
	collectInfo.GermplasmNameEn = args[4]
	collectInfo.SectionName = args[5]
	collectInfo.GenericName = args[6]
	collectInfo.ScientificName = args[7]
	collectInfo.ResourceType = args[8]
	collectInfo.CollectMethod = args[9]
	collectInfo.GermplasmSource = args[10]
	collectInfo.SourceCountry = args[11]
	collectInfo.SourceProvince = args[12]
	collectInfo.Source = args[13]
	collectInfo.SourceOrg = args[14]
	collectInfo.OriginCountry = args[15]
	collectInfo.OriginPlace = args[16]
	collectInfo.CollectPlaceLongitude = args[17]
	collectInfo.CollectPlaceLatitude = args[18]
	collectInfo.CollectPlaceAltitude = args[19]
	collectInfo.CollectPlaceSoilType = args[20]
	collectInfo.CollectPlaceEcologyType = args[21]
	collectInfo.CollectMaterialType = args[22]
	collectInfo.CollectPeople = args[23]
	collectInfo.CollectUnit = args[24]
	collectInfo.CollectTime = args[25]
	collectInfo.SpeciesName = args[26]
	collectInfo.Image = args[27]
	collectInfo.CollectRemark = args[28]
	collectInfo.CollectHash = util.GetSHA256String(args)
	//createTime, _ := ctx.GetStub().GetTxTimestamp()
	//fruitInfo.CreateTime = time.Unix(createTime.GetSeconds(), int64(createTime.GetNanos())).Local().Format("2006-01-02 15:04:05")
	//fruitInfo.TxId = ctx.GetStub().GetTxID()
	fruitInfo.CollectInfo = collectInfo
	fruitInfo.Status = "0"
	fruitInfo.IsContradict = "0"
	bytes, _ := json.Marshal(fruitInfo)
	return ctx.GetStub().PutState(fruitInfoID, bytes)
}

func (c *FruitInfoContract) SaveFruitInfo(ctx contractapi.TransactionContextInterface, fruitInfoID string, args []string) error {
	if len(args) != 14 {
		return fmt.Errorf("collect args num should be 14, but now is %d", len(args))
	}
	exists, err := c.FruitInfoExists(ctx, fruitInfoID)
	if err != nil {
		return fmt.Errorf("Could not read from world state. %s", err)
	} else if !exists {
		return fmt.Errorf("The asset %s does not exist", fruitInfoID)
	}

	fruitInfo := new(FruitInfo)
	fruitInfo, _ = c.ReadFruitInfo(ctx, fruitInfoID)
	if fruitInfo.IsLoaded != "0" {
		return fmt.Errorf("could not add saveInfo, because the info is already loaded")
	}
	if fruitInfo.Status != "0" {
		return fmt.Errorf("we expect status from 0 to 1, but the status now is %s", fruitInfo.Status)
	}

	var saveInfo SaveInfo
	saveInfo.MainPreference = args[0]
	saveInfo.MainUse = args[1]
	saveInfo.PreservationFacility = args[2]
	saveInfo.GermplasmType = args[3]
	saveInfo.SaveQuantity = args[4]
	saveInfo.MeasuringUnit = args[5]
	saveInfo.SaveUnit = args[6]
	saveInfo.SaveVault = args[7]
	saveInfo.SavePlace = args[8]
	saveInfo.WarehousingYear = args[9]
	saveInfo.SaveProperty = args[10]
	saveInfo.ResourceDescription = args[11]
	saveInfo.ResourceRemark = args[12]
	saveInfo.GermplasmImage = args[13]
	saveInfo.SaveHash = util.GetSHA256String(args)
	fruitInfo.SaveInfo = saveInfo

	//createTime, _ := ctx.GetStub().GetTxTimestamp()
	//fruitInfo.CreateTime = time.Unix(createTime.GetSeconds(), int64(createTime.GetNanos())).Local().Format("2006-01-02 15:04:05")
	//fruitInfo.TxId = ctx.GetStub().GetTxID()

	fruitInfo.Status = "1"
	fruitInfo.IsContradict = "0"
	bytes, _ := json.Marshal(fruitInfo)
	return ctx.GetStub().PutState(fruitInfoID, bytes)
}

func (c *FruitInfoContract) RejectSave(ctx contractapi.TransactionContextInterface, fruitInfoID string) error {
	exists, err := c.FruitInfoExists(ctx, fruitInfoID)
	if err != nil {
		return fmt.Errorf("Could not read from world state. %s", err)
	} else if !exists {
		return fmt.Errorf("The asset %s does not exist", fruitInfoID)
	}
	fruitInfo := new(FruitInfo)
	fruitInfo, _ = c.ReadFruitInfo(ctx, fruitInfoID)
	if fruitInfo.IsLoaded != "0" {
		return fmt.Errorf("could not reject saveInfo, because the info is already loaded")
	}
	if fruitInfo.Status != "1" {
		return fmt.Errorf("Could not reject save, because the status is %s, but we except 1.", fruitInfo.Status)
	}

	createTime, _ := ctx.GetStub().GetTxTimestamp()
	fruitInfo.CreateTime = time.Unix(createTime.GetSeconds(), int64(createTime.GetNanos())).Local().Format("2006-01-02 15:04:05")
	fruitInfo.TxId = ctx.GetStub().GetTxID()

	fruitInfo.IsContradict = "1"
	bytes, _ := json.Marshal(fruitInfo)
	return ctx.GetStub().PutState(fruitInfoID, bytes)
}

func (c *FruitInfoContract) ModifySaveFruitInfo(ctx contractapi.TransactionContextInterface, fruitInfoID string, args []string) error {
	if len(args) != 14 {
		return fmt.Errorf("collect args num should be 14, but now is %d", len(args))
	}
	exists, err := c.FruitInfoExists(ctx, fruitInfoID)
	if err != nil {
		return fmt.Errorf("Could not read from world state. %s", err)
	} else if !exists {
		return fmt.Errorf("The asset %s does not exist", fruitInfoID)
	}
	fruitInfo := new(FruitInfo)
	fruitInfo, _ = c.ReadFruitInfo(ctx, fruitInfoID)
	if fruitInfo.IsLoaded != "0" {
		return fmt.Errorf("could not modify saveInfo, because the info is already loaded")
	}

	var saveInfo SaveInfo
	saveInfo.MainPreference = args[0]
	saveInfo.MainUse = args[1]
	saveInfo.PreservationFacility = args[2]
	saveInfo.GermplasmType = args[3]
	saveInfo.SaveQuantity = args[4]
	saveInfo.MeasuringUnit = args[5]
	saveInfo.SaveUnit = args[6]
	saveInfo.SaveVault = args[7]
	saveInfo.SavePlace = args[8]
	saveInfo.WarehousingYear = args[9]
	saveInfo.SaveProperty = args[10]
	saveInfo.ResourceDescription = args[11]
	saveInfo.ResourceRemark = args[12]
	saveInfo.GermplasmImage = args[13]
	saveInfo.SaveHash = util.GetSHA256String(args)
	fruitInfo.SaveInfo = saveInfo

	//createTime, _ := ctx.GetStub().GetTxTimestamp()
	//fruitInfo.CreateTime = time.Unix(createTime.GetSeconds(), int64(createTime.GetNanos())).Local().Format("2006-01-02 15:04:05")
	//fruitInfo.TxId = ctx.GetStub().GetTxID()

	fruitInfo.Status = "1"
	fruitInfo.IsContradict = "0"
	bytes, _ := json.Marshal(fruitInfo)
	return ctx.GetStub().PutState(fruitInfoID, bytes)
}

func (c *FruitInfoContract) EnterFruitInfo(ctx contractapi.TransactionContextInterface, fruitInfoID string, args []string) error {
	if len(args) != 6 {
		return fmt.Errorf("collect args num should be 6, but now is %d", len(args))
	}
	exists, err := c.FruitInfoExists(ctx, fruitInfoID)
	if err != nil {
		return fmt.Errorf("Could not read from world state. %s", err)
	} else if !exists {
		return fmt.Errorf("The asset %s does not exist", fruitInfoID)
	}

	fruitInfo := new(FruitInfo)
	fruitInfo, _ = c.ReadFruitInfo(ctx, fruitInfoID)
	if fruitInfo.IsLoaded != "0" {
		return fmt.Errorf("could not add enterinfo, because the info is already loaded")
	}
	if fruitInfo.Status != "1" {
		return fmt.Errorf("we expect status from 1 to 2, but the status now is %s", fruitInfo.Status)
	}

	var enterInfo EnterInfo
	enterInfo.Certifier = args[0]
	enterInfo.CertifyOrg = args[1]
	enterInfo.CertifyPlace = args[2]
	enterInfo.CertifyYear = args[3]
	enterInfo.OperationRange = args[4]
	enterInfo.EnterRemark = args[5]
	enterInfo.EnterHash = util.GetSHA256String(args)
	fruitInfo.EnterInfo = enterInfo

	//createTime, _ := ctx.GetStub().GetTxTimestamp()
	//fruitInfo.CreateTime = time.Unix(createTime.GetSeconds(), int64(createTime.GetNanos())).Local().Format("2006-01-02 15:04:05")
	//fruitInfo.TxId = ctx.GetStub().GetTxID()

	fruitInfo.Status = "2"
	fruitInfo.IsContradict = "0"
	bytes, _ := json.Marshal(fruitInfo)
	return ctx.GetStub().PutState(fruitInfoID, bytes)
}

func (c *FruitInfoContract) RejectEnter(ctx contractapi.TransactionContextInterface, fruitInfoID string) error {
	exists, err := c.FruitInfoExists(ctx, fruitInfoID)
	if err != nil {
		return fmt.Errorf("Could not read from world state. %s", err)
	} else if !exists {
		return fmt.Errorf("The asset %s does not exist", fruitInfoID)
	}
	fruitInfo := new(FruitInfo)
	fruitInfo, _ = c.ReadFruitInfo(ctx, fruitInfoID)
	if fruitInfo.IsLoaded != "0" {
		return fmt.Errorf("could not reject enterinfo, because the info is already loaded")
	}
	if fruitInfo.Status != "2" {
		return fmt.Errorf("Could not reject enter, because the status is %s, but we except 2.", fruitInfo.Status)
	}

	//createTime, _ := ctx.GetStub().GetTxTimestamp()
	//fruitInfo.CreateTime = time.Unix(createTime.GetSeconds(), int64(createTime.GetNanos())).Local().Format("2006-01-02 15:04:05")
	//fruitInfo.TxId = ctx.GetStub().GetTxID()

	fruitInfo.IsContradict = "1"
	bytes, _ := json.Marshal(fruitInfo)
	return ctx.GetStub().PutState(fruitInfoID, bytes)
}

func (c *FruitInfoContract) ModifyEnterFruitInfo(ctx contractapi.TransactionContextInterface, fruitInfoID string, args []string) error {
	if len(args) != 6 {
		return fmt.Errorf("collect args num should be 6, but now is %d", len(args))
	}
	exists, err := c.FruitInfoExists(ctx, fruitInfoID)
	if err != nil {
		return fmt.Errorf("Could not read from world state. %s", err)
	} else if !exists {
		return fmt.Errorf("The asset %s does not exist", fruitInfoID)
	}
	fruitInfo := new(FruitInfo)
	fruitInfo, _ = c.ReadFruitInfo(ctx, fruitInfoID)
	if fruitInfo.IsLoaded != "0" {
		return fmt.Errorf("could not modify enterInfo, because the info is already loaded")
	}

	var enterInfo EnterInfo
	enterInfo.Certifier = args[0]
	enterInfo.CertifyOrg = args[1]
	enterInfo.CertifyPlace = args[2]
	enterInfo.CertifyYear = args[3]
	enterInfo.OperationRange = args[4]
	enterInfo.EnterRemark = args[5]
	enterInfo.EnterHash = util.GetSHA256String(args)
	fruitInfo.EnterInfo = enterInfo

	//createTime, _ := ctx.GetStub().GetTxTimestamp()
	//fruitInfo.CreateTime = time.Unix(createTime.GetSeconds(), int64(createTime.GetNanos())).Local().Format("2006-01-02 15:04:05")
	//fruitInfo.TxId = ctx.GetStub().GetTxID()

	fruitInfo.Status = "2"
	fruitInfo.IsContradict = "0"
	bytes, _ := json.Marshal(fruitInfo)
	return ctx.GetStub().PutState(fruitInfoID, bytes)
}

func (c *FruitInfoContract) ShareFruitInfo(ctx contractapi.TransactionContextInterface, fruitInfoID string, args []string) error {
	if len(args) != 7 {
		return fmt.Errorf("collect args num should be 7, but now is %d", len(args))
	}
	exists, err := c.FruitInfoExists(ctx, fruitInfoID)
	if err != nil {
		return fmt.Errorf("Could not read from world state. %s", err)
	} else if !exists {
		return fmt.Errorf("The asset %s does not exist", fruitInfoID)
	}
	fruitInfo := new(FruitInfo)
	fruitInfo, _ = c.ReadFruitInfo(ctx, fruitInfoID)
	if fruitInfo.IsLoaded != "0" {
		return fmt.Errorf("could not add shareinfo, because the info is already loaded")
	}
	if fruitInfo.Status != "2" {
		return fmt.Errorf("we expect status from 2 to 3, but the status now is %s", fruitInfo.Status)
	}

	var shareInfo ShareInfo
	shareInfo.ShareObj = args[0]
	shareInfo.ContactInfo = args[1]
	shareInfo.ShareMode = args[2]
	shareInfo.ShareUse = args[3]
	shareInfo.ShareNum = args[4]
	shareInfo.ShareBeginTime = args[5]
	shareInfo.ShareEndTime = args[6]
	shareInfo.ShareHash = util.GetSHA256String(args)
	fruitInfo.ShareInfo = shareInfo

	//createTime, _ := ctx.GetStub().GetTxTimestamp()
	//fruitInfo.CreateTime = time.Unix(createTime.GetSeconds(), int64(createTime.GetNanos())).Local().Format("2006-01-02 15:04:05")
	//fruitInfo.TxId = ctx.GetStub().GetTxID()

	fruitInfo.Status = "3"
	fruitInfo.IsContradict = "0"
	bytes, _ := json.Marshal(fruitInfo)
	return ctx.GetStub().PutState(fruitInfoID, bytes)
}

func (c *FruitInfoContract) RejectShare(ctx contractapi.TransactionContextInterface, fruitInfoID string) error {
	exists, err := c.FruitInfoExists(ctx, fruitInfoID)
	if err != nil {
		return fmt.Errorf("Could not read from world state. %s", err)
	} else if !exists {
		return fmt.Errorf("The asset %s does not exist", fruitInfoID)
	}
	fruitInfo := new(FruitInfo)
	fruitInfo, _ = c.ReadFruitInfo(ctx, fruitInfoID)
	if fruitInfo.IsLoaded != "0" {
		return fmt.Errorf("could not reject shareinfo, because the info is already loaded")
	}
	if fruitInfo.Status != "3" {
		return fmt.Errorf("Could not reject share, because the status is %s, but we except 3.", fruitInfo.Status)
	}

	//createTime, _ := ctx.GetStub().GetTxTimestamp()
	//fruitInfo.CreateTime = time.Unix(createTime.GetSeconds(), int64(createTime.GetNanos())).Local().Format("2006-01-02 15:04:05")
	//fruitInfo.TxId = ctx.GetStub().GetTxID()

	fruitInfo.IsContradict = "1"
	bytes, _ := json.Marshal(fruitInfo)
	return ctx.GetStub().PutState(fruitInfoID, bytes)
}

func (c *FruitInfoContract) ModifyShareFruitInfo(ctx contractapi.TransactionContextInterface, fruitInfoID string, args []string) error {
	if len(args) != 7 {
		return fmt.Errorf("collect args num should be 7, but now is %d", len(args))
	}
	exists, err := c.FruitInfoExists(ctx, fruitInfoID)
	if err != nil {
		return fmt.Errorf("Could not read from world state. %s", err)
	} else if !exists {
		return fmt.Errorf("The asset %s does not exist", fruitInfoID)
	}

	fruitInfo := new(FruitInfo)
	fruitInfo, _ = c.ReadFruitInfo(ctx, fruitInfoID)
	if fruitInfo.IsLoaded != "0" {
		return fmt.Errorf("could not modify shareInfo, because the info is already loaded")
	}

	var shareInfo ShareInfo
	shareInfo.ShareObj = args[0]
	shareInfo.ContactInfo = args[1]
	shareInfo.ShareMode = args[2]
	shareInfo.ShareUse = args[3]
	shareInfo.ShareNum = args[4]
	shareInfo.ShareBeginTime = args[5]
	shareInfo.ShareEndTime = args[6]
	shareInfo.ShareHash = util.GetSHA256String(args)
	fruitInfo.ShareInfo = shareInfo

	//createTime, _ := ctx.GetStub().GetTxTimestamp()
	//fruitInfo.CreateTime = time.Unix(createTime.GetSeconds(), int64(createTime.GetNanos())).Local().Format("2006-01-02 15:04:05")
	//fruitInfo.TxId = ctx.GetStub().GetTxID()

	fruitInfo.Status = "3"
	fruitInfo.IsContradict = "0"
	bytes, _ := json.Marshal(fruitInfo)
	return ctx.GetStub().PutState(fruitInfoID, bytes)
}

func (c *FruitInfoContract) LoadFruitInfo(ctx contractapi.TransactionContextInterface, fruitInfoID string) error {
	exists, err := c.FruitInfoExists(ctx, fruitInfoID)
	if err != nil {
		return fmt.Errorf("Could not read from world state. %s", err)
	} else if !exists {
		return fmt.Errorf("The asset %s does not exist", fruitInfoID)
	}
	fruitInfo := new(FruitInfo)
	fruitInfo, _ = c.ReadFruitInfo(ctx, fruitInfoID)
	if fruitInfo.IsLoaded != "0" {
		return fmt.Errorf("could not load, because the info is already loaded")
	}
	if fruitInfo.Status != "3" {
		return fmt.Errorf("we expect status from 3 to 4, but the status now is %s", fruitInfo.Status)
	}

	//createTime, _ := ctx.GetStub().GetTxTimestamp()
	//fruitInfo.CreateTime = time.Unix(createTime.GetSeconds(), int64(createTime.GetNanos())).Local().Format("2006-01-02 15:04:05")
	//fruitInfo.TxId = ctx.GetStub().GetTxID()
	fruitInfo.IsLoaded = "1"
	bytes, _ := json.Marshal(fruitInfo)
	//ctx.GetStub().GetTxID()
	return ctx.GetStub().PutState(fruitInfoID, bytes)
}

// SetCollectInfo retrieves an instance of FruitInfo from the world state and set collection information for it
//func (c *FruitInfoContract) SetCollectInfo(ctx contractapi.TransactionContextInterface, fruitInfoID string, args []string) error {
//	if len(args) != 11 {
//		return fmt.Errorf("collect args num should be 11, but now is %d", len(args))
//	}
//	exists, err := c.FruitInfoExists(ctx, fruitInfoID)
//	if err != nil {
//		return fmt.Errorf("Could not read from world state. %s", err)
//	} else if !exists {
//		return fmt.Errorf("The asset %s does not exist", fruitInfoID)
//	}
//	var collectInfo CollectInfo
//	collectInfo.CollectPlaceLongitude = args[0]
//	collectInfo.CollectPlaceLatitude = args[1]
//	collectInfo.CollectPlaceAltitude = args[2]
//	collectInfo.CollectPlaceSoilType = args[3]
//	collectInfo.CollectPlaceEcologyType = args[4]
//	collectInfo.CollectMaterialType = args[5]
//	collectInfo.CollectPeople = args[6]
//	collectInfo.CollectUnit = args[7]
//	collectInfo.CollectTime = args[8]
//	collectInfo.SpeciesName = args[9]
//	collectInfo.Image = args[10]
//	fruitInfo := new(FruitInfo)
//	fruitInfo, _ = c.ReadFruitInfo(ctx, fruitInfoID)
//	fruitInfo.CollectInfo = collectInfo
//	bytes, _ := json.Marshal(fruitInfo)
//	return ctx.GetStub().PutState(fruitInfoID, bytes)
//}

// SetSpeciesInfo retrieves an instance of FruitInfo from the world state and set species information for it
//func (c *FruitInfoContract) SetSpeciesInfo(ctx contractapi.TransactionContextInterface, fruitInfoID string, args []string) error {
//	if len(args) != 10 {
//		return fmt.Errorf("speciesInfo args num should be 10, but now is %d", len(args))
//	}
//	exists, err := c.FruitInfoExists(ctx, fruitInfoID)
//	if err != nil {
//		return fmt.Errorf("could not read from world state. %s", err)
//	} else if !exists {
//		return fmt.Errorf("the asset %s does not exist", fruitInfoID)
//	}
//	var speciesInfo SpeciesInfo
//	speciesInfo.Type = args[0]
//	speciesInfo.Name = args[1]
//	speciesInfo.GermplasmName = args[2]
//	speciesInfo.GermplasmNameEn = args[3]
//	speciesInfo.SectionName = args[4]
//	speciesInfo.GenericName = args[5]
//	speciesInfo.ScientificName = args[6]
//	speciesInfo.ResourceType = args[7]
//	speciesInfo.CollectMethod = args[8]
//	speciesInfo.GermplasmSource = args[9]
//	fruitInfo := new(FruitInfo)
//	fruitInfo, _ = c.ReadFruitInfo(ctx, fruitInfoID)
//	fruitInfo.SpeciesInfo = speciesInfo
//	bytes, _ := json.Marshal(fruitInfo)
//	return ctx.GetStub().PutState(fruitInfoID, bytes)
//}

// SetSourceInfo retrieves an instance of FruitInfo from the world state and set source information for it
//func (c *FruitInfoContract) SetSourceInfo(ctx contractapi.TransactionContextInterface, fruitInfoID string, args []string) error {
//	if len(args) != 6 {
//		return fmt.Errorf("speciesInfo args num should be 6, but now is %d", len(args))
//	}
//	exists, err := c.FruitInfoExists(ctx, fruitInfoID)
//	if err != nil {
//		return fmt.Errorf("could not read from world state. %s", err)
//	} else if !exists {
//		return fmt.Errorf("the asset %s does not exist", fruitInfoID)
//	}
//	var sourceInfo SourceInfo
//	sourceInfo.SourceCountry = args[0]
//	sourceInfo.SourceProvince = args[1]
//	sourceInfo.Source = args[2]
//	sourceInfo.SourceOrg = args[3]
//	sourceInfo.OriginCountry = args[4]
//	sourceInfo.OriginPlace = args[5]
//	fruitInfo := new(FruitInfo)
//	fruitInfo, _ = c.ReadFruitInfo(ctx, fruitInfoID)
//	fruitInfo.SourceInfo = sourceInfo
//	bytes, _ := json.Marshal(fruitInfo)
//	return ctx.GetStub().PutState(fruitInfoID, bytes)
//}

// UpdateProcessID retrieves an instance of FruitInfo from the world state and update its processID
//func (c *FruitInfoContract) UpdateProcessID(ctx contractapi.TransactionContextInterface, fruitInfoID string, processID string) error {
//	exists, err := c.FruitInfoExists(ctx, fruitInfoID)
//	if err != nil {
//		return fmt.Errorf("Could not read from world state. %s", err)
//	} else if !exists {
//		return fmt.Errorf("The asset %s does not exist", fruitInfoID)
//	}
//	fruitInfo := new(FruitInfo)
//	fruitInfo, _ = c.ReadFruitInfo(ctx, fruitInfoID)
//	fruitInfo.ProcessInstanceID = processID
//	bytes, _ := json.Marshal(fruitInfo)
//	return ctx.GetStub().PutState(fruitInfoID, bytes)
//}

// UpdateCollectID retrieves an instance of FruitInfo from the world state and update its collectID
//func (c *FruitInfoContract) UpdateCollectID(ctx contractapi.TransactionContextInterface, fruitInfoID string, collectID string) error {
//	exists, err := c.FruitInfoExists(ctx, fruitInfoID)
//	if err != nil {
//		return fmt.Errorf("Could not read from world state. %s", err)
//	} else if !exists {
//		return fmt.Errorf("The asset %s does not exist", fruitInfoID)
//	}
//	fruitInfo := new(FruitInfo)
//	fruitInfo, _ = c.ReadFruitInfo(ctx, fruitInfoID)
//	fruitInfo.CollectID = collectID
//	bytes, _ := json.Marshal(fruitInfo)
//	return ctx.GetStub().PutState(fruitInfoID, bytes)
//}

// ReadFruitInfo retrieves an instance of FruitInfo from the world state
func (c *FruitInfoContract) ReadFruitInfo(ctx contractapi.TransactionContextInterface, fruitInfoID string) (*FruitInfo, error) {
	exists, err := c.FruitInfoExists(ctx, fruitInfoID)
	if err != nil {
		return nil, fmt.Errorf("Could not read from world state. %s", err)
	} else if !exists {
		return nil, fmt.Errorf("The asset %s does not exist", fruitInfoID)
	}

	bytes, _ := ctx.GetStub().GetState(fruitInfoID)

	fruitInfo := new(FruitInfo)

	err = json.Unmarshal(bytes, fruitInfo)

	if err != nil {
		return nil, fmt.Errorf("Could not unmarshal world state data to type FruitInfo")
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
		fruitInfo := new(FruitInfo)
		err = json.Unmarshal(queryResponse.Value, fruitInfo)
		if err != nil {
			return nil, fmt.Errorf("unmarshl error, %s", err)
		}
		fruitInfos = append(fruitInfos, fruitInfo)
	}
	return fruitInfos, nil
}

// ReadHistory retrieves an instance of FruitInfo from the world state and return its alter history
func (c *FruitInfoContract) ReadHistory(ctx contractapi.TransactionContextInterface, fruitInfoID string) ([]*FruitInfo, error) {
	historyIterator, err := ctx.GetStub().GetHistoryForKey(fruitInfoID)
	if err != nil {
		return nil, fmt.Errorf("can not get history for key, %s", err)
	}
	defer func(historyIterator shim.HistoryQueryIteratorInterface) {
		err := historyIterator.Close()
		if err != nil {
			return
		}
	}(historyIterator)
	var fruitInfos []*FruitInfo
	//var records []HistoryQueryResult
	for historyIterator.HasNext() {
		queryResponse, err := historyIterator.Next()
		if err != nil {
			return nil, fmt.Errorf("ite.next dont exists, %s", err)
		}
		var fruitInfo FruitInfo
		if len(queryResponse.Value) > 0 {
			err = json.Unmarshal(queryResponse.Value, &fruitInfo)
			if err != nil {
				return nil, fmt.Errorf("unmarshl error, %s", err)
			}
		} else {
			fruitInfo = FruitInfo{
				ID:        fruitInfoID,
				IsDeleted: "1",
			}
		}
		//timestamp, err := ptypes.Timestamp(queryResponse.Timestamp)
		//if err != nil {
		//	return nil, err
		//}
		//record := HistoryQueryResult{
		//	TxId:      queryResponse.TxId,
		//	Timestamp: timestamp,
		//	Record:    &fruitInfo,
		//	IsDelete:  queryResponse.IsDelete,
		//}
		//records = append(records, record)
		fruitInfos = append(fruitInfos, &fruitInfo)
	}
	return fruitInfos, nil
}

// DeleteFruitInfo deletes an instance of FruitInfo from the world state
func (c *FruitInfoContract) DeleteFruitInfo(ctx contractapi.TransactionContextInterface, fruitInfoID string) error {
	exists, err := c.FruitInfoExists(ctx, fruitInfoID)
	if err != nil {
		return fmt.Errorf("Could not read from world state. %s", err)
	} else if !exists {
		return fmt.Errorf("The asset %s does not exist", fruitInfoID)
	}
	fruitInfo := new(FruitInfo)
	fruitInfo, _ = c.ReadFruitInfo(ctx, fruitInfoID)
	if fruitInfo.IsLoaded != "0" {
		return fmt.Errorf("you cannot delete info which is already loaded")
	}
	return ctx.GetStub().DelState(fruitInfoID)
}

// ProgrammerDeleteFruitInfo programmer delete the whole info
func (c *FruitInfoContract) ProgrammerDeleteFruitInfo(ctx contractapi.TransactionContextInterface, fruitInfoID string) error {
	exists, err := c.FruitInfoExists(ctx, fruitInfoID)
	if err != nil {
		return fmt.Errorf("Could not read from world state. %s", err)
	} else if !exists {
		return fmt.Errorf("The asset %s does not exist", fruitInfoID)
	}
	return ctx.GetStub().DelState(fruitInfoID)
}
