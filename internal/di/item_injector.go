package di

import (
	"database/sql"

	"github.com/alphatechnolog/purplish-items/delivery/http"
	"github.com/alphatechnolog/purplish-items/infrastructure/database"
	"github.com/alphatechnolog/purplish-items/internal/usecase"
	"github.com/gin-gonic/gin"
)

type ItemInjector struct {
	db *sql.DB
}

func NewItemInjector(db *sql.DB) ModuleInjector {
	return &ItemInjector{db}
}

func (ii *ItemInjector) Inject(routerGroup *gin.RouterGroup) {
	sqliteRepo := database.NewSQLiteRepository(ii.db)
	itemUseCase := usecase.NewItemUsecase(sqliteRepo)
	itemHandler := http.NewItemHandler(itemUseCase)

	routerGroup.GET("/", http.APIGatewayScopeCheck("r:items"), itemHandler.GetItems)
	routerGroup.GET("/:id/", http.APIGatewayScopeCheck("r:items"), itemHandler.GetItem)
	routerGroup.GET("/for-warehouse/:id/", http.APIGatewayScopeCheck("r:items"), itemHandler.GetItemsByWarehouse)
	routerGroup.POST("/", http.APIGatewayScopeCheck("c:items"), itemHandler.CreateItem)
	routerGroup.POST("/assign-to-warehouse/", http.APIGatewayScopeCheck("c:items"), itemHandler.AssignToWarehouse)
	routerGroup.PUT("/:id/", http.APIGatewayScopeCheck("u:items"), itemHandler.UpdateItem)
	routerGroup.DELETE("/:id/", http.APIGatewayScopeCheck("d:items"), itemHandler.DeleteItem)
}
