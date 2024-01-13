package storage

// DefaultPageSize represents the default page size for query results.
const DefaultPageSize = 5

// Parameters represents parametrisation of the queries.
type Parameters struct {
	Where    []string
	Page     uint64
	PageSize uint64
}
