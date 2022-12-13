package main

// response
type PaymentMethod struct {
	ID      string      `json:"id"`
	Type    string      `json:"type"`
	Saved   bool        `json:"saved"`
	Title   string      `json:"title,omitempty"`
	Details interface{} `json:"details,omitempty"`
}

/*кошелек ЮMoney;
банковская карта — произвольная карта и карты, привязанные к кошельку ЮMoney или к Сбер ID;
SberPay;
СБП.*/

type PaymentMethodData struct {
	Type string `json:"type"`
}

type paymentMethod struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Saved bool   `json:"saved"`
	Title string `json:"title"`
}

type BankCardPaymentDetails struct {
	Card BankCardInfo `json:"card"`
}

type BankCardInfo struct {
	First6        string `json:"first_6"`
	Last4         string `json:"last_4"`
	ExpiryYear    string `json:"expiry_year"`
	ExpiryMonth   string `json:"expiry_month"`
	CardType      string `json:"card_type"`
	IssuerCountry string `json:"issuer_country"`
	IssuerName    string `json:"issuer_name"`
	Source        string `json:"source"`
}

type YooMoneyPaymentDetails struct {
	AccountNumber string `json:"account_number"`
}

type SberPayPaymentDetails struct {
	CardInfo CardInfo `json:"card"`
	Phone    string   `json:"phone"`
}

// для сбп
type CardInfo struct {
	First6      string `json:"first_6,omitempty"`
	Last4       string `json:"last_4"`
	ExpiryYear  string `json:"expiry_year"`
	ExpiryMonth string `json:"expiry_month"`
	CardType    string `json:"card_type"`
}

/*
func (pm *PaymentMethod) BindPaymentMethod(data []byte) error {
	var p paymentMethod
	err := json.Unmarshal(data, &p)
	if err != nil {
		//TODO ВЫКИНУТЬ ОШИБКУ
	}
	pm.ID = p.ID
	pm.Type = p.Type
	pm.Saved = p.Saved
	pm.Title = p.Title

	switch p.Type {
	case "bank_card":
		pm.Details = &BankCardPaymentDetails{}
	case "yoo_money":
		pm.Details = &YooMoneyPaymentDetails{}
	case "sberbank":
		pm.Details = &SberPayPaymentDetails{}
	case "sbp":
		pm.Details = nil
	default:
		//TODO МБ ошибку выкинуть
		pm.Details = &map[string]interface{}{}
	}

	err = json.Unmarshal(data, pm.Details)

	if err != nil {
		return err
	}

	return nil
}
*/
