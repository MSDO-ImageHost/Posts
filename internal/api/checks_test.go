package api

import (
	"testing"
	"time"
)

func TestNoPostHistoryStructHasStringHeader(t *testing.T) {

	post := NoPostHistoryStruct{Header: "Header is a string"}
	if post.HasStringHeader() != true {
		t.Error("Failed to validate header as string, when it is.")
	}

	post.Header = nil
	if post.HasStringHeader() != false {
		t.Error("Incorrectly validated header as a string.")
	}

	post.Header = int(123)
	if post.HasStringHeader() != false {
		t.Error("Incorrectly validated header as a string.")
	}

	post.Header = float32(45.87)
	if post.HasStringHeader() != false {
		t.Error("Incorrectly validated header as a string.")
	}

	post.Header = false
	if post.HasStringHeader() != false {
		t.Error("Incorrectly validated header as a string.")
	}

	post.Header = true
	if post.HasStringHeader() != false {
		t.Error("Incorrectly validated header as a string.")
	}

	post.Header = time.Now()
	if post.HasStringHeader() != false {
		t.Error("Incorrectly validated header as a string.")
	}
}

func TestNoPostHistoryStructHasStringBody(t *testing.T) {
	post := NoPostHistoryStruct{Body: "Body is a string"}
	if post.HasStringBody() != true {
		t.Error("Failed to validate body as string, when it is.")
	}

	post.Body = nil
	if post.HasStringBody() != false {
		t.Error("Incorrectly validated body as a string.")
	}

	post.Body = int(123)
	if post.HasStringBody() != false {
		t.Error("Incorrectly validated body as a string.")
	}

	post.Body = float32(45.87)
	if post.HasStringBody() != false {
		t.Error("Incorrectly validated body as a string.")
	}

	post.Body = false
	if post.HasStringBody() != false {
		t.Error("Incorrectly validated body as a string.")
	}

	post.Body = true
	if post.HasStringBody() != false {
		t.Error("Incorrectly validated body as a string.")
	}

	post.Body = time.Now()
	if post.HasStringBody() != false {
		t.Error("Incorrectly validated body as a string.")
	}
}

func TestNoPostHistoryStructHasImageData(t *testing.T) {

	imageData := []byte("asidfjasimagedataiusdhajsd")

	post := NoPostHistoryStruct{ImageData: &imageData}
	if post.HasImageData() != true {
		t.Error("Incorrectly validated image data to be absent.")
	}

	post.ImageData = nil
	if post.HasImageData() != false {
		t.Error("Incorrectly validated image data to be present.")
	}
}

func TestNoPostHistoryStructHasRequiredFields(t *testing.T) {

	imageData := []byte("asidfjasimagedataiusdhajsd")

	post := NoPostHistoryStruct{Header: "This is the header", Body: "This is the body", ImageData: &imageData}
	if err := post.HasRequiredFields(); err != nil {
		t.Error("Failed to assert required fields to be present and have correct type", err)
	}

	post.Header = nil
	if err := post.HasRequiredFields(); err == nil {
		t.Error("Incorrectly validated fields to be present and have correct type.", err)
	}

	post.Body = nil
	if err := post.HasRequiredFields(); err == nil {
		t.Error("Incorrectly validated fields to be present and have correct type.", err)
	}

	post.ImageData = nil
	if err := post.HasRequiredFields(); err == nil {
		t.Error("Incorrectly validated fields to be present and have correct type.", err)
	}

	post.Header = int(123)
	if err := post.HasRequiredFields(); err == nil {
		t.Error("Incorrectly validated fields to be present and have correct type.", err)
	}

	post.Header = float32(123)
	if err := post.HasRequiredFields(); err == nil {
		t.Error("Incorrectly validated fields to be present and have correct type.", err)
	}

	post.Body = int(123)
	if err := post.HasRequiredFields(); err == nil {
		t.Error("Incorrectly validated fields to be present and have correct type.", err)
	}

	post.Body = float32(123)
	if err := post.HasRequiredFields(); err == nil {
		t.Error("Incorrectly validated fields to be present and have correct type.", err)
	}
}
