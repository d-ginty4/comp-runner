package routes

// @Summary Add a Category to a Competition
// @Tags Competitions
// @Accept json
// @Produce json
// @Success 200 {object} models.Category
// @Param category body models.Category true "Category to add"
// @Router /v1/competition/{competitionId}/category [post]
func addCompetitionCategory() {

}

// @Summary Update a Category in a Competition
// @Tags Competitions
// @Accept json
// @Produce json
// @Success 200 {object} models.Category
// @Param category body models.Category true "Category to update"
// @Router /v1/competition/{competitionId}/category/{categoryId} [put]
func updateCompetitionCategory() {

}

// @Summary Update a Category in a Competition
// @Tags Competitions
// @Accept json
// @Produce json
// @Success 202
// @Router /v1/competition/{competitionId}/category/{categoryId} [delete]
func deleteCompetitionCategory() {

}
