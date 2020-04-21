# TODO

Memcached needs to be running in order for the cache to work currently.

```bash
docker run -p 11211:11211 -d memcached memcached -I 30m
```

Up next: Finishing caching

Web acceleration features

- TTL on cached items
- Obey various headers, such as Cache-Control
- Compression of responses (in progress)
- Various additional functionality that you would need from a reverse proxy
  and cache
- Pre-heating/auto-warming of cached objects
- Pre-fetching of linked resources e.g. static assets like css, js, images
- Autocrawling
- ETags support
- Hostname resolution
- Code optimization
- Queuing duplicated requests that are in progress
- Experiment with some go features for lower latency requests

Other features

- Default in-memory heap cache?
- Healthcheck/status endpoint
- Stats
- Cache stats
- Various other configuration controls (e.g. port) from flags / a config file
- Reconfiguration without restarts
- IP blacklisting
- Rate limiting based on configurable policies
- Private object caching
- AB testing
- Request tracing

General tasks

- Test coverage
- Refactoring
- Benchmark tests
- Handle failures gracefully.

Bugs

- Only GET/HEAD requests should be cached.
- Handle cache failures gracefully
- Various other configuration controls (e.g. port) from flags / a config file
- Compression of various kinds, gzip, image compression, avoiding compressing
  data twice.
- Take advantage of HTTP/2 for further acceleration. HTTP/2 can send
  multiple requests for data in parallel over a single TCP connection.
  So, Shiny could supply the JS and CSS for an HTML page when a request comes in.
- Via header, and explore other headers to add.

Principles/design intents

- Stateless, should leave distributed work to the chosen store (cassandra, redis, memcached)
- convention, a simpler system with a simple config
- for greater control, I could permit Go plugins for users to hook into the state machine
- 
