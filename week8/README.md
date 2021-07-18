1. 使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能。
```
redis-benchmark -n 100000 -q -t get,set -d 10
SET: 5636.66 requests per second, p50=7.727 msec
GET: 5171.97 requests per second, p50=8.407 msec

redis-cli flushall
OK
redis-benchmark -n 100000 -q -t get,set -d 20
SET: 5601.61 requests per second, p50=7.727 msec
GET: 5783.36 requests per second, p50=7.519 msec

redis-cli flushall
OK
redis-benchmark -n 100000 -q -t get,set -d 50
SET: 5496.32 requests per second, p50=7.895 msec
GET: 5460.30 requests per second, p50=7.935 msec

redis-cli flushall
OK
redis-benchmark -n 100000 -q -t get,set -d 100
SET: 5445.44 requests per second, p50=8.055 msec
GET: 5757.72 requests per second, p50=7.631 msec

redis-cli flushall
OK
redis-benchmark -n 100000 -q -t get,set -d 200
SET: 4794.55 requests per second, p50=8.807 msec
GET: 5298.29 requests per second, p50=8.159 msec

redis-cli flushall
OK
redis-benchmark -n 100000 -q -t get,set -d 1024
SET: 4668.97 requests per second, p50=9.439 msec
GET: 4452.56 requests per second, p50=9.623 msec

redis-cli flushall
OK
redis-benchmark -n 100000 -q -t get,set -d 5120
SET: 3642.99 requests per second, p50=12.471 msec
GET: 4020.42 requests per second, p50=11.263 msec

redis-cli flushall
OK
```
2. 写入一定量的 kv 数据, 根据数据大小 1w-50w 自己评估, 结合写入前后的 info memory 信息, 分析上述不同 value 大小下，平均每个 key 的占用内存空间。
Memory usage for 100000 keys with 10.00 B (10) value
        used_memory: total 7.87 MB (8248576), each key 82.00 B (82)
        used_memory_rss: total 8.27 MB (8675328), each key 86.00 B (86)

go run cmd/gen_redis_data/*.go --size=20
Memory usage for 100000 keys with 20.00 B (20) value
        used_memory: total 8.63 MB (9048576), each key 90.00 B (90)
        used_memory_rss: total 6.32 MB (6623232), each key 66.00 B (66)

go run cmd/gen_redis_data/*.go -size=50
Memory usage for 100000 keys with 50.00 B (50) value
        used_memory: total 11.68 MB (12248576), each key 122.00 B (122)
        used_memory_rss: total 7.61 MB (7983104), each key 79.00 B (79)

go run cmd/gen_redis_data/*.go -size=100
Memory usage for 100000 keys with 100.00 B (100) value
        used_memory: total 17.02 MB (17848576), each key 178.00 B (178)
        used_memory_rss: total 11.77 MB (12341248), each key 123.00 B (123)

go run cmd/gen_redis_data/*.go -size=200
Memory usage for 100000 keys with 200.00 B (200) value
        used_memory: total 27.70 MB (29048576), each key 290.00 B (290)
        used_memory_rss: total 23.09 MB (24215552), each key 242.00 B (242)

go run cmd/gen_redis_data/*.go -size=1024
Memory usage for 100000 keys with 1.00 KB (1024) value
        used_memory: total 128.41 MB (134648576), each key 1.31 KB (1346)
        used_memory_rss: total 123.90 MB (129916928), each key 1.27 KB (1299)

go run cmd/gen_redis_data/*.go -size=5120
Memory usage for 100000 keys with 5.00 KB (5120) value
        used_memory: total 592.28 MB (621048576), each key 6.06 KB (6210)
        used_memory_rss: total 593.52 MB (622354432), each key 6.08 KB (6223)