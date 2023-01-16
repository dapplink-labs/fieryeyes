package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net"
	"net/http"
	"strconv"
)

const metricsNamespace = "indexer"

type Metrics struct {
	SyncAddressInfo *prometheus.GaugeVec
}

func NewMetrics() *Metrics {
	return &Metrics{
		SyncAddressInfo: promauto.NewGaugeVec(prometheus.GaugeOpts{
			Name:      "sync address info",
			Help:      "address info for nft holder",
			Namespace: metricsNamespace,
		}, []string{
			"chain",
		}),
	}
}

func (m *Metrics) SetAddressInfo(height uint64) {
	m.SyncAddressInfo.WithLabelValues("nft").Set(1)
}

func (m *Metrics) Serve(hostname string, port uint64) (*http.Server, error) {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	srv := new(http.Server)
	srv.Addr = net.JoinHostPort(hostname, strconv.FormatUint(port, 10))
	srv.Handler = mux
	err := srv.ListenAndServe()
	return srv, err
}
