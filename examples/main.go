package main

import (
	"context"

	"github.com/hardstylez72/okex"
	"github.com/hardstylez72/okex/api"
)

func main() {
	api.NewClient(context.Background(), "", "", "", okex.NormalServer, nil)
}
