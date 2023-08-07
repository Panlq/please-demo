//go:build tools
// +build tools

package tools

// see https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module
import (
	_ "github.com/envoyproxy/protoc-gen-validate"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2"
	_ "github.com/segmentio/golines"
	_ "github.com/twitchtv/twirp/protoc-gen-twirp"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)

//go:generate go install github.com/envoyproxy/protoc-gen-validate
//go:generate go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
//go:generate go install github.com/twitchtv/twirp/protoc-gen-twirp
//go:generate go install google.golang.org/protobuf/cmd/protoc-gen-go
//go:generate go install github.com/segmentio/golines
