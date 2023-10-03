package datasource_handler

import (
	"fmt"
	data_source_dto_in "github.com/hramov/gvc-bi/backend/dashboard/internal/domain/data_source/dto/in"
	"github.com/hramov/gvc-bi/backend/dashboard/pkg/utils"
	"net/http"
)

func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	source, err := utils.GetBody[data_source_dto_in.Datasource](r)
	if err != nil {
		utils.SendError(http.StatusBadRequest, fmt.Sprintf("wrong body format: %v", err.Error()), w)
		return
	}

	id, err := h.service.Create(source)
	if err != nil {
		utils.SendError(http.StatusInternalServerError, fmt.Sprintf("cannot fetch data from database: %v", err.Error()), w)
		return
	}

	utils.SendResponse(http.StatusCreated, id, w)
}
