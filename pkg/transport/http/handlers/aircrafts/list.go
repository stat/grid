package aircrafts

import (
	"net/http"

	"events/pkg/models"
	"events/pkg/transport/http/respond"
	"events/pkg/utils"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	uriParams := &URIParams{}

	if err := c.ShouldBindUri(uriParams); err != nil {
		respond.WithError(c, http.StatusBadRequest, err)
		return
	}

	// lookup

	locations := []*models.LocationEvent{}
	// locations, err := cache.Backend.GetAircraftsLocations()

	// if err != nil {
	//   respond.WithError(c, http.StatusNotFound, err)
	//   return
	// }

	// success

	respond.With(c, http.StatusOK, utils.Ref(locations))
}
