package datasource_handler

import (
	"context"
	"fmt"
	data_source_dto_out "github.com/hramov/gvc-bi/backend/dashboard/internal/domain/data_source/dto/out"
	"github.com/hramov/gvc-bi/backend/dashboard/pkg/utils"
	"net/http"
	"time"
)

func (h *Handler) getDrivers(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 15*time.Second)
	defer cancel()

	r = r.WithContext(ctx)

	drivers, err := h.service.GetDrivers(ctx)
	if err != nil {
		utils.SendError(http.StatusInternalServerError, fmt.Sprintf("cannot fetch data from database: %v", err.Error()), w)
		return
	}

	response := []*data_source_dto_out.Driver{}

	for _, v := range drivers {
		response = append(response, &data_source_dto_out.Driver{
			Id:          v.Id,
			Title:       v.Title,
			Code:        v.Code,
			DateCreated: v.DateCreated,
			DbNeed:      v.DbNeed,
		})
	}
	utils.SendResponse(http.StatusOK, response, w)
}
