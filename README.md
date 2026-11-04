# FAL

Interview practice in Go, Odin, and Elixir, covering data structures, concurrency, data processing, and financial algorithms. The goal is hands-on language practice across a broad range of problems.

## Setup

Install the tools for the language you want to practice:

| Language | Requirements                                | Setup                                                                            |
| -------- | ------------------------------------------- | -------------------------------------------------------------------------------- |
| Go       | Go 1.23+                                    | From `go/`, run `go mod download` and `go install gotest.tools/gotestsum@latest` |
| Odin     | [Odin](https://odin-lang.org/docs/install/) | No dependency setup required                                                     |
| Elixir   | See [Elixir setup](elixir/README.md#setup)  | See the language README                                                          |

Install [just](https://github.com/casey/just) to use the command recipes below.

## Project Structure

```
fal/
├── go/
│   ├── data_structures/    # Classic data structure problems
│   ├── concurrency/        # Go concurrency patterns
│   ├── data_processing/    # Stream and batch processing
│   ├── financial/          # Financial domain algorithms
│   ├── justfile            # Go commands
│   ├── go.mod
│   └── go.sum
├── odin/
│   ├── data_structures/    # Grid and data structure problems
│   └── justfile            # Odin commands
├── elixir/
│   ├── lib/                # Elixir modules
│   ├── test/               # ExUnit tests
│   ├── justfile            # Elixir commands
│   └── mix.exs             # Mix project (no external dependencies)
└── README.md
```

## Running Tests

Each language has its own `justfile`; there is no root justfile. Run `just test` from `go/`, `odin/`, or `elixir/`. Direct language commands also run from their respective directories.

From the repository root, select a justfile explicitly:

```bash
just --justfile go/justfile test
just --justfile odin/justfile test
just --justfile elixir/justfile test
```

In any language directory, `just --list` lists commands. Tests for unsolved exercises are expected to fail.

### Go Commands

Go tests use [testify](https://github.com/stretchr/testify) assertions and [gotestsum](https://github.com/gotestyourself/gotestsum) for test output and watch mode.

Run these from `go/`:

```bash
just test-dots         # Compact test output
just test-ds           # Data structures
just test-concurrency
just test-data         # Data processing
just test-financial
just test-one TestLRUCache
just test-race
just watch             # Rerun all tests on changes
just watch-ds          # Watch data structure tests
just test-coverage     # Generate coverage report
just bench
just fmt
just lint              # Requires golangci-lint
```

### Odin Commands

Run `just test-ds` from `odin/` to test data structure exercises with Odin's built-in test runner.

### Elixir Commands

See the [Elixir command reference](elixir/README.md#commands) for ExUnit tests, formatting, compilation checks, and interactive sessions.

## Practice Workflow

1. Pick an exercise from the source files in the language directory.
2. Read its problem description and hints.
3. Run its tests to see the expected behavior.
4. Implement the solution and rerun the tests as you work. For Go, watch mode automates reruns.

## License

MIT
