package helios

type users struct {
	Users []user `json:"users"`
}

type user struct {
	ID              int    `json:"id"`
	InstagramHandle string `json:"instagram_handle"`
}
