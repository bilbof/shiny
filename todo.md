# TODO

Memcached needs to be running in order for the cache to work currently.

```bash
docker run -p 11211:11211 -d memcached memcached -I 30m
```

- Should only cache HEAD and GET requests that return certain status codes
- Make it more robust to failure in general
- Refactor it a bit, as it's currently a bit messy.
- Add unit tests
- I suppose it should fail gracefully if memcached isn't reachable.
- It'd be nice for there to be a healthcheck/status endpoint for shiny.
- Various additional functionality that you would need from a reverse proxy
  and cache
- Various other configuration controls (e.g. port) from flags / a config file
