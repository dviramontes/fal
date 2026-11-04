defmodule Falcon.DataStructures.ElapsedMinutes do
  @moduledoc """
  Exercise: Elapsed Minutes Across Daylight Saving Time

  Given a chronological log of local `HH:MM` timestamps from the night daylight
  saving time ends, return the actual number of minutes between the first and
  last entries. The clock rolls back exactly one hour, from 02:00 to 01:00.
  A consecutive pair in the log reveals this rollback by going backward in
  local time, even though the entries are in chronological order.

  Examples:

    * `elapsed_minutes(["00:50", "01:10", "01:50", "01:10", "02:10"])` -> `140`
    * `elapsed_minutes(["01:55", "01:05", "01:25", "01:40"])` -> `45`

  Constraints:
    * The input contains at least two valid, zero-padded, 24-hour timestamps.
    * All entries belong to the same local calendar date; no midnight crossing
      occurs between entries.
    * Exactly one consecutive pair has a strictly decreasing local time,
      revealing the single one-hour rollback.
    * Equal timestamps are allowed elsewhere in the log.
    * Return an integer number of minutes. Input validation is not required.

  Hints:
    * Convert an `HH:MM` timestamp into minutes since local midnight.
    * How does the rollback affect the difference between local clock times?
  """

  @spec elapsed_minutes([String.t()]) :: non_neg_integer()
  def elapsed_minutes(timestamps) do
    minutes =
      timestamps
      |> dbg()
      |> Enum.map(fn ts ->
        hhmm = String.split(ts, ":")
        [hh, mm] = Enum.map(hhmm, &String.to_integer/1)
        hh * 60 + mm
      end)

    [first | _] = minutes
    last = List.last(minutes)

    last - first
  end
end
