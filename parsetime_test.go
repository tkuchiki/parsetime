package parsetime

import (
	"testing"
	"time"
)

func TestNewParseTimeWithLocalTimezone(test *testing.T) {
	p := NewParseTimeWithLocalTimezone()

	var t time.Time
	t, _ = p.Parse("1970-01-01T00:00:00")

	zone, offset := time.Now().In(time.Local).Zone()

	zone2, offset2 := t.Zone()
	if offset != offset2 || zone != zone2 {
		test.Error("Incorrect timezone")
	}
}

func TestNewParseTimeWithTimezone(test *testing.T) {
	p, err := NewParseTimeWithTimezone("America/Atikokan")

	if err != nil {
		test.Error(err)
	}

	t, _ := p.Parse("1970-01-01T00:00:00")

	zone, offset := t.Zone()
	if zone != "EST" || offset != -18000 {
		test.Error("Incorrect timezone")
	}
}

func TestNewParseTimeWithOffset(test *testing.T) {
	p := NewParseTimeWithOffset(-18000, "EST")

	var t time.Time
	t, _ = p.Parse("1970-01-01T00:00:00")
	zone, offset := t.Zone()

	loc, _ := time.LoadLocation("America/Atikokan")
	zone2, offset2 := time.Now().In(loc).Zone()

	if offset != offset2 || zone != zone2 {
		test.Error("Incorrect timezone")
	}
}

func TestNewParseTimeWithLocation(test *testing.T) {
	loc, _ := time.LoadLocation("America/Atikokan")
	p := NewParseTimeWithLocation(loc)

	var t time.Time
	t, _ = p.Parse("1970-01-01T00:00:00")

	_, offset := time.Now().In(loc).Zone()

	_, offset2 := t.Zone()
	if offset != offset2 {
		test.Error("Incorrect timezone")
	}
}

func TestSetFormats(test *testing.T) {
	p := NewParseTime()

	p.SetFormats([]string{"02/Jan/2006:15:04:05 -0700"})
	formats := p.GetFormats()

	if formats[0] != "02/Jan/2006:15:04:05 -0700" {
		test.Error("Does not match time formats")
	}
}

func TestLocation(test *testing.T) {
	p := NewParseTime()
	loc, _ := time.LoadLocation("America/Atikokan")
	p.SetLocation(loc)
	loc2 := p.GetLocation()

	if loc.String() != loc2.String() {
		test.Error("Does not match time location")
	}

	p.ClearLocation()
	if p.GetLocation() != nil {
		test.Error("Failed clear time location")
	}
}

func TestParse(test *testing.T) {
	_, err := Parse("2015-09-06T05:58:05+00:00")
	if err != nil {
		test.Error("Invalid date/time")
	}
}

func TestParseWithFormats(test *testing.T) {
	_, err := ParseWithFormats("17/Sep/2014:21:14:24 +0000", []string{"02/Jan/2006:15:04:05 -0700"})
	if err != nil {
		test.Error("Invalid date/time")
	}
}

func TestIso8601(test *testing.T) {
	_, err := Iso8601("2015-09-06 05:58:05+00:00")
	if err != nil {
		test.Error("Invalid date/time")
	}
}

func TestIso8601T(test *testing.T) {
	_, err := Iso8601T("2015-09-06T05:58:05+00:00")
	if err != nil {
		test.Error("Invalid date/time")
	}
}

func TestRfc850(test *testing.T) {
	_, err := Rfc850("Sunday, 25-Oct-15 00:03:45 UTC")
	if err != nil {
		test.Error("Invalid date/time")
	}
}

func TestRfc822(test *testing.T) {
	_, err := Rfc822("25 Oct 15 00:03 UTC")
	if err != nil {
		test.Error("Invalid date/time")
	}
}

func TestRfc1123(test *testing.T) {
	_, err := Rfc1123("Sun, 25 Oct 2015 00:03:45 UTC")
	if err != nil {
		test.Error("Invalid date/time")
	}
}

func TestRfc3339(test *testing.T) {
	_, err := Rfc3339("2015-10-25T00:03:45+09:00")
	if err != nil {
		test.Error("Invalid date/time")
	}
}

func TestGoStandard(test *testing.T) {
	_, err := GoStandard("Sun Oct 25 00:03:45 +0000 2015")
	if err != nil {
		test.Error("Invalid date/time")
	}
}

func TestLog(test *testing.T) {
	_, err := Log("17/Sep/2014:21:14:24 +0000")
	if err != nil {
		test.Error("Invalid date/time")
	}
}
