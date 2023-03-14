/*
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const getStateError = "world state get error"

type MockStub struct {
	shim.ChaincodeStubInterface
	mock.Mock
}

func (ms *MockStub) GetState(key string) ([]byte, error) {
	args := ms.Called(key)

	return args.Get(0).([]byte), args.Error(1)
}

func (ms *MockStub) PutState(key string, value []byte) error {
	args := ms.Called(key, value)

	return args.Error(0)
}

func (ms *MockStub) DelState(key string) error {
	args := ms.Called(key)

	return args.Error(0)
}

type MockContext struct {
	contractapi.TransactionContextInterface
	mock.Mock
}

func (mc *MockContext) GetStub() shim.ChaincodeStubInterface {
	args := mc.Called()

	return args.Get(0).(*MockStub)
}

func configureStub() (*MockContext, *MockStub) {
	var nilBytes []byte

	testFruitInfo := new(FruitInfo)
	testFruitInfo.ID = "set value"
	fruitInfoBytes, _ := json.Marshal(testFruitInfo)

	ms := new(MockStub)
	ms.On("GetState", "statebad").Return(nilBytes, errors.New(getStateError))
	ms.On("GetState", "missingkey").Return(nilBytes, nil)
	ms.On("GetState", "existingkey").Return([]byte("some value"), nil)
	ms.On("GetState", "fruitInfokey").Return(fruitInfoBytes, nil)
	ms.On("PutState", mock.AnythingOfType("string"), mock.AnythingOfType("[]uint8")).Return(nil)
	ms.On("DelState", mock.AnythingOfType("string")).Return(nil)

	mc := new(MockContext)
	mc.On("GetStub").Return(ms)

	return mc, ms
}

func TestFruitInfoExists(t *testing.T) {
	var exists bool
	var err error

	ctx, _ := configureStub()
	c := new(FruitInfoContract)

	exists, err = c.FruitInfoExists(ctx, "statebad")
	assert.EqualError(t, err, getStateError)
	assert.False(t, exists, "should return false on error")

	exists, err = c.FruitInfoExists(ctx, "missingkey")
	assert.Nil(t, err, "should not return error when can read from world state but no value for key")
	assert.False(t, exists, "should return false when no value for key in world state")

	exists, err = c.FruitInfoExists(ctx, "existingkey")
	assert.Nil(t, err, "should not return error when can read from world state and value exists for key")
	assert.True(t, exists, "should return true when value for key in world state")
}

func TestCreateFruitInfo(t *testing.T) {
	var err error
	var fruitInfo FruitInfo
	fruitInfo.ID = "missingkey"
	fruitInfo.CollectID = "some value"
	fruitInfo.ProcessInstanceID = "some value"
	bytes, _ := json.Marshal(fruitInfo)

	ctx, stub := configureStub()
	c := new(FruitInfoContract)

	err = c.CreateFruitInfo(ctx, "statebad", "some value", "some value")
	assert.EqualError(t, err, fmt.Sprintf("Could not read from world state. %s", getStateError), "should error when exists errors")

	err = c.CreateFruitInfo(ctx, "existingkey", "some value", "some value")
	assert.EqualError(t, err, "The asset existingkey already exists", "should error when exists returns true")

	err = c.CreateFruitInfo(ctx, "missingkey", "some value", "some value")
	stub.AssertCalled(t, "PutState", "missingkey", bytes)
}

func TestReadFruitInfo(t *testing.T) {
	var fruitInfo *FruitInfo
	var err error

	ctx, _ := configureStub()
	c := new(FruitInfoContract)

	fruitInfo, err = c.ReadFruitInfo(ctx, "statebad")
	assert.EqualError(t, err, fmt.Sprintf("Could not read from world state. %s", getStateError), "should error when exists errors when reading")
	assert.Nil(t, fruitInfo, "should not return FruitInfo when exists errors when reading")

	fruitInfo, err = c.ReadFruitInfo(ctx, "missingkey")
	assert.EqualError(t, err, "The asset missingkey does not exist", "should error when exists returns true when reading")
	assert.Nil(t, fruitInfo, "should not return FruitInfo when key does not exist in world state when reading")

	fruitInfo, err = c.ReadFruitInfo(ctx, "existingkey")
	assert.EqualError(t, err, "Could not unmarshal world state data to type FruitInfo", "should error when data in key is not FruitInfo")
	assert.Nil(t, fruitInfo, "should not return FruitInfo when data in key is not of type FruitInfo")

	fruitInfo, err = c.ReadFruitInfo(ctx, "fruitInfokey")
	expectedFruitInfo := new(FruitInfo)
	expectedFruitInfo.ID = "set value"
	assert.Nil(t, err, "should not return error when FruitInfo exists in world state when reading")
	assert.Equal(t, expectedFruitInfo, fruitInfo, "should return deserialized FruitInfo from world state")
}

func TestFruitInfoContract_SetCollectInfo(t *testing.T) {
	var err error

	ctx, stub := configureStub()
	c := new(FruitInfoContract)
	var collectString = []string{"", "", "",
		"6", "4", "3", "horatio",
		"中国热带农业科学院热带作物品种资源研究所", "2020-01-01T00:00:00.000Z", "null", ""}
	var collectInfo = CollectInfo{"", "", "",
		"6", "4", "3", "horatio",
		"中国热带农业科学院热带作物品种资源研究所", "2020-01-01T00:00:00.000Z", "null", ""}

	err = c.SetCollectInfo(ctx, "statebad", collectString)
	assert.EqualError(t, err, fmt.Sprintf("Could not read from world state. %s", getStateError), "should error when exists errors when updating")

	err = c.SetCollectInfo(ctx, "missingkey", collectString)
	assert.EqualError(t, err, "The asset missingkey does not exist", "should error when exists returns true when updating")

	err = c.SetCollectInfo(ctx, "fruitInfokey", collectString)
	expectedFruitInfo := new(FruitInfo)
	expectedFruitInfo.CollectInfo = collectInfo
	expectedFruitInfoBytes, _ := json.Marshal(expectedFruitInfo)
	assert.Nil(t, err, "should not return error when FruitInfo exists in world state when updating")
	stub.AssertCalled(t, "PutState", "fruitInfokey", expectedFruitInfoBytes)
}

func TestFruitInfoContract_SetSourceInfo(t *testing.T) {
	var err error

	ctx, stub := configureStub()
	c := new(FruitInfoContract)
	var sourceInfo = []string{"中国", "广西", "", "", "", ""}

	err = c.SetSourceInfo(ctx, "statebad", sourceInfo)
	assert.EqualError(t, err, fmt.Sprintf("Could not read from world state. %s", getStateError), "should error when exists errors when updating")

	err = c.SetSourceInfo(ctx, "missingkey", sourceInfo)
	assert.EqualError(t, err, "The asset missingkey does not exist", "should error when exists returns true when updating")

	err = c.SetSourceInfo(ctx, "fruitInfokey", sourceInfo)
	expectedFruitInfo := new(FruitInfo)
	expectedFruitInfo.ID = "new value"
	expectedFruitInfoBytes, _ := json.Marshal(expectedFruitInfo)
	assert.Nil(t, err, "should not return error when FruitInfo exists in world state when updating")
	stub.AssertCalled(t, "PutState", "fruitInfokey", expectedFruitInfoBytes)
}

func TestFruitInfoContract_SetSpeciesInfo(t *testing.T) {
	var err error

	ctx, stub := configureStub()
	c := new(FruitInfoContract)
	var sourceInfo = []string{"产胶作物", "橡胶树", "AC/T/4-13/23", "", "Euphobiaceae（大戟科）",
		"Hevea（橡胶树）", "H. brasiliensis", "野生资源", "引种", "1"}

	err = c.SetSourceInfo(ctx, "statebad", sourceInfo)
	assert.EqualError(t, err, fmt.Sprintf("Could not read from world state. %s", getStateError), "should error when exists errors when updating")

	err = c.SetSourceInfo(ctx, "missingkey", sourceInfo)
	assert.EqualError(t, err, "The asset missingkey does not exist", "should error when exists returns true when updating")

	err = c.SetSourceInfo(ctx, "fruitInfokey", sourceInfo)
	expectedFruitInfo := new(FruitInfo)
	expectedFruitInfo.ID = "new value"
	expectedFruitInfoBytes, _ := json.Marshal(expectedFruitInfo)
	assert.Nil(t, err, "should not return error when FruitInfo exists in world state when updating")
	stub.AssertCalled(t, "PutState", "fruitInfokey", expectedFruitInfoBytes)
}

func TestFruitInfoContract_AddTransportInfo(t *testing.T) {
	var err error

	ctx, stub := configureStub()
	c := new(FruitInfoContract)
	var transportInfo = []string{"2023-02-10T00:00:00.000Z", "2023-02-12T00:00:00.000Z", "0", "广西", "北京",
		"空运", "新华物流", "孙一峰"}

	err = c.AddTransportInfo(ctx, "statebad", transportInfo)
	assert.EqualError(t, err, fmt.Sprintf("Could not read from world state. %s", getStateError), "should error when exists errors when updating")

	err = c.AddTransportInfo(ctx, "missingkey", transportInfo)
	assert.EqualError(t, err, "The asset missingkey does not exist", "should error when exists returns true when updating")

	err = c.AddTransportInfo(ctx, "fruitInfokey", transportInfo)
	expectedFruitInfo := new(FruitInfo)
	expectedFruitInfo.ID = "new value"
	expectedFruitInfoBytes, _ := json.Marshal(expectedFruitInfo)
	assert.Nil(t, err, "should not return error when FruitInfo exists in world state when updating")
	stub.AssertCalled(t, "PutState", "fruitInfokey", expectedFruitInfoBytes)
}

func TestFruitInfoContract_UpdateCollectID(t *testing.T) {
	var err error

	ctx, stub := configureStub()
	c := new(FruitInfoContract)

	err = c.UpdateCollectID(ctx, "statebad", "new value")
	assert.EqualError(t, err, fmt.Sprintf("Could not read from world state. %s", getStateError), "should error when exists errors when updating")

	err = c.UpdateCollectID(ctx, "missingkey", "new value")
	assert.EqualError(t, err, "The asset missingkey does not exist", "should error when exists returns true when updating")

	err = c.UpdateCollectID(ctx, "fruitInfokey", "new value")
	expectedFruitInfo := new(FruitInfo)
	expectedFruitInfo.CollectID = "new value"
	expectedFruitInfoBytes, _ := json.Marshal(expectedFruitInfo)
	assert.Nil(t, err, "should not return error when FruitInfo exists in world state when updating")
	stub.AssertCalled(t, "PutState", "fruitInfokey", expectedFruitInfoBytes)
}

func TestFruitInfoContract_UpdateProcessID(t *testing.T) {
	var err error

	ctx, stub := configureStub()
	c := new(FruitInfoContract)

	err = c.UpdateProcessID(ctx, "statebad", "new value")
	assert.EqualError(t, err, fmt.Sprintf("Could not read from world state. %s", getStateError), "should error when exists errors when updating")

	err = c.UpdateProcessID(ctx, "missingkey", "new value")
	assert.EqualError(t, err, "The asset missingkey does not exist", "should error when exists returns true when updating")

	err = c.UpdateProcessID(ctx, "fruitInfokey", "new value")
	expectedFruitInfo := new(FruitInfo)
	expectedFruitInfo.ProcessInstanceID = "new value"
	expectedFruitInfoBytes, _ := json.Marshal(expectedFruitInfo)
	assert.Nil(t, err, "should not return error when FruitInfo exists in world state when updating")
	stub.AssertCalled(t, "PutState", "fruitInfokey", expectedFruitInfoBytes)
}

func TestDeleteFruitInfo(t *testing.T) {
	var err error

	ctx, stub := configureStub()
	c := new(FruitInfoContract)

	err = c.DeleteFruitInfo(ctx, "statebad")
	assert.EqualError(t, err, fmt.Sprintf("Could not read from world state. %s", getStateError), "should error when exists errors")

	err = c.DeleteFruitInfo(ctx, "missingkey")
	assert.EqualError(t, err, "The asset missingkey does not exist", "should error when exists returns true when deleting")

	err = c.DeleteFruitInfo(ctx, "fruitInfokey")
	assert.Nil(t, err, "should not return error when FruitInfo exists in world state when deleting")
	stub.AssertCalled(t, "DelState", "fruitInfokey")
}
