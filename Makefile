table.go: main/makeconversiontable/makeconversiontable.go date_table.txt
	go run main/makeconversiontable/makeconversiontable.go date_table.txt date_table.go
	gofmt -w date_table.go
