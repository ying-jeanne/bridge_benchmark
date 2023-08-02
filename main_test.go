package main

import (
	"context"
	"testing"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
)

func BenchmarkTracerSpanOperations(b *testing.B) {
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
