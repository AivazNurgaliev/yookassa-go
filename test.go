package main

/*
func mawin() {
	shopId := "964684"
	key := "test_Zmo-H81YdVEZkwZvvrLC5iKju5xBBulaYKjguhB5paA"
	pR := PaymentRequest{
		Amount: Amount{
			Value:    "2.00",
			Currency: "RUB",
		},
		PaymentMethodID: "2b2103a5-000f-5000-8000-1a4ef98aaab7",
		PaymentMethodData: PaymentMethodData{
			Type: "bank_card",
		},
		Confirmation: RedirectConfirmation{
			Type:      "redirect",
			ReturnURL: "https://www.example.com/return_url",
		},
		Description: "Заказ №72",
	}

	testClient := NewClient(shopId, key)
	idempotenceKey, err := NewIdempotenceKey()
	fmt.Println()
	fmt.Println(idempotenceKey)
	if err != nil {
		fmt.Errorf("idempotencekey creation exception")
	}

	p := &pR
	fmt.Println()
	fmt.Printf("%+v\n", p)
	payment, err := testClient.CreatePayment(idempotenceKey, p)
	if err != nil {
		fmt.Errorf("payment creation exception")
	}
	fmt.Println()
	fmt.Printf("%+v\n", payment)
}
*/
