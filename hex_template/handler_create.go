package hex_template

/*
.NamePackage
.ModuleName
.EntityVariableLowerCase
.EntityName

*/

const HandlerCreate = `
	package {{.NamePackage}}

	import (
		"net/http"

		"github.com/gin-gonic/gin"
		"{{.ModuleName}}/internal/domain"
	)

	func (h Handler) Create{{.EntityName}}(c *gin.Context) {
		var {{.EntityVariableLowerCase}} domain.{{.EntityName}}
		if err := c.BindJSON(&{{.EntityVariableLowerCase}}); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		insertedId, err := h.{{.EntityName}}Service.Create({{.EntityVariableLowerCase}})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "oops!"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"{{.EntityVariableLowerCase}}_id": insertedId})
	}
`
