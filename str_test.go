package rrule

import (
	"testing"
)

func TestStrToRRule(t *testing.T) {
	str := "FREQ=WEEKLY;DTSTART=20120201T093000Z;INTERVAL=5;WKST=TU;COUNT=2;UNTIL=20130130T230000Z;BYSETPOS=2;BYMONTH=3;BYYEARDAY=95;BYWEEKNO=1;BYDAY=MO,+2FR;BYHOUR=9;BYMINUTE=30;BYSECOND=0;BYEASTER=-1"
	r, _ := StrToRRule(str)
	if s := r.String(); s != str {
		t.Errorf("StrToRRule(%q).String() = %q, want %q", str, s, str)
	}
}

func TestInvalidString(t *testing.T) {
	cases := []string{
		"",
		"FREQ",
		"FREQ=HELLO",
		"BYMONTH=",
		"FREQ=WEEKLY;HELLO=WORLD",
		"FREQ=WEEKLY;BYMONTHDAY=I",
		"FREQ=WEEKLY;BYDAY=M",
		"FREQ=WEEKLY;BYDAY=MQ",
		"FREQ=WEEKLY;BYDAY=+MO",
	}
	for _, item := range cases {
		if _, e := StrToRRule(item); e == nil {
			t.Errorf("StrToRRule(%q) = nil, want error", item)
		}
	}
}

func TestStrToRRuleSet(t *testing.T) {
	str := `RRULE:FREQ=WEEKLY;DTSTART=20120201T093000Z;BYDAY=MO,TU,WE,TH,FR;COUNT=10
RDATE:20121201T093000Z
EXRULE:FREQ=WEEKLY;DTSTART=20120208T093000Z;BYDAY=MO,TU,WE,TH,FR;COUNT=3
EXDATE:20120203T093000Z`
	r, err := StrToRRuleSet(str)
	if err != nil {
		t.Errorf("StrToRRule(%q) error: %v", str, err)
	}

	if len(r.All()) != 7 {
		t.Errorf("len(StrToRRule(%q)) = %d, want %d", str, len(r.All()), 7)
	}
}
