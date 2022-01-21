package promtheus

import (
	"github.com/go-kit/kit/metrics"
	kitprom "github.com/go-kit/kit/metrics/prometheus"
	"github.com/prometheus/client_golang/prometheus"
)

type Metrics struct {
	metrics.Counter
	metrics.Histogram
}

func NewCounter(subsystem string, methodName string, helpMsg string) metrics.Counter {
	return kitprom.NewCounterFrom(prometheus.CounterOpts{
		Subsystem:   subsystem,
		Name:        methodName + "_count",
		Help:        helpMsg,
	}, []string{subsystem + "_counter"})
}

func NewHistogram(subsystem string, methodName string, helpMsg string) metrics.Histogram {
	return kitprom.NewHistogramFrom(prometheus.HistogramOpts{
		Subsystem:   subsystem,
		Name:        methodName + "_consume",
		Help:        helpMsg,
	}, []string{subsystem + "_histogram"})
}
