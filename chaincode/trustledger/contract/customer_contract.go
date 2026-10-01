package contract

import (
	"encoding/json"
	"fmt"

	"trustledger/model"
	"trustledger/utils"

	"github.com/hyperledger/fabric-contract-api-go/v2/contractapi"
)


// CreateCustomer adds a new customer to the blockchain.
func (s *SmartContract) CreateCustomer(
	ctx contractapi.TransactionContextInterface,
	customerID string,
	fullName string,
	dateOfBirth string,
	email string,
	phone string,
	address string,
	nationalID string,
	issuingBank string,
	documentHash string,
) error {

	exists, err := s.CustomerExists(ctx, customerID)
	if err != nil {
		return err
	}

	if exists {
		return fmt.Errorf("customer %s already exists", customerID)
	}

	customer := model.Customer{
		CustomerID:      customerID,
		FullName:        fullName,
		DateOfBirth:     dateOfBirth,
		Email:           email,
		Phone:           phone,
		Address:         address,
		NationalID:      nationalID,
		IssuingBank:     issuingBank,
		KYCStatus:       model.KYCStatusPending,
		ConsentGranted:  false,
		DocumentHash:    documentHash,
		CreatedAt:       utils.GetCurrentTimestamp(),
		UpdatedAt:       utils.GetCurrentTimestamp(),
	}

	if err := utils.ValidateCustomer(customer); err != nil {
		return err
	}

	customerJSON, err := json.Marshal(customer)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(customerID, customerJSON)
}
