package main

var testCases = []struct {
	s1            string
	s2            string
	want          int
	expectedError bool
}{
	{
		"",
		"",
		0,
		false,
	},
	{
		"A",
		"A",
		0,
		false,
	},
	{
		"GGACTGAAATCTG",
		"GGACTGAAATCTG",
		0,
		false,
	},
	{
		"AATG",
		"AAA",
		0,
		true,
	},
	{
		"ATA",
		"AGTG",
		0,
		true,
	},
}
