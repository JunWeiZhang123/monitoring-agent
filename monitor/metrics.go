package monitor
import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
	// prom_v1 "github.com/vinkdong/monitoring-agent/prometheus/v1"
)

var (
	http_request_count = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Help: "http_count",
			Name: "http_count"}, []string{"endpoint"},)
)

func prometheusInit() {
	prometheus.MustRegister(http_request_count)
}

func startPrometheus()  {
	http.Handle("/metrics", promhttp.Handler())
} 