package repository

type ElasticsearchRepo interface {
	ListIndexNames() (result []string)
}
