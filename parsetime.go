package parsetime

import (
	"errors"
	"time"
)

var invalidDateTimeError error = errors.New("Invalid date/time")

func _parse(timestr string, timeFormats []string) (t time.Time, err error) {
	for _, format := range timeFormats {
		t, err = time.Parse(format, timestr)
		if err == nil {
			return t, err
		}
	}

	return t, invalidDateTimeError
}

func Parse(timestr string) (t time.Time, err error) {
	var timeFormats = []string{
		Date, MMDDYYYY, ShortTime, LongTime, ANSICAsctime,
		RFC850NoWeekday, RFC850NoTime, BrokenRFC850,
		BrokenRFC850NoWeekday, BrokenRFC850NoTime,
		CommonLog, CommonLogNoTime, LsDate, LsTime, WindowsDir,
		ISO8601DateHour, ISO8601DateHourMin, ISO8601Datetime,
		ISO8601Date, ISO8601DatetimeCompact, ISO8601DatetimeCompactUTC,
		ISO8601DatetimeOffset, ISO8601DatetimeOffsetTZD,
		ISO8601DatetimeFracSecOffset, ISO8601DatetimeFracSecUTC,
		ISO8601DatetimeFracSec,
		ISO8601TDatetime, ISO8601TDatetimeCompact,
		ISO8601TDatetimeCompactUTC, ISO8601TDateHour, ISO8601TDateHourMin,
		ISO8601TDatetimeOffset, ISO8601TDatetimeOffsetTZD,
		ISO8601TDatetimeFracSecOffset, ISO8601TDatetimeFracSecUTC,
		ISO8601TDatetimeFracSec, ISO8601Time, ISO8601TimeOffset,
		MMDDYYYYDatetime,
		time.ANSIC, time.UnixDate, time.RubyDate, time.RFC822,
		time.RFC822Z, time.RFC850, time.RFC1123, time.RFC1123Z,
		time.RFC3339, time.RFC3339Nano, time.Kitchen,
		time.Stamp, time.StampMilli, time.StampMicro, time.StampNano,
	}

	t, err = _parse(timestr, timeFormats)

	return t, err
}

func ParseWithFormats(timestr string, timeFormats []string) (t time.Time, err error) {
	t, err = _parse(timestr, timeFormats)

	return t, err
}

func Iso8601(timestr string) (t time.Time, err error) {
	var timeFormats = []string{
		ISO8601DateHour, ISO8601DateHourMin, ISO8601Datetime,
		ISO8601Date, ISO8601DatetimeCompact, ISO8601DatetimeCompactUTC,
		ISO8601DatetimeOffset, ISO8601DatetimeOffsetTZD,
		ISO8601DatetimeFracSecOffset, ISO8601DatetimeFracSecUTC,
		ISO8601DatetimeFracSec,
		ISO8601Date, ISO8601Time, ISO8601TimeOffset,
	}

	t, err = _parse(timestr, timeFormats)

	return t, err
}

func Iso8601T(timestr string) (t time.Time, err error) {
	var timeFormats = []string{
		ISO8601TDateHour, ISO8601TDateHourMin, ISO8601TDatetime,
		ISO8601TDatetimeCompact, ISO8601TDatetimeCompactUTC,
		ISO8601TDatetimeOffset, ISO8601TDatetimeOffsetTZD,
		ISO8601TDatetimeFracSecOffset, ISO8601TDatetimeFracSecUTC,
		ISO8601TDatetimeFracSec,
		ISO8601Date, ISO8601Time, ISO8601TimeOffset,
	}

	t, err = _parse(timestr, timeFormats)

	return t, err
}

func Rfc850(timestr string) (t time.Time, err error) {
	var timeFormats = []string{
		RFC850NoWeekday, RFC850NoTime, BrokenRFC850,
		BrokenRFC850NoWeekday, BrokenRFC850NoTime, time.RFC850,
	}

	t, err = _parse(timestr, timeFormats)

	return t, err
}

func Rfc822(timestr string) (t time.Time, err error) {
	var timeFormats = []string{
		time.RFC822, time.RFC822Z,
	}

	t, err = _parse(timestr, timeFormats)

	return t, err
}

func Rfc1123(timestr string) (t time.Time, err error) {
	var timeFormats = []string{
		time.RFC1123, time.RFC1123Z,
	}

	t, err = _parse(timestr, timeFormats)

	return t, err
}

func Rfc3339(timestr string) (t time.Time, err error) {
	var timeFormats = []string{
		time.RFC3339, time.RFC3339Nano,
	}

	t, err = _parse(timestr, timeFormats)

	return t, err
}

func GoStandard(timestr string) (t time.Time, err error) {
	var timeFormats = []string{
		time.ANSIC, time.UnixDate, time.RubyDate, time.RFC822,
		time.RFC822Z, time.RFC850, time.RFC1123, time.RFC1123Z,
		time.RFC3339, time.RFC3339Nano, time.Kitchen,
		time.Stamp, time.StampMilli, time.StampMicro, time.StampNano,
	}

	t, err = _parse(timestr, timeFormats)

	return t, err
}

func Log(timestr string) (t time.Time, err error) {
	var timeFormats = []string{
		ISO8601TDatetimeOffset, ISO8601TDatetime, ISO8601TDatetimeOffsetTZD,
		ISO8601DatetimeOffset, ISO8601Datetime, ISO8601DatetimeOffsetTZD,
		ISO8601DatetimeOffset2, CommonLog, ISO8601DatetimeOffsetTZD2,
	}

	t, err = _parse(timestr, timeFormats)

	return t, err
}
