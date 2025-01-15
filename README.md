# Hello Action

A simple GitHub Action written in Go that prints a message.

## Inputs

### `message`

**Optional** The message to print. Default `"Hello, World!"`.

## Example usage

```yaml
uses: your-username/hello-action@v1
with:
  message: 'Hello from my workflow!'
```

## Development

This action is written in Go and uses Docker to run. To build locally:

1. Build the Go binary:
   ```bash
   go build -o main .
   ```

2. Build the Docker image:
   ```bash
   docker build -t hello-action .
   ```
