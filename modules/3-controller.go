package modules

import (
	"encoding/json"
	"github.com/purawaktra/semeru3-go/dto"
	"github.com/purawaktra/semeru3-go/utils"
)

type Semeru3Controller struct {
	uc Semeru3Usecase
}

func CreateSemeru3Controller(uc Semeru3Usecase) Semeru3Controller {
	return Semeru3Controller{
		uc: uc,
	}
}

func (ctrl Semeru3Controller) SelectProfileById(req []byte) (any, error, string) {
	// unmarshal to struct and check err
	var requestData dto.BodyProfile
	err := json.Unmarshal(req, &requestData)
	if err != nil {
		utils.Error(err, "SelectProfileById", requestData)
		return nil, err, "FM"
	}

	// call usecase for the profile and check err
	profile, err, code := ctrl.uc.SelectProfileById(requestData.AccountId)
	if err != nil {
		utils.Error(err, "SelectProfileById", requestData)
		return nil, err, code
	}

	// create return
	return profile, nil, code
}

func (ctrl Semeru3Controller) InsertProfile(req []byte) (error, string) {
	// unmarshal to struct and check err
	var requestData dto.BodyProfile
	err := json.Unmarshal(req, &requestData)
	if err != nil {
		utils.Error(err, "InsertProfile", requestData)
		return err, "FM"
	}

	// convert request to dto
	query := Profile{
		AccountId:   requestData.AccountId,
		FirstName:   requestData.FirstName,
		LastName:    requestData.LastName,
		Address:     requestData.Address,
		City:        requestData.City,
		Province:    requestData.Province,
		Zipcode:     requestData.Zipcode,
		PhoneNumber: requestData.PhoneNumber,
	}

	// call usecase for the profile and check err
	err, code := ctrl.uc.InsertProfile(query)
	if err != nil {
		utils.Error(err, "InsertProfile", requestData)
		return err, code
	}

	// create return
	return nil, code
}

func (ctrl Semeru3Controller) UpdateProfileById(req []byte) (error, string) {
	// unmarshal to struct and check err
	var requestData dto.BodyProfile
	err := json.Unmarshal(req, &requestData)
	if err != nil {
		utils.Error(err, "UpdateProfileById", requestData)
		return err, "FM"
	}

	// convert request to dto
	query := Profile{
		AccountId:   requestData.AccountId,
		FirstName:   requestData.FirstName,
		LastName:    requestData.LastName,
		Address:     requestData.Address,
		City:        requestData.City,
		Province:    requestData.Province,
		Zipcode:     requestData.Zipcode,
		PhoneNumber: requestData.PhoneNumber,
	}

	// call usecase for the profile and check err
	err, code := ctrl.uc.UpdateProfileById(query)
	if err != nil {
		utils.Error(err, "UpdateProfileById", requestData)
		return err, code
	}

	// create return
	return nil, code
}

func (ctrl Semeru3Controller) DeleteProfileById(req []byte) (error, string) {
	// unmarshal to struct and check err
	var requestData dto.BodyProfile
	err := json.Unmarshal(req, &requestData)
	if err != nil {
		utils.Error(err, "DeleteProfileById", requestData)
		return err, "FM"
	}

	// call usecase for the profile and check err
	err, code := ctrl.uc.DeleteProfileById(requestData.AccountId)
	if err != nil {
		utils.Error(err, "DeleteProfileById", requestData)
		return err, code
	}

	// create return
	return nil, code
}
