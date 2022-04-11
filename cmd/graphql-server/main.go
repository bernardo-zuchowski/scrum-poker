package main

import (
	graphqlserver "multipoker/api/graphql"
)

func main() {
	graphqlserver.Start(":8080")
}
