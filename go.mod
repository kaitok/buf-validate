module buf-validate-example

go 1.24

toolchain go1.24.3

require (
	buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go v1.36.6-20250425153114-8976f5be98c1.1
	github.com/go-chi/chi/v5 v5.2.1
	google.golang.org/protobuf v1.36.6
)

require github.com/bufbuild/connect-go v1.10.0 // indirect
