package main

import "testing"

func Test_Main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			run()
		})
	}

    _,err := run()
    if err != nil {
        t.Error("failed running main")
    }
}
