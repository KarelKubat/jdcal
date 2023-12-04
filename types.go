package jdcal

import "time"

type Type int

const (
	Gregorian Type = iota
	Julian
)

type Date struct {
	Year  int
	Month time.Month
	Day   int
	Type  Type
}

type ConversionEntry struct {
	JDate, GDate Date
}
