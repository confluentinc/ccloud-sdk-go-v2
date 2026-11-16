module github.com/confluentinc/ccloud-sdk-go-v2/flink-gateway

go 1.26.10

require golang.org/x/oauth2 v0.38.0

require (
	github.com/golang/protobuf v1.5.0 // indirect
	golang.org/x/net v0.0.0-20200822124328-c89045814202 // indirect
	google.golang.org/appengine v1.6.6 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
)

// v0.22.0 was tagged on a commit that predates v0.20.0 and v0.21.0 and does not include their changes.
retract v0.22.0
