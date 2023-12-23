package jdcal

/*
BeforeOrEqual is true when the date in question occurs earlier than the other date or exactly on that other date. Note that different date types raise an error, see After().
*/
func (d Date) BeforeOrEqual(other Date) (bool, error) {
	bf, err := d.Before(other)
	if err != nil {
		return false, err
	}
	eq, err := d.Equal(other)
	if err != nil {
		return false, err
	}
	return bf || eq, nil
}
