package account

type Account struct {
	ID      string  `json:"id"` //stripe subscription id
	Name    string  `json:"name"`
	Balance float32 `json:"balance"`
}
