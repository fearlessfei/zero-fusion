package initialize

import (
	"zero-fusion/share/metric/otelmetric"
)

func InitOTELMeter(enable bool) {
	if enable {
		otelmetric.InitPrometheusMeterProvider()
	}
}
