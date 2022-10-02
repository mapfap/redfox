package telemetry

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
	"log"
)

const (
	service     = "redfox"
	environment = "local"
	version     = "0.0.1"
	//url         = "http://localhost:14268/api/traces"
	url = "http://localhost:14278/api/traces"
)

func Init() (*tracesdk.TracerProvider, error) {
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	tp := tracesdk.NewTracerProvider(
		// TODO: Sampler in production
		tracesdk.WithSampler(tracesdk.AlwaysSample()),
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// Record information about this application in a Resource.
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(service),
			semconv.ServiceVersionKey.String(version),
			attribute.String("environment", environment),
		)),
	)
	otel.SetTracerProvider(tp)

	//ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()

	// TODO: Make sure cleanly shutdown and flush telemetry when the application exits.
	//defer func(ctx context.Context) {
	//	log.Println("[otel] ?????????")
	//	ctx, cancel = context.WithTimeout(ctx, 5*time.Second)
	//	defer cancel()
	//	if err := tp.Shutdown(ctx); err != nil {
	//		log.Fatal(err)
	//	} else {
	//		log.Println("[otel] gracefully shutdown")
	//	}
	//}(ctx)
	return tp, nil
}
