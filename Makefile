gen:
	for f in *; do \
		if [ -d "$$f" ]; then \
			protoc --go_out=. --go_opt=paths=source_relative \
				--go-grpc_out=. --go-grpc_opt=paths=source_relative \
				$$f/$$f.proto ; \
		fi \
	done
