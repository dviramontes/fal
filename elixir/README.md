# Fal — Elixir

## Setup

Requires Elixir 1.20+ and a compatible Erlang/OTP runtime. There are no external dependencies to install.

## Commands

Run these from `elixir/`:

```bash
just test             # Run ExUnit tests (mix test)
just test-ds          # Run data structure exercises
just test-one test/fal/data_structures/minutes_until_apocalypse_test.exs:6
just test-coverage    # Run tests with coverage
just fmt             # Format Elixir source and tests
just check           # Check formatting, compile strictly, and run tests
just iex             # Start an interactive session with the project loaded
```

Mix equivalents are `mix test`, `mix format`, and `iex -S mix`. For commands from the repository root, see [Running Tests](../README.md#running-tests).

## Exercise Layout

Place exercise modules under `lib/fal/`, grouped by topic (for example, `data_structures`, `concurrency`, `data_processing`, or `financial`). Put matching `*_test.exs` files under `test/fal/` with the same topic layout.

Put exercise descriptions in module documentation and use ExUnit assertions in tests. Follow the shared [exercise authoring guidance](../AGENTS.md#adding-exercises).
