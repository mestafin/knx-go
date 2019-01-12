// Copyright 2017 Ole Krüger.
// Licensed under the MIT license which can be found in the LICENSE file.

package dpt

import (
	"fmt"
)

// A DatapointValue is a value of a datapoint.
type DatapointValue interface {
	// Pack the datapoint to a byte array.
	Pack() []byte

	// Unpack a the datapoint value from a byte array.
	Unpack(data []byte) error
}

// DatapointMeta gives meta information about a datapoint type.
type DatapointMeta interface {
	// Unit returns the unit of this datapoint type or empty string if it doesn't have a unit.
	Unit() string
}

// DPT_1001 represents DPT 1.001 / Switch.
type DPT_1001 bool

func (d DPT_1001) Pack() []byte {
	return packB1(bool(d))
}

func (d *DPT_1001) Unpack(data []byte) error {
	return unpackB1(data, (*bool)(d))
}

func (d DPT_1001) Unit() string {
	return ""
}

func (d DPT_1001) String() string {
	if d {
		return "On"
	} else {
		return "Off"
	}
}

// DPT_1002 represents DPT 1.002 / Bool.
type DPT_1002 bool

func (d DPT_1002) Pack() []byte {
	return packB1(bool(d))
}

func (d *DPT_1002) Unpack(data []byte) error {
	return unpackB1(data, (*bool)(d))
}

func (d DPT_1002) Unit() string {
	return ""
}

func (d DPT_1002) String() string {
	if d {
		return "True"
	} else {
		return "False"
	}
}

// DPT_1003 represents DPT 1.003 / Enable.
type DPT_1003 bool

func (d DPT_1003) Pack() []byte {
	return packB1(bool(d))
}

func (d *DPT_1003) Unpack(data []byte) error {
	return unpackB1(data, (*bool)(d))
}

func (d DPT_1003) Unit() string {
	return ""
}

func (d DPT_1003) String() string {
	if d {
		return "Enable"
	} else {
		return "Disable"
	}
}

// DPT_1009 represents DPT 1.009 / OpenClose.
type DPT_1009 bool

func (d DPT_1009) Pack() []byte {
	return packB1(bool(d))
}

func (d *DPT_1009) Unpack(data []byte) error {
	return unpackB1(data, (*bool)(d))
}

func (d DPT_1009) Unit() string {
	return ""
}

func (d DPT_1009) String() string {
	if d {
		return "Close"
	} else {
		return "Open"
	}
}

// DPT_1010 represents DPT 1.010 / Start.
type DPT_1010 bool

func (d DPT_1010) Pack() []byte {
	return packB1(bool(d))
}

func (d *DPT_1010) Unpack(data []byte) error {
	return unpackB1(data, (*bool)(d))
}

func (d DPT_1010) Unit() string {
	return ""
}

func (d DPT_1010) String() string {
	if d {
		return "Start"
	} else {
		return "Stop"
	}
}

// DPT_5001 represents DPT 5.001 / Scaling.
type DPT_5001 float32

func (d DPT_5001) Pack() []byte {
	if d <= 0 {
		return packU8(0)
	} else if d >= 100 {
		return packU8(255)
	} else {
		return packU8(uint8(d * 2.55))
	}
}

func (d *DPT_5001) Unpack(data []byte) (err error) {
	var value uint8
	if err := unpackU8(data, &value); err != nil {
		return err
	}

	*d = DPT_5001(value) / 2.55

	return nil
}

func (d DPT_5001) Unit() string {
	return "%"
}

func (d DPT_5001) String() string {
	return fmt.Sprintf("%.2f%%", float32(d))
}

// DPT_5003 represents DPT 5.003 / Angle.
type DPT_5003 float32

func (d DPT_5003) Pack() []byte {
	value := d

	// Clip the value to a valid range
	if value < 0 {
		value = 0
	}
	if value > 360 {
		value = 360
	}

	// Scale the value to the target range
	value /= 360
	value *= 255

	return packU8(uint8(value))
}

func (d *DPT_5003) Unpack(data []byte) error {
	var buf uint8

	err := unpackU8(data, &buf)
	if err == nil {
		value := DPT_5003(buf)

		// Scale the value
		value *= 360
		value /= 255

		// Check the value for valid range
		if value > 360 {
			return fmt.Errorf("value \"%.2f\" outside range [0, 360]", value)
		}

		*d = value
	}

	return nil
}

func (d DPT_5003) Unit() string {
	return "°"
}

func (d DPT_5003) String() string {
	return fmt.Sprintf("%.2f °", float32(d))
}

// DPT_5004 represents DPT 5.004 / Percent_U8.
type DPT_5004 float32

func (d DPT_5004) Pack() []byte {
	value := d

	// Clip the value to a valid range
	if value < 0 {
		value = 0
	}
	if value > 255 {
		value = 255
	}

	return packU8(uint8(value))
}

func (d *DPT_5004) Unpack(data []byte) error {
	var buf uint8

	err := unpackU8(data, &buf)
	if err == nil {
		value := DPT_5004(buf)

		// Check the value for valid range
		if value > 255 {
			return fmt.Errorf("value \"%.2f\" outside range [0, 255]", value)
		}

		*d = value
	}

	return nil
}

func (d DPT_5004) Unit() string {
	return "%"
}

func (d DPT_5004) String() string {
	return fmt.Sprintf("%.2f %%", float32(d))
}

// DPT_9001 represents DPT 9.001 / Temperature.
type DPT_9001 float32

func (d DPT_9001) Pack() []byte {
	value := d

	// Clip the value to a valid range
	if value < -273 {
		value = -273
	}
	if value > 670760 {
		value = 670760
	}

	return packF16(float32(value))
}

func (d *DPT_9001) Unpack(data []byte) error {
	var buf float32

	err := unpackF16(data, &buf)
	if err == nil {
		value := DPT_9001(buf)

		// Check the value for valid range
		if value < -273 {
			return fmt.Errorf("value \"%.2f\" outside range [-273, 670760]", value)
		}
		if value > 670760 {
			return fmt.Errorf("value \"%.2f\" outside range [-273, 670760]", value)
		}

		*d = value
	}

	return nil
}

func (d DPT_9001) Unit() string {
	return "°C"
}

func (d DPT_9001) String() string {
	return fmt.Sprintf("%.2f °C", float32(d))
}

// DPT_9004 represents DPT 9.004 / Illumination.
type DPT_9004 float32

func (d DPT_9004) Pack() []byte {
	value := d

	// Clip the value to a valid range
	if value < 0 {
		value = 0
	}
	if value > 670760 {
		value = 670760
	}

	return packF16(float32(value))
}

func (d *DPT_9004) Unpack(data []byte) error {
	var buf float32

	err := unpackF16(data, &buf)
	if err == nil {
		value := DPT_9004(buf)

		// Check the value for valid range
		if value > 670760 {
			return fmt.Errorf("value \"%.2f\" outside range [0, 670760]", value)
		}

		*d = value
	}

	return nil
}

func (d DPT_9004) Unit() string {
	return "lux"
}

func (d DPT_9004) String() string {
	return fmt.Sprintf("%.2f lux", float32(d))
}
