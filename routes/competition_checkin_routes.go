package routes

// @Summary Add Checkin Requirements to a Competition
// @Tags Competitions
// @Accept json
// @Produce json
// @Success 200 {object} models.Checkin
// @Param checkin body models.Checkin true "Checkin requirements for the competition"
// @Router /v1/competition/{competitionId}/checkin [post]
func addCompetitionCheckin() {

}

// @Summary Update Checkin Requirements for a Competition
// @Tags Competitions
// @Accept json
// @Produce json
// @Success 200 {object} models.Checkin
// @Param checkin body models.Checkin true "Updated checkin requirements for the competition"
// @Router /v1/competition/{competitionId}/checkin [put]
func updateCompetitionCheckin() {

}

// @Summary Delete the checkin for a Competition
// @Tags Competitions
// @Accept json
// @Produce json
// @Success 202
// @Router /v1/competition/{competitionId}/checkin [delete]
func deleteCompetitionCheckin() {

}
