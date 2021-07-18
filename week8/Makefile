bench_redis:
	docker run --name redis_for_benchmark --rm -p 6379:6379 -d redis:alpine
	redis-benchmark -n 100000 -q -t get,set -d 10
	redis-cli flushall
	redis-benchmark -n 100000 -q -t get,set -d 20
	redis-cli flushall
	redis-benchmark -n 100000 -q -t get,set -d 50
	redis-cli flushall
	redis-benchmark -n 100000 -q -t get,set -d 100
	redis-cli flushall
	redis-benchmark -n 100000 -q -t get,set -d 200
	redis-cli flushall
	redis-benchmark -n 100000 -q -t get,set -d 1024
	redis-cli flushall
	redis-benchmark -n 100000 -q -t get,set -d 5120
	redis-cli flushall
	docker stop redis_for_benchmark

generate_values:
	docker run --name redis_for_benchmark --rm -p 6379:6379 -d redis:alpine
	go run cmd/gen_redis_data/*.go -size=10
	go run cmd/gen_redis_data/*.go --size=20
	go run cmd/gen_redis_data/*.go -size=50
	go run cmd/gen_redis_data/*.go -size=100
	go run cmd/gen_redis_data/*.go -size=200
	go run cmd/gen_redis_data/*.go -size=1024
	go run cmd/gen_redis_data/*.go -size=5120
	docker stop redis_for_benchmark