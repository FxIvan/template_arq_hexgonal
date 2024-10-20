#!/bin/bash

echo "Ingrese nombre de modulo, por ejemplo github.com/tu-usuario/tu-repositorio"
read module_name
echo "confirmar nombre? yes/no"
read creation_confirmation

rm -rf go.mod go.sum

if [ "$creation_confirmation" = "yes" ]
then
    echo "creando go module..."
    go mod init $module_name
else
    echo "Operación cancelada."
fi

# Función para crear directorios
echo "Ingresar el nombre de la entidad"
read nameEntity
rm -rf cmd internal
mkdir cmd cmd/api cmd/api/handlers cmd/api/handlers/$nameEntity
mkdir internal internal/domain internal/ports internal/repositories internal/repositories/mongo internal/repositories/mongo/$nameEntity internal/services internal/services/$nameEntity
###########################
#- cmd
#    - api
#        - main.go
#    - cli
#    - worker
#    - migration
#    - handlers
#        - $nameEntity
#           - handler_create.go   
#           - handler_handler.go
#
###########################
# cp hex_template/main.go cmd/api/
# cp hex_template/handler_create.go cmd/handlers/$nameEntity/
# cp hex_template/handler_handler.go cmd/handlers/$nameEntity/
###########################
#- internal
#    - domain
#        - internal_domain.go
#    - ports
#        - internal_ports.go
#    - repositories
#        - mongo
#           - $nameEntity
#               - internal_repositories_mongo_insert.go
#               - internal_repositories_mongo_repository.go
#           - internal_repositories_mongo_connect_client.go
#    - services
#       - services_create.go
#       - services_service.go
###########################

# cp hex_template/internal_domain.go internal/domain/
# cp hex_template/internal_ports.go internal/ports/
# cp hex_template/internal_repositories_mongo_insert.go internal/repositories/mongo/$nameEntity/
# cp hex_template/internal_repositories_mongo_repository.go internal/repositories/mongo/$nameEntity/
# cp hex_template/internal_repositories_mongo_connect_client.go internal/repositories/mongo/
# cp hex_template/services_create.go internal/services/$nameEntity/
# cp hex_template/services_service.go internal/services/$nameEntity/

go get github.com/joho/godotenv
go get -u github.com/gin-gonic/gin
go get go.mongodb.org/mongo-driver/mongo

go run init.go
