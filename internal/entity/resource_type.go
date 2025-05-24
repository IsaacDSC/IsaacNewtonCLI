package entity

type ResourceType string

const (
	ResourceTypeGraphQLAPI ResourceType = "graphql_api"
	ResourceTypeWebApp     ResourceType = "web_app"
	ResourceTypeJob        ResourceType = "job"
	ResourceTypeQueue      ResourceType = "queue"
	ResourceSqlDb          ResourceType = "sql_db"
	ResourceTypeNoSqlDb    ResourceType = "nosql_db"
	ResourceTypeCache      ResourceType = "cache"
	ResourceTypeStorage    ResourceType = "storage"
)

func (r ResourceType) String() string {
	return string(r)
}
