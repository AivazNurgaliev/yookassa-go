package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
)

type Client struct {
	baseURL string
	shopID  string
	apiKEY  string
	http    *http.Client
}

func NewClient(id, key string) *Client {
	client := &Client{
		baseURL: "https://api.yookassa.ru/v3/payments",
		shopID:  id,
		apiKEY:  key,
		http:    &http.Client{},
	}

	return client
}

func (c *Client) CreatePayment(idemptKey string, p *PaymentRequest) (*Payment, error) {
	return c.CreatePaymentContext(context.Background(), idemptKey, p)
}

func (c *Client) CreatePaymentContext(ctx context.Context, idemptKey string, p *PaymentRequest) (*Payment, error) {
	data, err := json.Marshal(p)
	if err != nil {
		return nil, fmt.Errorf("request encode error: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "https://api.yookassa.ru/v3/payments", bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("cannot prepare http request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("Idempotence-Key", idemptKey)
	req.SetBasicAuth(c.shopID, c.apiKEY)

	return c.GetPaymentByRequest(req)
}

func (c *Client) GetPaymentById(id string, idemptKey string) (*Payment, error) {
	return c.GetPaymentByIdContext(context.Background(), id, idemptKey)
}

func (c *Client) GetPaymentByIdContext(ctx context.Context, id string, idemptKey string) (*Payment, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL+"/"+id, nil)

	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("Idempotence-Key", idemptKey)
	req.SetBasicAuth(c.shopID, c.apiKEY)

	if err != nil {
		return nil, fmt.Errorf("cannot prepare http request: %w", err)
	}

	return c.GetPaymentByRequest(req)
}

func (c *Client) GetPaymentByRequest(req *http.Request) (*Payment, error) {
	resp, err := c.http.Do(req)

	if err != nil {
		return nil, fmt.Errorf("request error: %w", err)
	}
	defer resp.Body.Close()

	var p *Payment
	if err = json.NewDecoder(resp.Body).Decode(&p); err != nil {
		return nil, fmt.Errorf("response decode error: %w", err)
	}
	return p, nil
}

func NewIdempotenceKey() (string, error) {
	key, err := uuid.NewRandom()
	if err != nil {
		return "", fmt.Errorf("indepotence key generation error: %w", err)
	}
	return key.String(), nil
}
