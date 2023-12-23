package jdcal

/*
IsSet returns true when any of the fields Year, Month or Day of a given date have a non-zero value. IsSet is false when a date is not initialized. Example:

	d := Date{}
	fmt.Println(d.IsSet())  // false

	d.Day = 1
	fmt.Println(d.IsSet())  // true
*/
func (d Date) IsSet() bool {
	return d.Year != 0 || d.Month != 0 || d.Day != 0
}
