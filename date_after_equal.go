package jdcal

/*
AfterOrEqual is true when the date in question occurs later than the other date.
Comparing dates of different types raises an error, see Equal().
*/
func (d Date) AfterOrEqual(other Date) (bool, error) {
	af, err := d.After(other)
	if err != nil {
		return false, err
	}
	eq, err := d.Equal(other)
	if err != nil {
		return false, err
	}
	return af || eq, nil
}
