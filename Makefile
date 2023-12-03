table.go: main/maketable/maketable.go date_table.txt
	go run main/maketable/maketable.go date_table.txt date_table.go
	gofmt -w date_table.go

