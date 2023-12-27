package models

type PhonesInput struct {
	Home   GenericPhoneInput `json:"home_phone,omitempty"`
	Mobile GenericPhoneInput `json:"mobile_phone,omitempty"`
}

type GenericPhoneInput struct {
	CountryCode string `json:"country_code"`
	AreaCode    string `json:"area_code"`
	Number      string `json:"number"`
}
