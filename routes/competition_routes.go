package routes

// @Summary List Competitions
// @Tags Competitions
// @Accept json
// @Produce json
// @Success 200 {object} models.competitions.Competition
// @Router /v1/competition [get]
func listCompetitions() {

}

// @Summary Create a Competition
// @Tags Competitions
// @Accept json
// @Produce json
// @Success 200 {object} models.competitions.Competition
// @Param competition body models.competitions.CompetitionRequest true "Competition to create"
// @Router /v1/competition [post]
func createCompetition() {

}

// @Summary Get a Competition
// @Tags Competitions
// @Accept json
// @Produce json
// @Success 200 {object} models.competitions.Competition
// @Router /v1/competition/{competitionId} [get]
func getCompetition() {

}

// @Summary Update a Competition
// @Tags Competitions
// @Accept json
// @Produce json
// @Success 200 {object} models.competitions.Competition
// @Param competition body models.competitions.Competition true "Competition to update"
// @Router /v1/competition/{competitionId} [put]
func updateCompetition() {

}
