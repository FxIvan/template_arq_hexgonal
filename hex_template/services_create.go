package hex_template

const ServicesCreate = `
	package player

	import (
		"time"

		"{{.ModuleName}}/internal/domain"
	)

	func (s Service) Create({{.EntityVariableLowerCase}} domain.{{.EntityName}}) (id interface{}, err error) {
		{{.EntityVariableLowerCase}}.CreationTime = time.Now().UTC()

		insertResult, err := s.Repo.Insert({{.EntityVariableLowerCase}})
		if err != nil {
			return nil, err
		}

		return insertResult, nil
	}
`
