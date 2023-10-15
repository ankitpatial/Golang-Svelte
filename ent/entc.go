//go:build ignore

package main

import (
	"log"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	ex, err := entgql.NewExtension(
		entgql.WithWhereFilters(true),
		entgql.WithConfigPath("../server/gqlgen.yml"),
		// Generate the filters to a separate schema file and load it in the gqlgen.yml config.
		entgql.WithSchemaPath("../server/ent.graphql"),
	)
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}

	// feature ref: https://entgo.io/docs/feature-flags
	//
	conf := &gen.Config{
		Features: []gen.Feature{
			gen.FeatureModifier,
			gen.FeatureUpsert,
		},
	}
	opts := []entc.Option{
		entc.Extensions(ex),
	}

	if err := entc.Generate("./schema", conf, opts...); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
