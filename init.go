package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	template_library "text/template"

	"github.com/fxivan/arq_hexagonal/hex_template"
)

type TemplateRoute struct {
	NameTemplate string
	Route        string
	NameFile     string
}

func firstLetterUpper(entity string) string {
	return strings.ToUpper(entity[:1]) + entity[1:]
}

func firstLetterLower(entity string) string {
	return strings.ToLower(entity[:1]) + entity[1:]
}

func main() {
	mt := hex_template.MainTemplate
	hc := hex_template.HandlerCreate
	hh := hex_template.HandlerHandler
	id := hex_template.InternalDomain
	ip := hex_template.InternalPorts
	mc := hex_template.InternalRepositoriesMongoConnectClient
	mi := hex_template.InternalRepositoriesMongoInsert
	mr := hex_template.InternalRepositoriesMongoRepository
	sc := hex_template.ServicesCreate
	ss := hex_template.ServiceService

	type Recipient struct {
		ModuleName              string
		Entity                  string
		EntityVariableLowerCase string
		EntityEndpointPOST      string
		PortService             int
		NamePackage             string
		EntityName              string
	}

	var lowerCaseEntity, upperCaseEntity, moduleName_scanf, entity_cmd string
	var port_scanf int

	fmt.Println("Enter the module name:")
	fmt.Scanf("%s", &moduleName_scanf)

	fmt.Println("Ingresa el nombre de la entidad:")
	fmt.Scanf("%s", &entity_cmd)
	lowerCaseEntity = firstLetterLower(entity_cmd)
	upperCaseEntity = firstLetterUpper(entity_cmd)

	fmt.Println("Enter the port:")
	fmt.Scanf("%d", &port_scanf)

	var recipients = []Recipient{
		{
			moduleName_scanf,
			upperCaseEntity,
			lowerCaseEntity,
			lowerCaseEntity,
			port_scanf,
			lowerCaseEntity,
			upperCaseEntity,
		},
	}

	templateWithRoutes := []TemplateRoute{
		{
			NameTemplate: mt,
			Route:        "cmd/api",
			NameFile:     "main",
		},
		{
			NameTemplate: hc,
			Route:        "cmd/api/handlers/" + upperCaseEntity,
			NameFile:     "create",
		},
		{
			NameTemplate: hh,
			Route:        "cmd/api/handlers/" + upperCaseEntity,
			NameFile:     "handler",
		},
		{
			NameTemplate: id,
			Route:        "internal/domain",
			NameFile:     upperCaseEntity,
		},
		{
			NameTemplate: ip,
			Route:        "internal/ports",
			NameFile:     upperCaseEntity,
		},
		{
			NameTemplate: mc,
			Route:        "internal/repositories/mongo",
			NameFile:     "connect_client",
		},
		{
			NameTemplate: mi,
			Route:        "internal/repositories/mongo/" + upperCaseEntity,
			NameFile:     "insert",
		},
		{
			NameTemplate: mr,
			Route:        "internal/repositories/mongo/" + upperCaseEntity,
			NameFile:     "repository",
		},
		{
			NameTemplate: sc,
			Route:        "internal/services/" + upperCaseEntity,
			NameFile:     "create",
		},
		{
			NameTemplate: ss,
			Route:        "internal/services/" + upperCaseEntity,
			NameFile:     "service",
		},
	}

	for _, template := range templateWithRoutes {
		tmpl, err := template_library.New("template").Parse(template.NameTemplate)
		if err != nil {
			log.Fatalf("parsing: %s", err)
		}

		f, err := os.Create(template.Route + "/" + template.NameFile + ".go")
		if err != nil {
			log.Fatalf("creating file: %s", err)
		}

		err = tmpl.Execute(f, recipients[0])
		if err != nil {
			log.Fatalf("execution: %s", err)
		}

		f.Close()

		log.Printf("Template processed successfully. Output written to %s/%s.go", template.Route, template.NameFile)
	}
}
