package account

//  Account model fields
type Account struct {
	ID      string `json:"id"` //stripe subscription id
	Name    string `json:"name"`
	Balance string `json:"balance"`
}
