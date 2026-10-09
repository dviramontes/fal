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

    * `[[2, 1, 1], [1, 1, 0], [0, 1, 1]]` -> `4`
    * `[[2, 1, 1], [0, 1, 1], [1, 0, 1]]` -> `-1`

  Assume all rows have the same length and cells contain only 0, 1, or 2.
  The grid is a list of lists. Elixir data is immutable.

  Hints:

    * Track all initial zombies as simultaneous sources.
    * Process infection one minute at a time.
    * Consider `:queue` for a frontier and `MapSet` for coordinates.
  """

  @spec minutes_until_apocalypse([[0 | 1 | 2]]) :: integer()
  def minutes_until_apocalypse(_grid) do
    # TODO: Implement this.
    0
  end
end
