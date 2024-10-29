package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/schlauerlauer/liesl/db"
)

func recursiveLeafs(ctx context.Context, queries *db.Queries, leafs []string, level int) error {
	if level >= 10 {
		fmt.Println("level limit reached, returning")
		return nil
	}

	for _, leaf := range leafs {
		fmt.Printf("%s [ ] %s\n", strings.Repeat(" ", (1+level)*2), leaf)

		// if leaf == rootNode // TODO

		if subLeafs, err := queries.GetEdges(ctx, leaf); err != nil {
			return err
		} else {
			recursiveLeafs(ctx, queries, subLeafs, level+1)
		}
	}
	return nil
}
