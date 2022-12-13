package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strings"
)

// TODO ПОЛУЧЕНИЕ  RESPONSE PAYMENT ИЗ BYTE(UNMARSHAL) + OMITEMPTY + API METHODS
func main() {
	client := NewClient("964684", "test_Zmo-H81YdVEZkwZvvrLC5iKju5xBBulaYKjguhB5paA")
	//fmt.Println(NewIdempotenceKey())
	//var p *Payment
	/*	idemptKey, err := NewIdempotenceKey()
		if err != nil {
			fmt.Println(err)
		}
		p, err := client.GetPaymentById("2b2a97c7-000f-5000-8000-1e26b3a87572", idemptKey)
		if err != nil {
			return
		}
		fmt.Printf("%+v\n", p)*/
	pR := PaymentRequest{
		Amount: Amount{
			Value:    "150.00",
			Currency: "RUB",
		},
		Confirmation: ConfirmationRequest{
			Type: "embedded",
		},
		//Description: "Заказ №144",
		Receipt: &ReceiptRequest{
			/*			Customer: Customer{
						FullName: "Иванов Иван Иванович",
						Phone:    "79000000000",
					},*/
			Items: []Item{
				{
					Description: "Наименование товара 1",
					Amount: Amount{
						Value:    "1250.00",
						Currency: "RUB",
					},
					VatCode:  2,
					Quantity: "1.00",
					//PaymentSubject: "commodity",
					//PaymentMode:    "full_payment",
				},
			},
		},
		Capture: true,
	}

	idemptKey, err := NewIdempotenceKey()
	if err != nil {
		fmt.Println(err)
	}
	var p *Payment
	p, err = client.CreatePayment(idemptKey, &pR)
	if err != nil {
		fmt.Errorf("Error creating payment, ")
	}

	fmt.Printf("%+v\n", p)
	confirmationToken := p.Confirmation.ConfirmationToken
	fmt.Println(confirmationToken)
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/payment", func(c *gin.Context) {
		c.HTML(http.StatusOK, "payment.tmpl", gin.H{
			"confirmation_token": confirmationToken,
		})
	})
	router.Run(":8080")
}

/*
func createPayment(p PaymentRequest) Payment {
	url := "https://api.yookassa.ru/v3/payments"
	method := "POST"

	marshalledPaymentRequest, err := json.Marshal(p)

	if err != nil {
		fmt.Println(err)
		//return nil
	}
	payload := strings.NewReader(string(marshalledPaymentRequest))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	fmt.Println(req)
	if err != nil {
		fmt.Println(err)
		//return nil
	}
	k, err := NewIdempotenceKey()
	if err != nil {
		fmt.Println(err)
		//return nil
	}

	req.Header.Add("Idempotence-key", k)
	req.Header.Add("Authorization", "Basic OTY0Njg0OnRlc3RfWm1vLUg4MVlkVkVaa3dadnZyTEM1aUtqdTV4QkJ1bGFZS2pndWhCNXBhQQ==")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		//return nil
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		//return nil
	}
	//fmt.Println(string(body))

	//return body
	var paymentResponse Payment
	err = json.Unmarshal(body, &paymentResponse)

	if err != nil {
		fmt.Println(err)
	}
	return paymentResponse
}

// TODO BINDING CONFIRMATION DETAILS
func getPayment(id string) Payment {
	url := "https://api.yookassa.ru/v3/payments"
	method := "GET"

	bytes := make([]byte, 2048)
	payload := strings.NewReader(string(bytes))

	client := &http.Client{}
	req, err := http.NewRequest(method, url+"/"+id, payload)
	//fmt.Println(req)
	if err != nil {
		fmt.Println(err)
		//return nil
	}

	idempotenceKey, err := NewIdempotenceKey()
	if err != nil {
		//return nil
	}

	req.Header.Add("Idempotence-key", idempotenceKey)
	req.Header.Add("Authorization", "Basic OTY0Njg0OnRlc3RfWm1vLUg4MVlkVkVaa3dadnZyTEM1aUtqdTV4QkJ1bGFZS2pndWhCNXBhQQ==")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		//return nil
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		//return nil
	}
	//fmt.Println()
	//fmt.Println(string(body))
	var p Payment
	err = json.Unmarshal(body, &p)
	//fmt.Println("P:")
	//fmt.Println()
	//fmt.Printf("%+v\n", p)
	if err != nil {
		fmt.Println(err)
		//return nil
	}
	return p
}
*/
// todo перемещения по курсору
func getPayments() (*PaymentHistoryResponse, error) {
	url := "https://api.yookassa.ru/v3/payments"
	method := "GET"

	bytes := make([]byte, 2048)
	payload := strings.NewReader(string(bytes))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	//fmt.Println(req)
	if err != nil {
		fmt.Println(err)
		//return nil
	}

	idempotenceKey, err := NewIdempotenceKey()
	if err != nil {
		//return nil
	}

	req.Header.Add("Idempotence-key", idempotenceKey)
	req.Header.Add("Authorization", "Basic OTY0Njg0OnRlc3RfWm1vLUg4MVlkVkVaa3dadnZyTEM1aUtqdTV4QkJ1bGFZS2pndWhCNXBhQQ==")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		//return nil
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		//return nil
	}
	fmt.Println()
	fmt.Println(string(body))
	var p PaymentHistoryResponse
	err = json.Unmarshal(body, &p)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong")
	}

	return &p, nil
}
