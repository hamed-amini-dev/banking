package accounts

// Use params for sending item need to service logic function

type TransferParams struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Balance string `json:"balance"`
}
