package api

import "fmt"

func (s NoPostHistoryStruct) HasRequiredFields() error {

	if !s.HasStringHeader() {
		return fmt.Errorf("Header does not contain a string")
	}
	if !s.HasStringBody() {
		return fmt.Errorf("Body does not contain a string")
	}

	if !s.HasImageData() {
		return fmt.Errorf("No image data")
	}
	return nil
}

func (s NoPostHistoryStruct) HasStringHeader() bool {
	_, ok := s.Header.(string)
	return ok && s.Header != nil
}

func (s NoPostHistoryStruct) HasStringBody() bool {
	_, ok := s.Body.(string)
	return ok && s.Body != nil
}

func (s NoPostHistoryStruct) HasImageData() bool {
	return s.ImageData != nil
}
