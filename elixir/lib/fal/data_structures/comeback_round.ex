defmodule Fal.DataStructures.ComebackRound do
  @moduledoc """
  Exercise: Comeback Round

  Given rounds in a board game, where each round is a list of points earned by
  each player, find the earliest round from which the eventual winner is the
  sole leader at the end of every remaining round. Scores accumulate across
  rounds. Return the round number (1-based), or -1 if the final score is tied.

  Examples:

    * `comeback_round([[5, 3, 2], [0, 5, 2], [4, 2, 2], [1, 3, 0]])` -> `2`
    * `comeback_round([[4, 2], [1, 3], [2, 1]])` -> `3`
    * `comeback_round([[4, 2], [1, 3]])` -> `-1`

  Constraints:
    * Input is a nonempty list of rounds, each containing at least one player.
    * Every round contains the same number of players, in the same order.
    * Points are non-negative integers; zero-point rounds are allowed.
    * A sole leader has a cumulative score strictly greater than every other
      player's score. A tie for the lead interrupts a winning streak.
    * A tie below the highest final score does not count as a tied final result.
    * Inputs satisfy these constraints; no input validation is required.

  Hints:
    * Distinguish points earned in a round from cumulative scores after it.
    * What must be true about the eventual winner in every round starting
    * your proposed answer?
  """

  @spec comeback_round([[non_neg_integer()]]) :: pos_integer() | -1
  def comeback_round(rounds) do
    [first | rest] = rounds

    {_totals, leaders} =
      rest
      |> Enum.reduce({first, [find_leader(first)]}, fn round, {totals, leaders} = _acc ->
        new_totals = Enum.zip_with(round, totals, &+/2)
        new_winner = find_leader(new_totals)
        {new_totals, [new_winner | leaders]}
      end)

    final_leader = hd(leaders)

    if final_leader == nil do
      -1
    else
      streak = leaders |> Enum.take_while(&(&1 == final_leader)) |> length()
      length(leaders) - streak + 1
    end
  end

  defp find_leader(round) do
    max = Enum.max(round)
    tie? = Enum.count(round, &(&1 == max)) > 1
    new_leader = Enum.find_index(round, &(&1 == max))
    if tie?, do: nil, else: new_leader
  end
end
