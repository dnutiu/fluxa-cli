<img src="fluxa-logo.svg" alt="Fluxa logo" width="96">

# Fluxa CLI

**Your money, one command away.** 💸

Fluxa CLI brings accounts, transactions, subscriptions, and categories into your terminal. It uses Fluxa's public `/api/v1` API, Cobra commands, and Viper settings.

## 🚀 Get started

You need Go 1.26 or newer to build the CLI and a Fluxa Pro API key from **Tools → API access**. The current API does not accept Basic accounts.

```sh
go build -o fluxa .
export FLUXA_API_KEY='fluxa_live_...'
./fluxa auth status
./fluxa entities list
./fluxa config set-entity 7
```

Examples below use `fluxa`; use `./fluxa` if the binary is not on your `PATH`. Replace example IDs with IDs from your own entity.

## ✨ A day with Fluxa

Start by checking access and choosing the entity you want to work with:

```sh
fluxa auth status
fluxa entities list
fluxa config set-entity 7
```

Find an account ID and an expense category ID, then record an expense:

```sh
fluxa accounts list
fluxa categories list
# If the account list is empty, create one and use its returned ID:
printf '%s\n' '{"name":"Everyday cash","kind":"cash","currency":"RON"}' | fluxa accounts add --file -
cat > coffee.json <<'JSON'
{
  "category_id": 12,
  "account_id": 4,
  "amount": "18.50",
  "currency": "RON",
  "date": "2026-09-27",
  "description": "Coffee with the team"
}
JSON
fluxa transactions add --file coffee.json --idempotency-key coffee-2026-09-27
```

The category determines whether a transaction is income or expense. Money stays a decimal string, so `"18.50"` keeps its exact value. Reuse the same idempotency key and identical JSON if a response is lost and you need to retry.

Now review this month's activity and check recurring payments:

```sh
fluxa transactions list --kind expense --date-from 2026-09-01 --date-to 2026-09-30
fluxa subscriptions list --active=true
fluxa transactions list --updated-since 2026-09-27T00:00:00Z --output json
```

The last command gives you the API's `data` and `meta` envelope for scripts. You can also pause a subscription with `fluxa subscriptions pause 12` and resume it with `fluxa subscriptions resume 12`.

## 🧭 Shared flags and arguments

Global flags work before or after a command. Use `fluxa --help` or `fluxa transactions list --help` for terminal help.

| Flag or argument | What it does |
|---|---|
| `--config PATH` | Read or save settings at `PATH`. Default: the OS user config directory's `fluxa/config.yaml`. |
| `--base-url URL` | Talk to another Fluxa host. Default: `https://fluxa.nuculabs.dev`. HTTP is allowed only for loopback development hosts. |
| `--entity ID` | Use a positive entity ID for this command, overriding the saved default. Required for accounts, categories, transactions, and subscriptions unless an entity is saved. |
| `--output table\|json` | Choose a readable table (default) or the JSON `data` envelope and pagination `meta` when present. |
| `--file PATH` | Read a create or edit request from a JSON file. Use `--file -` for stdin. |
| `--yes` | Confirm account archive, transaction delete, or subscription delete. |
| `--idempotency-key VALUE` | Choose the retry key for `transactions add`; otherwise the CLI generates one. |
| `-h`, `--help` | Show help for the current command. |
| `ID` | A positive integer identifying an account, transaction, or subscription. |
| `URL` | A Fluxa host URL without a path, query, or embedded credentials. |
| `PATH` | A file path. For `--file`, use `-` to read JSON from stdin. |

Viper also reads `FLUXA_BASE_URL`, `FLUXA_ENTITY`, and `FLUXA_OUTPUT`. An explicit flag takes precedence over the environment and saved settings. Keep your key in `FLUXA_API_KEY`; the CLI does not put it in the config file. Saved config files use mode `0600`.

## 🧰 Command reference

### 🔑 Access and settings

| Command | Friendly explanation |
|---|---|
| `fluxa auth status` | Check your API key and show how many entities it can access. With `--output json`, list those entities. |
| `fluxa config show` | Show the active base URL, entity, output format, and config file path. It never displays the API key. |
| `fluxa config set-entity ID` | Remember an entity as the default for future commands. |
| `fluxa config set-url URL` | Remember a Fluxa server URL for future commands. |
| `fluxa entities list` | Discover the entities your key can access. Choose one with `config set-entity` or `--entity`. |

### 🏦 Accounts

| Command | Friendly explanation |
|---|---|
| `fluxa accounts list [--include-archived]` | See active accounts and balances. Add `--include-archived` to include archived accounts. |
| `fluxa accounts show ID` | Inspect one account, including an archived one. |
| `fluxa accounts add --file PATH` | Create an account from a JSON object. |
| `fluxa accounts edit ID --file PATH` | Update only the fields supplied in a JSON object. |
| `fluxa accounts archive ID --yes` | Archive an account and keep its transaction history. `--yes` confirms the action. |
| `fluxa accounts restore ID` | Make an archived account active again. |

For `accounts add`, `name` is required. `kind` defaults to `cash`, `currency` to `RON`, and `opening_balance` to `0`. For example:

```json
{"name":"Operations EUR","kind":"bank","currency":"EUR","opening_balance":"100.00"}
```

### 🏷️ Categories

