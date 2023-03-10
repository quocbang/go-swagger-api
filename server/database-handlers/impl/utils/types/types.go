package types

import (
	"database/sql"
	"database/sql/driver"
	"time"

	"github.com/shopspring/decimal"
)

// TimeNano is the number of nanoseconds elapsed since January 1, 1970 UTC.
type TimeNano int64

// ToTimeNano returns the time nano corresponding to t.
func ToTimeNano(t time.Time) TimeNano {
	return TimeNano(t.UnixNano())
}

// Time returns the local Time corresponding to tn nanoseconds.
func (tn TimeNano) Time() time.Time {
	return time.Unix(0, int64(tn))
}

// NullTimeNano represents an TimeNano that may be null.
// NullTimeNano implements the Scanner interface so
// it can be used as a scan destination, similar to sql.NullInt64.
type NullTimeNano struct {
	TimeNano TimeNano
	Valid    bool // Valid is true if TimeNano is not NULL
}

func (nt *NullTimeNano) Scan(value interface{}) error {
	var i sql.NullInt64
	if err := i.Scan(value); err != nil {
		return err
	}

	*nt = NullTimeNano{
		TimeNano: TimeNano(i.Int64),
		Valid:    i.Valid,
	}
	return nil
}

// Value implements database/sql/driver Valuer interface.
func (nt NullTimeNano) Value() (driver.Value, error) {
	if nt.Valid {
		return int64(nt.TimeNano), nil
	}
	return nil, nil
}

func (nt NullTimeNano) Time() time.Time {
	if nt.Valid {
		return nt.TimeNano.Time()
	}
	return time.Time{}
}

func ToNullTimeNano(t time.Time) NullTimeNano {
	if t.IsZero() {
		return NullTimeNano{
			TimeNano: 0,
			Valid:    false,
		}
	}
	return NullTimeNano{
		TimeNano: ToTimeNano(t),
		Valid:    true,
	}
}

// DecimalNewer definition.
type DecimalNewer interface {
	// RequireFromString returns a Decimal pointer from a string representation or panics.
	RequireFromString(value string) *decimal.Decimal
	// NewFromString returns a new Decimal from a string representation. Trailing zeroes are not trimmed.
	NewFromString(value string) (*decimal.Decimal, error)
	// NewFromInt16 converts a int16 to a Decimal pointer.
	NewFromInt16(value int16) *decimal.Decimal
	// NewFromInt32 converts a int32 to a Decimal pointer.
	NewFromInt32(value int32) *decimal.Decimal
	// NewFromInt64 converts a int64 to a Decimal pointer.
	NewFromInt64(value int64) *decimal.Decimal
	// NewFromFloat32 converts a float32 to a Decimal pointer.
	NewFromFloat32(value float32) *decimal.Decimal
	// NewFromFloat64 converts a float64 to a Decimal pointer.
	NewFromFloat64(value float64) *decimal.Decimal
	// NewFromDecimal converts a Decimal to a Decimal pointer.
	NewFromDecimal(value decimal.Decimal) *decimal.Decimal
}

// Decimal definition.
var Decimal DecimalNewer = &decimalNewer{}

type decimalNewer struct{}

// RequireFromString returns a Decimal pointer from a string representation or panics.
func (*decimalNewer) RequireFromString(value string) *decimal.Decimal {
	res := decimal.RequireFromString(value)
	return &res
}

// NewFromString returns a new Decimal from a string representation. Trailing zeroes are not trimmed.
func (*decimalNewer) NewFromString(value string) (*decimal.Decimal, error) {
	res, err := decimal.NewFromString(value)
	return &res, err
}

// NewFromInt16 converts a int16 to a Decimal pointer.
func (*decimalNewer) NewFromInt16(value int16) *decimal.Decimal {
	res := decimal.NewFromInt32(int32(value))
	return &res
}

// NewFromInt32 converts a int32 to a Decimal pointer.
func (*decimalNewer) NewFromInt32(value int32) *decimal.Decimal {
	res := decimal.NewFromInt32(value)
	return &res
}

// NewFromInt64 converts a int64 to a Decimal pointer.
func (*decimalNewer) NewFromInt64(value int64) *decimal.Decimal {
	res := decimal.NewFromInt(value)
	return &res
}

// NewFromFloat32 converts a float32 to a Decimal pointer.
func (*decimalNewer) NewFromFloat32(value float32) *decimal.Decimal {
	res := decimal.NewFromFloat32(value)
	return &res
}

// NewFromFloat64 converts a float64 to a Decimal pointer.
func (*decimalNewer) NewFromFloat64(value float64) *decimal.Decimal {
	res := decimal.NewFromFloat(value)
	return &res
}

// NewFromDecimal converts a Decimal to a Decimal pointer.
func (*decimalNewer) NewFromDecimal(value decimal.Decimal) *decimal.Decimal {
	return &value
}

// Btoi converts boolean to integer.
func Btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}
