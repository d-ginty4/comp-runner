package category

type Category struct {
	Id            int    `json:"id"`
	CompetitionId int    `json:"competitionId"`
	Name          string `json:"name"`
}
