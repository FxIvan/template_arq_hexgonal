package hex_template

const HandlerHandler = `
	package {{.NamePackage}}

	import "{{.ModuleName}}/internal/ports"

	type Handler struct {
		{{.EntityName}}Service ports.{{.EntityName}}Service
}
`
