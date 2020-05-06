package main

import "testing"

func TestGetHalfYearDaysMap(t *testing.T) {
	m := getHalfYearDaysMap()
	if len(m) < 180 {
		t.Fatalf("%v", m)
	}
}

func TestProcessRepositories(t *testing.T) {
	c := processRepositories("zhangwei412827@163.com")
	if _, ok := c["2020/04/10"]; !ok {
		t.Fatalf("2020-04-10 commits not found.")
	}
	if c["2020/04/10"] != 2 {
		t.Fatalf("2020-04-10 should has two commits.")
	}
}

func BenchmarkGetHalfYearDaysMap(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		getHalfYearDaysMap()
	}
}
