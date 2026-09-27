<img src="fluxa-logo.svg" alt="Fluxa logo" width="96">

# Fluxa CLI

**Your money, one command away.** 💸

Fluxa CLI brings accounts, transactions, subscriptions, and categories into your terminal. It uses Fluxa's public `/api/v1` API, Cobra commands, and Viper settings.

## 🚀 Install

You need Go 1.26 or newer. Install the latest published version from GitHub:

```sh
go install github.com/dnutiu/fluxa-cli/cmd/fluxa@latest
fluxa --help
```

Go puts the `fluxa` binary in `GOBIN` if set, or `$(go env GOPATH)/bin`
otherwise. Add that directory to your `PATH` if `fluxa` is not found. Use
`go install` for an executable; [`go get` manages dependencies](https://go.dev/doc/go-get-install-deprecation).

To build the repository yourself instead:

```sh
git clone https://github.com/dnutiu/fluxa-cli.git
cd fluxa-cli
go build -o fluxa ./cmd/fluxa
```

## 🔑 Connect

Create a Fluxa Pro API key in **Tools → API access**. The current API does not
accept Basic accounts.

```sh
export FLUXA_API_KEY='fluxa_live_...'
fluxa auth status
fluxa entities list
fluxa config set-entity 7
```

If you built the binary locally, use `./fluxa` for the examples until it is
on your `PATH`. Replace example IDs with IDs from your own entity.

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
fluxa accounts add -n 'Everyday cash' -k cash -c RON
fluxa transactions add -g 12 -a 4 -m 18.50 -d 2026-09-27 -n 'Coffee with the team'
```

Use the returned account ID in place of `4`, and an expense category ID in place of `12`. The category determines whether a transaction is income or expense. Amounts stay exact decimal strings. If a create response is lost, retry the same field values with the same idempotency key; add `-i coffee-2026-09-27` to choose one yourself.

Now review this month's activity and check recurring payments:

```sh
fluxa transactions list -k expense -f 2026-09-01 -t 2026-09-30
fluxa transactions edit 42 -n 'Coffee and breakfast'
fluxa subscriptions add -n 'Cloud storage' -m 12.50 -s 2026-09-27 -i monthly
fluxa subscriptions pause 12
fluxa subscriptions resume 12
fluxa transactions list -s 2026-09-27T00:00:00Z -o json
```

The last command gives you the API's `data` and `meta` envelope for scripts. Every command can print a ready-to-adapt example without credentials: `fluxa transactions add --example` or `fluxa accounts edit -X`.

## 🧭 Shared flags and arguments

Global flags work before or after a command. Use `fluxa --help` or `fluxa transactions list --help` for terminal help.

| Flag or argument | What it does |
|---|---|
| `-C`, `--config PATH` | Read or save settings at `PATH`. Default: the OS user config directory's `fluxa/config.yaml`. |
| `-u`, `--base-url URL` | Talk to another Fluxa host. Default: `https://fluxa.nuculabs.dev`. HTTP is allowed only for loopback development hosts. |
| `-e`, `--entity ID` | Use a positive entity ID for this command, overriding the saved default. Required for resource commands unless an entity is saved. |
| `-o`, `--output table\|json` | Choose a readable table (default) or the API's JSON `data` and optional pagination `meta`. |
| `-X`, `--example` | Print an example for this command, including command groups, without connecting to Fluxa. |
| `-h`, `--help` | Show help for the current command. |
| `ID` | A positive integer identifying an entity, account, transaction, subscription, category, or account group as appropriate. |
| `URL` | A Fluxa host URL without a path, query, or embedded credentials. |
| `PATH` | A configuration file path. |

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
| `fluxa accounts list [-a\|--include-archived]` | See accounts and balances; include archived accounts when requested. |
| `fluxa accounts show ID` | Inspect one account, including an archived one. |
| `fluxa accounts add -n NAME [flags]` | Open an account; the name is required. |
| `fluxa accounts edit ID [flags]` | Change only the fields you supply; include at least one. |
| `fluxa accounts archive ID -y` | Archive an account and keep its transaction history. |
| `fluxa accounts restore ID` | Make an archived account active again. |

Account add and edit flags:

| Flag | Friendly explanation |
|---|---|
| `-n`, `--name NAME` | Give the account a name; required on add. |
| `-k`, `--kind KIND` | Choose `cash`, `bank`, `card`, `savings`, or `investment`; defaults to `cash`. |
| `-c`, `--currency CODE` | Choose `RON`, `EUR`, `USD`, `GBP`, or `CHF`; defaults to `RON`. |
| `-b`, `--opening-balance AMOUNT` | Set a decimal opening balance, such as `100.00`; defaults to `0`. |
| `-i`, `--annual-interest RATE` | Set the annual interest percentage as a decimal. |
| `-t`, `--interest-tax RATE` | Set the interest tax percentage as a decimal. |
| `-g`, `--account-group-id ID` | Link an account group by positive ID. |
| `-G`, `--clear-account-group` | On edit, remove the linked group. |

```sh
fluxa accounts add -n 'Operations EUR' -k bank -c EUR -b 100.00
fluxa accounts edit 4 -G
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
| `fluxa transactions add -g ID -m AMOUNT (-d DATE\|-t TIME) [flags]` | Record income or an expense; category, nonzero amount, and date or time are required. |
| `fluxa transactions edit ID [flags]` | Change only supplied fields; include at least one. |
| `fluxa transactions delete ID -y` | Soft-delete a transaction. |

Transaction add and edit flags:

| Flag | Friendly explanation |
|---|---|
| `-g`, `--category-id ID` | Pick the income or expense category; required on add. |
| `-a`, `--account-id ID` | Link an account by positive ID. |
| `-m`, `--amount AMOUNT` | Set a nonzero decimal amount, such as `18.50`; required on add. |
| `-c`, `--currency CODE` | Choose `RON`, `EUR`, `USD`, `GBP`, or `CHF`. |
| `-r`, `--exchange-rate RATE` | Set a positive decimal exchange rate. |
| `-d`, `--date YYYY-MM-DD` | Set the transaction date; use this or `--occurred-at` on add. |
| `-t`, `--occurred-at TIMESTAMP` | Set an exact RFC 3339 time. |
| `-n`, `--description TEXT` | Add a note. |
| `-x`, `--exclude-from-analytics[=true\|false]` | Include or exclude this transaction from analytics; explicit `false` works on edit. |
| `-A`, `--clear-account` | On edit, unlink the account. |
| `-N`, `--clear-description` | On edit, remove the note. |
| `-i`, `--idempotency-key VALUE` | On add, supply a retry key (1–200 characters). Otherwise the CLI generates one and prints it to stderr. |

`transactions list` filters:

| Flag | What it does |
|---|---|
| `-p`, `--page N` | Start at page `N`; default `1`. |
| `-l`, `--per-page N` | Return `N` records per page; default `25`, maximum `100`. |
| `-g`, `--category-id ID` | Show one category's transactions. |
| `-a`, `--account-id ID` | Show one account's transactions. |
| `-k`, `--kind income\|expense` | Show income or expenses. |
| `-f`, `--date-from YYYY-MM-DD` | Include transactions from this date onward. |
| `-t`, `--date-to YYYY-MM-DD` | Include transactions through this date. |
| `-s`, `--updated-since TIMESTAMP` | Include records updated after an RFC 3339 timestamp, such as `2026-09-27T00:00:00Z`. |
| `-D`, `--include-deleted` | Include soft-deleted transactions. |

Supply only changed fields for `edit`. Retries must use the same idempotency key and the exact same field values. The CLI does not automatically retry writes.

### 🔁 Subscriptions

| Command | Friendly explanation |
|---|---|
| `fluxa subscriptions list [filters]` | Browse recurring payments. |
| `fluxa subscriptions show ID` | Inspect one subscription and its billing schedule. |
| `fluxa subscriptions add -n NAME -m AMOUNT -s DATE [flags]` | Set up a recurring payment; name, positive amount, and start date are required. |
| `fluxa subscriptions edit ID [flags]` | Change supplied fields; include at least one. |
| `fluxa subscriptions pause ID` | Stop future charges for this subscription. |
| `fluxa subscriptions resume ID` | Resume the billing schedule. |
| `fluxa subscriptions delete ID -y` | Delete a subscription while retaining generated transactions. |

Subscription add and edit flags:

| Flag | Friendly explanation |
|---|---|
| `-n`, `--name NAME` | Name the recurring payment; required on add. |
| `-a`, `--account-id ID` | Link an account by positive ID. |
| `-m`, `--amount AMOUNT` | Set a positive decimal amount; required on add. |
| `-c`, `--currency CODE` | Choose `RON`, `EUR`, `USD`, `GBP`, or `CHF`; defaults to `RON`. |
| `-r`, `--exchange-rate RATE` | Set a positive decimal exchange rate. |
| `-s`, `--start-date YYYY-MM-DD` | Pick the first billing date; required on add. |
| `-i`, `--recurrence-interval monthly\|yearly` | Choose the billing interval; defaults to `monthly`. |
| `-A`, `--active[=true\|false]` | Start active by default, or set `--active=false` to pause. |
| `-v`, `--include-vat[=true\|false]` | Include VAT in the calculated amount. |
| `-f`, `--informative[=true\|false]` | Mark as informative. |
| `-G`, `--clear-account` | On edit, unlink the account. |

`subscriptions list` filters:

| Flag | What it does |
|---|---|
| `-p`, `--page N` | Start at page `N`; default `1`. |
| `-l`, `--per-page N` | Return `N` records per page; default `25`, maximum `100`. |
| `-a`, `--active[=true\|false]` | Show active or paused subscriptions. Omit it to show both. |
| `-s`, `--updated-since TIMESTAMP` | Include records updated after an RFC 3339 timestamp. |

```sh
fluxa subscriptions add -n 'Cloud storage' -a 4 -m 12.50 -c EUR -s 2026-09-27 -i monthly
fluxa subscriptions edit 12 --active=false --include-vat=true
```

### ✅ Confirmation and output

`accounts archive`, `transactions delete`, and `subscriptions delete` require `-y` or `--yes` so an accidental invocation cannot remove anything. The Rails application's `docs/api.md` is the field-level API reference.

Tables show selected fields and pagination. `--output json` prints `data` and, for paginated collections, `meta`. API rate limits still apply; the current key limit is 10 requests per minute.

### 🐚 Help and completion

| Command | Friendly explanation |
|---|---|
| `fluxa help [command]` | Read help for any command, such as `fluxa help transactions list`. |
| `fluxa completion bash` | Generate Bash completion. |
| `fluxa completion fish` | Generate Fish completion. |
| `fluxa completion powershell` | Generate PowerShell completion. |
| `fluxa completion zsh` | Generate Zsh completion. |

All commands, including `help` and completion commands, accept `-X` or `--example`. Add `--no-descriptions` to a completion shell command for a smaller script. For one session, run `source <(fluxa completion zsh)` or `source <(fluxa completion bash)`; each completion command's `--help` explains persistent installation.

## 🧪 Build and test

```sh
go test ./...
go vet ./...
go build -o fluxa ./cmd/fluxa
```

Tests cover command flags and examples, typed input validation, partial updates, config secrecy, request headers and paths, API errors, and account balance mapping. They use local HTTP fixtures, so no Fluxa account is needed.

## 🏗️ Code layout and future API work

`cmd` wires Cobra commands and parses terminal flags; `internal/application` holds one use case and narrow port per operation; `internal/domain` owns models and validation; `internal/infrastructure` implements HTTP and Viper settings; `internal/presentation` formats results. JSON request encoding stays in the HTTP adapter.

This CLI covers the current bearer-authenticated `/api/v1` surface. More Fluxa workflows need new Rails endpoints; their proposed contracts are in `docs/future_plans/fluxa_cli.md` in the Rails repository.
