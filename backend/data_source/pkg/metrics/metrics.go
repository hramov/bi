package metrics

import (
	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	Namespace = "gvc"
	Subsystem = "dashboard"
)

var (
	RequestsTotal         *prometheus.CounterVec   // Распределение запросов по роутам
	InFlightRequestsTotal prometheus.Gauge         // Количество одновременно выполняющихся запросов
	SummaryResponseTime   *prometheus.SummaryVec   // Время ответа на запрос (квантили)
	HistogramResponseTime *prometheus.HistogramVec // Время ответа на запрос (гистограмма)
	UsersQueueTotal       prometheus.Gauge         // Размер очереди пользовательских запросов

	HistogramWorkerResponseTime *prometheus.HistogramVec // Время обработки сообщения воркером (гистограмма)
)

func InitMetrics() {

	// RequestsTotal Распределение запросов по роутам
	RequestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: Namespace,
			Subsystem: Subsystem,
			Name:      "requests_total",
		},
		[]string{"route", "step", "is_error"},
	)

	// SummaryResponseTime Время ответа на запрос (квантили)
	SummaryResponseTime = promauto.NewSummaryVec(
		prometheus.SummaryOpts{
			Namespace: Namespace,
			Subsystem: Subsystem,
			Name:      "summary_response_time_seconds",
			Objectives: map[float64]float64{
				0.5:  0.1,
				0.9:  0.01,
				0.99: 0.001,
			},
		},
		[]string{"route", "step", "is_error"},
	)

	// InFlightRequestsTotal Количество одновременно выполняющихся запросов
	InFlightRequestsTotal = promauto.NewGauge(
		prometheus.GaugeOpts{
			Namespace: Namespace,
			Subsystem: Subsystem,
			Name:      "in_flight_requests_total",
		},
	)

	// HistogramResponseTime Время ответа на запрос (гистограмма)
	HistogramResponseTime = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: Namespace,
			Subsystem: Subsystem,
			Name:      "histogram_response_time_seconds",
			Buckets:   []float64{0.0001, 0.0005, 0.001, 0.005, 0.01, 0.05, 0.1, 0.5, 1, 2},
		},
		[]string{"route", "step", "is_error"},
	)

	// UsersQueueTotal Размер очереди пользовательских запросов
	UsersQueueTotal = promauto.NewGauge(
		prometheus.GaugeOpts{
			Namespace: Namespace,
			Subsystem: Subsystem,
			Name:      "users_queue_total",
		},
	)

}

func HandleMetrics(r chi.Router) {
	r.Handle("/metrics", promhttp.Handler())
}
