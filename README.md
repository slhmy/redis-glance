# RedisGlance

A repository to help you quickly get to know with Redis

## Environment

To try this out, it will be nice if you already have some Redis server Running.
Otherwise, you can use **docker** for environment setup.

### Setup Single Redis Server

``` bash
docker run -e "IP=0.0.0.0" -p 6379:6379 redis
```

### Setup Redis Cluster

Currently, we use [grokzen/redis-cluster](https://github.com/Grokzen/docker-redis-cluster)
to set up the testing Redis Cluster.

``` bash
docker run -e "IP=0.0.0.0" -p 7000-7005:7000-7005 grokzen/redis-cluster:latest
```

## Usage

There is two simple test in `test` directory.

You can control the `goroutine` and repeat `key` number in the tests,
to try various cases,
finding out in which case the cluster/single Redis server is in better performance.

### Conclusion

In brief, when there is a higher request rate, the cluster is in better performance,
the bound of this rate depends on the running environment.

## Contributing

Any ideas/issues or suggestions are welcome.

Tell us anything you want to find out with Redis.