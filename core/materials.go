package core

import (
	"database/sql"
	"net/http"

	"github.com/alphatechnolog/purplish-materials/database"
	"github.com/gin-gonic/gin"
)

func getMaterials(d *sql.DB, c *gin.Context) error {
	materials, err := database.GetMaterials(d)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, gin.H{
		"materials": materials,
	})

	return nil
}

// Creates routes for the prefixed module `/materials`
func CreateMaterialsRoutes(d *sql.DB, r *gin.RouterGroup) {
	r.GET("/", WrapError(WithDB(d, getMaterials)))
}
