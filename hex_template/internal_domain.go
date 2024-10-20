package hex_template

const InternalDomain = `
	package domain

	import "time"

	type {{.EntityName}} struct {
		Name         string    ` + "`json:\"name\" binding:\"required\"`" + `
		Age          int       ` + "`json:\"age\" binding:\"required\"`" + `
		CreationTime time.Time ` + "`json:\"creation_time\"`" + `
	}
`
