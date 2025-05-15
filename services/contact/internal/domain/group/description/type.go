package description

import "github.com/pkg/errors"

var (
	MaxLength      = 1000
	ErrWrongLength = errors.Errorf("name must be less than or equal to %d characters", MaxLength)
)

type Description string

func (d Description) String() string {
	return string(d)
}

func New(description string) (*Description, error) {
	if len([]rune(description)) > MaxLength {
		return nil, ErrWrongLength
	}
	d := Description(description)
	return &d, nil
}
