/*
 * SPDX-License-Identifier: Apache-2.0
 */

package main

// FruitInfo stores a value
type FruitInfo struct {
	ID                string          `json:"id"`                  // id
	ProcessInstanceID string          `json:"process_instance_id"` // 实例id
	CollectID         string          `json:"collect_id"`          // 采集号
	SpeciesInfo       SpeciesInfo     `json:"species_info"`        // 物种信息
	SourceInfo        SourceInfo      `json:"source_info"`         // 来源信息
	CollectInfo       CollectInfo     `json:"collect_info"`        // 采集信息
	TransportInfo     []TransportInfo `json:"transport_info"`      //  物流信息
}

type SpeciesInfo struct {
	Type            string `json:"type"`              // 果树类型
	Name            string `json:"name"`              // 名称
	GermplasmName   string `json:"germplasm_name"`    // 种质名称
	GermplasmNameEn string `json:"germplasm_Name_En"` // 种质名
	SectionName     string `json:"section_name"`      // 科名
	GenericName     string `json:"generic_name"`      // 属名
	ScientificName  string `json:"scientific_name"`   // 学名
	ResourceType    string `json:"resource_type"`     // 产地类型
	CollectMethod   string `json:"collect_method"`    // 采集方式
	GermplasmSource string `json:"germplasm_source"`  // 种质资源
}

type SourceInfo struct {
	SourceCountry  string `json:"source_country"`  // 前产国
	SourceProvince string `json:"source_province"` // 前产地
	Source         string `json:"source"`          // 来源
	SourceOrg      string `json:"source_org"`      // 来源组织
	OriginCountry  string `json:"origin_country"`  // 原产国
	OriginPlace    string `json:"origin_place"`    // 原产地
}

type CollectInfo struct {
	CollectPlaceLongitude   string `json:"collect_place_longitude"`    // 经度
	CollectPlaceLatitude    string `json:"collect_place_latitude"`     // 纬度
	CollectPlaceAltitude    string `json:"collect_place_altitude"`     // 海拔
	CollectPlaceSoilType    string `json:"collect_place_soil_type"`    // 土壤类型
	CollectPlaceEcologyType string `json:"collect_place_ecology_type"` // 生态类型
	CollectMaterialType     string `json:"collect_material_type"`      // 材料类型
	CollectPeople           string `json:"collect_people"`             // 采集人
	CollectUnit             string `json:"collect_unit"`               // 采集单位
	CollectTime             string `json:"collect_time"`               // 采集时间
	SpeciesName             string `json:"species_name"`               // 物种名称
	Image                   string `json:"image"`                      // 图片
}

type TransportInfo struct {
	TransportDepartureTime  string `json:"transport_departure_time"`  // 离开时间
	TransportArrivalTime    string `json:"transport_arrival_time"`    // 到达时间
	TransportMission        string `json:"transport_mission"`         // 运输任务（运输中 or 运输完成）
	TransportDeparturePlace string `json:"transport_departure_place"` // 运输起点
	TransportDestination    string `json:"transport_destination"`     // 运输终点
	TransportMethod         string `json:"transport_method"`          // 运输方式
	TransportDepartmentName string `json:"transport_department_name"` // 物流公司名称
	TransporterName         string `json:"transporter_name"`          // 负责人姓名
}
