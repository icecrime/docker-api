# docker-api

The reference implementation of the Docker Remote API.

# Rationale

The goal of this repository is to provide:

- The `api` package which defines the interfaces and types used by the Docker API. This package is entirely **transport agnostic**.
- The `client` package which implements the `api` interfaces in terms of calling out to a remote JSON/HTTP endpoint. It essentially behaves as a decorator.
- The `server` package which implements the `api` interfaces in terms of serving requests received over a JSON/HTTP socket. The package does not provide request handling code: it simply exposes the interface over the network, but delegates the behavior to another instance of the `api` interfaces. It essentially behaves as a decorator.

# Usage

- Docker Engine implements the daemon behavior in terms of the `api` package interfaces, providing this implementation as a backend to the `server` package.
- Docker Swarm implements the daemon behavior in terms of the `api` package interfaces, providing this implementation as a backend to the `server` package.
- Docker Swarm uses the `client` package when proxying the call to a remote API server.
