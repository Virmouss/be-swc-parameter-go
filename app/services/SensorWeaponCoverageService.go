package services

import (
	"be-swc-parameter/app/db"
	"be-swc-parameter/app/db/repository"
	"be-swc-parameter/app/generated/swc"
	"be-swc-parameter/app/model"
	"context"
	"log"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type SensorWeaponCoverageService struct {
	db *gorm.DB
}

func NewSensorWeaponCoverageService() *SensorWeaponCoverageService {

	db, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	return &SensorWeaponCoverageService{db: db}
}

func (s *SensorWeaponCoverageService) AddSWCParameter(ctx context.Context, swc_request *swc.AddSWCParameterReq) (*swc.AddSWCParameterRes, error) {

	saved := model.SensorWeaponCoverage{
		Type:         swc_request.Type,
		GroupSWC:     swc_request.Group,
		Item:         swc_request.Item,
		Environment:  swc_request.Environment,
		Value:        swc_request.Value,
		DefaultValue: swc_request.DefaultSwc,
		Unit:         swc_request.Unit,
		CreatedAt:    time.Now().String(),
		UpdatedAt:    time.Now().String(),
	}

	result := s.db.Model(&model.SensorWeaponCoverage{}).Create(&saved)

	if result.Error != nil {
		log.Fatal(result.Error)
		response := swc.AddSWCParameterRes{
			Id:        saved.ID,
			Message:   "Failed to Save Data, Item Value Already Store in Database",
			CreatedAt: saved.CreatedAt,
		}

		return &response, result.Error
	}
	response := swc.AddSWCParameterRes{
		Id:        saved.ID,
		Message:   "Data saved successfully",
		CreatedAt: saved.CreatedAt,
	}
	return &response, nil
}

func (s *SensorWeaponCoverageService) UpdateSWCParameter(ctx context.Context, swc_request *swc.UpdateSWCParameterReq) (*swc.UpdateSWCParameterRes, error) {

	saved := model.SensorWeaponCoverage{
		ID:          swc_request.Id,
		Type:        swc_request.Type,
		GroupSWC:    swc_request.Group,
		Item:        swc_request.Item,
		Environment: swc_request.Environment,
		Value:       swc_request.Value,
		Unit:        swc_request.Unit,
		UpdatedAt:   time.Now().String(),
		IsDeleted:   false,
	}

	result := s.db.Model(&model.SensorWeaponCoverage{}).Where("id = ?", saved.ID).Updates(&saved)

	if result.Error != nil {
		//log.Fatal(result.Error)
		response := swc.UpdateSWCParameterRes{
			Id:        saved.ID,
			Message:   "Failed to Save Data, Item Value Already Store in Database",
			UpdatedAt: saved.UpdatedAt,
		}

		return &response, result.Error
	}
	response := swc.UpdateSWCParameterRes{
		Id:        saved.ID,
		Message:   "Data saved successfully",
		UpdatedAt: saved.UpdatedAt,
	}
	return &response, nil
}

func (s *SensorWeaponCoverageService) DeleteSWCParameter(ctx context.Context, swc_request *swc.DeleteSWCParameterReq) (*swc.DeleteSWCParameterRes, error) {

	result := repository.DeleteByIds(s.db, true, swc_request.Id)

	if result != nil {
		log.Fatal(result)
		response := swc.DeleteSWCParameterRes{
			Message: "Failed to Delete Data",
		}

		return &response, result
	}
	response := swc.DeleteSWCParameterRes{
		Message: "Data Delete successfully",
	}
	return &response, nil
}

func (s *SensorWeaponCoverageService) DeleteAllSWCParameter(ctx context.Context, swc_request *swc.EmptyRequest) (*swc.DeleteAllSWCParameterRes, error) {
	result := repository.DeleteAll(s.db, true)

	if result != nil {
		log.Fatal(result)
		response := swc.DeleteAllSWCParameterRes{
			Message: "Failed to Delete Data",
		}

		return &response, result
	}
	response := swc.DeleteAllSWCParameterRes{
		Message: "Data Delete successfully",
	}
	return &response, nil
}

func (s *SensorWeaponCoverageService) GetAllSWCParameter(ctx context.Context, swc_requestuest *swc.GetAllSWCParameterReq) (*swc.GetAllSWCParameterRes, error) {

	pageNumber := swc_requestuest.PageNumber
	pageSize := swc_requestuest.PageSize
	group := swc_requestuest.Group
	field := swc_requestuest.Field
	ascending := swc_requestuest.Ascending

	ascendingStr := strconv.FormatBool(ascending)

	if field != "" {
		if field == "default_swc" {
			field = "defaultValue"
		}
		if field == "type" {
			field = "type"
		}
	}

	var usersPage []model.SensorWeaponCoverage
	var swcData []*swc.GetAllSWCParameterRes_SWCParameterData
	var count int64
	if group != "" {
		if ascendingStr != "" {
			usersPage = repository.FindAllByParamsSort(s.db, group, pageNumber, pageSize, ascending)
			count = repository.GetCountDataByGroup(s.db, group)
		} else {
			usersPage = repository.FindAllByParams(s.db, group, pageNumber, pageSize)
			count = repository.GetCountDataByGroup(s.db, group)
		}
	} else {
		usersPage = repository.FindAll(s.db, pageNumber, pageSize)
		count = repository.GetCountDataAll(s.db)
	}

	for _, result := range usersPage {
		if !result.IsDeleted {
			swcData = append(swcData, &swc.GetAllSWCParameterRes_SWCParameterData{
				Id:          result.ID,
				Type:        result.Type,
				Group:       result.GroupSWC,
				Item:        result.Item,
				Environment: result.Environment,
				Value:       result.Value,
				DefaultSwc:  result.DefaultValue,
				Unit:        result.Unit,
				UpdatedAt:   result.UpdatedAt,
				CreatedAt:   result.CreatedAt,
			})
		}
	}
	response := &swc.GetAllSWCParameterRes{
		SwcParameterData: swcData,
		TotalData:        int32(count),
	}
	return response, nil
}

// func (s *SensorWeaponCoverageService) GetCountAllData() int {
// 	var totalData int

// 	datas := repository.GetSWCCount(s.db)

// 	for _, result := range datas {
// 		item := result.Item
// 		items := strings.Split(item, ",")

// 		if len(items) == 4 {
// 			totalData++
// 		}
// 	}
// 	return totalData

// }
