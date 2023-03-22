package main

import (
	"context"
	"log"
	"time"
	"io"
	"bytes"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

func dummy(c chan int) {
	time.Sleep(3 * time.Second)
	random := <-c
	log.Println("receiving--", random)
}

func traceProvider() (*tracesdk.TracerProvider, error) {
	var foo bytes.Buffer

	w := io.Writer(&foo)
	// Create the stdout exporter
	exp, err := stdouttrace.New(
		stdouttrace.WithWriter(w),
		// Use human-readable output.
		stdouttrace.WithPrettyPrint(),
		// Do not print timestamps for the demo.
		stdouttrace.WithoutTimestamps(),
	)
	if err != nil {
		return nil, err
	}

	tp := tracesdk.NewTracerProvider(
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// Record information about this application in a Resource.
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName("go-sample-app"),
			attribute.String("environment", "DEV"),
			attribute.Int64("ID", 123),
		)),
	)

	return tp, err;
}

func main() {
	tp, err := traceProvider()
	if err != nil {
		log.Fatal(err)
	}

	// Register our TracerProvider as the global so any imported
	// instrumentation in the future will default to using it.
	otel.SetTracerProvider(tp)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tr := tp.Tracer("component-main")
	ctx, span := tr.Start(ctx, "foo")
	defer span.End()

	ch := make(chan int)
	go dummy(ch)
	log.Println("sending--")
	ch <- 100
	log.Println("received--")
}
