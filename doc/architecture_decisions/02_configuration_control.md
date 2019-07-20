# 02 Configuration control

It is important for the software to be highly configurable.

At the moment it's just a proxy server. One option would be to set various
controls with flags. For example, `shiny --port 80 --proxy-port 3000`.

While OK for these cases, this isn't a very expressive way to
configure the server.

Varnish has a configuration language, VCL, which lets you hook into a request
at various states. You can also override defaults like the hash function by
which requests are fingerprinted.

I can't see a reason for wanting to hook into caching requests at the moment.
There may be a need for this in the future, but I'd like to keep things simple
with flags for now.

## Middleware

I think the best solution would be to start out with flags, which could
later be replaced with a config file.

For things where a user wants to introduce their custom handlers, I suppose
supporting middleware like Rack would be the best solution. If the core
software is built to support that case, that'd make it easier to extend.

I will try to write the caching component as separate middleware.
