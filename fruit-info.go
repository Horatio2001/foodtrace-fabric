/*
 * SPDX-License-Identifier: Apache-2.0
 */

package main

// FruitInfo stores a value
type FruitInfo struct {
	ID           string      `json:"id"`            // id
	Status       string      `json:"status"`        // 当前状态 0:待保存 1:待录入 2：待共享 3：待上链
	CreateTime   string      `json:"create_time"`   // 创建时间
	TxId         string      `json:"tx_id"`         // 交易id
	IsContradict string      `json:"is_contradict"` // 是否被驳回 0:normal 1：被驳回
	IsDeleted    string      `json:"is_deleted"`    // 是否被删除 0:未删除 1：已删除
	IsLoaded     string      `json:"is_loaded"`     // 是否存证 0：否 1：已存证
	CollectInfo  CollectInfo `json:"collect_info"`  // 采集信息
	SaveInfo     SaveInfo    `json:"save_info"`     // 保存信息
	EnterInfo    EnterInfo   `json:"enter_info"`    // 录入信息
	ShareInfo    ShareInfo   `json:"share_info"`    // 共享信息
}

type CollectInfo struct {
	CollectID               string `json:"collect_id"`                 // 收集号（自动生成
	Type                    string `json:"type"`                       // 果树类型
	Name                    string `json:"name"`                       // 名称
	GermplasmName           string `json:"germplasm_name"`             // 种质名称
	GermplasmNameEn         string `json:"germplasm_Name_En"`          // 种质外文名
	SectionName             string `json:"section_name"`               // 科名
	GenericName             string `json:"generic_name"`               // 属名
	ScientificName          string `json:"scientific_name"`            // 学名
	ResourceType            string `json:"resource_type"`              // 种质资源类型
	CollectMethod           string `json:"collect_method"`             // 采集方式
	GermplasmSource         string `json:"germplasm_source"`           // 种质来源
	SourceCountry           string `json:"source_country"`             // 来源国
	SourceProvince          string `json:"source_province"`            // 来源省
	Source                  string `json:"source"`                     // 来源地
	SourceOrg               string `json:"source_org"`                 // 来源机构
	OriginCountry           string `json:"origin_country"`             // 原产国
	OriginPlace             string `json:"origin_place"`               // 原产地
	CollectPlaceLongitude   string `json:"collect_place_longitude"`    // 经度
	CollectPlaceLatitude    string `json:"collect_place_latitude"`     // 纬度
	CollectPlaceAltitude    string `json:"collect_place_altitude"`     // 海拔
	CollectPlaceSoilType    string `json:"collect_place_soil_type"`    // 土壤类型
	CollectPlaceEcologyType string `json:"collect_place_ecology_type"` // 生态类型
	CollectMaterialType     string `json:"collect_material_type"`      // 材料类型
	CollectPeople           string `json:"collect_people"`             // 采集人
	CollectUnit             string `json:"collect_unit"`               // 采集单位
	CollectTime             string `json:"collect_time"`               // 采集时间
	SpeciesName             string `json:"species_name"`               // 项目归口
	Image                   string `json:"image"`                      // 原生境图片
	CollectRemark           string `json:"collect_remark"`             // 备注
	CollectHash             string `json:"collect_hash"`               // 收集哈希
}

type SaveInfo struct {
	MainPreference       string `json:"main_preference"`       //主要特征
	MainUse              string `json:"main_use"`              //主要用途
	PreservationFacility string `json:"preservation_facility"` //保存设施
	GermplasmType        string `json:"germplasm_type"`        //种质类型
	SaveQuantity         string `json:"save_quantity"`         //保存数量
	MeasuringUnit        string `json:"measuring_unit"`        //计量单位
	SaveUnit             string `json:"save_unit"`             //保存单位
	SaveVault            string `json:"save_vault"`            //保存库
	SavePlace            string `json:"save_place"`            //保存地点
	WarehousingYear      string `json:"ware_housing_year"`     //入库年份
	SaveProperty         string `json:"save_property"`         //保存性质
	ResourceDescription  string `json:"resource_description"`  //资源描述
	ResourceRemark       string `json:"resource_remark"`       //资源备注
	GermplasmImage       string `json:"germplasm_image"`       //种质图片
	SaveHash             string `json:"save_hash"`             //保存哈希
}

type EnterInfo struct {
	Certifier      string `json:"certifier"`       //鉴定人
	CertifyOrg     string `json:"certify_org"`     //鉴定机构
	CertifyPlace   string `json:"certify_place"`   //鉴定地点
	CertifyYear    string `json:"certify_year"`    //鉴定年份
	OperationRange string `json:"operation_range"` //操作范围
	EnterRemark    string `json:"enter_remark"`    //录入备注
	EnterHash      string `json:"enter_hash"`      //录入哈希
}

type ShareInfo struct {
	ShareObj       string `json:"share_obj"`        //共享对象
	ContactInfo    string `json:"contact_info"`     //联系方式
	ShareMode      string `json:"share_mode"`       //共享方式
	ShareUse       string `json:"share_use"`        //共享用语
	ShareNum       string `json:"share_num"`        //共享份次
	ShareBeginTime string `json:"share_begin_time"` //共享开始时间
	ShareEndTime   string `json:"share_end_time"`   //共享结束时间
	ShareHash      string `json:"share_hash"`       //共享哈希
}

//type SpeciesInfo struct {
//	Type            string `json:"type"`              // 果树类型
//	Name            string `json:"name"`              // 名称
//	GermplasmName   string `json:"germplasm_name"`    // 种质名称
//	GermplasmNameEn string `json:"germplasm_Name_En"` // 种质外文名
//	SectionName     string `json:"section_name"`      // 科名
//	GenericName     string `json:"generic_name"`      // 属名
//	ScientificName  string `json:"scientific_name"`   // 学名
//	ResourceType    string `json:"resource_type"`     // 种质资源类型
//	CollectMethod   string `json:"collect_method"`    // 采集方式
//	GermplasmSource string `json:"germplasm_source"`  // 种质资源
//}
//
//type SourceInfo struct {
//	SourceCountry  string `json:"source_country"`  // 前产国
//	SourceProvince string `json:"source_province"` // 前产地
//	Source         string `json:"source"`          // 来源
//	SourceOrg      string `json:"source_org"`      // 来源组织
//	OriginCountry  string `json:"origin_country"`  // 原产国
//	OriginPlace    string `json:"origin_place"`    // 原产地
//}
//
//type CollectInfo struct {
//	CollectPlaceLongitude   string `json:"collect_place_longitude"`    // 经度
//	CollectPlaceLatitude    string `json:"collect_place_latitude"`     // 纬度
//	CollectPlaceAltitude    string `json:"collect_place_altitude"`     // 海拔
//	CollectPlaceSoilType    string `json:"collect_place_soil_type"`    // 土壤类型
//	CollectPlaceEcologyType string `json:"collect_place_ecology_type"` // 生态类型
//	CollectMaterialType     string `json:"collect_material_type"`      // 材料类型
//	CollectPeople           string `json:"collect_people"`             // 采集人
//	CollectUnit             string `json:"collect_unit"`               // 采集单位
//	CollectTime             string `json:"collect_time"`               // 采集时间
//	SpeciesName             string `json:"species_name"`               // 物种名称
//	Image                   string `json:"image"`                      // 图片
//}
//
//// HistoryQueryResult structure used for returning result of history query
//type HistoryQueryResult struct {
//	Record    *FruitInfo `json:"record"`
//	TxId      string     `json:"txId"`
//	Timestamp time.Time  `json:"timestamp"`
//	IsDelete  bool       `json:"isDelete"`
//}
