package modules

import (
	"errors"
	"github.com/purawaktra/semeru3-go/entities"
	"github.com/purawaktra/semeru3-go/utils"
)

type Semeru3Usecase struct {
	repo Semeru3Repo
}

func CreateSemeru3Usecase(repo Semeru3Repo) Semeru3Usecase {
	return Semeru3Usecase{
		repo: repo,
	}
}

func (uc Semeru3Usecase) SelectProfileById(accountId int) (Profile, error, string) {
	// create check input on account id
	if accountId < 1 {
		utils.Error(errors.New("accountId can not be zero or negative"), "SelectProfileById", accountId)
		return Profile{}, errors.New("accountId can not be nil, negative or zero"), "FC"
	}

	// convert input to entity
	query := entities.Profile{AccountId: uint(accountId)}

	// call repo for the account id and check err
	profile, err, code := uc.repo.SelectProfileById(query)
	if err != nil {
		utils.Error(err, "SelectProfileById", query)
		return Profile{}, err, code
	}

	// convert a result to dto and create return
	result := Profile{
		AccountId:   int(profile.AccountId),
		FirstName:   profile.FirstName,
		LastName:    profile.LastName,
		Address:     profile.Address,
		City:        int(profile.City),
		Province:    int(profile.Province),
		Zipcode:     profile.Zipcode,
		PhoneNumber: profile.PhoneNumber,
	}
	return result, nil, code
}

func (uc Semeru3Usecase) InsertProfile(query Profile) (error, string) {
	// create check input on account id, city, and province
	if query.AccountId < 1 {
		utils.Error(errors.New("accountId can not be zero or negative"), "InsertProfile", query.AccountId)
		return errors.New("accountId can not be nil, negative or zero"), "FC"
	}
	if query.City < 0 {
		utils.Error(errors.New("city can not be negative"), "InsertProfile", query.City)
		return errors.New("city can not be negative"), "FC"
	}
	if query.Province < 0 {
		utils.Error(errors.New("province can not be negative"), "InsertProfile", query.City)
		return errors.New("province can not be negative"), "FC"
	}

	// convert input to entity
	profile := entities.Profile{
		AccountId:   uint(query.AccountId),
		FirstName:   query.FirstName,
		LastName:    query.LastName,
		Address:     query.Address,
		City:        uint(query.City),
		Province:    uint(query.Province),
		Zipcode:     query.Zipcode,
		PhoneNumber: query.PhoneNumber,
	}

	// call usecase and check err
	err, code := uc.repo.InsertProfile(profile)
	if err != nil {
		utils.Error(err, "InsertProfile", "")
		return err, code
	}

	// create return
	return nil, code
}

func (uc Semeru3Usecase) UpdateProfileById(query Profile) (error, string) {
	// create check input on account id, city, and province
	if query.AccountId < 1 {
		utils.Error(errors.New("accountId can not be zero or negative"), "UpdateProfileById", query.AccountId)
		return errors.New("accountId can not be nil, negative or zero"), "FC"
	}
	if query.City < 0 {
		utils.Error(errors.New("city can not be negative"), "UpdateProfileById", query.City)
		return errors.New("city can not be negative"), "FC"
	}
	if query.Province < 0 {
		utils.Error(errors.New("province can not be negative"), "UpdateProfileById", query.City)
		return errors.New("province can not be negative"), "FC"
	}

	// convert input to entity
	profile := entities.Profile{
		AccountId:   uint(query.AccountId),
		FirstName:   query.FirstName,
		LastName:    query.LastName,
		Address:     query.Address,
		City:        uint(query.City),
		Province:    uint(query.Province),
		Zipcode:     query.Zipcode,
		PhoneNumber: query.PhoneNumber,
	}

	// call usecase and check err
	err, code := uc.repo.UpdateProfileById(profile)
	if err != nil {
		utils.Error(err, "UpdateProfileById", "")
		return err, code
	}

	// create return
	return nil, code
}

func (uc Semeru3Usecase) DeleteProfileById(accountId int) (error, string) {
	// create check input on account id, city, and province
	if accountId < 1 {
		utils.Error(errors.New("accountId can not be zero or negative"), "DeleteProfileById", accountId)
		return errors.New("accountId can not be nil, negative or zero"), "FC"
	}

	// convert input to entity
	query := entities.Profile{AccountId: uint(accountId)}

	// call repo for the account id and check err
	err, code := uc.repo.DeleteProfileById(query)
	if err != nil {
		utils.Error(err, "DeleteProfileById", query)
		return err, code
	}

	// create return
	return nil, code
}
