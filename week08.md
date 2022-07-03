1、使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能。

从 Docker 容器中运行的redis测试 benchmark：

root@9cb4fb69938a:/data# redis-benchmark -n 1000 -q -d 10 -t get,set
SET: 55555.56 requests per second, p50=0.439 msec
GET: 62500.00 requests per second, p50=0.407 msec

root@9cb4fb69938a:/data# redis-benchmark -n 1000 -q -d 20 -t get,set
SET: 52631.58 requests per second, p50=0.439 msec
GET: 55555.56 requests per second, p50=0.479 msec

root@9cb4fb69938a:/data# redis-benchmark -n 1000 -q -d 50 -t get,set
SET: 62500.00 requests per second, p50=0.391 msec
GET: 55555.56 requests per second, p50=0.455 msec

root@9cb4fb69938a:/data# redis-benchmark -n 1000 -q -d 100 -t get,set
SET: 55555.56 requests per second, p50=0.479 msec
GET: 55555.56 requests per second, p50=0.447 msec

root@9cb4fb69938a:/data# redis-benchmark -n 1000 -q -d 200 -t get,set
SET: 50000.00 requests per second, p50=0.503 msec
GET: 62500.00 requests per second, p50=0.383 msec

root@9cb4fb69938a:/data# redis-benchmark -n 1000 -q -d 1024 -t get,set
SET: 55555.56 requests per second, p50=0.447 msec
GET: 71428.57 requests per second, p50=0.343 msec

root@9cb4fb69938a:/data# redis-benchmark -n 1000 -q -d 5120 -t get,set
SET: 50000.00 requests per second, p50=0.543 msec
GET: 66666.67 requests per second, p50=0.335 msec

2、写入一定量的 kv 数据, 根据数据大小 1w-50w 自己评估, 结合写入前后的 info memory 信息 , 分析上述不同 value 大小下，平均每个 key 的占用内存空间。

info memory 信息含义

| 指标                      | 含义                                                         |
| ------------------------- | ------------------------------------------------------------ |
| used_memory               | 由 Redis 分配器分配的内存总量，以字节（byte）为单位，即当前redis使用内存大小。 |
| used_memory_human         | 已更直观的单位展示分配的内存总量。                           |
| used_memory_rss           | 向操作系统申请的内存大小，即redis使用的物理内存大小。        |
| used_memory_rss_human     | 已更直观的单位展示向操作系统申请的内存大小。                 |
| used_memory_peak          | redis的内存消耗峰值，以字节为单位，即历史使用记录中redis使用内存峰值。 |
| used_memory_peak_human    | 以更直观的格式返回redis的内存消耗峰值                        |
| used_memory_peak_perc     | 使用内存达到峰值内存的百分比                                 |
| used_memory_overhead      | Redis为了维护数据集的内部机制所需的内存开销，包括所有客户端输出缓冲区、查询缓冲区、AOF重写缓冲区和主从复制的backlog。 |
| used_memory_startup       | Redis服务器启动时消耗的内存                                  |
| used_memory_dataset       | 数据实际占用的内存大小，即 used_memory-used_memory_overhead  |
| used_memory_dataset_perc  | 数据占用的内存大小的百分比                                   |
| total_system_memory       | 整个系统内存                                                 |
| total_system_memory_human | 以更直观的格式显示整个系统内存                               |
| used_memory_lua           | Lua脚本存储占用的内存                                        |
| used_memory_lua_human     | 以更直观的格式显示Lua脚本存储占用的内存                      |
| maxmemory                 | Redis实例的最大内存配置                                      |
| maxmemory_human           | 以更直观的格式显示Redis实例的最大内存配置                    |
| maxmemory_policy          | 当达到maxmemory时的淘汰策略                                  |
| mem_fragmentation_ratio   | 碎片率，used_memory_rss/ used_memory。ratio指数>1表明有内存碎片，越大表明越多，<1表明正在使用虚拟内存，虚拟内存其实就是硬盘，性能比内存低得多，这是应该增强机器的内存以提高性能。一般来说，mem_fragmentation_ratio的数值在1 ~ 1.5之间是比较健康的。 |
| mem_allocator             | 内存分配器                                                   |
| active_defrag_running     | 表示没有活动的defrag(表示内存碎片整理)任务正在运行，1表示有活动的defrag任务正在运行 |
| lazyfree_pending_objects  | 0表示不存在延迟释放的挂起对象                                |



####  Before 插入之前

```
# Memory
...
used_memory:1402440
used_memory_human:1.34M
used_memory_rss:11956224
used_memory_rss_human:11.40M
used_memory_peak:414964784
used_memory_peak_human:395.74M
used_memory_peak_perc:0.34%
used_memory_overhead:851120
used_memory_startup:809824
used_memory_dataset:551320
used_memory_dataset_perc:93.03%
...
```

####  插入 100000 条数据（每个数据大小为 100 bytes) 后的 memory 信息

```
# Memory
...
used_memory:18726696
used_memory_human:17.86M
used_memory_rss:29011968
used_memory_rss_human:27.67M
used_memory_peak:414964784
used_memory_peak_human:395.74M
used_memory_peak_perc:4.51%
used_memory_overhead:5375368
used_memory_startup:809824
used_memory_dataset:13351328
used_memory_dataset_perc:74.52%
...
```

每个 key 的平均占用内存空间 (18726696 - 1402440) / 100000 - 100 = 73.2(byte)



####  插入 100000 条数据（每个数据大小为 5120 bytes) 后的 memory 信息

```
# Memory
...
used_memory:622451136
used_memory_human:593.62M
used_memory_rss:637300736
used_memory_rss_human:607.78M
used_memory_peak:622472136
used_memory_peak_human:593.64M
used_memory_peak_perc:100.00%
used_memory_overhead:5899656
used_memory_startup:809824
used_memory_dataset:616551480
used_memory_dataset_perc:99.18%
...
```

每个 key 的平均占用内存空间 (622451136 - 1402440) / 100000 - 5120 = 1090.5(byte)