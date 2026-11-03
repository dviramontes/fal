defmodule Falcon.DataStructures.ElapsedMinutesTest do
  use ExUnit.Case, async: true

  import Falcon.DataStructures.ElapsedMinutes

  test "first supplied example" do
    assert elapsed_minutes(["00:50", "01:10", "01:50", "01:10", "02:10"]) == 140
  end

  test "second supplied example" do
    assert elapsed_minutes(["01:55", "01:05", "01:25", "01:40"]) == 45
  end

  test "two entries one minute apart across the rollback" do
    assert elapsed_minutes(["01:59", "01:00"]) == 1
  end

  test "a small backward clock difference can represent nearly an hour" do
    assert elapsed_minutes(["01:30", "01:29"]) == 59
  end

  test "first and last local times can be equal" do
    assert elapsed_minutes(["01:30", "01:50", "01:10", "01:30"]) == 60
  end

  test "last local time can precede the first" do
    assert elapsed_minutes(["01:50", "01:59", "01:00", "01:10"]) == 20
  end

  test "rollback can be the first pair" do
    assert elapsed_minutes(["01:45", "01:15", "02:00", "03:05"]) == 140
  end

  test "rollback can be the last pair" do
    assert elapsed_minutes(["00:00", "00:45", "01:55", "01:05"]) == 125
  end

  test "duplicate entries do not add elapsed time" do
    assert elapsed_minutes(["00:50", "00:50", "01:50", "01:10", "01:10", "02:10"]) ==
             140
  end

  test "log can span most of the local date" do
    assert elapsed_minutes(["00:00", "01:59", "01:00", "12:00", "23:59"]) == 1499
  end
end
