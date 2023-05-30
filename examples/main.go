package main

import (
	"context"

	"github.com/amir-the-h/okex"
	"github.com/amir-the-h/okex/api"
)

func main() {
	api.NewClient(context.Background(), "", "", "", okex.NormalServer, nil)
}
