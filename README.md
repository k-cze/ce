## Local

Load environment variables from `.env` and run the app:

```bash
make local
```

Or run explicitly:

```bash
export $(grep -v '^#' .env | xargs)
go run ./cmd/ce
```

## Docker

Build the Docker image:

```bash
make docker-build
```

Run the Docker container, forwarding port 8080 and loading env vars from `.env`:

```bash
make docker-run
```

To stop and remove running containers for this image:

```bash
make docker-clean
```

## Notes

* The app reads configuration from `.env` use `.env.example` as template
* Data files (like CSVs) should be placed inside the `/data` directory
