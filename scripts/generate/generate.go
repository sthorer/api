package main

import (
	"log"

	"github.com/facebookincubator/ent/entc"
	"github.com/facebookincubator/ent/entc/gen"
	"github.com/facebookincubator/ent/schema/field"
)

func main() {
	err := entc.Generate("./ent/schema", &gen.Config{
		Header: "// github.com/sthorer/api",
		IDType: &field.TypeInfo{Type: field.TypeInt},
	})

	if err != nil {
		log.Fatal("running ent codegen:", err)
	}
}
