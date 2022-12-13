package main

//request

type SberPayPayment struct {
	Type  string `json:"type"`
	Phone string `json:"phone"`
}

type SBPPayment struct {
	Type string `json:"type"`
}

type YooMoneyPayment struct {
	Type string `json:"type"`
}

// Card - если будем хранить данные карточек
type BankCardPayment struct {
	Type string `json:"type"`
	//Card Card   `json:"card"`
}

type Card struct {
	Number      string `json:"number"`
	ExpiryYear  string `json:"expiry_year"`
	ExpiryMonth string `json:"expiry_month"`
	CSC         string `json:"csc"`
	CardHolder  string `json:"cardholder"`
}

//response

/*func (pm *PaymentMethod) BindPaymentMethod(data []byte) error {
	var p paymentMethod
	err := json.Unmarshal(data, &p)

	if err != nil {
		//TODO ВЫКИНУТЬ ОШИБКУ
	}

	pm.ID = p.ID
	pm.Type = p.Type
	pm.Saved = p.Saved

	switch pm.Type {
	case ""
	}
}*/
