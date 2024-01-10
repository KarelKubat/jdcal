package jdcal

// Generated file, don't edit by hand. Look for a nearby Makefile.

import (
	"time"
)

/*
ZonesTable defines matching geographical zones to the date of adoption of the
Gregorian calendar. Some zones temporarily reverted to the Julian calendar, these
have multiple cutover dates.

This table reflects https://en.wikipedia.org/wiki/List_of_adoption_dates_of_the_Gregorian_calendar_by_country.

Note: The Cutovers entries are the dates where that period STOPPED. So: "{Year: 1918, Month: time.April, Day: 17, Type: Julian}" is the Julian date 1918/04/16, when the Julian calendar was abandoned in favor of "the other one", being the Gregorian.

Another note: All symbols of the table are exported (upper case). You can skip reading the table; rather, have a look at the "Zones*()" functions that can do the lifting.
*/
var ZonesTable = [...]ZoneEntry{
	// [...] is syntactic sugar to let the compiler figure out the array size. That way
	// we get a fixed size array and not a slice.
	{
		Name: "Albania (non-Catholic)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1912, Month: time.November, Day: 14, Type: Julian},
		},
	},
	{
		Name: "Albania (Catholic)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1583, Month: time.October, Day: 5, Type: Julian},
		},
	},
	{
		Name: "Armenia",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1918, Month: time.April, Day: 17, Type: Julian},
		},
	},
	{
		Name: "Austria (Brixen, Salzburg, Tyrol)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1583, Month: time.October, Day: 5, Type: Julian},
		},
	},
	{
		Name: "Austria (Carinthia)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1584, Month: time.January, Day: 5, Type: Julian},
		},
	},
	{
		Name: "Azerbaijan",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1918, Month: time.April, Day: 17, Type: Julian},
		},
	},
	{
		Name: "Belarus (Lithuania Governorate)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1585, Month: time.December, Day: 25, Type: Julian},
			{Year: 1800, Month: time.January, Day: 11, Type: Gregorian},
			{Year: 1915, Month: time.August, Day: 22, Type: Julian},
		},
	},
	{
		Name: "Belarus (Soviet Russia)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1918, Month: time.January, Day: 31, Type: Julian},
		},
	},
	{
		Name: "Belgium (Flanders)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1582, Month: time.December, Day: 14, Type: Julian},
		},
	},
	{
		Name: "Belgium (Liège)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1583, Month: time.February, Day: 10, Type: Julian},
		},
	},
	{
		Name: "Belgium (Southern Netherlands)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1582, Month: time.December, Day: 20, Type: Julian},
		},
	},
	{
		Name: "Bulgaria",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1916, Month: time.March, Day: 31, Type: Julian},
		},
	},
	{
		Name: "Cambodia",
		Cutovers: []Date{
			{Year: 1863, Month: time.January, Day: 1, Type: Julian},
		},
	},
	{
		Name: "Canada (French colonial empire)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1582, Month: time.December, Day: 9, Type: Julian},
		},
	},
	{
		Name: "Canada (Nova Scotia)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1582, Month: time.December, Day: 9, Type: Julian},
			{Year: 1710, Month: time.October, Day: 3, Type: Gregorian},
			{Year: 1752, Month: time.September, Day: 2, Type: Julian},
		},
	},
	{
		Name: "Canada (Britsh empire)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1752, Month: time.September, Day: 2, Type: Julian},
		},
	},
	{
		Name: "China",
		Cutovers: []Date{
			{Year: 1911, Month: time.December, Day: 1, Type: Julian},
		},
	},
	{
		Name: "Czech Republic (Bohemia)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1584, Month: time.January, Day: 6, Type: Julian},
		},
	},
	{
		Name: "Czech Republic (Moravia)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1584, Month: time.October, Day: 3, Type: Julian},
		},
	},
	{
		Name: "Czech Republic (Silesia)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1584, Month: time.January, Day: 12, Type: Julian},
		},
	},
	{
		Name: "Denmark",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1700, Month: time.February, Day: 18, Type: Julian},
		},
	},
	{
		Name: "Egypt",
		Cutovers: []Date{
			{Year: 1875, Month: time.June, Day: 11, Type: Julian},
		},
	},
	{
		Name: "Estonia",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1918, Month: time.February, Day: 15, Type: Julian},
		},
	},
	{
		Name: "Faroe Islands",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1700, Month: time.November, Day: 16, Type: Julian},
		},
	},
	{
		Name: "France (largest part, note: different calendar 1793-1805)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1582, Month: time.December, Day: 9, Type: Julian},
		},
	},
	{
		Name: "France (Sedan)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1582, Month: time.December, Day: 9, Type: Julian},
		},
	},
	{
		Name: "France (Austrian Upper Alsace, Breisgau)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1583, Month: time.October, Day: 13, Type: Julian},
		},
	},
	{
		Name: "France (Catholic Strasbourg)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1583, Month: time.November, Day: 16, Type: Julian},
		},
	},
	{
		Name: "France (Alsace, Protestant parts)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1682, Month: time.February, Day: 5, Type: Julian},
		},
	},
	{
		Name: "France (Lorraine)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1582, Month: time.December, Day: 9, Type: Julian},
			{Year: 1735, Month: time.January, Day: 1, Type: Gregorian},
			{Year: 1760, Month: time.February, Day: 16, Type: Julian},
		},
	},
	{
		Name: "France (Mulhouse)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1700, Month: time.December, Day: 31, Type: Julian},
		},
	},
	{
		Name: "Germany (Aachen)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1582, Month: time.December, Day: 31, Type: Julian},
		},
	},
	{
		Name: "Germany (Augsburg)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1583, Month: time.February, Day: 13, Type: Julian},
		},
	},
	{
		Name: "Germany (Baden-Baden)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1583, Month: time.November, Day: 16, Type: Julian},
		},
	},
	{
		Name: "Germany (Bavaria, Regensburg, Freising)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1583, Month: time.October, Day: 5, Type: Julian},
		},
	},
	{
		Name: "Germany (Breisgau-Hochschwarzwald)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1583, Month: time.October, Day: 5, Type: Julian},
		},
	},
	{
		Name: "Germany (Cologne)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1583, Month: time.November, Day: 3, Type: Julian},
		},
	},
	{
		Name: "Germany (Jülich-Berg)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1583, Month: time.November, Day: 2, Type: Julian},
		},
	},
	{
		Name: "Germany (Mainz)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1583, Month: time.November, Day: 11, Type: Julian},
		},
	},
	{
		Name: "Germany (Münster, Duchy of Cleves)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1583, Month: time.November, Day: 17, Type: Julian},
		},
	},
	{
		Name: "Germany (Osnabrück)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1624, Month: time.January, Day: 1, Type: Julian},
		},
	},
	{
		Name: "Germany (Paderborn)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1585, Month: time.June, Day: 16, Type: Julian},
		},
	},
	{
		Name: "Germany (Pfalz-Neuburg)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1585, Month: time.December, Day: 13, Type: Julian},
		},
	},
	{
		Name: "Germany (Silesia)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1584, Month: time.January, Day: 12, Type: Julian},
		},
	},
	{
		Name: "Germany (Trier)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1583, Month: time.October, Day: 4, Type: Julian},
		},
	},
	{
		Name: "Germany (Westfalia)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1584, Month: time.June, Day: 1, Type: Julian},
		},
	},
	{
		Name: "Germany (Würzburg)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1583, Month: time.November, Day: 4, Type: Julian},
		},
	},
	{
		Name: "Germany (Prince-Bishopric of Hildesheim)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1631, Month: time.March, Day: 15, Type: Julian},
		},
	},
	{
		Name: "Germany (Bishopric of Minden)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1668, Month: time.February, Day: 1, Type: Julian},
		},
	},
	{
		Name: "Germany (Protestant parts)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1700, Month: time.February, Day: 18, Type: Julian},
		},
	},
	{
		Name: "Georgia",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1918, Month: time.April, Day: 17, Type: Julian},
		},
	},
	{
		Name: "Greece (excl. Mt. Athos)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1923, Month: time.February, Day: 15, Type: Julian},
		},
	},
	{
		Name: "Hungary",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1587, Month: time.October, Day: 21, Type: Julian},
		},
	},
	{
		Name: "Iceland",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1700, Month: time.November, Day: 16, Type: Julian},
		},
	},
	{
		Name: "Ireland (O'Neill and O'Donnell Gaelic Lordships in Ulster)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1584, Month: time.January, Day: 1, Type: Julian},
			{Year: 1630, Month: time.January, Day: 1, Type: Gregorian},
			{Year: 1641, Month: time.January, Day: 1, Type: Julian},
		},
	},
	{
		Name: "Ireland",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1752, Month: time.September, Day: 2, Type: Julian},
		},
	},
	{
		Name: "Italy",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1582, Month: time.October, Day: 4, Type: Julian},
		},
	},
	{
		Name: "Italy (Tirol)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1583, Month: time.October, Day: 5, Type: Julian},
		},
	},
	{
		Name: "Japan",
		Cutovers: []Date{
			{Year: 1872, Month: time.December, Day: 2, Type: Julian},
		},
	},
	{
		Name: "Laos",
		Cutovers: []Date{
			{Year: 1889, Month: time.January, Day: 1, Type: Julian},
		},
	},
	{
		Name: "Latvia (Courland)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1617, Month: time.January, Day: 1, Type: Julian},
			{Year: 1796, Month: time.January, Day: 28, Type: Gregorian},
			{Year: 1915, Month: time.May, Day: 11, Type: Julian},
		},
	},
	{
		Name: "Lithuania (Grand Dutchy of Litnuania)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1585, Month: time.December, Day: 21, Type: Julian},
		},
	},
	{
		Name: "Latvia (Livland)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1915, Month: time.August, Day: 22, Type: Julian},
		},
	},
	{
		Name: "Lithuania (Lithuania Governorate)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1585, Month: time.December, Day: 25, Type: Julian},
			{Year: 1800, Month: time.January, Day: 11, Type: Gregorian},
			{Year: 1915, Month: time.August, Day: 22, Type: Julian},
		},
	},
	{
		Name: "Lithuania (Kovno and Vilna Governorates)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1915, Month: time.May, Day: 11, Type: Julian},
		},
	},
	{
		Name: "Lithuania (Dutchy of Prussia)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1612, Month: time.May, Day: 11, Type: Julian},
		},
	},
	{
		Name: "Luxembourg",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1582, Month: time.December, Day: 20, Type: Julian},
		},
	},
	{
		Name: "Montenegro",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1919, Month: time.January, Day: 14, Type: Julian},
		},
	},
	{
		Name: "Myanmar (Burma, British Empire)",
		Cutovers: []Date{
			{Year: 1885, Month: time.January, Day: 1, Type: Julian},
		},
	},
	{
		Name: "Netherlands (Brabant)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1582, Month: time.December, Day: 14, Type: Julian},
		},
	},
	{
		Name: "Netherlands (Drenthe)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1701, Month: time.April, Day: 30, Type: Julian},
		},
	},
	{
		Name: "Netherlands (Frisia)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1701, Month: time.December, Day: 31, Type: Julian},
		},
	},
	{
		Name: "Netherlands (Gelderland)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1700, Month: time.June, Day: 12, Type: Julian},
		},
	},
	{
		Name: "Netherlands (Groningen City)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1583, Month: time.January, Day: 1, Type: Julian},
			{Year: 1594, Month: time.November, Day: 10, Type: Gregorian},
			{Year: 1700, Month: time.December, Day: 31, Type: Julian},
		},
	},
	{
		Name: "Groningen (Province)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1700, Month: time.December, Day: 31, Type: Julian},
		},
	},
	{
		Name: "Netherlands (Holland)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1583, Month: time.January, Day: 1, Type: Julian},
		},
	},
	{
		Name: "Netherlands (Utrecht, Overijssel)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1700, Month: time.November, Day: 30, Type: Julian},
		},
	},
	{
		Name: "Netherlands (Zeeland, States General)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1582, Month: time.December, Day: 14, Type: Julian},
		},
	},
	{
		Name: "North Korea",
		Cutovers: []Date{
			{Year: 1896, Month: time.January, Day: 1, Type: Julian},
		},
	},
	{
		Name: "North Macedonia",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1919, Month: time.January, Day: 14, Type: Julian},
		},
	},
	{
		Name: "Norway",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1700, Month: time.February, Day: 18, Type: Julian},
		},
	},
	{
		Name: "Poland (excl. Dutchy of Prussia, Silesia)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1582, Month: time.October, Day: 4, Type: Julian},
		},
	},
	{
		Name: "Poland (Dutchy of Prussia)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1612, Month: time.August, Day: 22, Type: Julian},
		},
	},
	{
		Name: "Poland (Silesia)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1584, Month: time.January, Day: 12, Type: Julian},
		},
	},
	{
		Name: "Portugal",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1582, Month: time.October, Day: 4, Type: Julian},
		},
	},
	{
		Name: "Romania (Transylvania)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1590, Month: time.December, Day: 14, Type: Julian},
		},
	},
	{
		Name: "Romania (excl. Transylvania)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1919, Month: time.March, Day: 29, Type: Julian},
		},
	},
	{
		Name: "Russia (Dutchy of Prussia)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1612, Month: time.August, Day: 22, Type: Julian},
		},
	},
	{
		Name: "Russia (excl. Dutchy of Prussia)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1918, Month: time.January, Day: 31, Type: Julian},
		},
	},
	{
		Name: "Saudi Arabia",
		Cutovers: []Date{
			{Year: 2016, Month: time.October, Day: 1, Type: Julian},
		},
	},
	{
		Name: "Serbia",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1919, Month: time.January, Day: 14, Type: Julian},
		},
	},
	{
		Name: "Slovenia",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1583, Month: time.December, Day: 14, Type: Julian},
		},
	},
	{
		Name: "Spain",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1582, Month: time.October, Day: 4, Type: Julian},
		},
	},
	{
		Name: "South Korea",
		Cutovers: []Date{
			{Year: 1896, Month: time.January, Day: 1, Type: Julian},
		},
	},
	{
		Name: "Sweden",
		Cutovers: []Date{
			{Year: 1712, Month: time.February, Day: 29, Type: Gregorian},
			{Year: 1753, Month: time.February, Day: 17, Type: Julian},
		},
	},
	{
		Name: "Switzerland (Basel, Roman Catholic Diocese)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1583, Month: time.October, Day: 20, Type: Julian},
		},
	},
	{
		Name: "Switzerland (Luzern, Uri, Schwyz, Zug, Freiburg, Solothurn, Le Landeron)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1584, Month: time.January, Day: 11, Type: Julian},
		},
	},
	{
		Name: "Switzerland (Obwalden, Nidwalden)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1584, Month: time.February, Day: 11, Type: Julian},
		},
	},
	{
		Name: "Switzerland (Thurgau, Appenzell Innerrhoden)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1584, Month: time.January, Day: 1, Type: Julian},
		},
	},
	{
		Name: "Switzerland (Appenzell Ausserrhoden)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1584, Month: time.January, Day: 1, Type: Julian},
			{Year: 1597, Month: time.January, Day: 1, Type: Gregorian},
			{Year: 1798, Month: time.December, Day: 25, Type: Julian},
		},
	},
	{
		Name: "Switzerland (Lower Valais)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1623, Month: time.January, Day: 1, Type: Julian},
		},
	},
	{
		Name: "Switzerland (Valais)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1655, Month: time.February, Day: 28, Type: Julian},
		},
	},
	{
		Name: "Switzerland (Protestant parts of Basel, Bern, Neuchâtel, Sargans, Schaffhausen, Geneva, Zürich, Glarus)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1700, Month: time.December, Day: 31, Type: Julian},
		},
	},
	{
		Name: "Switzerland (Graubünden, Catholic parts)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1623, Month: time.January, Day: 1, Type: Julian},
		},
	},
	{
		Name: "Switzerland (Graubünden, Oberengadin and Bergel)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1783, Month: time.January, Day: 1, Type: Julian},
		},
	},
	{
		Name: "Switzerland (Graubünden, Schiers and Grüsh)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1811, Month: time.December, Day: 25, Type: Julian},
		},
	},
	{
		Name: "Thailand (Siam)",
		Cutovers: []Date{
			{Year: 1889, Month: time.January, Day: 1, Type: Julian},
		},
	},
	{
		Name: "Turkey",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1917, Month: time.February, Day: 15, Type: Julian},
		},
	},
	{
		Name: "Uganda (British Empire)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1893, Month: time.December, Day: 31, Type: Julian},
		},
	},
	{
		Name: "Ukraine",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1918, Month: time.February, Day: 15, Type: Julian},
		},
	},
	{
		Name: "United Kingdom (Kingdom of Great Britain, Kingdom of Ireland)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1752, Month: time.September, Day: 2, Type: Julian},
		},
	},
	{
		Name: "United States of America (French & Spanish colonial empires)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1582, Month: time.December, Day: 9, Type: Julian},
		},
	},
	{
		Name: "United States of America (British Empire)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1752, Month: time.September, Day: 2, Type: Julian},
		},
	},
	{
		Name: "United States of America (Russian Empire: Alaska)",
		Cutovers: []Date{
			First(Gregorian),
			{Year: 1867, Month: time.October, Day: 6, Type: Julian},
		},
	},
}
