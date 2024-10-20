package hex_template

// connect_client.go

const InternalRepositoriesMongoConnectClient = `
	package mongo

	import (
		"context"
		"time"

		"go.mongodb.org/mongo-driver/mongo"
		"go.mongodb.org/mongo-driver/mongo/options"
	)

	func ConnectClient(DB_URI string) (client *mongo.Client, err error) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		client, err = mongo.Connect(ctx, options.Client().ApplyURI(DB_URI))
		if err != nil {
			return nil, err
		}

		err = client.Ping(ctx, nil)
		if err != nil {
			return nil, err
		}

		return client, nil
	}
`
