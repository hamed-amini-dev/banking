package accounts

type TransferParams struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Balance string `json:"balance"`
}
