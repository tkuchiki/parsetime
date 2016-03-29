package parsetime

import (
	"errors"
	"time"
)

var invalidDateTimeError error = errors.New("Invalid date/time")

type ParseTime struct {
	location    *time.Location
	timeFormats []string
}

func NewParseTime() ParseTime {
	return ParseTime{}
}

func NewParseTimeWithLocalTimezone() ParseTime {
	zone, offset := time.Now().In(time.Local).Zone()

	return ParseTime{
		location: time.FixedZone(zone, offset),
	}
}

func NewParseTimeWithTimezone(timezone string) (p ParseTime, err error) {
	var loc *time.Location
	loc, err = time.LoadLocation(timezone)
	if err != nil {
		return p, err
	}

	p = ParseTime{
		location: loc,
	}

	return p, err
}

func NewParseTimeWithOffset(offset int, zone string) ParseTime {
	return ParseTime{
		location: time.FixedZone(zone, offset),
	}
}

func NewParseTimeWithLocation(loc *time.Location) ParseTime {
	return ParseTime{
		location: loc,
	}
}

func (pt *ParseTime) GetFormats() []string {
	return pt.timeFormats
}

func (pt *ParseTime) SetFormats(timeFormats []string) {
	pt.timeFormats = timeFormats
}

func (pt *ParseTime) GetLocation() *time.Location {
	return pt.location
}

func (pt *ParseTime) SetLocation(loc *time.Location) {
	pt.location = loc
}

func (pt *ParseTime) ClearLocation() {
	pt.location = nil
}

func (pt *ParseTime) Parse(timestr string) (t time.Time, err error) {
	if len(pt.timeFormats) == 0 {
		t, err = Parse(timestr)
	} else {
		t, err = _parse(timestr, pt.timeFormats)
	}
	if err != nil {
		return t, err
	}

	if pt.location != nil {
		return t.In(pt.location), err
	} else {
		return t, err
	}
}

func (pt *ParseTime) Iso8601(timestr string) (t time.Time, err error) {
	t, err = Iso8601(timestr)
	if err != nil {
		return t, err
	}

	if pt.location != nil {
		return t.In(pt.location), err
	} else {
		return t, err
	}
}

func (pt *ParseTime) Iso8601T(timestr string, localTimezone ...bool) (t time.Time, err error) {
	t, err = Iso8601T(timestr)
	if err != nil {
		return t, err
	}

	if pt.location != nil {
		return t.In(pt.location), err
	} else {
		return t, err
	}
}

func (pt *ParseTime) Rfc850(timestr string) (t time.Time, err error) {
	t, err = Rfc850(timestr)
	if err != nil {
		return t, err
	}

	if pt.location != nil {
		return t.In(pt.location), err
	} else {
		return t, err
	}
}

func (pt *ParseTime) Rfc822(timestr string) (t time.Time, err error) {
	t, err = Rfc822(timestr)
	if err != nil {
		return t, err
	}

	if pt.location != nil {
		return t.In(pt.location), err
	} else {
		return t, err
	}
}

func (pt *ParseTime) Rfc1123(timestr string) (t time.Time, err error) {
	t, err = Rfc1123(timestr)
	if err != nil {
		return t, err
	}

	if pt.location != nil {
		return t.In(pt.location), err
	} else {
		return t, err
	}
}

func (pt *ParseTime) Rfc3339(timestr string) (t time.Time, err error) {
	t, err = Rfc3339(timestr)
	if err != nil {
		return t, err
	}

	if pt.location != nil {
		return t.In(pt.location), err
	} else {
		return t, err
	}
}

func (pt *ParseTime) GoStandard(timestr string) (t time.Time, err error) {
	t, err = GoStandard(timestr)
	if err != nil {
		return t, err
	}

	if pt.location != nil {
		return t.In(pt.location), err
	} else {
		return t, err
	}
}

func (pt *ParseTime) Log(timestr string) (t time.Time, err error) {
	t, err = Log(timestr)
	if err != nil {
		return t, err
	}

	if pt.location != nil {
		return t.In(pt.location), err
	} else {
		return t, err
	}
}

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
		ISO8601DatetimeOffset2, ISO8601DatetimeOffsetTZD2,
		ISO8601DatetimeFracSecOffset2,
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
