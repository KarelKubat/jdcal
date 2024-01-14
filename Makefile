foo: date_table.go zones_table.go spring-eq-errors.png

date_table.go: main/makeconversiontable/makeconversiontable.go date_table.txt Makefile
	go run main/makeconversiontable/makeconversiontable.go date_table.txt date_table.go
	gofmt -w date_table.go

zones_table.go: main/makezonestable/makezonestable.go zones_table.txt Makefile
	go run main/makezonestable/makezonestable.go zones_table.txt zones_table.go
	gofmt -w zones_table.go

spring-eq-errors.png: main/daysdiff/spring-equinox.csv Makefile
	gnuplot < main/daysdiff/gnuplot.in

main/daysdiff/spring-equinox.csv: main/daysdiff/daysdiff.go Makefile
	go run main/daysdiff/daysdiff.go > main/daysdiff/spring-equinox.csv.tmp
	echo '0, 0, 0, 0' >> main/daysdiff/spring-equinox.csv.tmp
	sort -n main/daysdiff/spring-equinox.csv.tmp > main/daysdiff/spring-equinox.csv
	rm main/daysdiff/spring-equinox.csv.tmp

.PHONY: fullmoons.go
fullmoons.go:
	go run main/makefullmoons/makefullmoons.go fullmoons.out
	mv fullmoons.out fullmoons.go
	gofmt -w fullmoons.go

install: zones_table.go date_table.go Makefile
	go install main/jdcal/jdcal.go

.PHONY: test
test:
	go test ./...
	go run main/demo1/demo1.go
	go run main/demo2/demo2.go
	go run main/demo3/demo3.go
	go run main/demo4/demo4.go
