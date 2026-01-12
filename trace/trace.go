package trace

import (
	"github.com/kweaver-ai/TelemetrySDK-Go/exporter/v2/ar_trace"
	"github.com/kweaver-ai/TelemetrySDK-Go/exporter/v2/public"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func InitTracer(anyRobotURL, name, version, instance string) *sdktrace.TracerProvider {
	traceClient := public.NewHTTPClient(public.WithAnyRobotURL(anyRobotURL))
	traceExporter := ar_trace.NewExporter(traceClient)
	public.SetServiceInfo(name, version, instance)
	tracerProvider := sdktrace.NewTracerProvider(sdktrace.WithBatcher(traceExporter), sdktrace.WithResource(ar_trace.TraceResource()))
	otel.SetTracerProvider(tracerProvider)
	tracerPropagator := propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{})
	otel.SetTextMapPropagator(tracerPropagator)
	return tracerProvider
}
