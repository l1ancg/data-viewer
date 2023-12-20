package handler

import (
	"github.com/l1ancg/data-viewer/backend/internal/application"
	"github.com/l1ancg/data-viewer/backend/internal/repository"
	"net/http"

	"github.com/l1ancg/data-viewer/backend/pkg"
	"github.com/l1ancg/data-viewer/backend/pkg/connect"
	"github.com/l1ancg/data-viewer/backend/pkg/log"
	"github.com/l1ancg/data-viewer/backend/pkg/utils"
	"github.com/labstack/echo/v4"
)

type QueryHandler struct {
	pkg.AbstractHandler
}

type QueryParams struct {
	Ql         string                 `json:"ql"`
	ResourceId int                    `json:"resourceId"`
	Parameter  map[string]interface{} `json:"parameter"`
}

func QueryHandlerProvider(dispatcher *connect.ConnectDispatcher, db *repository.Database) *QueryHandler {
	return &QueryHandler{
		AbstractHandler: pkg.AbstractHandler{
			Method: "POST",
			Path:   "/query",
			Handler: func(c echo.Context) error {
				params := new(QueryParams)
				if err := c.Bind(params); err != nil {
					return err
				}
				log.Logger.Info("query params:", params)
				sql, err := utils.Parse(params.Ql, params.Parameter)
				if err != nil {
					return err
				}

				var res application.Resource
				db.Select(&res, params.ResourceId)

				log.Logger.Info("query sql:", sql)
				result, err := dispatcher.Query("MySQL", params.ResourceId, res.Data, sql)
				if err != nil {
					return err
				}
				return c.JSON(http.StatusOK, result)
			},
		},
	}
}
