package jdcal

func New(dt Date) (Date, error) {
	return dt, dt.Valid()
}
