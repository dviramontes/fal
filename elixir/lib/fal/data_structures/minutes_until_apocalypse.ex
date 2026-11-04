defmodule Fal.DataStructures.MinutesUntilApocalypse do
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
  """

  @spec minutes_until_apocalypse([[0 | 1 | 2]]) :: integer()
  def minutes_until_apocalypse(grid) do
    zombies =
      for {row, x} <- Enum.with_index(grid),
          {cell, y} <- Enum.with_index(row),
          cell == 2,
          into: MapSet.new() do
        {x, y}
      end

    living =
      for {row, x} <- Enum.with_index(grid),
          {cell, y} <- Enum.with_index(row),
          cell == 1,
          into: MapSet.new() do
        {x, y}
      end

    recur_minutes_until_apocalypse(zombies, 0, living)
  end

  defp recur_minutes_until_apocalypse(zombies, minutes, living) do
    cond do
      MapSet.size(living) == 0 ->
        minutes

      MapSet.size(zombies) == 0 ->
        -1

      true ->
        next_infection =
          zombies
          |> Enum.reduce(MapSet.new(), fn {x, y}, zombies ->
            [{x - 1, y}, {x + 1, y}, {x, y - 1}, {x, y + 1}]
            |> Enum.filter(fn coors ->
              MapSet.member?(living, coors)
            end)
            |> Enum.reduce(zombies, fn coors, collected ->
              MapSet.put(collected, coors)
            end)
          end)

        remaining = MapSet.difference(living, next_infection)

        if MapSet.size(next_infection) > 0 do
          recur_minutes_until_apocalypse(next_infection, minutes + 1, remaining)
        else
          recur_minutes_until_apocalypse(next_infection, minutes, remaining)
        end
    end
  end
end
