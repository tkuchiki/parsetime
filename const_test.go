package parsetime

import (
	"testing"
)

func TestConstDate(test *testing.T) {
	_, err := ParseWithFormats("20151024", []string{Date})
	if err != nil {
		test.Error("Invalid date/time")
	}
}

func TestConstMMDDYYYY(test *testing.T) {
	_, err := ParseWithFormats("10/24/2015", []string{MMDDYYYY})
	if err != nil {
		test.Error("Invalid date/time")
	}
}

func TestConstShortTime(test *testing.T) {
	_, err := ParseWithFormats("05:58 PM", []string{ShortTime})
	if err != nil {
		test.Error("Invalid date/time")
	}

	_, err = ParseWithFormats("05:58 AM", []string{ShortTime})
	if err != nil {
		test.Error("Invalid date/time")
	}
}

func TestConstLongTime(test *testing.T) {
	_, err := ParseWithFormats("05:58:03 PM", []string{LongTime})
	if err != nil {
		test.Error("Invalid date/time")
	}

	_, err = ParseWithFormats("05:58:03 AM", []string{LongTime})
	if err != nil {
		test.Error("Invalid date/time")
	}
}

func TestConstANSICAsctime(test *testing.T) {
	_, err := ParseWithFormats("Sun Oct 25 00:03:45 2015", []string{ANSICAsctime})
	if err != nil {
		test.Error("Invalid date/time")
	}
}

func TestConstRFC850NoWeekday(test *testing.T) {
	_, err := ParseWithFormats("25-Oct-15 00:03:45 UTC", []string{RFC850NoWeekday})
	if err != nil {
		test.Error("Invalid date/time")
	}
}

func TestConstRFC850NoTime(test *testing.T) {
	_, err := ParseWithFormats("25-Oct-15", []string{RFC850NoTime})
	if err != nil {
		test.Error("Invalid date/time")
	}
}

func TestConstBrokenRFC850(test *testing.T) {
	_, err := ParseWithFormats("Sunday, 25-Oct-2015 00:03:45 UTC", []string{BrokenRFC850})
	if err != nil {
		test.Error("Invalid date/time")
	}
}

func TestConstBrokenRFC850NoWeekday(test *testing.T) {
	_, err := ParseWithFormats("25-Oct-2015 00:03:45 UTC", []string{BrokenRFC850NoWeekday})
	if err != nil {
		test.Error("Invalid date/time")
	}
}

func TestConstBrokenRFC850NoTime(test *testing.T) {
	_, err := ParseWithFormats("25-Oct-2015", []string{BrokenRFC850NoTime})
	if err != nil {
		test.Error("Invalid date/time")
	}
}

func TestConstCommonLog(test *testing.T) {
	_, err := ParseWithFormats("17/Sep/2014:21:14:24 +0000", []string{CommonLog})
	if err != nil {
		test.Error("Invalid date/time")
	}
}

func TestConstCommonLogNoTime(test *testing.T) {
	_, err := ParseWithFormats("17/Sep/2014", []string{CommonLogNoTime})
	if err != nil {
		test.Error("Invalid date/time")
	}
}

func TestConstLsDate(test *testing.T) {
	_, err := ParseWithFormats("Oct 25 2015", []string{LsDate})
	if err != nil {
		test.Error("Invalid date/time")
	}
}

func TestConstLsTime(test *testing.T) {
	_, err := ParseWithFormats("Oct 25 00:03", []string{LsTime})
	if err != nil {
		test.Error("Invalid date/time")
	}
}

func TestConstWindowsDir(test *testing.T) {
	_, err := ParseWithFormats("10-25-15 00:03AM", []string{WindowsDir})
	if err != nil {
		test.Error("Invalid date/time")
	}
}

func TestConstISO8601DateHour(test *testing.T) {
	_, err := ParseWithFormats("2015-10-25 00", []string{ISO8601DateHour})
	if err != nil {
		test.Error("Invalid date/time")
	}
}

func TestConstISO8601DateHourMin(test *testing.T) {
	_, err := ParseWithFormats("2015-10-25 00:03", []string{ISO8601DateHourMin})
	if err != nil {
		test.Error("Invalid date/time")
	}
}

func TestConstISO8601Datetime(test *testing.T) {
	_, err := ParseWithFormats("2015-10-25 00:03:45", []string{ISO8601Datetime})
	if err != nil {
		test.Error("Invalid date/time")
	}
}

func TestConstISO8601DatetimeCompact(test *testing.T) {
	_, err := ParseWithFormats("20151025 000345", []string{ISO8601DatetimeCompact})
	if err != nil {
		test.Error("Invalid date/time")
	}
}

func TestConstISO8601DatetimeCompactUTC(test *testing.T) {
	_, err := ParseWithFormats("20151025 000345Z", []string{ISO8601DatetimeCompactUTC})
	if err != nil {
		test.Error("Invalid date/time")
	}
}

func TestConstISO8601DatetimeOffset(test *testing.T) {
	_, err := ParseWithFormats("2015-10-25 00:03:45+00:00", []string{ISO8601DatetimeOffset})
	if err != nil {
		test.Error("Invalid date/time")
	}
}

func TestConstISO8601DatetimeOffsetTZD(test *testing.T) {
	_, err := ParseWithFormats("2015-10-25 00:03:45+00:00 UTC", []string{ISO8601DatetimeOffsetTZD})
	if err != nil {
		test.Error("Invalid date/time")
	}
}

