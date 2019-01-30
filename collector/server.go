package collector

import (
        "log"
        "net"
        "time"
        "io"

        "google.golang.org/grpc"

        pb "github.com/luanguimaraesla/memoir-gateway/metrics"
        "github.com/luanguimaraesla/memoir-gateway/exporter"
)

type collectorServer struct {
        measures []*pb.Measure
}

func (p *collectorServer) AddMeasure(stream pb.Metrics_AddMeasureServer) error {
        var measureCount int32
        startTime := time.Now()
        for {
                measure, err := stream.Recv()
                if err == io.EOF {
                        endTime := time.Now()
                        log.Printf("sending response and closing connection")
                        return stream.SendAndClose(&pb.GatewaySummary{
                                MeasureCount: measureCount,
                                ElapsedTime: int32(endTime.Sub(startTime).Seconds()),
                        })
                }
                if err != nil {
                        return err
                }
                measureCount++
                log.Printf("Received metric: {Name: %s, Value: %f}", measure.Name, measure.Value)
                exporter.AddMetric(measure)
        }
}

func RunCollectorServer(addr string) {
        lis, err := net.Listen("tcp", addr)
        if err != nil {
                log.Fatalf("failed to listen: %v", err)
        }
        log.Printf("listening gRPC on %s", addr)

        grpcServer := grpc.NewServer()
        pb.RegisterMetricsServer(grpcServer, &collectorServer{})
        if err := grpcServer.Serve(lis); err != nil {
                log.Fatalf("failed to serve: %v", err)
        }
}
