# Shiny

Shiny is a work in progress web accelerator.

The aim of the project is to build a configurable program to accelerate delivery of content to clients.

Shiny will act as a reverse proxy that also caches requests to a data store such as memcached.

## Installation

This will install and start up memcache (the cache) and shiny (the web accelerator). It presumes you already have docker installed.

```
docker run -p 11211:11211 -d memcached memcached -I 30m
go get github.com/bilbof/shiny
cd $GOPATH/src/github.com/bilbof/shiny
go install && $GOPATH/bin/shiny
```
