package main

import "testing"

func TestHamming(t *testing.T) {
	for _, tc := range testCases {
		got, err := Distance(tc.s1, tc.s2)
		if tc.expectedError {
			var _ error = err

			if err == nil {
				t.Fatalf("Distance(%q, %q); expected error, got nil", tc.s1, tc.s2)
			}
		} else {
			if got != tc.want {
				t.Fatalf("Distance(%q, %q) = %d, want %d.", tc.s1, tc.s2, got, tc.want)
			}

			if err != nil {
				t.Fatalf("Distance(%q, %q) returned unexpected error: %v", tc.s1, tc.s1, err)
			}
		}
	}
}

func BenchmrkHamming(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			Distance(tc.s1, tc.s2)
		}
	}
}
