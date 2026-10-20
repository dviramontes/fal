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

    zombies =
      for {row, x} <- Enum.with_index(grid),
          {cell, y} <- Enum.with_index(row),
          cell == 2 do
        {x, y}
      end

    living =
      for {row, x} <- Enum.with_index(grid),
          {cell, y} <- Enum.with_index(row),
          cell == 1,
          into: MapSet.new() do
        {x, y}
      end
      |> dbg()

    recur_minutes_until_apocalypse(grid, {zombies, 0, living})
  end

  # defp recur_minutes_until_apocalypse(_grid, {[], 0, living} = _acc) do
  #   -1
  # end

  # defp recur_minutes_until_apocalypse(_grid, {[2], _minutes, _living} = _acc) do
  #     0
  #   end
  #
  #   defp recur_minutes_until_apocalypse(_grid, {[1], _minutes, _living} = _acc) do
  #     -1
  #   end
  #
  #   defp recur_minutes_until_apocalypse(_grid, {[0], _minutes, _living} = _acc) do
  #     0
  #   end
  #
  #   defp recur_minutes_until_apocalypse(_grid, {[[], []], _minutes, _living} = _acc) do
  #     0
  #   end

  defp recur_minutes_until_apocalypse(_grid, {[], minutes, living} = _acc) do
    if MapSet.size(living) == 0 do
      minutes
    else
      -1
    end
  end

  defp recur_minutes_until_apocalypse(grid, {acc, minutes, living} = _acc) do
    next_infection =
      acc
      |> Enum.reduce([], fn {x, y}, zombies ->
        check_infected(grid, x, y)
        |> Enum.reduce(zombies, fn coor, collected ->
          [coor | collected]
        end)
        |> Enum.uniq()
      end)

    remaining = MapSet.difference(living, MapSet.new(next_infection))

    if Enum.count(next_infection) > 0 do
      recur_minutes_until_apocalypse(grid, {next_infection, minutes + 1, remaining})
    else
      recur_minutes_until_apocalypse(grid, {next_infection, minutes, remaining})
    end
  end

  defp check_infected(grid, x, y) do
    for dir <- [:top, :left, :right, :bottom],
        {:ok, x, y} <- [lookup(grid, x, y, dir)],
        neighbor = lookup(grid, x, y),
        neighbor == 1 do
      {x, y}
    end
  end

  defp lookup(grid, x, y) do
    grid |> Enum.at(x) |> Enum.at(y)
  end

  defp lookup(grid, x, y, direction) do
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

  defp direction(x, y, direction) do
    case direction do
      :top -> top(x, y)
      :bottom -> bottom(x, y)
      :left -> left(x, y)
      :right -> right(x, y)
      _ -> :invalid_direction
    end
  end

  defp top(x, y) do
    {x - 1, y}
  end

  defp bottom(x, y) do
    {x + 1, y}
  end

  defp left(x, y) do
    {x, y - 1}
  end

  defp right(x, y) do
    {x, y + 1}
  end
end
