package handler

import (
	"context"
	"github.com/hramov/gvc-bi/backend/datasource/internal/domain/data_source/connections"
	"github.com/hramov/gvc-bi/backend/datasource/pkg/utils"
	"net/http"
	"time"
)

func (h *Handler) recallDataSources(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 15*time.Second)
	defer cancel()

	ds, err := h.service.Get(ctx)
	if err != nil {
		utils.SendError(http.StatusInternalServerError, err.Error(), w)
		return
	}

	var rc []connections.RawConnection

	for _, v := range ds {
		rc = append(rc, connections.RawConnection{
			SourceId:   v.Id,
			DriverId:   v.DriverId,
			PluginName: v.PluginName,
			Dsn:        v.Dsn,
		})
	}

	errs := connections.Connect(ctx, rc)

	errStr := ""

	for _, e := range errs {
		errStr += e.Error() + ":"
	}

	if len(errs) > 0 {
		utils.SendError(http.StatusInternalServerError, errStr, w)
		return
	}

	utils.SendResponse(http.StatusOK, "OK", w)
}
