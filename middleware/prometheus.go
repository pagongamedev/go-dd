// Develope by MongMX

package middleware

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

// Prometheus of prometheus monitor
type Prometheus struct {
	Namespace   string
	Subsystem   string
	MetricPath  string
	reqCount    *prometheus.CounterVec
	reqDuration *prometheus.HistogramVec
}

// HandlerPrometheus for prometheus collect data
func (m *Prometheus) HandlerPrometheus() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if c.Path() == m.MetricPath {
			return c.Next()
		}

		start := time.Now()

		c.Next()

		r := c.Route()

		statusCode := strconv.Itoa(c.Context().Response.StatusCode())
		elapsed := float64(time.Since(start)) / float64(time.Second)

		m.reqCount.WithLabelValues(statusCode, c.Method(), r.Path).Inc()
		m.reqDuration.WithLabelValues(c.Method(), r.Path).Observe(elapsed)
		return nil
	}
}

// Register app for metrics collect
func (m *Prometheus) Register(app *fiber.App) {
	m.registerDefaultMetrics()
	app.Use(m.HandlerPrometheus())
}

// SetupPath for metrics view
func (m *Prometheus) SetupPath(app *fiber.App) {
	app.Get(m.MetricPath, m.metricHandler)
}

func (m *Prometheus) metricHandler(c *fiber.Ctx) error {
	p := fasthttpadaptor.NewFastHTTPHandler(promhttp.Handler())
	p(c.Context())
	return nil
}

func (m *Prometheus) registerDefaultMetrics() {
	m.reqCount = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name:      "requests_total",
			Namespace: m.Namespace,
			Subsystem: m.Subsystem,
			Help:      "Number of HTTP requests",
		},
		[]string{"status_code", "method", "path"},
	)

	m.reqDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:      "request_duration_seconds",
		Namespace: m.Namespace,
		Subsystem: m.Subsystem,
		Help:      "Duration of HTTP requests",
	}, []string{"method", "handler"})
}

// NewPrometheus for prometheus
func NewPrometheus(namespace string, subsystem string) *Prometheus {
	return &Prometheus{
		Namespace: namespace,
		Subsystem: subsystem,
	}
}
