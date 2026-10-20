defmodule Falcon.DataStructures.MinutesUntilApocalypseTest do
  use ExUnit.Case, async: true

  import Falcon.DataStructures.MinutesUntilApocalypse

  test "reachable example" do
    assert minutes_until_apocalypse([[2, 1, 1], [1, 1, 0], [0, 1, 1]]) == 4
  end

  test "unreachable example" do
    assert minutes_until_apocalypse([[2, 1, 1], [0, 1, 1], [1, 0, 1]]) == -1
  end

  test "empty grid" do
    assert minutes_until_apocalypse([]) == 0
  end

  test "zero columns" do
    assert minutes_until_apocalypse([[], []]) == 0
  end

  test "single empty lot" do
    assert minutes_until_apocalypse([[0]]) == 0
  end

  test "single zombie" do
    assert minutes_until_apocalypse([[2]]) == 0
  end

  test "single person" do
    assert minutes_until_apocalypse([[1]]) == -1
  end

  test "no living people" do
    assert minutes_until_apocalypse([[2, 0, 2], [0, 2, 0]]) == 0
  end

  test "no zombies" do
    assert minutes_until_apocalypse([[1, 1], [1, 1]]) == -1
  end

  test "infection takes one minute" do
    assert minutes_until_apocalypse([[2, 1]]) == 1
  end

  test "multiple people are infected simultaneously in one minute" do
    assert minutes_until_apocalypse([[0, 1, 0], [1, 2, 1], [0, 1, 0]]) == 1
  end

  test "single row spreads right" do
    assert minutes_until_apocalypse([[2, 1, 1, 1, 1]]) == 4
  end

  test "single column spreads upward" do
    assert minutes_until_apocalypse([[1], [1], [1], [2]]) == 3
  end

  test "single row spreads left" do
    assert minutes_until_apocalypse([[1, 1, 1, 2]]) == 3
  end

  test "multiple sources spread simultaneously" do
    assert minutes_until_apocalypse([[2, 1, 1, 1, 2]]) == 2
  end

  test "infection spreads in all four directions" do
    assert minutes_until_apocalypse([[1, 1, 1], [1, 2, 1], [1, 1, 1]]) == 2
  end

  test "diagonal person is unreachable" do
    assert minutes_until_apocalypse([[2, 0], [0, 1]]) == -1
  end

  test "empty lot blocks spread" do
    assert minutes_until_apocalypse([[2, 0, 1]]) == -1
  end

  test "separate regions are infected simultaneously" do
    assert minutes_until_apocalypse([[2, 1, 0, 1, 1, 2]]) == 2
  end

  test "rectangular grid" do
    assert minutes_until_apocalypse([[2, 1, 1, 1], [1, 1, 1, 1]]) == 4
  end

  test "infection follows a winding path" do
    assert minutes_until_apocalypse([[2, 1, 1], [0, 0, 1], [1, 1, 1]]) == 6
  end
end
