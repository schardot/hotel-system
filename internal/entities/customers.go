package entities

type Customer struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	TelNum    string `json:"tel_number"`
	CelNum    string `json:"cel_number"`
	IdType    string `json:"id_type"`
	IdNum     string `json:"id_number"`
	Address   string `json:"address"`
	ZipCode   string `json:"zip_code"`
	City      string `json:"city"`
	State     string `json:"state"`
	Country   string `json:"country"`
	CreatedAt string `json:"created_at"`
}
