package main

import (
	"context"
	"testing"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func BenchmarkTracerSpanOperations(b *testing.B) {
	b.Setenv("JAEGER_AGENT_HOST", "jaeger")
	b.Setenv("JAEGER_AGENT_PORT", "6831")
	b.Setenv("JAEGER_SAMPLER_TYPE", "const")
	b.Setenv("JAEGER_SAMPLER_PARAM", "1")
	b.Setenv("JAEGER_TAGS", "app=mimir1")
	InitOpenTracingTracer()
	for i := 0; i < b.N; i++ {
		ctx := context.Background()
		sp := opentracing.GlobalTracer().StartSpan("test1", opentracing.Tag{Key: "organization", Value: "user1"})
		defer sp.Finish()
		ctx = opentracing.ContextWithSpan(ctx, sp)
		sp2, _ := opentracing.StartSpanFromContext(ctx, "test2")
		defer sp2.Finish()
		sp2.LogFields(log.String("event", "soft error"))
	}
}

func BenchmarkTracerSpanOperationsWithOtelBridge(b *testing.B) {
	b.Setenv("JAEGER_AGENT_HOST", "localhost")
	b.Setenv("JAEGER_AGENT_PORT", "6831")
	b.Setenv("JAEGER_SAMPLER_TYPE", "const")
	b.Setenv("JAEGER_SAMPLER_PARAM", "1")
	b.Setenv("JAEGER_TAGS", "app=mimir2")
	InitOpentelemetryTracer()
	for i := 0; i < b.N; i++ {
		ctx := context.Background()
		sp := opentracing.GlobalTracer().StartSpan("test1", opentracing.Tag{Key: "organization", Value: "user1"})
		defer sp.Finish()
		ctx = opentracing.ContextWithSpan(ctx, sp)
		sp2, _ := opentracing.StartSpanFromContext(ctx, "test2")
		defer sp2.Finish()
		sp2.LogFields(log.String("event", "soft error"))
	}
}

func BenchmarkOtelTracerSpanOperationsWithOtelBridge(b *testing.B) {
	b.Setenv("JAEGER_AGENT_HOST", "localhost")
	b.Setenv("JAEGER_AGENT_PORT", "6831")
	b.Setenv("JAEGER_SAMPLER_TYPE", "const")
	b.Setenv("JAEGER_SAMPLER_PARAM", "1")
	b.Setenv("JAEGER_TAGS", "app=mimir3")
	InitOpentelemetryTracer()
	for i := 0; i < b.N; i++ {
		ctx := context.Background()
		tracer := otel.Tracer("")
		ctx, sp := tracer.Start(ctx, "test1", trace.WithAttributes(attribute.String("organization", "user1")))
		defer sp.End()
		_, sp2 := tracer.Start(ctx, "test2")
		defer sp2.End()
		sp2.AddEvent("soft error")
	}
}
