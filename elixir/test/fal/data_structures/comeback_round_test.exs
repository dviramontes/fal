defmodule Fal.DataStructures.ComebackRoundTest do
  use ExUnit.Case, async: true

  import Fal.DataStructures.ComebackRound

  test "winner takes a lasting lead in the second round" do
    assert comeback_round([[5, 3, 2], [0, 5, 2], [4, 2, 2], [1, 3, 0]]) == 2
  end

  test "a tied intermediate score interrupts an earlier lead" do
    assert comeback_round([[4, 2], [1, 3], [2, 1]]) == 3
  end

  test "a tied final score has no winner" do
    assert comeback_round([[4, 2], [1, 3]]) == -1
  end

  test "winner leads from the first round" do
    assert comeback_round([[3, 1], [2, 1], [0, 0]]) == 1
  end

  test "winner loses and then regains the lead" do
    assert comeback_round([[5, 1], [0, 6], [4, 0], [1, 0]]) == 3
  end

  test "winner first leads in the final round" do
    assert comeback_round([[0, 5], [2, 0], [4, 0]]) == 3
  end

  test "single round with a unique winner" do
    assert comeback_round([[1, 7, 3]]) == 1
  end

  test "single round with a tie for the highest score" do
    assert comeback_round([[7, 7, 3]]) == -1
  end

  test "all players finish with zero points" do
    assert comeback_round([[0, 0, 0], [0, 0, 0]]) == -1
  end

  test "one player is always the sole leader" do
    assert comeback_round([[0], [0], [5]]) == 1
  end

  test "a tie among losing players does not prevent a winner" do
    assert comeback_round([[2, 1, 1], [2, 1, 1]]) == 1
  end

  test "ties with different opponents interrupt the winner's lead" do
    assert comeback_round([[5, 5, 0], [1, 0, 6], [1, 0, 0]]) == 3
  end

  test "round points do not replace cumulative scores" do
    assert comeback_round([[10, 0], [0, 3], [0, 3]]) == 1
  end

  test "zero-point rounds preserve a tie until it is broken" do
    assert comeback_round([[1, 1], [0, 0], [0, 1], [0, 0]]) == 3
  end
end
