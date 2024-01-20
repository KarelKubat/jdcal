.PHONY: fullmoons
foo:
	@echo
	@echo 'Make what?' 2>&1
	@echo '  make tables            - create lookup tables (except fullmoons)'
	@echo '  make test              - run tests'
	@echo '  make install           - install jdcal CLI'
	@echo
	@echo '  make fullmoons         - scrape full moons dates'
	@echo '  make bigconversiondata - huge list of J/G dates for testing'
	@echo
	@echo '  make all               - all of the above, except fullmoon re-scraping and big conversion dates'
	@echo
	@exit 1

.PHONY: all
all: tables install test

.PHONY: tables
tables: date_table.go zones_table.go spring-eq-errors.png progressiontable.go

progressiontable.go: main/makeprogressiontable/makeprogressiontable.go main/makeprogressiontable/makeprogressiontable.sh
	sh main/makeprogressiontable/makeprogressiontable.sh

date_table.go: main/makeconversiontable/makeconversiontable.go main/makeconversiontable/date_table.txt Makefile
	go run main/makeconversiontable/makeconversiontable.go main/makeconversiontable/date_table.txt date_table.go
	gofmt -w date_table.go

zones_table.go: main/makezonestable/makezonestable.go main/makezonestable/zones_table.txt Makefile
	go run main/makezonestable/makezonestable.go main/makezonestable/zones_table.txt zones_table.go
	gofmt -w zones_table.go

spring-eq-errors.png: main/daysdiff/spring-equinox.csv Makefile
	gnuplot < main/daysdiff/gnuplot.in

main/daysdiff/spring-equinox.csv: main/daysdiff/daysdiff.go Makefile
	go run main/daysdiff/daysdiff.go > main/daysdiff/spring-equinox.csv.tmp
	echo '0, 0, 0, 0' >> main/daysdiff/spring-equinox.csv.tmp
	sort -n main/daysdiff/spring-equinox.csv.tmp > main/daysdiff/spring-equinox.csv
	rm main/daysdiff/spring-equinox.csv.tmp

.PHONY: fullmoons
fullmoons:
	sh main/makefullmoons/makefullmoons.sh

install: zones_table.go date_table.go Makefile
	go install main/jdcal/jdcal.go

.PHONY: test
test:
	go test ./...
	for f in main/demo*/demo?.go ; do \
		echo ; \
		echo ---------- "$$f" ---------- ; \
		go run "$$f" || exit 1 ; \
	done

.PHONY: bigconversiondata
bigconversiondata:
	sh main/bigconversiontest/make.sh
