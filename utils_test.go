package main

import "testing"

func TestGetUUID(t *testing.T) {
	for _, test := range []struct {
		in  string
		out string
	}{
		{"/aaa/", ""},
		{"/5466df56-14f3-4df6-83d6-3deff2e931eb/", "5466df56-14f3-4df6-83d6-3deff2e931eb"},
		{"5466df56-14f3-4df6-83d6-3deff2e931eb/", "5466df56-14f3-4df6-83d6-3deff2e931eb"},
		{"/adasda/5466df56-14f3-4df6-83d6-3deff2e931eb/", "5466df56-14f3-4df6-83d6-3deff2e931eb"},
		{"/adasda/5466df56-14f3-4df6-83d6-3deff2e931eb", "5466df56-14f3-4df6-83d6-3deff2e931eb"},
	} {
		if o := getUUID(test.in); o != test.out {
			t.Error("Expected:", o, "to be:", test.out, "with input:", test.in)
		}
	}
}
