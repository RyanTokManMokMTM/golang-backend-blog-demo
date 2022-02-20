// Package convert - Convert the string type to Int
package convert

import "strconv"

type StrTo string

//String
func (str StrTo) String() string {
	return string(str)
}

//Int Convert string to int with error
func (str StrTo) Int() (int, error) {
	num, err := strconv.Atoi(str.String())
	return num, err
}

//MustInt Convert string to int without error(must be an int)
func (str StrTo) MustInt() int {
	v, _ := str.Int()
	return v
}

//UInt32 Convert string to uint32 with error
func (str StrTo) UInt32() (uint32, error) {
	v, err := strconv.Atoi(str.String())
	return uint32(v), err
}

//MustUInt32 Convert string to uint32 without error(must be an uint32)
func (str StrTo) MustUInt32() uint32 {
	v, _ := str.UInt32()
	return v
}
