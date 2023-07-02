package modules

import (
	"fmt"
	"github.com/purawaktra/semeru3-go/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Semeru3Repo struct {
	db *gorm.DB
}

func CreateSemeru3Repo(gorm *gorm.DB) Semeru3Repo {
	return Semeru3Repo{
		db: gorm,
	}
}

type Semeru3RepoInterface interface {
	SelectAccountById(query entities.Accounts, offset uint) ([]entities.Accounts, error)
	SelectAccountByFirstName(query entities.Accounts, offset uint) ([]entities.Accounts, error)
	SelectAccountByLastName(query entities.Accounts, offset uint) ([]entities.Accounts, error)
	SelectAccountByEmail(query entities.Accounts, offset uint) ([]entities.Accounts, error)
	SelectAllAccount(offset uint) ([]entities.Accounts, error)
	InsertSingleAccount(query entities.Accounts) (entities.Accounts, error)
	InsertMultipleAccount(queries []entities.Accounts) ([]entities.Accounts, []error)
	UpdateSingleAccountById(query entities.Accounts) (entities.Accounts, error)
	UpdateMultipleAccountById(queries []entities.Accounts) ([]entities.Accounts, []error)
	DeleteSingleAccountById(query entities.Accounts) (entities.Accounts, error)
	DeleteMultipleAccountById(queries []entities.Accounts) ([]entities.Accounts, []error)
}

func (br Semeru3Repo) SelectAccountById(query entities.Accounts, offset uint) ([]entities.Accounts, error) {
	var accounts = make([]entities.Accounts, 0)
	tx := br.db.Raw(
		fmt.Sprint("SELECT * FROM customers WHERE account_id = ", query.AccountId, " LIMIT 10 OFFSET ", offset)).Scan(
		&accounts)
	err := tx.Error
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

func (br Semeru3Repo) SelectAccountByFirstName(query entities.Accounts, offset uint) ([]entities.Accounts, error) {
	var accounts []entities.Accounts
	tx := br.db.Raw(
		fmt.Sprint("SELECT * FROM customers WHERE first_name LIKE \"%", query.FirstName, "%\" LIMIT 10 OFFSET ", offset)).Scan(
		&accounts)
	err := tx.Error
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

func (br Semeru3Repo) SelectAccountByLastName(query entities.Accounts, offset uint) ([]entities.Accounts, error) {
	var accounts []entities.Accounts
	tx := br.db.Raw(
		fmt.Sprint("SELECT * FROM customers WHERE first_name LIKE \"%", query.LastName, "%\" LIMIT 10 OFFSET ", offset)).Scan(
		&accounts)
	err := tx.Error
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

func (br Semeru3Repo) SelectAccountByEmail(query entities.Accounts, offset uint) ([]entities.Accounts, error) {
	var accounts []entities.Accounts
	tx := br.db.Raw(
		fmt.Sprint("SELECT * FROM customers WHERE first_name LIKE \"%", query.LastName, "%\" LIMIT 10 OFFSET ", offset)).Scan(
		&accounts)
	err := tx.Error
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

func (br Semeru3Repo) SelectAllAccount(offset uint) ([]entities.Accounts, error) {
	var accounts []entities.Accounts
	tx := br.db.Raw(
		fmt.Sprint("SELECT * FROM customers LIMIT 10 OFFSET ", offset)).Scan(
		&accounts)
	err := tx.Error
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

func (br Semeru3Repo) InsertSingleAccount(query entities.Accounts) (entities.Accounts, error) {
	var account entities.Accounts
	tx := br.db.Model(&account).Clauses(clause.Returning{}).Create(query)
	err := tx.Error
	if err != nil {
		return entities.Accounts{}, err
	}
	return account, nil
}

func (br Semeru3Repo) InsertMultipleAccount(queries []entities.Accounts) ([]entities.Accounts, []error) {
	var account entities.Accounts
	var accounts []entities.Accounts
	var errors []error
	for _, query := range queries {
		tx := br.db.Model(&account).Clauses(clause.Returning{}).Create(query)
		accounts = append(accounts, account)
		err := tx.Error
		if err != nil {
			errors = append(errors, err)
		} else {
			errors = append(errors, nil)
		}
	}
	return accounts, errors
}

func (br Semeru3Repo) UpdateSingleAccountById(query entities.Accounts) (entities.Accounts, error) {
	var account entities.Accounts
	tx := br.db.Model(&account).Clauses(clause.Returning{}).Updates(query)
	err := tx.Error
	if err != nil {
		return entities.Accounts{}, err
	}
	return account, nil
}

func (br Semeru3Repo) UpdateMultipleAccountById(queries []entities.Accounts) ([]entities.Accounts, []error) {
	var account entities.Accounts
	var accounts []entities.Accounts
	var errors []error
	for _, query := range queries {
		tx := br.db.Model(&account).Clauses(clause.Returning{}).Updates(query)
		accounts = append(accounts, account)
		err := tx.Error
		if err != nil {
			errors = append(errors, err)
		} else {
			errors = append(errors, nil)
		}
	}
	return accounts, errors
}

func (br Semeru3Repo) DeleteSingleAccountById(query entities.Accounts) (entities.Accounts, error) {
	tx := br.db.Clauses(clause.Returning{}).Delete(query)
	account := query
	err := tx.Error
	if err != nil {
		return entities.Accounts{}, err
	}
	return account, nil
}

func (br Semeru3Repo) DeleteMultipleAccountById(queries []entities.Accounts) ([]entities.Accounts, []error) {
	var errors []error
	var accounts []entities.Accounts
	for _, query := range queries {
		tx := br.db.Clauses(clause.Returning{}).Delete(&query)
		accounts = append(accounts, query)
		err := tx.Error
		if err != nil {
			errors = append(errors, err)
		} else {
			errors = append(errors, nil)
		}
	}
	return accounts, errors
}
