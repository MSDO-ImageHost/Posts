package api

import (
	"fmt"
	"time"
)

func isString(e interface{}) bool {
	_, ok := e.(string)
	return ok
}

func isUInt(e interface{}) bool {
	_, ok := e.(uint)
	return ok
}

func isTime(e interface{}) bool {
	ts, ok := e.(string)
	if !ok {
		return ok
	}
	tt, err := time.Parse(time.RFC3339, ts)
	if err != nil {
		return false
	}
	return !tt.IsZero()
}

func parseString(e interface{}) (string, error) {
	if !isString(e) {
		return "", fmt.Errorf("Could no parse as string")
	}
	str, _ := e.(string)
	return str, nil
}

func parseTime(e interface{}) (time.Time, error) {
	if !isTime(e) {
		return time.Time{}, fmt.Errorf("Timestamp is not a valid RFC3339 format")
	}
	ts, _ := e.(string)
	tt, _ := time.Parse(time.RFC3339, ts)
	return tt, nil
}

func parseUint(e interface{}) (uint, error) {
	if !isUInt(e) {
		return 0, fmt.Errorf("Can't cast to unsigned integer")
	}
	ui, _ := e.(uint)
	return ui, nil
}

// NoPostHistoryStruct methods
func (s NoPostHistoryStruct) HasRequiredFields() error {

	if !s.HasStringHeader() {
		return fmt.Errorf("Payload \"header\" field does not contain a string")
	}
	if !s.HasStringBody() {
		return fmt.Errorf("Payload \"Body\" field does not contain a string")
	}

	if !s.HasImageData() {
		return fmt.Errorf("No image data")
	}
	return nil
}

func (s NoPostHistoryStruct) HasStringHeader() bool {
	return isString(s.Header)
}

func (s NoPostHistoryStruct) HasStringBody() bool {
	return isString(s.Body)
}

func (s NoPostHistoryStruct) HasImageData() bool {
	return s.ImageData != nil
}

// PageStruct methods
func (p PagingStruct) StartIsUInt() bool {
	return isUInt(p.Start)
}

func (p PagingStruct) EndIsUInt() bool {
	return isUInt(p.End)
}

func (p PagingStruct) StartIsTime() bool {
	return isTime(p.Start)
}

func (p PagingStruct) EndIsTime() bool {
	return isTime(p.End)
}

func (p PagingStruct) LimitIsUInt() bool {
	return isUInt(p.Limit)
}

func (p PagingStruct) ParseStartTime() (time.Time, error) {
	return parseTime(p.Start)
}

func (p PagingStruct) ParseEndTime() (time.Time, error) {
	return parseTime(p.End)
}

func (p PagingStruct) ParseStartUInt() (uint, error) {
	return parseUint(p.Start)
}

func (p PagingStruct) ParseEndUInt() (uint, error) {
	return parseUint(p.End)
}

func (p PagingStruct) TimeBased() bool {
	return p.StartIsTime() && p.EndIsTime()
}

func (p PagingStruct) NumberBased() bool {
	return p.StartIsUInt() && p.EndIsUInt()
}

func (p PagingStruct) ValidSetting() error {
	if p.NumberBased() {
		start, err := p.ParseStartUInt()
		if err != nil {
			return fmt.Errorf("Could not parse Start as unsigned integer")
		}
		end, err := p.ParseEndUInt()
		if err != nil {
			return fmt.Errorf("Could not parse End as unsigned integer")
		}
		if start < end {
			return fmt.Errorf("Start can't be greator than End")
		}
	}

	if p.TimeBased() {
		startTime, err := p.ParseStartTime()
		if err != nil {
			return fmt.Errorf("Could not parse Start as time")
		}
		endTime, err := p.ParseEndTime()
		if err != nil {
			return fmt.Errorf("Could not parse End as time")
		}

		if !startTime.Before(endTime) {
			return fmt.Errorf("Start time can't be earlier than End time")
		}
	}
	return nil
}
