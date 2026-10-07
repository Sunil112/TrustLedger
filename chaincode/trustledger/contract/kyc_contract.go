package contract

import (
	"fmt"

	"trustledger/model"
	"trustledger/utils"

	"github.com/hyperledger/fabric-contract-api-go/v2/contractapi"
)

// IssueKYC marks a customer's KYC as VERIFIED.
func (s *SmartContract) IssueKYC(
	ctx contractapi.TransactionContextInterface,
	customerID string,
) error {

	customer, err := s.ReadCustomer(ctx, customerID)
	if err != nil {
		return err
	}

	if customer.KYCStatus == model.KYCStatusVerified {
		return fmt.Errorf("customer %s is already KYC verified", customerID)
	}

	customer.KYCStatus = model.KYCStatusVerified
	customer.UpdatedAt = utils.GetCurrentTimestamp()

	return s.saveCustomer(ctx, customer)
}
