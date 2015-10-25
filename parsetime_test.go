package parsetime

import (
	"testing"
)

func TestParse(test *testing.T) {
	_, err := Parse("2015-09-06T05:58:05+00:00")
	if err != nil {
		test.Errorf("Invalid date/time")
	}
}

func TestParseWithFormats(test *testing.T) {
	_, err := ParseWithFormats("17/Sep/2014:21:14:24 +0000", []string{"02/Jan/2006:15:04:05 -0700"})
	if err != nil {
		test.Errorf("Invalid date/time")
	}
}

func TestIso8601(test *testing.T) {
	_, err := Iso8601("2015-09-06 05:58:05+00:00")
	if err != nil {
		test.Errorf("Invalid date/time")
	}
}

func TestIso8601T(test *testing.T) {
	_, err := Iso8601T("2015-09-06T05:58:05+00:00")
	if err != nil {
		test.Errorf("Invalid date/time")
	}
}

func TestRfc850(test *testing.T) {
	_, err := Rfc850("Sunday, 25-Oct-15 00:03:45 UTC")
	if err != nil {
		test.Errorf("Invalid date/time")
	}
}

func TestRfc822(test *testing.T) {
	_, err := Rfc822("25 Oct 15 00:03 UTC")
	if err != nil {
		test.Errorf("Invalid date/time")
	}
}

func TestRfc1123(test *testing.T) {
	_, err := Rfc1123("Sun, 25 Oct 2015 00:03:45 UTC")
	if err != nil {
		test.Errorf("Invalid date/time")
	}
}

func TestRfc3339(test *testing.T) {
	_, err := Rfc3339("2015-10-25T00:03:45+09:00")
	if err != nil {
		test.Errorf("Invalid date/time")
	}
}

func TestGoStandard(test *testing.T) {
	_, err := GoStandard("Sun Oct 25 00:03:45 +0000 2015")
	if err != nil {
		test.Errorf("Invalid date/time")
	}
}

func TestLog(test *testing.T) {
	_, err := Log("17/Sep/2014:21:14:24 +0000")
	if err != nil {
		test.Errorf("Invalid date/time")
	}
}
