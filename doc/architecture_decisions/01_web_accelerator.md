# Web accelerator

Shiny will be a web accelerator.

The aim is for the software to become a flexible reverse proxy focused
on being a web server accelerator.

Features I'd like Shiny to have:

- Request proxying
- Configuration controls
- Caching (using e.g. memcached or cassandra as distributed cache store)
- Prefetching (and caching)
- Document compression and optimization
- Persist connections with clients
- TCP acceleration
- Simplicity (don't implement anything other than web accelerator features)

For now I'll focus on getting the first two done, then caching. Other features
are nice but not really necessary.
