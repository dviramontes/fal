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
    |> Enum.reduce({[], 0}, fn {x, y}, {zombies, minute} ->
      infected = check_infected(grid, x, y)

      new_infections =
        Enum.reduce(infected, [], fn {x, y}, acc ->
          [check_infected(grid, x, y) | acc]
        end)

      {new_infections, minute + Enum.count(new_infections)}
    end)
    |> then(fn {_, minutes} -> minutes end)
  end

  def check_infected(grid, x, y) do
    for dir <- [:top, :left, :right, :bottom],
        {:ok, x, y} <- [lookup(grid, x, y, dir)],
        neighbor = lookup(grid, x, y),
        neighbor == 1 do
      {x, y}
    end
  end

  def lookup(grid, x, y) do
    grid |> Enum.at(x) |> Enum.at(y)
  end

  def lookup(grid, x, y, direction) do
    x_len = length(grid)
    y_len = Enum.at(grid, 0, []) |> length

    case direction(x, y, direction) do
      {x, y} ->
        case x >= 0 and x < x_len and y >= 0 and y < y_len do
          true ->
            {:ok, x, y}

          false ->
            :out_of_bounds
        end

      err ->
        err
    end
  end

  def direction(x, y, direction) do
    case direction do
      :top -> top(x, y)
      :bottom -> bottom(x, y)
      :left -> left(x, y)
      :right -> right(x, y)
      _ -> :invalid_direction
    end
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
