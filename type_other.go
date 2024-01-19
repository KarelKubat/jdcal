package jdcal

/*
Other returns the "other" calendar type.
*/
func (tp Type) Other() Type {
	if tp == Gregorian {
		return Julian
	}
	return Gregorian
}
