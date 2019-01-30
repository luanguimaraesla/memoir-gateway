// Copyright 2015 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// A minimal example of how to include Prometheus instrumentation.
package exporter

import (
        "log"
        "net/http"

        "github.com/prometheus/client_golang/prometheus"
        "github.com/prometheus/client_golang/prometheus/promhttp"

        "github.com/luanguimaraesla/memoir-gateway/err"
        pb "github.com/luanguimaraesla/memoir-gateway/metrics"
)

func RunPrometheusServer(addr string) {
        http.Handle("/metrics", promhttp.Handler())
        log.Fatal(http.ListenAndServe(addr, nil))
}

func AddMetric(m *pb.Measure) error {
        switch m.Kind {
        case pb.Measure_GAUGE:
                addGauge(m)
        case pb.Measure_COUNTER:
                return err.NewError("counter parser not implemented")
        case pb.Measure_HISTOGRAM:
                return err.NewError("histogram parser not implemented")
        case pb.Measure_SUMMARY:
                return err.NewError("summary not implemented")
        default:
                return err.NewError("invalid metric type")
        }
        return nil
}

func addGauge(m *pb.Measure) {
        log.Printf("adding gauge metric: %v", m)
        gauges := prometheus.NewGaugeVec(
                        prometheus.GaugeOpts{
                                Name: m.Name,
                                Help: m.Help,
                        },
                        []string{"agent"},
                )
	prometheus.MustRegister(gauges)
        gauges.With(prometheus.Labels{"agent": "telegram"}).Set(float64(m.Value))
}
