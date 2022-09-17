package campaignimages

import "time"

type CampaignImages struct {
	ID          int       `json:"id"`
	Campaign_id int       `json:"campaigns_id"`
	Filename    string    `json:"filename"`
	Is_primary  int       `json:"is_primary"`
	Created_at  time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
}
