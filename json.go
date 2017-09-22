package rrule

import "encoding/json"

// MarshalJSON serializes the given RRule in string format
func (r *RRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

// UnmarshalJSON reads an RRule from it's JSON representation
func (r *RRule) UnmarshalJSON(data []byte) error {
	var str string
	err := json.Unmarshal(data, &str)
	if err != nil {
		return err
	}
	rule, err := StrToRRule(str)
	if err != nil {
		return err
	}
	r.dupFrom(rule)
	return nil
}

// MarshalJSON serializes the given Set in []string format
func (s *Set) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.spec())
}

// UnmarshalJSON reads a Set from it's JSON representation
func (s *Set) UnmarshalJSON(data []byte) error {
	var lines []string
	err := json.Unmarshal(data, &lines)
	if err != nil {
		return err
	}
	for _, line := range lines {
		err = s.add(line)
		if err != nil {
			return err
		}
	}
	return nil
}
