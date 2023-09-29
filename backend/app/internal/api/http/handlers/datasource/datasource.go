package datasource

import (
	"database/sql"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/hramov/gvc-bi/backend/pkg/database/postgres"
	"github.com/hramov/gvc-bi/backend/pkg/utils"
	"net/http"
	"strconv"
	"time"
)

type Repository interface {
	GetDrivers() ([]*Driver, error)
	GetDriverById(id int) (*Driver, error)
	Get() ([]*Datasource, error)
	GetById(id int) (*Datasource, error)
	Create(driver Datasource) (*int, error)
	Update(driver Datasource, id int) (*int, error)
	Delete(id int) (*int, error)
}

type Driver struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Code        string    `json:"code"`
	DateCreated time.Time `json:"date_created"`
	DbNeed      bool      `json:"db_need"`
}

type Datasource struct {
	Id          int       `json:"id"`
	Driver      string    `json:"driver"`
	DriverId    int       `json:"driver_id"`
	Title       string    `json:"title"`
	Dsn         string    `json:"dsn"`
	Checked     bool      `json:"checked"`
	DateCreated time.Time `json:"date_created"`
}

func (d *Datasource) Connect() (*sql.DB, error) {
	return postgres.New(nil, d.Dsn)
}

type Handler struct {
	Repository Repository
}

func New(repo Repository) *Handler {
	return &Handler{Repository: repo}
}

func (h *Handler) Register(r chi.Router) {
	r.Get("/driver", h.getDrivers)
	r.Get("/driver/{id}", h.getDriverById)
	r.Get("/", h.get)
	r.Get("/{id}", h.getById)
	r.Post("/", h.create)
	r.Put("/{id}", h.update)
	r.Delete("/{id}", h.delete)
}

func (h *Handler) getDrivers(w http.ResponseWriter, r *http.Request) {
	drivers, err := h.Repository.GetDrivers()
	if err != nil {
		utils.SendError(http.StatusInternalServerError, fmt.Sprintf("cannot fetch data from database: %v", err.Error()), w)
		return
	}
	utils.SendResponse(http.StatusOK, drivers, w)
}

func (h *Handler) getDriverById(w http.ResponseWriter, r *http.Request) {
	rawId := chi.URLParam(r, "id")
	if rawId == "" {
		utils.SendError(http.StatusBadRequest, "need to pass id parameter", w)
		return
	}

	id, err := strconv.Atoi(rawId)
	if err != nil {
		utils.SendError(http.StatusBadRequest, fmt.Sprintf("wrong id format: %v", err.Error()), w)
		return
	}

	driver, err := h.Repository.GetDriverById(id)
	if err != nil {
		utils.SendError(http.StatusInternalServerError, fmt.Sprintf("cannot fetch data from database: %v", err.Error()), w)
		return
	}
	utils.SendResponse(http.StatusOK, driver, w)
}

func (h *Handler) get(w http.ResponseWriter, r *http.Request) {
	drivers, err := h.Repository.Get()
	if err != nil {
		utils.SendError(http.StatusInternalServerError, fmt.Sprintf("cannot fetch data from database: %v", err.Error()), w)
		return
	}
	utils.SendResponse(http.StatusOK, drivers, w)
}

func (h *Handler) getById(w http.ResponseWriter, r *http.Request) {
	rawId := chi.URLParam(r, "id")
	if rawId == "" {
		utils.SendError(http.StatusBadRequest, "need to pass id parameter", w)
		return
	}

	id, err := strconv.Atoi(rawId)
	if err != nil {
		utils.SendError(http.StatusBadRequest, fmt.Sprintf("wrong id format: %v", err.Error()), w)
		return
	}

	driver, err := h.Repository.GetById(id)
	if err != nil {
		utils.SendError(http.StatusInternalServerError, fmt.Sprintf("cannot fetch data from database: %v", err.Error()), w)
		return
	}
	utils.SendResponse(http.StatusOK, driver, w)
}

func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	source, err := utils.GetBody[Datasource](r)
	if err != nil {
		utils.SendError(http.StatusBadRequest, fmt.Sprintf("wrong body format: %v", err.Error()), w)
		return
	}

	id, err := h.Repository.Create(source)
	if err != nil {
		utils.SendError(http.StatusInternalServerError, fmt.Sprintf("cannot fetch data from database: %v", err.Error()), w)
		return
	}

	utils.SendResponse(http.StatusCreated, id, w)
}

func (h *Handler) update(w http.ResponseWriter, r *http.Request) {
	rawId := chi.URLParam(r, "id")
	if rawId == "" {
		utils.SendError(http.StatusBadRequest, "need to pass id parameter", w)
		return
	}

	id, err := strconv.Atoi(rawId)
	if err != nil {
		utils.SendError(http.StatusBadRequest, fmt.Sprintf("wrong id format: %v", err.Error()), w)
		return
	}

	source, err := utils.GetBody[Datasource](r)
	if err != nil {
		utils.SendError(http.StatusBadRequest, fmt.Sprintf("wrong body format: %v", err.Error()), w)
		return
	}

	updatedId, err := h.Repository.Update(source, id)
	if err != nil {
		utils.SendError(http.StatusInternalServerError, fmt.Sprintf("cannot fetch data from database: %v", err.Error()), w)
		return
	}

	utils.SendResponse(http.StatusCreated, updatedId, w)
}

func (h *Handler) delete(w http.ResponseWriter, r *http.Request) {
	rawId := chi.URLParam(r, "id")
	if rawId == "" {
		utils.SendError(http.StatusBadRequest, "need to pass id parameter", w)
		return
	}

	id, err := strconv.Atoi(rawId)
	if err != nil {
		utils.SendError(http.StatusBadRequest, fmt.Sprintf("wrong id format: %v", err.Error()), w)
		return
	}

	deletedId, err := h.Repository.Delete(id)
	if err != nil {
		utils.SendError(http.StatusInternalServerError, fmt.Sprintf("cannot fetch data from database: %v", err.Error()), w)
		return
	}

	utils.SendResponse(http.StatusCreated, deletedId, w)
}
