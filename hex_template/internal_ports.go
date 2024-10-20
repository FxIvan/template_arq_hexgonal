package hex_template

const InternalPorts = `
	package ports

	import "{{.ModuleName}}/internal/domain"

	type {{.EntityName}}Service interface {
		Create({{.EntityVariableLowerCase}} domain.{{.EntityName}}) (id interface{}, err error)
	}

	type {{.EntityName}}Repository interface {
		Insert({{.EntityVariableLowerCase}} domain.{{.EntityName}}) (id interface{}, err error)
	}
`
