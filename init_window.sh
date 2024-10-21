#!/bin/bash

set -e  # Detener la ejecución del script si ocurre cualquier error

echo "Ingrese ruta donde se desplegara la arquitectura hexagonal:"
read path

# Verificar si el directorio existe
if [ ! -d "$path" ]; then
    echo "La ruta especificada no existe. Abortando."
    exit 1
fi

cd "$path" || { echo "No se pudo cambiar al directorio $path. Abortando."; exit 1; }

echo "Ingrese nombre de modulo, por ejemplo github.com/tu-usuario/tu-repositorio:"
read module_name
echo "confirmar nombre? yes/no"
read creation_confirmation

# Eliminar archivos si existen
rm -f go.mod go.sum

if [ "$creation_confirmation" = "yes" ]; then
    echo "Creando go module..."
    go mod init "$module_name" || { echo "Error al crear el módulo Go. Abortando."; exit 1; }
else
    echo "Operación cancelada."
    exit 0
fi

cd /home/almendra/hexaGoBuilder || { echo "No se pudo cambiar al directorio /home/almendra/hexaGoBuilder. Abortando."; exit 1; }

echo "Ingresar el nombre de la entidad:"
read nameEntity

# Eliminar y crear directorios con manejo de errores
rm -rf cmd internal
mkdir -p cmd/api/handlers/"$nameEntity" || { echo "Error al crear directorios cmd. Abortando."; exit 1; }
mkdir -p internal/domain internal/ports internal/repositories/mongo/"$nameEntity" internal/services/"$nameEntity" || { echo "Error al crear directorios internal. Abortando."; exit 1; }

cd "$path" || { echo "No se pudo cambiar al directorio $path. Abortando."; exit 1; }

# Instalar dependencias con manejo de errores
go get github.com/joho/godotenv || { echo "Error al instalar godotenv. Abortando."; exit 1; }
go get -u github.com/gin-gonic/gin || { echo "Error al instalar gin. Abortando."; exit 1; }
go get go.mongodb.org/mongo-driver/mongo || { echo "Error al instalar mongo-driver. Abortando."; exit 1; }

cd /home/almendra/hexaGoBuilder || { echo "No se pudo cambiar al directorio /home/almendra/hexaGoBuilder. Abortando."; exit 1; }

# Ejecutar el archivo Go con manejo de errores
go run init.go || { echo "Error al ejecutar init.go. Abortando."; exit 1; }

# Mover directorios al path especificado con manejo de errores
mv cmd "$path" || { echo "Error al mover el directorio cmd. Abortando."; exit 1; }
mv internal "$path" || { echo "Error al mover el directorio internal. Abortando."; exit 1; }

echo "Script ejecutado exitosamente."
