package datasource_handler

import (
	"fmt"
	data_source_dto_out "github.com/hramov/gvc-bi/backend/dashboard/internal/domain/data_source/dto/out"
	"github.com/hramov/gvc-bi/backend/dashboard/pkg/utils"
	"net/http"
)

func (h *Handler) get(w http.ResponseWriter, r *http.Request) {
	sources, err := h.service.Get()
	if err != nil {
		utils.SendError(http.StatusInternalServerError, fmt.Sprintf("cannot fetch data from database: %v", err.Error()), w)
		return
	}

	response := []*data_source_dto_out.Datasource{}

	for _, v := range sources {
		response = append(response, &data_source_dto_out.Datasource{
			Id:          v.Id,
			Driver:      v.Driver,
			DriverId:    v.DriverId,
			Title:       v.Title,
			Dsn:         v.Dsn,
			Checked:     v.Checked,
			DateCreated: v.DateCreated,
		})
	}

	utils.SendResponse(http.StatusOK, response, w)
}
