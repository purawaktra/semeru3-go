package modules

import (
	"fmt"
	"github.com/purawaktra/semeru3-go/entities"
	"github.com/purawaktra/semeru3-go/utils"
	"gorm.io/gorm"
)

type Semeru3Repo struct {
	db *gorm.DB
}

func CreateSemeru3Repo(gorm *gorm.DB) Semeru3Repo {
	return Semeru3Repo{
		db: gorm,
	}
}

func (sr Semeru3Repo) SelectProfileById(query entities.Profile) (entities.Profile, error, string) {
	utils.Debug("SelectProfileById", query)

	var profile entities.Profile
	tx := sr.db.Raw(
		fmt.Sprint("SELECT account_id, first_name, last_name, address, city, province, zipcode, phone_number FROM profiles WHERE account_id = ", query.AccountId, " LIMIT 1")).Scan(
		&profile)
	err := tx.Error
	if err != nil {
		utils.Error(err, "SelectProfileById", query)
		return entities.Profile{}, err, "DB"
	}

	return profile, nil, "00"
}

func (sr Semeru3Repo) InsertProfile(query entities.Profile) (error, string) {
	utils.Debug("InsertProfile", query)

	tx := sr.db.Exec(
		"INSERT INTO profiles (account_id, first_name, last_name, address, city, province, zipcode, phone_number) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		query.AccountId,
		query.FirstName,
		query.LastName,
		query.Address,
		query.City,
		query.Province,
		query.Zipcode,
		query.PhoneNumber)
	err := tx.Error
	if err != nil {
		utils.Error(err, "InsertProfile", query)
		return err, "DB"
	}
	return nil, "00"
}

func (sr Semeru3Repo) UpdateProfileById(query entities.Profile) (error, string) {
	utils.Debug("UpdateProfileById", query)

	tx := sr.db.Exec(
		"UPDATE profiles SET first_name = ?, last_name = ?, address = ?, city = ?, province = ?, zipcode = ?, phone_number = ? WHERE account_id = ?",
		query.FirstName,
		query.LastName,
		query.Address,
		query.City,
		query.Province,
		query.Zipcode,
		query.PhoneNumber,
		query.AccountId)
	err := tx.Error
	if err != nil {
		utils.Error(err, "UpdateProfileById", query)
		return err, "DB"
	}
	return nil, "00"
}

func (sr Semeru3Repo) DeleteProfileById(query entities.Profile) (error, string) {
	utils.Debug("DeleteProfileById", query)

	tx := sr.db.Exec(
		"DELETE FROM profiles WHERE account_id = ?",
		query.AccountId)
	err := tx.Error
	if err != nil {
		utils.Error(err, "DeleteProfileById", query)
		return err, "DB"
	}
	return nil, "00"
}
