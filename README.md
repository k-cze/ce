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

## API Endpoints

| Method | Endpoint        | Description                    |
| ------ | --------------- | ------------------------------ |
| GET    | `/api/rates`    | Returns exchange rates relative to USD for specified currencies. Example: `/api/rates?currencies=USD,GBP,EUR,PLN` |
| GET    | `/api/exchange` | Performs currency conversion. Example: `/api/exchange?from=WBTC&to=USDT&amount=1.0` |

### Required Query Parameters

- **`/api/rates`**  
  - `currencies` — Comma-separated list of currency codes to fetch rates for (e.g., `USD,GBP,EUR,PLN`).

- **`/api/exchange`**  
  - `from` — Source currency code (e.g., `WBTC`).  
  - `to` — Target currency code (e.g., `USDT`).  
  - `amount` — Amount to convert (e.g., `1.0`).

## Notes

* The app reads configuration from `.env` use `.env.example` as template
* You must provide an **Open Exchange Rates API key** in the `.env` file
* Data files (like CSVs) should be placed inside the `/data` directory