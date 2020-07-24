build-windows:
	go build -ldflags="-H windowsgui" -o build/dbt.exe

run-windows:
	$(MAKE) build-windows
	./build/dbt.exe