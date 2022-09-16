package schemas

type Users struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Occupation      string `json:"occupation"`
	Email           string `json:"email"`
	Password_hash   string `json:"password_hash"`
	Avatar_filename string `json:"avatar_filename"`
	Role            string `json:"role"`
	Token           string `json:"token"`
	Created_at      string `json:"create_at"`
	Updated_at      string `json:"updated_at"`
}
