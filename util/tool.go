package util

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"strings"
)

// WriteLedger 写入账本
func WriteLedger(obj interface{}, ctx contractapi.TransactionContextInterface, objectType string, keys []string) error {
	//创建复合主键
	var key string
	if val, err := ctx.GetStub().CreateCompositeKey(objectType, keys); err != nil {
		return fmt.Errorf("%s-创建复合主键出错 %s", objectType, err)
	} else {
		key = val
	}
	bytes, err := json.Marshal(obj)
	if err != nil {
		return fmt.Errorf("%s-序列化json数据失败出错: %s", objectType, err)
	}
	//写入区块链账本
	if err := ctx.GetStub().PutState(key, bytes); err != nil {
		return fmt.Errorf("%s-写入区块链账本出错: %s", objectType, err)
	}
	return nil
}

// DelLedger 删除账本
func DelLedger(ctx contractapi.TransactionContextInterface, objectType string, keys []string) error {
	//创建复合主键
	var key string
	if val, err := ctx.GetStub().CreateCompositeKey(objectType, keys); err != nil {
		return fmt.Errorf("%s-创建复合主键出错 %s", objectType, err)
	} else {
		key = val
	}
	//写入区块链账本
	if err := ctx.GetStub().DelState(key); err != nil {
		return fmt.Errorf("%s-删除区块链账本出错: %s", objectType, err)
	}
	return nil
}

// GetStateByPartialCompositeKeys 根据复合主键查询数据(适合获取全部，多个，单个数据)
// 将keys拆分查询
func GetStateByPartialCompositeKeys(ctx contractapi.TransactionContextInterface, objectType string, keys []string) (results [][]byte, err error) {
	if len(keys) == 0 {
		// 传入的keys长度为0，则查找并返回所有数据
		// 通过主键从区块链查找相关的数据，相当于对主键的模糊查询
		resultIterator, err := ctx.GetStub().GetStateByPartialCompositeKey(objectType, keys)
		if err != nil {
			return nil, fmt.Errorf("%s-获取全部数据出错: %s", objectType, err)
		}
		defer resultIterator.Close()

		//检查返回的数据是否为空，不为空则遍历数据，否则返回空数组
		for resultIterator.HasNext() {
			val, err := resultIterator.Next()
			if err != nil {
				return nil, fmt.Errorf("%s-返回的数据出错: %s", objectType, err)
			}

			results = append(results, val.GetValue())
		}
	} else {
		// 传入的keys长度不为0，查找相应的数据并返回
		for _, v := range keys {
			// 创建组合键
			key, err := ctx.GetStub().CreateCompositeKey(objectType, []string{v})
			if err != nil {
				return nil, fmt.Errorf("%s-创建组合键出错: %s", objectType, err)
			}
			// 从账本中获取数据
			bytes, err := ctx.GetStub().GetState(key)
			if err != nil {
				return nil, fmt.Errorf("%s-获取数据出错: %s", objectType, err)
			}

			if bytes != nil {
				results = append(results, bytes)
			}
		}
	}

	return results, nil
}

// GetStateByPartialCompositeKeys2 根据复合主键查询数据(适合获取全部或指定的数据)
func GetStateByPartialCompositeKeys2(ctx contractapi.TransactionContextInterface, objectType string, keys []string) (results [][]byte, err error) {
	// 通过主键从区块链查找相关的数据，相当于对主键的模糊查询
	resultIterator, err := ctx.GetStub().GetStateByPartialCompositeKey(objectType, keys)
	if err != nil {
		return nil, fmt.Errorf("%s-获取全部数据出错: %s", objectType, err)
	}
	defer resultIterator.Close()

	//检查返回的数据是否为空，不为空则遍历数据，否则返回空数组
	for resultIterator.HasNext() {
		val, err := resultIterator.Next()
		if err != nil {
			return nil, fmt.Errorf("%s-返回的数据出错: %s", objectType, err)
		}

		results = append(results, val.GetValue())
	}
	return results, nil
}

// func GetMD5String(content string) (result string) {
// 	hash := md5.New()
// 	hash.Write([]byte(content))
// 	return hex.EncodeToString(hash.Sum(nil))
// }

// func GetSHA1String(content string) (result string) {
// 	hash := sha1.New()
// 	hash.Write([]byte(content))
// 	return hex.EncodeToString(hash.Sum(nil))
// }

func GetSHA256String(content []string) (result string) {
	hash := sha256.New()
	hash.Write([]byte(strings.Join(content, "")))
	return hex.EncodeToString(hash.Sum(nil))
}
