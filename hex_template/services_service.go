package hex_template

const ServiceService = `
package {{.NamePackage}}

import (
	"{{.ModuleName}}/internal/ports"
)

type Service struct {
	Repo ports.{{.EntityName}}Repository
}

`
