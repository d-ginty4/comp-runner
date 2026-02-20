package routes

// @Summary Add an Event to a Competition
// @Tags Competitions
// @Accept json
// @Produce json
// @Success 200 {object} models.Event
// @Param event body models.Event true "Event to add"
// @Router /v1/competition/{competitionId}/event [post]
func createCompetitionEvent() {

}

// @Summary Update an Event in a Competition
// @Tags Competitions
// @Accept json
// @Produce json
// @Success 200 {object} models.Event
// @Param event body models.Event true "Event to update"
// @Router /v1/competition/{competitionId}/event/{eventId} [put]
func updateCompetitionEvent() {

}

// @Summary Delete an Event from a Competition
// @Tags Competitions
// @Accept json
// @Produce json
// @Success 202
// @Router /v1/competition/{competitionId}/event/{eventId} [delete]
func deleteCompetitionEvent() {

}
