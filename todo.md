# TODO

Memcached needs to be running in order for the cache to work currently.

```bash
docker run -p 11211:11211 -d memcached memcached -I 30m
```

- Should only cache HEAD and GET requests that return certain status codes
- Make it more robust to failure in general
- Refactor it a bit, as it's currently a bit messy.
- Handle failures gracefully.
- Add unit tests
- I suppose it should fail gracefully if memcached isn't reachable.
- It'd be nice for there to be a healthcheck/status endpoint for shiny.
- Various additional functionality that you would need from a reverse proxy
  and cache
- Various other configuration controls (e.g. port) from flags / a config file
- Compression of various kinds, gzip, image compression, avoiding compressing
  data twice.
- Take advantage of HTTP/2 for further acceleration. HTTP/2 can send
  multiple requests for data in parallel over a single TCP connection.
  So, Shiny could supply the JS and CSS for an HTML page when a request comes in.
- Via header, and explore other headers to add.
