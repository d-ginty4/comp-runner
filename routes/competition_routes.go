package routes

import (
	"apis/models/competition"
	"apis/services"
	"github.com/labstack/echo/v5"
	"net/http"
)

type CompetitionRoutes struct {
	ListCompetitions  func(c *echo.Context) error
	CreateCompetition func(c *echo.Context) error
	GetCompetition    func(c *echo.Context) error
	UpdateCompetition func(c *echo.Context) error
}

var competitionRoutes *CompetitionRoutes

func GetCompetitionRoutes() *CompetitionRoutes {
	if competitionRoutes == nil {
		return newCompetitionRoutes()
	}
	return competitionRoutes
}

func newCompetitionRoutes() *CompetitionRoutes {
	routes := &CompetitionRoutes{
		ListCompetitions:  listCompetitions,
		CreateCompetition: createCompetition,
		GetCompetition:    getCompetition,
		UpdateCompetition: updateCompetition,
	}

	competitionRoutes = routes
	return routes
}

func init() {
	GetCompetitionRoutes()
}

func registerCompetitionRoutes(g *echo.Group) {
	route := g.Group("/competition")
	route.GET("", competitionRoutes.ListCompetitions)
	route.POST("", competitionRoutes.CreateCompetition)

	routeId := route.Group("/:competitionId")
	routeId.GET("", competitionRoutes.GetCompetition)
	routeId.PUT("", competitionRoutes.UpdateCompetition)
}

// @Summary List Competitions
// @Tags Competitions
// @Accept json
// @Produce json
// @Success 200 {object} models.competitions.Competition
// @Router /v1/competition [get]
func listCompetitions(c *echo.Context) error {
	service := services.GetCompetitionService()

	competitions, err := service.ListCompetition(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}

	return c.JSON(200, competitions)
}

// @Summary Create a Competition
// @Tags Competitions
// @Accept json
// @Produce json
// @Success 200 {object} models.competitions.Competition
// @Param competition body models.competitions.CompetitionRequest true "Competition to create"
// @Router /v1/competition [post]
func createCompetition(c *echo.Context) error {
	service := services.GetCompetitionService()

	competitionRequest := new(competition.CompetitionRequest)
	if err := c.Bind(competitionRequest); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}

	competition, err := service.CreateCompetition(c.Request().Context(), competitionRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}

	return c.JSON(200, competition)
}

// @Summary Get a Competition
// @Tags Competitions
// @Accept json
// @Produce json
// @Success 200 {object} models.competitions.Competition
// @Router /v1/competition/{competitionId} [get]
func getCompetition(c *echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "missing goal id")
	}

	service := services.GetCompetitionService()

	competition, err := service.GetCompetition(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}

	return c.JSON(200, competition)
}

// @Summary Update a Competition
// @Tags Competitions
// @Accept json
// @Produce json
// @Success 200 {object} models.competitions.Competition
// @Param competition body models.competitions.Competition true "Competition to update"
// @Router /v1/competition/{competitionId} [put]
func updateCompetition(c *echo.Context) error {
	service := services.GetCompetitionService()

	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "missing goal id")
	}

	competitionRequest := new(competition.CompetitionRequest)
	if err := c.Bind(competitionRequest); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}

	competition, err := service.UpdateCompetition(c.Request().Context(), id, competitionRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}

	return c.JSON(200, competition)
}
