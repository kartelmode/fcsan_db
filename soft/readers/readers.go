package readers

import (
	"fmt"
	"main/apperrors"
	"main/console"
	"os"
	"strconv"
)

type IntReader struct{}

func (r *IntReader) Read() (int, error) {
	var line string = ""
	_, err := fmt.Scanln(&line)
	if err != nil {
		return 0, &apperrors.ErrInternal{
			Message: err.Error(),
		}
	}
	i, err := strconv.Atoi(line)
	if err != nil {
		return 0, &apperrors.ErrInternal{
			Message: err.Error(),
		}
	}
	return i, nil
}

func (r *IntReader) ReadRange(left, right int) (int, error) {
	var line string = ""
	_, err := fmt.Scanln(&line)
	if err != nil {
		return 0, &apperrors.ErrInternal{
			Message: err.Error(),
		}
	}
	i, err := strconv.Atoi(line)
	if err != nil {
		return 0, &apperrors.ErrInternal{
			Message: err.Error(),
		}
	}
	if left > i || right < i {
		return 0, &apperrors.ErrInternal{
			Message: fmt.Sprintf("Value must be between %d and %d", left, right),
		}
	}
	return i, nil
}

type StringReader struct{}

func (r *StringReader) Read() (string, error) {
	var text string = ""
	_, err := fmt.Scan(&text)
	if err != nil {
		return "", &apperrors.ErrInternal{
			Message: err.Error(),
		}
	}
	return text, nil
}

type PathReader struct{}

func (r *PathReader) Read() (string, error) {
	var text string = ""
	_, err := fmt.Scan(&text)
	if err != nil {
		return "", &apperrors.ErrInternal{
			Message: err.Error(),
		}
	}
	if _, err := os.OpenFile(text, os.O_RDONLY, os.ModeAppend); err != nil {
		return "", &apperrors.ErrInternal{
			Message: "Invalid filepath was provided",
		}
	}
	return text, nil
}

type Readers struct {
	intReader    *IntReader
	stringReader *StringReader
	pathReader   *PathReader
}

func (r *Readers) GetIntReader() console.IntReader {
	return r.intReader
}

func (r *Readers) GetStringReader() console.StringReader {
	return r.stringReader
}

func (r *Readers) GetPathReader() console.PathReader {
	return r.pathReader
}

func NewReaders() console.Readers {
	return &Readers{
		&IntReader{},
		&StringReader{},
		&PathReader{},
	}
}