func TestConstISO8601DatetimeFracSecOffset(test *testing.T) {
	_, err := ParseWithFormats("2015-10-25 00:03:45.123456789+00:00", []string{ISO8601DatetimeFracSecOffset})
	if err != nil {
		test.Error("Invalid date/time")
	}
}

func TestConstISO8601DatetimeFracSecUTC(test *testing.T) {
	_, err := ParseWithFormats("2015-10-25 00:03:45.123456789Z", []string{ISO8601DatetimeFracSecUTC})
	if err != nil {
		test.Error("Invalid date/time")
	}
}

func TestConstISO8601DatetimeFracSec(test *testing.T) {
	_, err := ParseWithFormats("2015-10-25 00:03:45.123456789", []string{ISO8601DatetimeFracSec})
	if err != nil {
		test.Error("Invalid date/time")
	}
}

func TestConstISO8601DatetimeOffset2(test *testing.T) {
	_, err := ParseWithFormats("2015-10-25 00:03:45 +00:00", []string{ISO8601DatetimeOffset2})
	if err != nil {
		test.Error("Invalid date/time")
	}
}

func TestConstISO8601DatetimeOffsetTZD2(test *testing.T) {
	_, err := ParseWithFormats("2015-10-25 00:03:45 +00:00 UTC", []string{ISO8601DatetimeOffsetTZD2})
	if err != nil {
		test.Error("Invalid date/time")
	}
}

func TestConstISO8601DatetimeFracSecOffset2(test *testing.T) {
	_, err := ParseWithFormats("2015-10-25 00:03:45.123456789 +00:00", []string{ISO8601DatetimeFracSecOffset2})
	if err != nil {
		test.Error("Invalid date/time")
	}
}

func TestConstISO8601TDateHour(test *testing.T) {
	_, err := ParseWithFormats("2015-10-25T00", []string{ISO8601TDateHour})
	if err != nil {
		test.Error("Invalid date/time")
	}
}

func TestConstISO8601TDateHourMin(test *testing.T) {
	_, err := ParseWithFormats("2015-10-25T00:03", []string{ISO8601TDateHourMin})
	if err != nil {
		test.Error("Invalid date/time")
	}
}

func TestConstISO8601TDatetime(test *testing.T) {
	_, err := ParseWithFormats("2015-10-25T00:03:45", []string{ISO8601TDatetime})
	if err != nil {
		test.Error("Invalid date/time")
	}
}

func TestConstISO8601TDatetimeCompact(test *testing.T) {
	_, err := ParseWithFormats("20151025T000345", []string{ISO8601TDatetimeCompact})
	if err != nil {
		test.Error("Invalid date/time")
	}
}

func TestConstISO8601TDatetimeCompactUTC(test *testing.T) {
	_, err := ParseWithFormats("20151025T000345Z", []string{ISO8601TDatetimeCompactUTC})
	if err != nil {
		test.Error("Invalid date/time")
	}
}

func TestConstISO8601TDatetimeOffset(test *testing.T) {
	_, err := ParseWithFormats("2015-10-25T00:03:45+00:00", []string{ISO8601TDatetimeOffset})
	if err != nil {
		test.Error("Invalid date/time")
	}
}

func TestConstISO8601TDatetimeOffsetTZD(test *testing.T) {
	_, err := ParseWithFormats("2015-10-25T00:03:45+00:00 UTC", []string{ISO8601TDatetimeOffsetTZD})
	if err != nil {
		test.Error("Invalid date/time")
	}
}

func TestConstISO8601TDatetimeFracSecOffset(test *testing.T) {
	_, err := ParseWithFormats("2015-10-25T00:03:45.123456789+00:00", []string{ISO8601TDatetimeFracSecOffset})
	if err != nil {
		test.Error("Invalid date/time")
	}
}

func TestConstISO8601TDatetimeFracSecUTC(test *testing.T) {
	_, err := ParseWithFormats("2015-10-25T00:03:45.123456789Z", []string{ISO8601TDatetimeFracSecUTC})
	if err != nil {
		test.Error("Invalid date/time")
	}
}

func TestConstISO8601TDatetimeFracSec(test *testing.T) {
	_, err := ParseWithFormats("2015-10-25T00:03:45.123456789", []string{ISO8601TDatetimeFracSec})
	if err != nil {
		test.Error("Invalid date/time")
	}
}

func TestConstISO8601Date(test *testing.T) {
	_, err := ParseWithFormats("2015-10-25", []string{ISO8601Date})
	if err != nil {
		test.Error("Invalid date/time")
	}
}

func TestConstISO8601Time(test *testing.T) {
	_, err := ParseWithFormats("00:03:45", []string{ISO8601Time})
	if err != nil {
		test.Error("Invalid date/time")
	}
}

func TestConstISO8601TimeOffset(test *testing.T) {
	_, err := ParseWithFormats("00:03:45+00:00", []string{ISO8601TimeOffset})
	if err != nil {
		test.Error("Invalid date/time")
	}
}

func TestConstMMDDYYYYDatetime(test *testing.T) {
	_, err := ParseWithFormats("10/25/2015 00:03:45", []string{MMDDYYYYDatetime})
	if err != nil {
		test.Error("Invalid date/time")
	}
}
