package cmd

import (
	"log"
	"os"

	"github.com/ennc0d3/yt-stats/internal/api"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/trace"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the API server",
	Long:  `Start the HTTP API server to serve YouTube statistics via REST endpoints.`,
	Run: func(cmd *cobra.Command, args []string) {
		runServer()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

func runServer() {
	// Configure zerolog
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	
	// Set log level from config
	level, err := zerolog.ParseLevel(viper.GetString("logLevel"))
	if err != nil {
		level = zerolog.InfoLevel
	}
	zerolog.SetGlobalLevel(level)
	
	logger := log.New(zerolog.ConsoleWriter{Out: os.Stderr}, "", 0)

	// Check API key
	apiKey := viper.GetString("apiKey")
	if apiKey == "" {
		logger.Fatal("YouTube API key is required. Set YTSTATS_API_KEY environment variable or use --api-key flag")
	}

	// Initialize OpenTelemetry tracing
	if err := initTracingExporter(); err != nil {
		logger.Fatal("failed to initialize OpenTelemetry tracing:", err)
	}

	// Initialize Prometheus metrics
	if err := initMetricExporter(); err != nil {
		logger.Fatal("failed to initialize Prometheus metrics:", err)
	}

	// Start the server
	logger.Printf("Starting server on port %d", viper.GetInt("port"))
	api.StartServer()
}

func initTracingExporter() error {
	traceExporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		return err
	}

	tp := trace.NewTracerProvider(
		trace.WithSyncer(traceExporter),
		trace.WithSampler(trace.AlwaysSample()),
	)

	otel.SetTracerProvider(tp)
	return nil
}

func initMetricExporter() error {
	metricExporter, err := prometheus.New()
	if err != nil {
		return err
	}

	provider := metric.NewMeterProvider(metric.WithReader(metricExporter))
	provider.Meter("youtube-info-api")
	return nil
}
