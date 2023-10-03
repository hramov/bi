package datasource_handler

import (
	"fmt"
	data_source_dto_out "github.com/hramov/gvc-bi/backend/dashboard/internal/domain/data_source/dto/out"
	"github.com/hramov/gvc-bi/backend/dashboard/pkg/utils"
	"net/http"
)

func (h *Handler) getDrivers(w http.ResponseWriter, r *http.Request) {
	drivers, err := h.service.GetDrivers()
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
