defmodule Falcon.DataStructures.MinutesUntilApocalypse do
  @moduledoc """
  Exercise: Minutes Until Apocalypse

  A rectangular grid contains empty lots (0), living people (1), and zombies (2).
  Every minute, all zombies simultaneously infect living people directly above,
  below, left, or right. Newly infected people spread infection the next minute.
  Empty lots cannot carry infection; diagonal neighbors are not adjacent.

  Return the minimum minutes until no living people remain, or -1 if any are
  unreachable. Return 0 if there are no living people, including an empty grid.

  Examples:

    * `[
    *     [2, 1, 1],
    *     [1, 1, 0],
    *     [0, 1, 1]
    *  ]` -> `4`
    * `[  [2, 1, 1],
    *     [0, 1, 1],
    *     [1, 0, 1]
    *  ]` -> `-1`

  Assume all rows have the same length and cells contain only 0, 1, or 2.
  The grid is a list of lists. Elixir data is immutable.

  Hints:

    * Track all initial zombies as simultaneous sources.
    * Process infection one minute at a time.
    * Consider `:queue` for a frontier and `MapSet` for coordinates.
  """

  @spec minutes_until_apocalypse([[0 | 1 | 2]]) :: integer()
  def minutes_until_apocalypse(grid) do
    #   reduce over the current minute’s zombie coordinates.
    #
    # For each coordinate:
    #
    # - Check its four neighbors for living people.
    # - Collect new infections into the next minute’s batch.
    # - Ensure the same person isn’t counted twice if two zombies reach them.
    #
    # After that reduce finishes, advance one minute and repeat with the new batch.
    zombies =
      for {row, i} <- Enum.with_index(grid),
          {col, ii} <- Enum.with_index(row),
          col == 2 do
        {i, ii}
      end

    zombies
    |> Enum.reduce({[], 0}, fn {x, y} = zombie, {newly_infected, minute} ->
      # TODO: check new infections
      # TODO: track new infections and advance minutes
      new_zombies = cond do
        1 == Enum.at(grid, top(x, y) |> dbg()) -> zombie
      end

      {[new_zombies | newly_infected], minute + 1}
    end)
    |> dbg()
  end

  def top(x, y) do
    {x - 1, y}
  end

  def bottom(x, y) do
    {x + 1, y}
  end

  def left(x, y) do
    {x, y - 1}
  end

  def right(x, y) do
    {x, y + 1}
  end
end
