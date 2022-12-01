package models

type Author struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Nickname   string `json:"nickname"`
	Speciality string `json:"speciality"`
}
