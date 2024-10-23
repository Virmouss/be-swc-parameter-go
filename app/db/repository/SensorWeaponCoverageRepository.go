package repository

import (
	"be-swc-parameter/app/model"
	"strings"

	"gorm.io/gorm"
)

func DeleteByIds(db *gorm.DB, newValue bool, ids []int32) error {

	query := "UPDATE swc_parameter SET is_deleted = ? WHERE id IN (?)"

	err := db.Exec(query, newValue, ids).Error

	if err != nil {
		return err
	}

	return nil
}

func DeleteAll(db *gorm.DB, newValue bool) error {

	query := "UPDATE swc_parameter SET is_deleted = ?"

	err := db.Exec(query, newValue).Error

	if err != nil {
		return err
	}

	return nil
}

func FindAll(db *gorm.DB, page, pageSize int32) []model.SensorWeaponCoverage {
	var results []model.SensorWeaponCoverage

	offset := (page - 1) * pageSize

	err := db.Where("is_deleted = ? AND group_swc != ''", false).
		Offset(int(offset)).
		Limit(int(pageSize)).
		Find(&results).Error

	if err != nil {
		return nil
	}
	return results
}

func FindAllByParams(db *gorm.DB, pattern string, page, pageSize int32) []model.SensorWeaponCoverage {
	var results []model.SensorWeaponCoverage

	searchPattern := "%" + strings.ToLower(pattern) + "%"
	offset := (page - 1) * pageSize

	err := db.Where("LOWER(group_swc) LIKE ? AND is_deleted = ? AND group_swc != ''", searchPattern, false).
		Offset(int(offset)).
		Limit(int(pageSize)).
		Find(&results).Error

	if err != nil {
		return nil
	}
	return results
}

func FindAllByParamsSort(db *gorm.DB, pattern string, page, pageSize int32, ascending bool) []model.SensorWeaponCoverage {
	var results []model.SensorWeaponCoverage

	searchPattern := "%" + strings.ToLower(pattern) + "%"
	offset := (page - 1) * pageSize

	order := "group_swc ASC"
	if !ascending {
		order = "group_swc DESC"
	}
	err := db.Where("LOWER(group_swc) LIKE ? AND is_deleted = ? AND group_swc != ''", searchPattern, false).
		Offset(int(offset)).
		Limit(int(pageSize)).
		Order(order).
		Find(&results).Error

	if err != nil {
		return nil
	}
	return results
}

func GetCountData(db *gorm.DB, group, searchCriteria, envi string) (int64, error) {
	var count int64

	err := db.Model(&model.SensorWeaponCoverage{}).
		Where("group_swc = ? AND group_swc != '' AND item = ? AND environment = ? AND is_deleted = ?", group, searchCriteria, envi, false).
		Count(&count).Error

	if err != nil {
		return 0, err
	}

	return count, nil
}

func GetCountDataAll(db *gorm.DB) int64 {
	var count int64

	result := db.Model(&model.SensorWeaponCoverage{}).
		Where("is_deleted = ? AND group_swc != ''", false).
		Count(&count)

	if result.Error != nil {
		return 0
	}

	return count

}

func GetCountDataByGroup(db *gorm.DB, group string) int64 {
	var count int64

	searchPattern := "%" + strings.ToLower(group) + "%"

	result := db.Model(&model.SensorWeaponCoverage{}).
		Where("LOWER(group_swc) LIKE ? AND is_deleted = ? AND group_swc != ''", searchPattern, false).
		Count(&count)

	if result.Error != nil {
		return 0
	}

	return count
}
