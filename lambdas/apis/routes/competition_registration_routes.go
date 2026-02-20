package routes

// @Summary Add an Athlete to a Competition
// @Tags Competitions
// @Accept json
// @Produce json
// @Success 200 {object} models.Registration
// @Param athlete body models.Registration true "Athlete to add"
// @Router /v1/competition/{competitionId}/register [post]
func createCompetitionAthlete() {

}

// @Summary Update an Athlete in a Competition
// @Tags Competitions
// @Accept json
// @Produce json
// @Success 200 {object} models.CompetitionAthlete
// @Param athlete body models.CompetitionAthlete true "Athlete to update"
// @Router /v1/competition/{competitionId}/athlete/{athleteId} [put]
func updateCompetitionAthlete() {

}

// @Summary Delete an Athlete from a Competition
// @Tags Competitions
// @Accept json
// @Produce json
// @Success 202
// @Router /v1/competition/{competitionId}/athlete/{athleteId} [delete]
func deleteCompetitionAthlete() {

}
