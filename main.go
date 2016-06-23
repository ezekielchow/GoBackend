package main

import(
    "encoding/json"
    "fmt"
    "log"
    "math/rand"

    "github.com/graphql-go/graphql"
)

func main() {
    // Schema
    fields := graphql.Fields{
        "ezekiel": &graphql.Field{
            Type: graphql.String,
            Resolve: func(p graphql.ResolveParams) (interface{}, error) {
                return "chow", nil
            },
        },
        "randINT": &graphql.Field{
            Type: graphql.Int,
            Resolve: func (p graphql.ResolveParams) (interface{}, error) {
                return rand.Intn(100), nil
            },
        },
    }
    rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
    schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
    schema, err := graphql.NewSchema(schemaConfig)
    if err != nil {
        log.Fatalf("failed to create new schema, error: %v", err)
    }

    // Query
    query := `
        {
            randINT
            ezekiel
        }
    `
    params := graphql.Params{Schema: schema, RequestString: query}
    r := graphql.Do(params)
    if len(r.Errors) > 0 {
        log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
    }
    rJSON, _ := json.Marshal(r)
    fmt.Printf("%s \n", rJSON) // {“data”:{“hello”:”world”}}
}
