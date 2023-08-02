# you can run the following command before launching your tests:
docker run -d --name jaeger \
  -p 5775:5775/udp \
  -p 6831:6831/udp \
  -p 6832:6832/udp \
  -p 5778:5778 \
  -p 16686:16686 \
  -p 14268:14268 \
  -p 9411:9411 \
  jaegertracing/all-in-one:latest

run `go test -bench=. -benchmem -count=3` to get the benchmark result

# result of the benchmark
there are 3 cases in the benchmark, using opentracing API, using opentracing API with opentelemetry implementation, using opentelemetry API with opentelemetry implementation

```
BenchmarkTracerSpanOperations-8                           274616              4612 ns/op            2017 B/op         25 allocs/op
BenchmarkTracerSpanOperations-8                           371631              3577 ns/op            2016 B/op         25 allocs/op
BenchmarkTracerSpanOperations-8                           364498              3430 ns/op            2016 B/op         25 allocs/op
BenchmarkTracerSpanOperationsWithOtelBridge-8             169641              6993 ns/op            4137 B/op         56 allocs/op
BenchmarkTracerSpanOperationsWithOtelBridge-8             189577              6316 ns/op            4140 B/op         56 allocs/op
BenchmarkTracerSpanOperationsWithOtelBridge-8             212000              7261 ns/op            4139 B/op         56 allocs/op
BenchmarkOtelTracerSpanOperationsWithOtelBridge-8         342580              3968 ns/op            2347 B/op         19 allocs/op
BenchmarkOtelTracerSpanOperationsWithOtelBridge-8         374302              3757 ns/op            2351 B/op         20 allocs/op
BenchmarkOtelTracerSpanOperationsWithOtelBridge-8         348621              3515 ns/op            2345 B/op         19 allocs/op
```

We can see that CPU/Memory usage of tracing library would be double at the beginning of migration with bridge (worst case), in reality it would be less than twice since we would remove the bridge from opentelemetry to opentracing that is current in Mimir at the beginning of migration [link](https://github.com/grafana/mimir/blob/main/pkg/mimir/tracing.go#L50). When migration is complet, the CPU/Memory would be slightly better on allocation per operation, but globally similar to the current consumption.

# result of using the bridge in Mimir
use GET /debug/pprof/profile to get runtime profiling data when running development/mimir-monolithic-mode, in order to profile cpu usage with bridge

get cpu profile of last 100s (I can't get it longer than this, because profile duration can't exceeds server's WriteTimeout)
go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30

result of code in main:
file:///private/var/folders/7x/yhw3mdn55nb4mt_qj8_jpp4r0000gn/T/pprof004.svg

result of code with bridge:
file:///private/var/folders/7x/yhw3mdn55nb4mt_qj8_jpp4r0000gn/T/pprof003.svg
<!-- /Users/ying-jeanne/pprof/pprof.mimir.samples.cpu.005.pb.gz -->

