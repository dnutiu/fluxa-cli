# Fluxa CLI agent guidance

- Build the CLI in Go with Cobra commands and Viper configuration.
- Use only Fluxa's bearer-authenticated `/api/v1` endpoints. Other JSON routes require a browser session.
- Keep money as decimal strings, use explicit entity IDs, and preserve API `data`, `meta`, and error envelopes.
- Read API keys from `FLUXA_API_KEY` or a secure credential source. Never print or store them in the configuration file.
- Add tests for HTTP behavior and command parsing when implementing a workflow.
- Keep this file and README accurate as commands are added.

## Architecture

- `internal/domain`: Fluxa resource models, IDs, filters, and validation. Keep this package independent of Cobra, Viper, and HTTP.
- `internal/application`: one use case per file. Each operation declares the smallest repository interface it needs and validates input before calling it.
- `internal/infrastructure/fluxa`: the bearer-authenticated HTTP adapter. Keep endpoint paths, query encoding, headers, and API error parsing here.
- `internal/infrastructure/config`: Viper settings and persistence. API keys must stay outside saved configuration.
- `internal/presentation`: JSON and table rendering.
- `cmd/fluxa`: the executable entry point so `go install github.com/dnutiu/fluxa-cli/cmd/fluxa@latest` produces `fluxa`.
- `cmd`: Cobra command wiring and terminal input only. Do not put API paths or request code in commands.

For a new API operation, add its domain model or filter if needed, an application use case with its own port, an HTTP adapter method, and a small Cobra command. Keep money as strings throughout.
