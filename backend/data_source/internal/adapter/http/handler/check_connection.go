package handler

import (
	"context"
	"fmt"
	data_sourse_dto_in "github.com/hramov/gvc-bi/backend/datasource/internal/domain/data_source/dto/in"
	data_sourse_dto_out "github.com/hramov/gvc-bi/backend/datasource/internal/domain/data_source/dto/out"
	"github.com/hramov/gvc-bi/backend/datasource/pkg/database"
	"github.com/hramov/gvc-bi/backend/datasource/pkg/utils"
	"net/http"
	"time"
)

func (h *Handler) checkConnection(w http.ResponseWriter, r *http.Request) {

	dataSource, err := utils.GetBody[data_sourse_dto_in.DataStorageOptions](r)
	if err != nil {
		utils.SendError(http.StatusBadRequest, fmt.Sprintf("cannot parse body: %v", err.Error()), w)
		return
	}

	storage, err := database.NewStorage(database.DataStorageOptions{
		Driver:   dataSource.Driver,
		Host:     dataSource.Host,
		Port:     dataSource.Port,
		User:     dataSource.User,
		Password: dataSource.Password,
		Database: dataSource.Database,
		Sslmode:  dataSource.Sslmode,
	})

	if err != nil {
		utils.SendError(http.StatusBadRequest, fmt.Sprintf("cannot check connection: %v", err.Error()), w)
		return
	}
	defer storage.Close()

	result := data_sourse_dto_out.CheckResult{}

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
