package my

import (
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
	"github.com/uber/jaeger-client-go"
	"io"
	"net/http"
	"testing"
	"time"
)
import jaegercfg "github.com/uber/jaeger-client-go/config"

func TestJaegerA(t *testing.T) {
	//cfg, err := jaegercfg.FromEnv()
	//if err != nil {
	//	t.Fatal(err)
	//	return
	//}
	//cfg.ServiceName = "example"

	cfg := jaegercfg.Configuration{
		ServiceName: "TestJaegerA",
		Sampler: &jaegercfg.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: "192.168.1.40:6831", // Docker 中的 Jaeger 容器名
			//AgentPort: 6831,     // Jaeger Agent 默认端口
		},
	}

	tracer, closer, err := cfg.NewTracer(jaegercfg.Logger(jaeger.StdLogger))
	if err != nil {
		t.Fatal(err)
		return
	}
	defer closer.Close()

	opentracing.SetGlobalTracer(tracer)
	http.HandleFunc("/test_a", handlera)
	t.Error(http.ListenAndServe(":8080", nil))
}

func handlera(w http.ResponseWriter, r *http.Request) {
	// 从请求中提取 SpanContext
	spanCtx, _ := opentracing.GlobalTracer().Extract(
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(r.Header))

	// 创建一个新的 Span
	span := opentracing.StartSpan("http-server", ext.RPCServerOption(spanCtx))
	defer span.Finish()

	// 模拟处理时间
	//time.Sleep(time.Millisecond * 100)

	reqB, _ := http.NewRequest("GET", "http://127.0.0.1:8081/test_b", nil)
	opentracing.GlobalTracer().Inject(span.Context(), opentracing.HTTPHeaders, reqB.Header)
	httpClient := &http.Client{}
	respB, _ := httpClient.Do(reqB)
	defer respB.Body.Close()

	// 读取响应
	b, _ := io.ReadAll(respB.Body)

	// 记录一些日志
	span.LogFields(
		log.String("event", "request_received"),
		log.String("path", r.URL.Path),
	)
	// 返回响应
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func TestJaegerB(t *testing.T) {
	cfg := jaegercfg.Configuration{
		ServiceName: "TestJaegerB",
		Sampler: &jaegercfg.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: "192.168.1.40:6831", // Docker 中的 Jaeger 容器名
			//AgentPort: 6831,     // Jaeger Agent 默认端口
		},
	}

	tracer, closer, err := cfg.NewTracer(jaegercfg.Logger(jaeger.StdLogger))
	if err != nil {
		t.Fatal(err)
		return
	}
	defer closer.Close()

	opentracing.SetGlobalTracer(tracer)
	http.HandleFunc("/test_b", handlerb)
	t.Error(http.ListenAndServe(":8081", nil))
}
func handlerb(w http.ResponseWriter, r *http.Request) {
	// 从请求中提取 SpanContext
	spanCtx, _ := opentracing.GlobalTracer().Extract(
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(r.Header))

	// 创建一个新的 Span
	span := opentracing.StartSpan("http-server", ext.RPCServerOption(spanCtx))
	defer span.Finish()

	// 模拟处理时间
	time.Sleep(time.Millisecond * 100)

	// 记录一些日志
	span.LogFields(
		log.String("event", "request_received"),
		log.String("path", r.URL.Path),
	)

	// 返回响应
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("This serverB"))
}
