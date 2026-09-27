package domain

type Entity struct {
	ID                 ID      `json:"id"`
	Name               string  `json:"name"`
	ReportName         *string `json:"report_name"`
	FiscalCode         *string `json:"fiscal_code"`
	RegistrationNumber *string `json:"registration_number"`
	Address            *string `json:"address"`
	PhoneNumber        *string `json:"phone_number"`
}
