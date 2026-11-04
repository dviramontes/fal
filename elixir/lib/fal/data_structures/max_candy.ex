defmodule Fal.DataStructures.MaxCandy do
  @moduledoc """
  Exercise: Maximum Candy

  You're trick-or-treating down a street. Each entry in a list gives the number
  of candy pieces available at that house. You cannot visit adjacent houses.
  Return the maximum total number of candy pieces you can collect.

  Examples:

    * `max_candy([2, 7, 9, 3, 1])` -> `12` (visit houses with 2, 9, and 1)
    * `max_candy([5, 1, 1, 5])` -> `10` (skip the middle two houses)
    * `max_candy([4])` -> `4`

  Constraints:
    * Input is a list of non-negative integers in street order.
    * The street is a straight line; the first and last houses are not adjacent
      unless they are the only two houses.
    * You may skip any house. Return 0 for an empty list.
    * Return only the maximum total, not the houses visited.

  Hints:
    * At each house, consider the consequences of visiting it or skipping it.
    * What information from earlier houses would help you make that choice?
  """

  @spec max_candy([non_neg_integer()]) :: non_neg_integer()
  def max_candy(houses) do
    {prev_house, _two_back_best} =
      Enum.reduce(houses, {0, 0}, fn current_candy, {prev_best, two_back_best} ->
        take_total = current_candy + two_back_best

        if prev_best > take_total do
          {prev_best, prev_best}
        else
          {take_total, prev_best}
        end
      end)

    prev_house
  end
end
