package main

import (
	"fmt"
	"strconv"
	"strings"
)

var cache = make(map[string]int)

func count(cfg string, nums []int) int {
	key := cfg
	for _, n := range nums {
		key += strconv.Itoa(n) + ","
	}
	if v, ok := cache[key]; ok {
		return v
	}

	if cfg == "" {
		if len(nums) == 0 {
			return 1
		}
		return 0
	}

	if strings.HasPrefix(cfg, "?") {
		return count(strings.Replace(cfg, "?", ".", 1), nums) +
			count(strings.Replace(cfg, "?", "#", 1), nums)
	}

	if strings.HasPrefix(cfg, ".") {
		res := count(strings.TrimPrefix(cfg, "."), nums)
		cache[key] = res
		return res
	}

	if strings.HasPrefix(cfg, "#") {
		if len(nums) == 0 {
			cache[key] = 0
			return 0
		}
		if len(cfg) < nums[0] {
			cache[key] = 0
			return 0
		}
		if strings.Contains(cfg[:nums[0]], ".") {
			cache[key] = 0
			return 0
		}
		if len(nums) > 1 {
			if len(cfg) < nums[0]+1 || cfg[nums[0]] == '#' {
				cache[key] = 0
				return 0
			}
			res := count(cfg[nums[0]+1:], nums[1:])
			cache[key] = res
			return res
		} else {
			res := count(cfg[nums[0]:], nums[1:])
			cache[key] = res
			return res
		}
	}

	return 0
}

func SolveA(input string) string {
	lines := Lines(input)

	total := 0
	for _, line := range lines {
		split := strings.Split(line, " ")
		arrangement := split[0]
		nums := MustParseSlice(split[1])
		total += count(arrangement, nums)
	}

	return fmt.Sprint(total)
}

func MustParseSlice(s string) []int {
	fields := strings.Split(s, ",")
	ints := make([]int, 0)
	for i := range fields {
		ints = append(ints, MustParseInteger[int](fields[i]))
	}
	return ints
}
