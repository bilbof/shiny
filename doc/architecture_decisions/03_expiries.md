# 03 Cache object expiries

Shiny will obey the Cache-Control header detailed in [RFC 7234][], provided
by origin servers. Therefore, rather than configuring the web accelerator
to cache responses, this control is delegated to origin servers.

This will make shiny more transparent with less magic, fewer bugs.

[RFC 7234]: https://tools.ietf.org/html/rfc7234
