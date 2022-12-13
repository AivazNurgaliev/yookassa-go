package main

type ConfirmationRequest struct {
	Type   string `json:"type"`
	Locale string `json:"locale,omitempty"`
}

type EmbeddedConfirmationDetails struct {
	Type              string `json:"type"`
	ConfirmationToken string `json:"confirmation_token"`
}

type RedirectConfirmationRequest struct {
	ReturnURL string `json:"return_url"`
	Locale    string `json:"locale,omitempty"`
	Enforce   bool   `json:"enforce,omitempty"`
}

//TODO WEB HOOK

//todo consts

// todo навесить omitempty non required полям
type ConfirmationResponse struct {
	Type                string      `json:"type"`
	ConfirmationDetails interface{} `json:"details,omitempty"`
}

/*type RedirectConfirmation struct {
	Type            string `json:"type"`
	ConfirmationURL string `json:"confirmation_url"`
	Enforce         bool   `json:"enforce,omitempty"`
	ReturnURL       string `json:"return_url,omitempty"`
}*/

/*type RedirectConfirmationDetails struct {
	Type            string `json:"type"`
	ConfirmationURL string `json:"confirmation_url"`
	Enforce         bool   `json:"enforce,omitempty"`
	ReturnURL       string `json:"return_url,omitempty"`
}*/

/*
type QRCodeConfirmationDetails struct {
	Type             string `json:"type"`
	ConfirmationData string `json:"confirmation_data"`
}*/

/*type MobileAppConfirmationDetails struct {
	Type            string `json:"type"`
	ConfirmationURL string `json:"confirmation_url"`
}*/

/*func (c *ConfirmationResponse) BindConfirmationDetails(data []byte) error {
	jsonMap := make(map[string]json.RawMessage)
	err := json.Unmarshal(data, &jsonMap)

	if err != nil {
		//TODO ВЫКИНУТЬ ОШИБКУ
	}

	c.Type = string(jsonMap["type"])

	switch c.Type {
	case "redirect":
		c.ConfirmationDetails = &RedirectConfirmationDetails{}
	case "embedded":
		c.ConfirmationDetails = &EmbeddedConfirmationDetails{}
	case "qr":
		c.ConfirmationDetails = &QRCodeConfirmationDetails{}
	case "external":
		c.ConfirmationDetails = nil
	default:
		return fmt.Errorf("Неверный тип подтверждения", c.Type)
	}

	err = json.Unmarshal(data, c.ConfirmationDetails)
	if err != nil {
		return fmt.Errorf("Что-то пошло не так", c.Type)
	}

	return nil
}
*/
