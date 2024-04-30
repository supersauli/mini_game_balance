package my

import (
	"context"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/sdk/trace"
	"log"
	"net/http"
	"testing"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
)

func TestA(t *testing.T) {
	// 创建和配置 OTLP exporter
	ctx := context.Background()
	exporter, err := otlptracehttp.New(ctx,
		//otlptracehttp.WithEndpointURL("http://192.168.1.40:4318"),
		otlptracehttp.WithEndpointURL("http://192.168.1.40:24318"),
		//otlptracehttp.WithEndpointURL("http://192.168.136.132:9092"),
		////otlptracehttp.WithEndpointURL("http://192.168.136.132:8080/otlp-http/v1/traces"),
		//otlptracehttp.WithHeaders(map[string]string{"xxxx": "aaa"}),
	)

	if err != nil {
		log.Fatalf("failed to create OTLP exporter: %v", err)
	}
	defer exporter.Shutdown(context.Background())

	// 初始化 HTTP 服务器

	tracerProvider := trace.NewTracerProvider(trace.WithBatcher(exporter))
	defer func() {
		if err := tracerProvider.Shutdown(ctx); err != nil {
			panic(err)
		}
	}()
	otel.SetTracerProvider(tracerProvider)
	http.HandleFunc("/test_a", handleRequest)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	tracer := otel.Tracer("adservice")
	ctx, span := tracer.Start(r.Context(), "handleRequest")
	defer span.End()
	_ = ctx
	span.SetAttributes(attribute.String("name", "this is a test"))
	span.SetStatus(codes.Ok, "this is ok")
	span.AddEvent("eeeeee")

	// 模拟处理时间
	time.Sleep(1 * time.Second)

	// Your code here

	w.Write([]byte("Hello, World!"))
}
