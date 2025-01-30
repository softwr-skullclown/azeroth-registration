# Developing Azeroth Registration

First it is highy recommended to install the following:

1. [Node 22](https://nodejs.org/en/download) for UI build
2. [Go 1.23+](https://go.dev/dl/) for Go build
3. [Task](https://taskfile.dev/installation) for (OPTIONAL) build/run workflows


## Developing the UI with Svelte 5

The UI is built on [Svelte 5](https://svelte.dev) and embedded into the Go http server to serve the assets directly.

## Developing the Backend with Go

The Go backend connects to the configured MySQL database for the private server authentication schema.