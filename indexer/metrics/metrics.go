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
	SyncHeight *prometheus.GaugeVec
}

func NewMetrics(monitoredTokens map[string]string) *Metrics {
	mts := make(map[string]string)
	mts["0x0000000000000000000000000000000000000000"] = "ETH"
	for addr, symbol := range monitoredTokens {
		mts[addr] = symbol
	}

	return &Metrics{
		SyncHeight: promauto.NewGaugeVec(prometheus.GaugeOpts{
			Name:      "sync_height",
			Help:      "The max height of the indexer's last batch of L1/L1 blocks.",
			Namespace: metricsNamespace,
		}, []string{
			"chain",
		}),
	}
}

func (m *Metrics) SetL1SyncHeight(height uint64) {
	m.SyncHeight.WithLabelValues("l1").Set(float64(height))
}

func (m *Metrics) SetL2SyncHeight(height uint64) {
	m.SyncHeight.WithLabelValues("l2").Set(float64(height))
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
