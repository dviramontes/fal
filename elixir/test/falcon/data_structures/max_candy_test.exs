defmodule Falcon.DataStructures.MaxCandyTest do
  use ExUnit.Case, async: true

  import Falcon.DataStructures.MaxCandy

  test "collects candy from the first example" do
    assert max_candy([2, 7, 9, 3, 1]) == 12
  end

  test "can skip two consecutive houses" do
    assert max_candy([5, 1, 1, 5]) == 10
  end

  test "single house" do
    assert max_candy([4]) == 4
  end

  test "empty street" do
    assert max_candy([]) == 0
  end

  test "single house with no candy" do
    assert max_candy([0]) == 0
  end

  test "all houses have no candy" do
    assert max_candy([0, 0, 0, 0]) == 0
  end

  test "two houses cannot both be visited" do
    assert max_candy([8, 3]) == 8
    assert max_candy([3, 8]) == 8
  end

  test "equal candy counts" do
    assert max_candy([3, 3, 3, 3, 3]) == 9
  end

  test "middle house can be better than both ends" do
    assert max_candy([4, 10, 4]) == 10
  end

  test "both ends can be better than the middle" do
    assert max_candy([4, 7, 4]) == 8
  end

  test "best houses need not all have the same index parity" do
    assert max_candy([10, 1, 1, 10, 1]) == 20
  end

  test "zero candy houses separate valuable houses" do
    assert max_candy([0, 6, 0, 7, 0]) == 13
  end

  test "increasing candy counts" do
    assert max_candy([1, 2, 3, 4, 5, 6]) == 12
  end

  test "decreasing candy counts" do
    assert max_candy([6, 5, 4, 3, 2, 1]) == 12
  end

  test "long street" do
    assert max_candy(List.duplicate(1, 10_000)) == 5_000
  end
end
