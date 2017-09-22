package rrule

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func timeToStr(time time.Time) string {
	return time.UTC().Format("20060102T150405Z")
}

func strToTime(str string) (time.Time, error) {
	return time.Parse("20060102T150405Z", str)
}

func (f Frequency) String() string {
	return [...]string{
		"YEARLY", "MONTHLY", "WEEKLY", "DAILY",
		"HOURLY", "MINUTELY", "SECONDLY"}[f]
}

func strToFreq(str string) (Frequency, error) {
	freqMap := map[string]Frequency{
		"YEARLY": YEARLY, "MONTHLY": MONTHLY, "WEEKLY": WEEKLY, "DAILY": DAILY,
		"HOURLY": HOURLY, "MINUTELY": MINUTELY, "SECONDLY": SECONDLY,
	}
	result, ok := freqMap[str]
	if !ok {
		return 0, errors.New("undefined frequency: " + str)
	}
	return result, nil
}

func (wday Weekday) String() string {
	s := [...]string{"MO", "TU", "WE", "TH", "FR", "SA", "SU"}[wday.weekday]
	if wday.n == 0 {
		return s
	}
	return fmt.Sprintf("%+d%s", wday.n, s)
}

func strToWeekday(str string) (Weekday, error) {
	if len(str) < 2 {
		return Weekday{}, errors.New("undefined weekday: " + str)
	}
	weekMap := map[string]Weekday{
		"MO": MO, "TU": TU, "WE": WE, "TH": TH,
		"FR": FR, "SA": SA, "SU": SU}
	result, ok := weekMap[str[len(str)-2:]]
	if !ok {
		return Weekday{}, errors.New("undefined weekday: " + str)
	}
	if len(str) > 2 {
		n, e := strconv.Atoi(str[:len(str)-2])
		if e != nil {
			return Weekday{}, e
		}
		result.n = n
	}
	return result, nil
}

func strToWeekdays(value string) ([]Weekday, error) {
	contents := strings.Split(value, ",")
	result := make([]Weekday, len(contents))
	var e error
	for i, s := range contents {
		result[i], e = strToWeekday(s)
		if e != nil {
			return nil, e
		}
	}
	return result, nil
}

func appendIntsOption(options []string, key string, value []int) []string {
	if len(value) == 0 {
		return options
	}
	valueStr := make([]string, len(value))
	for i, v := range value {
		valueStr[i] = strconv.Itoa(v)
	}
	return append(options, fmt.Sprintf("%s=%s", key, strings.Join(valueStr, ",")))
}

func strToInts(value string) ([]int, error) {
	contents := strings.Split(value, ",")
	result := make([]int, len(contents))
	var e error
	for i, s := range contents {
		result[i], e = strconv.Atoi(s)
		if e != nil {
			return nil, e
		}
	}
	return result, nil
}

func (option *ROption) String() string {
	result := []string{fmt.Sprintf("FREQ=%v", option.Freq)}
	if !option.Dtstart.IsZero() {
		result = append(result, fmt.Sprintf("DTSTART=%s", timeToStr(option.Dtstart)))
	}
	if option.Interval != 0 {
		result = append(result, fmt.Sprintf("INTERVAL=%v", option.Interval))
	}
	if option.Wkst != MO {
		result = append(result, fmt.Sprintf("WKST=%v", option.Wkst))
	}
	if option.Count != 0 {
		result = append(result, fmt.Sprintf("COUNT=%v", option.Count))
	}
	if !option.Until.IsZero() {
		result = append(result, fmt.Sprintf("UNTIL=%v", timeToStr(option.Until)))
	}
	result = appendIntsOption(result, "BYSETPOS", option.Bysetpos)
	result = appendIntsOption(result, "BYMONTH", option.Bymonth)
	result = appendIntsOption(result, "BYMONTHDAY", option.Bymonthday)
	result = appendIntsOption(result, "BYYEARDAY", option.Byyearday)
	result = appendIntsOption(result, "BYWEEKNO", option.Byweekno)
	if len(option.Byweekday) != 0 {
		valueStr := make([]string, len(option.Byweekday))
		for i, wday := range option.Byweekday {
			valueStr[i] = wday.String()
		}
		result = append(result, fmt.Sprintf("BYDAY=%s", strings.Join(valueStr, ",")))
	}
	result = appendIntsOption(result, "BYHOUR", option.Byhour)
	result = appendIntsOption(result, "BYMINUTE", option.Byminute)
	result = appendIntsOption(result, "BYSECOND", option.Bysecond)
	result = appendIntsOption(result, "BYEASTER", option.Byeaster)
	return strings.Join(result, ";")
}

