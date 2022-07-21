module github.com/th3g3ntl3m3n/go-micro/user-service

go 1.18

require (
	github.com/golang/protobuf v1.5.2
	github.com/micro/micro/v3 v3.11.0
	github.com/th3g3ntl3m3n/go-micro/email-service v0.0.0-20220721184533-8fb5dcd62667
	google.golang.org/protobuf v1.28.0
)

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.27.1
