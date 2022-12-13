package main

import "time"

//todo consts

// FraudData
// может структура конфирмации локаль и тайп а не интерфейс
// TODO receipt,RECEPIENT!!!
//ClientIP           string                 `json:"client_ip"`
//SavePaymentMethod  bool                   `json:"save_payment_method"`
//PaymentToken       string                 `json:"payment_token"`

type PaymentHistoryResponse struct {
	Type       string    `json:"type"`
	Items      []Payment `json:"items"`
	NextCursor string    `json:"next_cursor,omitempty"`
}

type PaymentRequest struct {
	Amount             Amount                 `json:"amount"`
	Description        string                 `json:"description,omitempty"`
	Receipt            *ReceiptRequest        `json:"receipt"`
	Recipient          *RecipientRequest      `json:"recipient,omitempty"`
	PaymentMethodID    string                 `json:"payment_method_id,omitempty"`
	PaymentMethodData  *PaymentMethod         `json:"payment_method_data"`
	Confirmation       ConfirmationRequest    `json:"confirmation"`
	Capture            bool                   `json:"capture"`
	Metadata           map[string]interface{} `json:"metadata,omitempty"`
	MerchantCustomerID string                 `json:"merchant_customer_id,omitempty"`
	//Receipt     ReceiptRequest   `json:"receipt"`

}

// todo навесить omitempty non required полям
// todo спросить про указатели

type Payment struct {
	ID                   string                      `json:"id"`
	Status               string                      `json:"status"`
	Amount               Amount                      `json:"amount"`
	IncomeAmount         Amount                      `json:"income_amount,omitempty"`
	RefundedAmount       Amount                      `json:"refunded_amount,omitempty"`
	Description          string                      `json:"description,omitempty"`
	Recipient            Recipient                   `json:"recipient"`
	CapturedAt           time.Time                   `json:"captured_at,omitempty"`
	CreatedAt            time.Time                   `json:"created_at"`
	ExpiresAt            time.Time                   `json:"expires_at,omitempty"`
	Test                 bool                        `json:"test"`
	Paid                 bool                        `json:"paid"`
	Refundable           bool                        `json:"refundable"`
	ReceiptRegistration  string                      `json:"receipt_registration"`
	Metadata             map[string]interface{}      `json:"metadata,omitempty"`
	CancellationDetails  CancellationDetails         `json:"cancellation_details,omitempty"`
	AuthorizationDetails AuthorizationDetails        `json:"authorization_details,omitempty"`
	MerchantCustomerID   string                      `json:"merchant_customer_id,omitempty"`
	PaymentMethod        PaymentMethod               `json:"payment_method"`
	Confirmation         EmbeddedConfirmationDetails `json:"confirmation"`
}

type Amount struct {
	Value    string `json:"value"`
	Currency string `json:"currency"`
}

type Recipient struct {
	AccountID string `json:"account_id"`
	GatewayID string `json:"gateway_id"`
}

type CancellationDetails struct {
	Party  string `json:"party"`
	Reason string `json:"reason"`
}

// Данные об авторизации платежа.
type AuthorizationDetails struct {
	RRN          string       `json:"rrn,omitempty"`
	AuthCode     string       `json:"auth_code,omitempty"`
	ThreeDSecure ThreeDSecure `json:"three_d_secure"`
}

type ThreeDSecure struct {
	Applied bool `json:"applied"`
}

type RecipientRequest struct {
	GatewayID string `json:"gateway_id"`
}

// tax system code если ффд 1.2
type ReceiptRequest struct {
	Customer Customer `json:"customer"`
	Items    []Item   `json:"items"`
	//ReceiptIndustryDetails    []PaymentSubjectIndustryDetails `json:"receipt_industry_details"`
	//ReceiptOperationalDetails *ReceiptOperationalDetails      `json:"receipt_operational_details"`
	//TaxSystemCode             int      `json:"tax_system_code"`

}

type Customer struct {
	FullName string `json:"full_name,omitempty"`
	INN      string `json:"inn,omitempty"`
	Email    string `json:"email,omitempty"`
	Phone    string `json:"phone,omitempty"`
}

// vat code - коды ставок ндс для самозанятых 1
type Item struct {
	Description                   string                          `json:"description"`
	Amount                        Amount                          `json:"amount"`
	VatCode                       int                             `json:"vat_code"`
	Quantity                      string                          `json:"quantity"`
	Measure                       string                          `json:"measure,omitempty"`
	PaymentSubject                string                          `json:"payment_subject,omitempty"`
	PaymentMode                   string                          `json:"payment_mode,omitempty"`
	CountryOfOriginCode           string                          `json:"country_of_origin_code,omitempty"`
	CustomsDeclarationNumber      string                          `json:"customs_declaration_number,omitempty"`
	Excise                        string                          `json:"excise,omitempty"`
	PaymentSubjectIndustryDetails []PaymentSubjectIndustryDetails `json:"payment_subject_industry_details,omitempty"`
	//MarkQuantity                  MarkQuantity `json:"mark_quantity"`
	//ProductCode                   string       `json:"product_code"`

}

type PaymentSubjectIndustryDetails struct {
	FederalID      string `json:"federal_id"`
	DocumentDate   string `json:"document_date"`
	DocumentNumber string `json:"document_number"`
	Value          string `json:"string"`
}

type ReceiptOperationalDetails struct {
	OperationID int       `json:"operation_id"`
	Value       string    `json:"value"`
	CreatedAt   time.Time `json:"created_at"`
}

type MarkQuantity struct {
	Numerator   int `json:"numerator"`
	Denominator int `json:"denominator"`
}

/*type Payment struct {
	ID                   string                `json:"id"`
	Status               string                `json:"status"`
	ReceiptRegistration  string                `json:"receipt_registration"`
	MerchantCustomerID   string                `json:"merchant_customer_id"`
	Description          string                `json:"description"`
	Test                 bool                  `json:"test"`
	Paid                 bool                  `json:"paid"`
	Refundable           bool                  `json:"refundable"`
	Amount               Amount                `json:"amount"`
	IncomeAmount         *IncomeAmount         `json:"income_amount"`
	RefundedAmount       *RefundedAmount       `json:"refunded_amount"`
	Recipient            Recipient             `json:"recipient"`
	PaymentMethod        PaymentMethod         `json:"payment_method"`
	CapturedAt           time.Time             `json:"captured_at"`
	CreatedAt            time.Time             `json:"created_at"`
	ExpiresAt            time.Time             `json:"expires_at"`
	CancellationDetails  *CancellationDetails  `json:"cancellation_details"`
	AuthorizationDetails *AuthorizationDetails `json:"authorization_details"`
	Confirmation         interface{}           `json:"confirmation"`

	Metadata map[string]interface{} `json:"metadata"`
}*/

/*
type Deal struct {
	ID          string       `json:"id"`
	Settlements []Settlement `json:"settlements"`
}*/

/*type Settlement struct {
	Type   string `json:"type"`
	Amount Amount `json:"amount"`
}
*/

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
