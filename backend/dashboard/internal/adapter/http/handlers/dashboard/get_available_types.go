package dashboard_handler

import (
	"github.com/hramov/gvc-bi/backend/dashboard/pkg/utils"
	"net/http"
)

func (h *Handler) getAvailableTypes(w http.ResponseWriter, r *http.Request) {
	data, err := h.service.GetAvailableTypes()
	if err != nil {
		utils.SendError(http.StatusInternalServerError, err.Error(), w)
		return
	}
	utils.SendResponse(http.StatusOK, data, w)
}
