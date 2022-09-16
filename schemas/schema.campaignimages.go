package schemas

type CampaignImages struct {
	ID          int    `json:"id"`
	Campaign_id int    `json:"campaigns_id"`
	Filename    string `json:"filename"`
	Is_primary  int    `json:"is_primary"`
	Created_at  string `json:"created_at"`
	Updated_at  string `json:"updated_at"`
}
