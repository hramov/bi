package handler

import (
	"context"
	"fmt"
	"github.com/hramov/gvc-bi/backend/datastorage/pkg/database"
	"github.com/hramov/gvc-bi/backend/datastorage/pkg/utils"
	"net/http"
	"time"
)

func (h *Handler) checkConnection(w http.ResponseWriter, r *http.Request) {

	dataSource, err := utils.GetBody[database.DataStorageOptions](r)
	if err != nil {
		utils.SendError(http.StatusBadRequest, fmt.Sprintf("cannot parse body: %v", err.Error()), w)
		return
	}

	storage, err := database.NewStorage(dataSource)
	if err != nil {
		utils.SendError(http.StatusBadRequest, fmt.Sprintf("cannot check connection: %v", err.Error()), w)
		return
	}
	defer storage.Close()

	result := CheckResult{}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = storage.PingContext(ctx)

	if err != nil {
		result.Status = "declined"
		result.Message = err.Error()
		utils.SendResponse(http.StatusOK, result, w)
		return
	}

	result.Status = "accepted"
	utils.SendResponse(http.StatusOK, result, w)
}
