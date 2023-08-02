export JAEGER_AGENT_HOST="jaeger"
export JAEGER_AGENT_PORT="6831"
export JAEGER_SAMPLER_TYPE="const"
export JAEGER_SAMPLER_PARAM="1"
export JAEGER_TAGS="app=mimir"

# you can run the following command before launching your tests:
# docker run -d --name jaeger \
#   -e JAEGER_AGENT_HOST=jaeger \
#   -e JAEGER_AGENT_PORT=6831 \
#   -e JAEGER_SAMPLER_TYPE=const \
#   -e JAEGER_SAMPLER_PARAM=1 \
#   -e JAEGER_TAGS=app=mimir \
#   -p 5775:5775/udp \
#   -p 6831:6831/udp \
#   -p 6832:6832/udp \
#   -p 5778:5778 \
#   -p 16686:16686 \
#   -p 14268:14268 \
#   -p 9411:9411 \
#   jaegertracing/all-in-one:latest