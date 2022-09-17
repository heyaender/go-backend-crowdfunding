package transactions

import "time"

type Transactions struct {
	ID          int       `json:"id"`
	User_id     int       `json:"users_id"`
	Campaign_id int       `json:"campaigns_id"`
	Amount      int       `json:"amount"`
	Status      string    `json:"status"`
	Code        string    `json:"code"`
	Created_at  time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
}
