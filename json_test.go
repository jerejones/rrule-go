package rrule

import (
	"encoding/json"
	"testing"
)

func TestRRule_MarshalJSON(t *testing.T) {
	str := "FREQ=WEEKLY;DTSTART=20120201T093000Z;INTERVAL=5;WKST=TU;COUNT=2;UNTIL=20130130T230000Z;BYSETPOS=2;BYMONTH=3;BYYEARDAY=95;BYWEEKNO=1;BYDAY=MO,+2FR;BYHOUR=9;BYMINUTE=30;BYSECOND=0;BYEASTER=-1"
	r, _ := StrToRRule(str)
	toEnc := struct{ Rule *RRule }{r}
	b, _ := json.Marshal(toEnc)
	expected := `{"Rule":"` + str + `"}`
	if string(b) != expected {
		t.Errorf("json.Marshal(StrToRRule(%q)) = %q, want %q", str, string(b), expected)
	}
}

func TestRRule_UnmarshalJSON(t *testing.T) {
	str := "FREQ=WEEKLY;DTSTART=20120201T093000Z;INTERVAL=5;WKST=TU;COUNT=2;UNTIL=20130130T230000Z;BYSETPOS=2;BYMONTH=3;BYYEARDAY=95;BYWEEKNO=1;BYDAY=MO,+2FR;BYHOUR=9;BYMINUTE=30;BYSECOND=0;BYEASTER=-1"
	j := []byte(`{"Rule":"` + str + `"}`)

	var r struct{ Rule *RRule }

	json.Unmarshal(j, &r)
	if s := r.Rule.String(); s != str {
		t.Errorf("json.Unmarshal(%q).String() = %q, want %q", string(j), s, str)
	}
}

func TestSet_MarshalJSON(t *testing.T) {
	str := `RRULE:FREQ=WEEKLY;DTSTART=20120201T093000Z;COUNT=10;BYDAY=MO,TU,WE,TH,FR
			RDATE:20121201T093000Z
			EXRULE:FREQ=WEEKLY;DTSTART=20120208T093000Z;COUNT=3;BYDAY=MO,TU,WE,TH,FR
			EXDATE:20120203T093000Z`
	s, _ := StrToRRuleSet(str)
	toEnc := struct{ RuleSet *Set }{s}
	b, _ := json.Marshal(toEnc)
	expected := `{"RuleSet":["RRULE:FREQ=WEEKLY;DTSTART=20120201T093000Z;COUNT=10;BYDAY=MO,TU,WE,TH,FR","RDATE:20121201T093000Z","EXRULE:FREQ=WEEKLY;DTSTART=20120208T093000Z;COUNT=3;BYDAY=MO,TU,WE,TH,FR","EXDATE:20120203T093000Z"]}`
	if string(b) != expected {
		t.Errorf("json.Marshal(StrToRRule(%q)) = %q, want %q", str, string(b), expected)
	}
}

func TestSet_UnmarshalJSON(t *testing.T) {
	str := `RRULE:FREQ=WEEKLY;DTSTART=20120201T093000Z;COUNT=10;BYDAY=MO,TU,WE,TH,FR
			RDATE:20121201T093000Z
			EXRULE:FREQ=WEEKLY;DTSTART=20120208T093000Z;COUNT=3;BYDAY=MO,TU,WE,TH,FR
			EXDATE:20120203T093000Z`
	j := []byte(`{"RuleSet":["RRULE:FREQ=WEEKLY;DTSTART=20120201T093000Z;COUNT=10;BYDAY=MO,TU,WE,TH,FR","RDATE:20121201T093000Z","EXRULE:FREQ=WEEKLY;DTSTART=20120208T093000Z;COUNT=3;BYDAY=MO,TU,WE,TH,FR","EXDATE:20120203T093000Z"]}`)

	var r struct{ RuleSet *Set }

	json.Unmarshal(j, &r)
	if s := r.RuleSet.String(); s != str {
		t.Errorf("json.Unmarshal(%q).String() = %q, want %q", string(j), s, str)
	}
}
