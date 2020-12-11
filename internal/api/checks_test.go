package api

import (
	"testing"
	"time"
)

func TestIsString(t *testing.T) {
	if isString("this is a string") != true {
		t.Error("Failed to verify string")
	}
	if isString(123) != false {
		t.Error("Incorrectly verified integer")
	}
	if isString(float32(123)) != false {
		t.Error("Incorrectly verified integer")
	}
	if isString(time.Now()) != false {
		t.Error("Incorrectly verified time.Time")
	}
	if isString([]byte("hehehe")) != false {
		t.Error("Incorrectly verified byte array")
	}
	if isString(map[int]string{2: "asd"}) != false {
		t.Error("Incorrectly verified map")
	}
}

func TestIsUint(t *testing.T) {
	if isUInt(uint(123)) != true {
		t.Error("Failed to verify unsigned integer")
	}
	if isUInt("this is a string") != false {
		t.Error("Incorrectly verified unsigned integer")
	}
	if isUInt(int(-123)) != false {
		t.Error("Incorrectly verified unsigned integer")
	}
	if isUInt(float32(123)) != false {
		t.Error("Incorrectly verified unsigned integer")
	}
	if isUInt(time.Now()) != false {
		t.Error("Incorrectly verified unsigned integer")
	}
	if isUInt([]byte("hehehe")) != false {
		t.Error("Incorrectly verified unsigned integer")
	}
	if isUInt(map[int]string{2: "asd"}) != false {
		t.Error("Incorrectly verified unsigned integer")
	}
}

func TestIsTime(t *testing.T) {
	if isTime(time.Now().Format(time.RFC3339)) != true {
		t.Error("Failed to verify RFC3339 timestamp")
	}
	if isTime("this is a string") != false {
		t.Error("Incorrectly verified RFC3339 timestamp")
	}
	if isTime("this is a string") != false {
		t.Error("Incorrectly verified RFC3339 timestamp")
	}
	if isTime(-123) != false {
		t.Error("Incorrectly verified RFC3339 timestamp")
	}
	if isTime(float32(123)) != false {
		t.Error("Incorrectly verified RFC3339 timestamp")
	}
	if isTime(time.Now()) != false {
		t.Error("Incorrectly verified RFC3339 timestamp")
	}
	if isTime([]byte("hehehe")) != false {
		t.Error("Incorrectly verified RFC3339 timestamp")
	}
	if isTime(map[int]string{2: "asd"}) != false {
		t.Error("Incorrectly verified RFC3339 timestamp")
	}
}

func TestPagingValidSetting(t *testing.T) {

	paging := PagingStruct{}
	if paging.ValidSetting() != nil {
		t.Error("Asserted valid setting as invalid")
	}

	// Testing for numbers
	paging.Start = 0
	paging.End = 9
	if paging.ValidSetting() != nil {
		t.Error("Asserted valid setting as invalid")
	}

	paging.Start = -10
	paging.End = -5
	if paging.ValidSetting() == nil {
		t.Error("Asserted invalid setting as valid")
	}

	// Testing for time formats
	paging.Start = time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC).Format(time.RFC3339)
	paging.End = time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC).Format(time.RFC3339)
	if paging.ValidSetting() != nil {
		t.Error("Asserted valid setting as invalid")
	}

	paging.Start = time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC).Format(time.RFC3339)
	paging.End = time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC).Format(time.RFC3339)
	if paging.ValidSetting() == nil {
		t.Error("Asserted invalid setting as valid")
	}

}