// StrToROption converts string to ROption
func StrToROption(rfcString string) (*ROption, error) {
	rfcString = strings.TrimSpace(rfcString)
	if len(rfcString) == 0 {
		return nil, errors.New("empty string")
	}
	result := ROption{}
	for _, attr := range strings.Split(rfcString, ";") {
		keyValue := strings.Split(attr, "=")
		if len(keyValue) != 2 {
			return nil, errors.New("wrong format")
		}
		key, value := keyValue[0], keyValue[1]
		if len(value) == 0 {
			return nil, errors.New(key + " option has no value")
		}
		var e error
		switch key {
		case "FREQ":
			result.Freq, e = strToFreq(value)
		case "DTSTART":
			result.Dtstart, e = strToTime(value)
		case "INTERVAL":
			result.Interval, e = strconv.Atoi(value)
		case "WKST":
			result.Wkst, e = strToWeekday(value)
		case "COUNT":
			result.Count, e = strconv.Atoi(value)
		case "UNTIL":
			result.Until, e = strToTime(value)
		case "BYSETPOS":
			result.Bysetpos, e = strToInts(value)
		case "BYMONTH":
			result.Bymonth, e = strToInts(value)
		case "BYMONTHDAY":
			result.Bymonthday, e = strToInts(value)
		case "BYYEARDAY":
			result.Byyearday, e = strToInts(value)
		case "BYWEEKNO":
			result.Byweekno, e = strToInts(value)
		case "BYDAY":
			result.Byweekday, e = strToWeekdays(value)
		case "BYHOUR":
			result.Byhour, e = strToInts(value)
		case "BYMINUTE":
			result.Byminute, e = strToInts(value)
		case "BYSECOND":
			result.Bysecond, e = strToInts(value)
		case "BYEASTER":
			result.Byeaster, e = strToInts(value)
		default:
			return nil, errors.New("unknown RRULE property: " + key)
		}
		if e != nil {
			return nil, e
		}
	}
	return &result, nil
}

func (r *RRule) String() string {
	return r.OrigOptions.String()
}

func (s *Set) String() string {
	return strings.Join(s.spec(), "\n")
}

func (s *Set) spec() []string {
	ret := make([]string, 0, len(s.rrule)+len(s.rdate)+len(s.exrule)+len(s.exdate))
	for _, r := range s.rrule {
		ret = append(ret, "RRULE:"+r.String())
	}
	for _, r := range s.rdate {
		ret = append(ret, "RDATE:"+timeToStr(r))
	}
	for _, r := range s.exrule {
		ret = append(ret, "EXRULE:"+r.String())
	}
	for _, r := range s.exdate {
		ret = append(ret, "EXDATE:"+timeToStr(r))
	}
	return ret
}

// StrToRRule converts string to RRule
func StrToRRule(rfcString string) (*RRule, error) {
	option, e := StrToROption(rfcString)
	if e != nil {
		return nil, e
	}
	return NewRRule(*option)
}

// StrToRRuleSet converts string to RRuleSet
func StrToRRuleSet(s string) (*Set, error) {
	s = strings.TrimSpace(strings.ToUpper(s))
	if s == "" {
		return nil, errors.New("empty string")
	}
	set := Set{}
	for _, line := range strings.Split(s, "\n") {
		err := set.add(line)
		if err != nil {
			return nil, err
		}
	}
	return &set, nil
}
