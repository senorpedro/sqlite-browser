run-dummy:
	go run . --db test.db

run-debug: 
	dlv debug -- --db test.db

