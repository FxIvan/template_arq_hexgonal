package hex_template

const MainTemplate = `
	package main

	import(
		"log"
		"os"
		"github.com/gin-gonic/gin"
		"github.com/joho/godotenv"

		"{{.ModuleName}}/cmd/api/handlers/{{.EntityVariableLowerCase}}"
		"{{.ModuleName}}/internal/repositories/mongo"
		{{.EntityVariableLowerCase}}Mongo "{{.ModuleName}}/internal/repositories/mongo/{{.EntityVariableLowerCase}}"
		{{.EntityVariableLowerCase}}Service "{{.ModuleName}}/internal/services/{{.EntityVariableLowerCase}}"
	)

	func main() {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		ginEngine := gin.Default()

		client, err := mongo.ConnectClient(os.Getenv("MONGO_URI"))
		if err != nil {
			log.Fatal(err.Error())
		}

		{{.EntityVariableLowerCase}}Repo := {{.EntityVariableLowerCase}}Mongo.Repository{
			Client: client,
		}

		{{.EntityVariableLowerCase}}Srv := {{.EntityVariableLowerCase}}Service.Service{
			Repo: {{.EntityVariableLowerCase}}Repo,
		}

		{{.EntityVariableLowerCase}}Handler := {{.EntityVariableLowerCase}}.Handler{
			{{.Entity}}Service: {{.EntityVariableLowerCase}}Srv,
		}

		ginEngine.POST("/{{.EntityEndpointPOST}}", {{.EntityVariableLowerCase}}Handler.Create{{.Entity}})

		log.Fatal(ginEngine.Run(":{{.PortService}}"))
	}
`
