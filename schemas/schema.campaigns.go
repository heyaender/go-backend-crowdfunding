package schemas

type Campaigns struct {
	ID                int    `json:"id"`
	Users_id          int    `json:"users_id"`
	Title             string `json:"title"`
	Short_description string `json:"short_description"`
	Description       string `json:"description"`
	Slug              string `json:"slug"`
	Backer_count      int    `json:"backer_count"`
	Perks             string `json:"perks"`
	Goal_amount       int    `json:"goal_amount"`
	Current_amount    int    `json:"current_amount"`
	Created_at        string `json:"created_at"`
	Updated_at        string `json:"updated_at"`
}
