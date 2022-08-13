package mock;

// # go-generate directives
// Just run `go generate ./...` in the ksql directory to generate the mocks
// See `go help generate` for more details

//go:generate mocker --prefix "" --dst api_clusters_ksqldbcm_v2__ClustersKsqldbcmV2Api.go --pkg mock ../v2/api_clusters_ksqldbcm_v2.go ClustersKsqldbcmV2Api
