package entx

import (
	"context"
	"database/sql"
	"log"

	"log/slog"
	"net/http"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.30.0"
)

func TestMySQL(t *testing.T) {
	initTracer()
	initMeter()

	driver, err := CreateDriver("mysql", mysqlDSN, true, true)
	if err != nil {
		t.Fatalf("failed opening connection to db: %v", err)
	}

	err = run(driver.DB())
	if err != nil {
		t.Fatalf("failed running query: %v", err)
	}

	slog.Info("Example finished updating, please visit localhost:2222/metrics")

	select {}
}

const instrumentationName = "github.com/moweilong/mo/entx/stdout"

var serviceName = semconv.ServiceNameKey.String("entx-test")

var mysqlDSN = "root:otel_password@tcp(127.0.0.1:33061)/db?parseTime=true"

func initTracer() {
	exporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		log.Fatal(err)
	}
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithSyncer(exporter),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			serviceName,
		)),
	)

	otel.SetTracerProvider(tp)
}

func initMeter() {
	// The exporter embeds a default OpenTelemetry Reader and
	// implements prometheus.Collector, allowing it to be used as
	// both a Reader and Collector.
	metricExporter, err := prometheus.New()
	if err != nil {
		log.Fatal(err)
	}
	meterProvider := metric.NewMeterProvider(
		metric.WithReader(metricExporter),
		metric.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			serviceName,
		)),
	)
	otel.SetMeterProvider(meterProvider)

	http.Handle("/metrics", promhttp.Handler())
	go func() {
		server := http.Server{
			Addr:              ":2222",
			ReadHeaderTimeout: 3 * time.Second,
		}
		_ = server.ListenAndServe()
	}()
	slog.Info("Prometheus server running on :2222")
}

func run(db *sql.DB) error {
	// Create a parent span (Optional)
	tracer := otel.GetTracerProvider()
	ctx, span := tracer.Tracer(instrumentationName).Start(context.Background(), "entx-example")
	defer span.End()

	err := query(ctx, db)
	if err != nil {
		span.RecordError(err)
		return err
	}
	return nil
}

func query(ctx context.Context, db *sql.DB) error {
	// Make a query
	rows, err := db.QueryContext(ctx, `SELECT CURRENT_TIMESTAMP`)
	if err != nil {
		return err
	}
	defer func() { _ = rows.Close() }()

	var currentTime time.Time
	for rows.Next() {
		err = rows.Scan(&currentTime)
		if err != nil {
			return err
		}
	}
	// Check for errors from iterating over rows
	if err = rows.Err(); err != nil {
		return err
	}
	slog.Info("Current time", "time", currentTime)
	return nil
}
