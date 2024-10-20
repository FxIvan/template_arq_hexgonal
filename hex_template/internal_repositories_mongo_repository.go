package hex_template

const InternalRepositoriesMongoRepository = `
	package {{.NamePackage}}

	import "go.mongodb.org/mongo-driver/mongo"

	type Repository struct {
		Client *mongo.Client
}
`
