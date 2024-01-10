foo: date_table.go zones_table.go main/daysdiff/spring-equinox.csv

date_table.go: main/makeconversiontable/makeconversiontable.go date_table.txt
	go run main/makeconversiontable/makeconversiontable.go date_table.txt date_table.go
	gofmt -w date_table.go

zones_table.go: main/makezonestable/makezonestable.go zones_table.txt
	go run main/makezonestable/makezonestable.go zones_table.txt zones_table.go
	gofmt -w zones_table.go

main/daysdiff/spring-equinox.csv: main/daysdiff/daysdiff.go
	go run main/daysdiff/daysdiff.go > main/daysdiff/spring-equinox.csv

install: zones_table.go date_table.go
	go install main/jdcal/jdcal.go
