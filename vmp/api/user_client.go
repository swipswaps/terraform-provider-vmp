// Copyright Verizon Media, Licensed under the terms of the Apache 2.0 license . See LICENSE file in project root for terms.

package api

import (
	"fmt"
)

// UserAPIClient interacts with the Verizon Media API
type UserAPIClient struct {
	BaseAPIClient *ApiClient
	PartnerID     *int
}

// NewUserAPIClient -
func NewUserAPIClient(baseAPIClient *ApiClient, partnerID *int) *UserAPIClient {
	return &UserAPIClient{
		BaseAPIClient: baseAPIClient,
		PartnerID:     partnerID,
	}
}

// CustomerUser -
type CustomerUser struct {
	Address1      string
	Address2      string
	City          string
	Country       string
	CustomID      string `json:"CustomId"` // Read-only
	Email         string
	Fax           string
	FirstName     string
	IsAdmin       int8
	LastName      string
	Mobile        string
	Password      *string // nullable
	Phone         string
	State         string
	TimeZoneID    *int `json:"TimeZoneId"` // nullable
	Title         string
	ZIP           string `json:"Zip"`
	LastLoginDate string // Read-only
}

// GetCustomerUser -
func (apiClient *UserAPIClient) GetCustomerUser(accountNumber string, customerUserID int) (*CustomerUser, error) {
	// TODO: support custom id types, not just Hex ID ANs
	baseURL := fmt.Sprintf("pcc/customers/users/%d?idtype=an&id=%s", customerUserID, accountNumber)
	relURL := FormatURLAddPartnerID(baseURL, apiClient.PartnerID)

	request, err := apiClient.BaseAPIClient.BuildRequest("GET", relURL, nil, false)
	InfoLogger.Printf("GetCustomerUser [GET] Url: %s\n", request.URL)

	if err != nil {
		return nil, err
	}

	parsedResponse := &CustomerUser{}

	_, err = apiClient.BaseAPIClient.SendRequest(request, &parsedResponse)

	if err != nil {
		return nil, err
	}

	return parsedResponse, nil
}

// AddCustomerUser -
func (apiClient *UserAPIClient) AddCustomerUser(accountNumber string, body *CustomerUser) (int, error) {
	// TODO: support custom id types, not just Hex ID ANs
	baseURL := fmt.Sprintf("pcc/customers/users?idtype=an&id=%s", accountNumber)
	relURL := FormatURLAddPartnerID(baseURL, apiClient.PartnerID)

	request, err := apiClient.BaseAPIClient.BuildRequest("POST", relURL, body, false)
	InfoLogger.Printf("AddCustomerUser [POST] Url: %s\n", request.URL)

	parsedResponse := &struct {
		CustomerUserID int `json:"CustomerUserId"`
	}{}

	_, err = apiClient.BaseAPIClient.SendRequest(request, &parsedResponse)

	if err != nil {
		return 0, err
	}

	return parsedResponse.CustomerUserID, err
}

// UpdateCustomerUser -
func (apiClient *UserAPIClient) UpdateCustomerUser(accountNumber string, customerUserID int, body *CustomerUser) error {
	// TODO: support custom ids for accounts
	baseURL := fmt.Sprintf("pcc/customers/users/%d?idtype=an&id=%s", customerUserID, accountNumber)
	relURL := FormatURLAddPartnerID(baseURL, apiClient.PartnerID)

	request, err := apiClient.BaseAPIClient.BuildRequest("PUT", relURL, body, false)
	InfoLogger.Printf("UpdateCustomerUser [PUT] Url: %s\n", request.URL)

	if err != nil {
		return err
	}

	_, err = apiClient.BaseAPIClient.SendRequest(request, nil)

	return err
}

// DeleteCustomerUser -
func (apiClient *UserAPIClient) DeleteCustomerUser(accountNumber string, customerUserID int) error {
	// TODO: support custom ids for accounts
	baseURL := fmt.Sprintf("pcc/customers/users/%d?idtype=an&id=%s", customerUserID, accountNumber)
	relURL := FormatURLAddPartnerID(baseURL, apiClient.PartnerID)

	request, err := apiClient.BaseAPIClient.BuildRequest("DELETE", relURL, nil, false)
	InfoLogger.Printf("DeleteCustomerUser [DELETE] Url: %s\n", request.URL)

	if err != nil {
		return err
	}

	_, err = apiClient.BaseAPIClient.SendRequest(request, nil)

	return err
}
