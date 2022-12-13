package main

import "time"

type Receipt struct {
	ID                   string    `json:"id"`
	Type                 string    `json:"type"`
	PaymentID            string    `json:"payment_id,omitempty"`
	RefundID             string    `json:"refund_id,omitempty"`
	Status               string    `json:"status"`
	FiscalDocumentNumber string    `json:"fiscal_document_number,omitempty"`
	FiscalStorageNumber  string    `json:"fiscal_storage_number,omitempty"`
	FiscalAttribute      string    `json:"fiscal_attribute,omitempty"`
	RegisteredAt         time.Time `json:"registered_at,omitempty"`
	FiscalProviderID     string    `json:"fiscal_provider_id"`
	//Items
}

/*type Settlement struct {
	Type   string `json:"type"`
	Amount Amount `json:"amount"`
}
*/
/*type ReceiptRequest struct {
	Customer                  Customer `json:"customer"`
	Items                     []Item   `json:"items"`
	TaxSystemCode             int      `json:"tax_system_code"`
	ReceiptIndustryDetails    []Psid   `json:"receipt_industry_details"`
	ReceiptOperationalDetails Rod      `json:"receipt_operational_details"`
}
*/
/*type Customer struct {
	FullName string `json:"full_name"`
	INN      string `json:"inn"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}
*/
//TODO vat code - коды ставок ндс для самозанятых 1 от 1 до 6
/*//todo paymentMode - full_payment
type Item struct {
	Description                   string       `json:"description"`
	Amount                        Amount       `json:"amount"`
	VatCode                       int          `json:"vat_code"`
	Quantity                      float32      `json:"quantity"`
	Measure                       string       `json:"measure"`
	MarkQuantity                  MarkQuantity `json:"mark_quantity"`
	PaymentSubject                string       `json:"payment_subject,omitempty"`
	PaymentMode                   string       `json:"payment_mode,omitempty"`
	CountryOfOriginCode           string       `json:"country_of_origin_code,omitempty"`
	CustomsDeclarationNumber      string       `json:"customs_declaration_number,omitempty"`
	Excise                        string       `json:"excise,omitempty"`
	ProductCode                   string       `json:"product_code"`
	PaymentSubjectIndustryDetails []Psid       `json:"payment_subject_industry_details"`
	Supplier                      Supplier     `json:"supplier,omitempty"`
}

type Supplier struct {
	Name  string `json:"name,omitempty"`
	Phone string `json:"phone,omitempty"`
	INN   string `json:"inn,omitempty"`
}
*/
/*type Psid struct {
	FederalID      string `json:"federal_id"`
	DocumentDate   string `json:"document_date"`
	DocumentNumber string `json:"document_number"`
	Value          string `json:"string"`
}*/
/*
type Rod struct {
	OperationID int       `json:"operation_id"`
	Value       string    `json:"value"`
	CreatedAt   time.Time `json:"created_at"`
}*/
/*
type MarkQuantity struct {
	Numerator   int `json:"numerator"`
	Denominator int `json:"denominator"`
}*/

/*type Transfer struct {
	AccountID         string                 `json:"account_id"`
	Amount            Amount                 `json:"amount"`
	Status            string                 `json:"status"`
	PlatformFeeAmount PlatformFeeAmount      `json:"platform_fee_amount"`
	Description       string                 `json:"description"`
	Metadata          map[string]interface{} `json:"metadata"`
}*/
/*
type PlatformFeeAmount struct {
	Value    string `json:"value"`
	Currency string `json:"currency"`
}*/
