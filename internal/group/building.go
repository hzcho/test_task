package group

import (
	"buildings/internal/domain/model"
	"buildings/internal/domain/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Building struct {
	bldUsecase usecase.Building
}

func NewBuilding(bldUsecase usecase.Building) *Building {
	return &Building{
		bldUsecase: bldUsecase,
	}
}

// @Summary get buildings
// @Tags buildings
// @Description Get a list of buildings by filter
// @ID get-buildings
// @Produce json
// @Param city query string false "City"
// @Param year query int false "Year"
// @Param floors query int false "number of floors"
// @Success 200 {array} model.Building
// @Failure 400 {string} string "Invalid request format"
// @Failure 500 {string} string "Server error"
// @Router /api/v1/buildings [get]
func (b *Building) Get(c *gin.Context) {
	city := c.Query("city")

	yearStr := c.Query("year")
	var year int
	if yearStr != "" {
		var err error
		year, err = strconv.Atoi(yearStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, "incorrect format of the 'year' field")
			return
		}
	}

	floorsStr := c.Query("floors")
	var floors int
	if floorsStr != "" {
		var err error
		floors, err = strconv.Atoi(floorsStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, "incorrect format of the 'floors' field")
			return
		}
	}

	filter := model.Filter{
		City:   city,
		Year:   year,
		Floors: floors,
	}

	buildings, err := b.bldUsecase.Get(c, filter)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "couldn't get the buildings")
		return
	}

	c.JSON(http.StatusOK, buildings)
}

// @Summary Save the building
// @Tags buildings
// @Description Save the new building
// @ID save-building
// @Accept json
// @Produce json
// @Param building body model.Building true "Building"
// @Success 200 {integer} int "ID of the saved building"
// @Failure 400 {string} string "Invalid request format"
// @Failure 500 {string} string "Server error"
// @Router /api/v1/buildings [post]
func (b *Building) Save(c *gin.Context) {
	building := model.Building{}

	if err := c.ShouldBindJSON(&building); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "bad request")
		return
	}

	id, err := b.bldUsecase.Save(c, building)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "couldn't save the building")
		return
	}

	c.JSON(http.StatusOK, id)
}
