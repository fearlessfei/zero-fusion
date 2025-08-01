package otelmetric

import (
	"fmt"

	runtimemetrics "go.opentelemetry.io/contrib/instrumentation/runtime"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/sdk/metric"
)

func InitPrometheusMeterProvider() {
	exporter, err := prometheus.New()
	if err != nil {
		panic(fmt.Sprintf("new prometheus exporter err: %v", err))
	}

	provider := metric.NewMeterProvider(
		metric.WithReader(exporter),
	)
	otel.SetMeterProvider(provider)

	if err = runtimemetrics.Start(); err != nil {
		panic(fmt.Sprintf("start runtime metrices err: %v", err))
	}
}