| Command | Friendly explanation |
|---|---|
| `fluxa categories list` | Find enabled income and expense categories and their IDs for transaction creation. |

### 💸 Transactions

| Command | Friendly explanation |
|---|---|
| `fluxa transactions list [filters]` | Browse transactions, one page at a time. Filters are listed below. |
| `fluxa transactions show ID` | Inspect one transaction. |
| `fluxa transactions add --file PATH [--idempotency-key VALUE]` | Create a transaction. If no key is supplied, the CLI generates one and prints it to stderr before sending. |
| `fluxa transactions edit ID --file PATH` | Update supplied fields on a transaction. |
| `fluxa transactions delete ID --yes` | Soft-delete a transaction. `--yes` confirms the action. |

`transactions list` filters:

| Flag | What it does |
|---|---|
| `--page N` | Start at page `N`; default `1`. |
| `--per-page N` | Return `N` records per page; default `25`, maximum `100`. |
| `--category-id ID` | Show one category's transactions. |
| `--account-id ID` | Show one account's transactions. |
| `--kind income\|expense` | Show income or expenses. |
| `--date-from YYYY-MM-DD` | Include transactions from this date onward. |
| `--date-to YYYY-MM-DD` | Include transactions through this date. |
| `--updated-since TIMESTAMP` | Include records updated after an RFC 3339 timestamp, such as `2026-09-27T00:00:00Z`. |
| `--include-deleted` | Include soft-deleted transactions. |

For `transactions add`, the JSON object needs `category_id`, a nonzero decimal-string `amount`, and either `date` or `occurred_at`. Other supported fields include `account_id`, `currency`, `exchange_rate`, `description`, and `exclude_from_analytics`. Supply only changed fields for `edit`. Use `--idempotency-key VALUE` (1–200 characters) when you want to choose the retry key yourself; retries must use the same key and exact request body. The CLI does not automatically retry writes.

### 🔁 Subscriptions

| Command | Friendly explanation |
|---|---|
| `fluxa subscriptions list [filters]` | Browse recurring payments. |
| `fluxa subscriptions show ID` | Inspect one subscription and its billing schedule. |
| `fluxa subscriptions add --file PATH` | Create a recurring payment from a JSON object. |
| `fluxa subscriptions edit ID --file PATH` | Update the supplied fields. |
| `fluxa subscriptions pause ID` | Stop future charges for this subscription. |
| `fluxa subscriptions resume ID` | Resume the billing schedule. |
| `fluxa subscriptions delete ID --yes` | Delete a subscription while retaining generated transactions. `--yes` confirms the action. |

`subscriptions list` filters:

| Flag | What it does |
|---|---|
| `--page N` | Start at page `N`; default `1`. |
| `--per-page N` | Return `N` records per page; default `25`, maximum `100`. |
| `--active` or `--active=true` | Show active subscriptions only. |
| `--active=false` | Show paused subscriptions only. Omit the flag to show both. |
| `--updated-since TIMESTAMP` | Include records updated after an RFC 3339 timestamp. |

For `subscriptions add`, provide `name`, a decimal-string `amount`, and `start_date`. `recurrence_interval` defaults to `monthly`, `currency` to `RON`, and `active` to `true`. Optional fields include `account_id`, `exchange_rate`, `include_vat`, and `informative`:

```json
{"name":"Cloud storage","account_id":4,"amount":"12.50","currency":"EUR","start_date":"2026-09-27","recurrence_interval":"monthly"}
```

### 📄 JSON files and output

`accounts add/edit`, `transactions add/edit`, and `subscriptions add/edit` all accept `--file PATH` or `--file -` for stdin. The file must contain one JSON object and be no larger than 1 MiB. The Rails application's `docs/api.md` is the field-level API reference.

Tables show selected fields and pagination. `--output json` prints `data` and, for paginated collections, `meta`. API rate limits still apply; the current key limit is 10 requests per minute.

### 🐚 Help and completion

| Command | Friendly explanation |
|---|---|
| `fluxa help [command]` | Read help for any command, such as `fluxa help transactions list`. |
| `fluxa completion bash` | Generate Bash completion. |
| `fluxa completion fish` | Generate Fish completion. |
| `fluxa completion powershell` | Generate PowerShell completion. |
| `fluxa completion zsh` | Generate Zsh completion. |

Add `--no-descriptions` to any `completion` shell command for a smaller script without descriptive text. For a one-session Zsh setup, run `source <(fluxa completion zsh)`; for Bash, use `source <(fluxa completion bash)`. Each completion command's `--help` explains persistent installation.

## 🧪 Build and test

```sh
go test ./...
go vet ./...
go build -o fluxa .
```

Tests cover command wiring and flags, config secrecy, operation validation, request headers and paths, API errors, and account balance mapping. They use local HTTP fixtures, so no Fluxa account is needed.

## 🏗️ Code layout and future API work

`cmd` wires Cobra commands; `internal/application` holds one use case and narrow port per operation; `internal/domain` owns models and filter rules; `internal/infrastructure` implements HTTP and Viper settings; `internal/presentation` formats results.

This CLI covers the current bearer-authenticated `/api/v1` surface. More Fluxa workflows need new Rails endpoints; their proposed contracts are in `docs/future_plans/fluxa_cli.md` in the Rails repository.
