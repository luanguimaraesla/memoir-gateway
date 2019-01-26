// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
        "log"

	"github.com/spf13/cobra"
        "github.com/luanguimaraesla/memoir-gateway/prometheus"
        "github.com/luanguimaraesla/memoir-gateway/collector"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run memoir prometheus matrics gateway server",
	Long: `Metrics-gateway exports listen to your gRPC to
create and expose custom timeseries using the
Open Metrics Prometheus format.`,
	Run: func(cmd *cobra.Command, args []string) {
                prometheus_addr, err := cmd.Flags().GetString("prometheus")
                if err != nil {
                        log.Panic("Can't start http server on", prometheus_addr)
                }

                grpc_addr, err := cmd.Flags().GetString("collector")
                if err != nil {
                        log.Panic("Can't start tcp server on", grpc_addr)
                }
                go func() {
                        prometheus.RunPrometheusServer(prometheus_addr)
                }()

                collector.RunCollectorServer(grpc_addr)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
        runCmd.Flags().StringP("prometheus", "p", ":9090", "prometheus bind address")
        runCmd.Flags().StringP("collector", "c", ":5000", "grpc bind address")
}
