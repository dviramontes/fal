package data_structures

import "core:testing"

@(test)
apocalypse_example_reachable :: proc(t: ^testing.T) {
	grid := [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}
	testing.expect_value(t, minutes_until_apocalypse(grid), 4)
}

@(test)
apocalypse_example_unreachable :: proc(t: ^testing.T) {
	grid := [][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}
	testing.expect_value(t, minutes_until_apocalypse(grid), -1)
}

@(test)
apocalypse_empty_grid :: proc(t: ^testing.T) {
	grid := [][]int{}
	testing.expect_value(t, minutes_until_apocalypse(grid), 0)
}

@(test)
apocalypse_zero_columns :: proc(t: ^testing.T) {
	grid := [][]int{{}, {}}
	testing.expect_value(t, minutes_until_apocalypse(grid), 0)
}

@(test)
apocalypse_empty_lot :: proc(t: ^testing.T) {
	grid := [][]int{{0}}
	testing.expect_value(t, minutes_until_apocalypse(grid), 0)
}

@(test)
apocalypse_single_zombie :: proc(t: ^testing.T) {
	grid := [][]int{{2}}
	testing.expect_value(t, minutes_until_apocalypse(grid), 0)
}

@(test)
apocalypse_single_person :: proc(t: ^testing.T) {
	grid := [][]int{{1}}
	testing.expect_value(t, minutes_until_apocalypse(grid), -1)
}

@(test)
apocalypse_no_living_people :: proc(t: ^testing.T) {
	grid := [][]int{{2, 0, 2}, {0, 2, 0}}
	testing.expect_value(t, minutes_until_apocalypse(grid), 0)
}

@(test)
apocalypse_no_zombies :: proc(t: ^testing.T) {
	grid := [][]int{{1, 1}, {1, 1}}
	testing.expect_value(t, minutes_until_apocalypse(grid), -1)
}

@(test)
apocalypse_one_minute :: proc(t: ^testing.T) {
	grid := [][]int{{2, 1}}
	testing.expect_value(t, minutes_until_apocalypse(grid), 1)
}

@(test)
apocalypse_single_row :: proc(t: ^testing.T) {
	grid := [][]int{{2, 1, 1, 1, 1}}
	testing.expect_value(t, minutes_until_apocalypse(grid), 4)
}

@(test)
apocalypse_single_column_upward :: proc(t: ^testing.T) {
	grid := [][]int{{1}, {1}, {1}, {2}}
	testing.expect_value(t, minutes_until_apocalypse(grid), 3)
}

@(test)
apocalypse_spread_left :: proc(t: ^testing.T) {
	grid := [][]int{{1, 1, 1, 2}}
	testing.expect_value(t, minutes_until_apocalypse(grid), 3)
}

@(test)
apocalypse_multiple_sources :: proc(t: ^testing.T) {
	grid := [][]int{{2, 1, 1, 1, 2}}
	testing.expect_value(t, minutes_until_apocalypse(grid), 2)
}

@(test)
apocalypse_four_directions :: proc(t: ^testing.T) {
	grid := [][]int{{1, 1, 1}, {1, 2, 1}, {1, 1, 1}}
	testing.expect_value(t, minutes_until_apocalypse(grid), 2)
}

@(test)
apocalypse_diagonal_is_unreachable :: proc(t: ^testing.T) {
	grid := [][]int{{2, 0}, {0, 1}}
	testing.expect_value(t, minutes_until_apocalypse(grid), -1)
}

@(test)
apocalypse_empty_lot_blocks_spread :: proc(t: ^testing.T) {
	grid := [][]int{{2, 0, 1}}
	testing.expect_value(t, minutes_until_apocalypse(grid), -1)
}

@(test)
apocalypse_separate_infected_regions :: proc(t: ^testing.T) {
	grid := [][]int{{2, 1, 0, 1, 1, 2}}
	testing.expect_value(t, minutes_until_apocalypse(grid), 2)
}

@(test)
apocalypse_rectangular_grid :: proc(t: ^testing.T) {
	grid := [][]int{{2, 1, 1, 1}, {1, 1, 1, 1}}
	testing.expect_value(t, minutes_until_apocalypse(grid), 4)
}

@(test)
apocalypse_winding_path :: proc(t: ^testing.T) {
	grid := [][]int{{2, 1, 1}, {0, 0, 1}, {1, 1, 1}}
	testing.expect_value(t, minutes_until_apocalypse(grid), 6)
}
