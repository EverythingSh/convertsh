con: 
	CGO_LDFLAGS="$(shell pkg-config --libs gio-2.0)" go build -o build/con cmd/main.go

clean:
	rm -f build/con
	rm -rf pkg/image/output_*

.PHONY: con clean