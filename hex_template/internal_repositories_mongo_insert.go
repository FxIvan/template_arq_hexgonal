package hex_template

const InternalRepositoriesMongoInsert = `
	package {{.NamePackage}}

	import (
		"context"
		"log"
		"os"

		"{{.ModuleName}}/internal/domain"
	)

	func (r Repository) Insert({{.EntityVariableLowerCase}} domain.{{.EntityName}}) (id interface{}, err error) {

		collection := r.Client.Database(os.Getenv("MONGO_NAME_DB")).Collection(os.Getenv("MONGO_NAME_COLLECTION"))
		insertResult, err := collection.InsertOne(context.Background(), {{.EntityVariableLowerCase}})
		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		return insertResult.InsertedID, nil
	}
`
