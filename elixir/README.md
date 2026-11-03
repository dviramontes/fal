# Falcon — Elixir

## Setup

Requires Elixir 1.20+ and a compatible Erlang/OTP runtime. There are no external dependencies to install.

## Commands

Run these from `elixir/`:

```bash
just test             # Run ExUnit tests (mix test)
just test-ds          # Run data structure exercises
just test-one test/falcon/data_structures/minutes_until_apocalypse_test.exs:6
just test-coverage    # Run tests with coverage
just fmt             # Format Elixir source and tests
just check           # Check formatting, compile strictly, and run tests
just iex             # Start an interactive session with the project loaded
```

Mix equivalents are `mix test`, `mix format`, and `iex -S mix`. For commands from the repository root, see [Running Tests](../README.md#running-tests).

## Exercise Layout

Place exercise modules under `lib/falcon/`, grouped by topic (for example, `data_structures`, `concurrency`, `data_processing`, or `financial`). Put matching `*_test.exs` files under `test/falcon/` with the same topic layout.

Put exercise descriptions in module documentation and use ExUnit assertions in tests. Follow the shared [exercise authoring guidance](../AGENTS.md#adding-exercises).

## Exercises

[Minutes Until Apocalypse](lib/falcon/data_structures/minutes_until_apocalypse.ex) is a grid infection problem. Implement `Falcon.DataStructures.MinutesUntilApocalypse.minutes_until_apocalypse/1` to satisfy the 20 tests in `test/falcon/data_structures/minutes_until_apocalypse_test.exs`.

Tests cover simultaneous infection, unreachable people, empty grids, and row/column boundaries.

[Maximum Candy](lib/falcon/data_structures/max_candy.ex) asks you to maximize candy
collected from a street without visiting adjacent houses. Implement
`Falcon.DataStructures.MaxCandy.max_candy/1` to satisfy the tests in
`test/falcon/data_structures/max_candy_test.exs`.

Tests cover the supplied examples, empty streets, zero candy, adjacency, choices
that skip consecutive houses, and a long street.

[Elapsed Minutes Across Daylight Saving Time](lib/falcon/data_structures/elapsed_minutes.ex)
asks you to calculate actual elapsed time from local timestamps across a one-hour
clock rollback. Implement `Falcon.DataStructures.ElapsedMinutes.elapsed_minutes/1`
to satisfy the tests in `test/falcon/data_structures/elapsed_minutes_test.exs`.

Tests cover the supplied examples, minimal rollback intervals, equal endpoints,
duplicate entries, rollback positions, and a log spanning most of the date.
