package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/schlauerlauer/liesl/db"
	"github.com/schlauerlauer/liesl/persistence"
	"github.com/urfave/cli/v2"
)

func main() {
	path := defaultString(os.Getenv("LIESL_DB"), "file:liesl.db")
	queries, err := persistence.NewRepository(path)
	if err != nil {
		slog.Error("error connecting repository", "err", err)
		os.Exit(1)
	}

	app := &cli.App{
		Name:  "liesl",
		Usage: "A recursive packing list cli",
		Commands: []*cli.Command{
			{
				Name: "get",
				Action: func(cCtx *cli.Context) error {
					ctx := context.Background() // TODO
					node := cCtx.Args().First()

					// print all nodes
					if node == "" {
						if res, err := queries.GetNodes(ctx); err != nil {
							fmt.Println(err)
							return err
						} else {
							for _, id := range res {
								fmt.Println(id)
							}
							return nil
						}
					}

					// print single node
					rootNode, err := queries.GetNode(ctx, node)
					if err != nil {
						return err
					}
					fmt.Println(rootNode)

					if rootLeafs, err := queries.GetEdges(ctx, node); err != nil {
						return err
					} else {
						recursiveLeafs(ctx, queries, rootLeafs, 0)
					}

					return nil
				},
			},
			{
				Name: "add",
				Action: func(cCtx *cli.Context) error {
					ctx := context.Background() // TODO
					id := cCtx.Args().First()

					err := queries.InsertNode(ctx, id)
					if err != nil {
						fmt.Println("error adding node")
						return err
					}

					nodes := cCtx.Args().Tail()
					for _, node := range nodes {
						if err := queries.InsertNode(ctx, node); err != nil {
							fmt.Println("error inserting node", err)
						}

						if err := queries.InsertEdge(ctx, db.InsertEdgeParams{
							Source: id,
							Target: node,
						}); err != nil {
							fmt.Println("error inserting edge", err)
						}
					}

					fmt.Println(id, "added with", len(nodes), "edges")
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		slog.Error("err", "err", err)
		os.Exit(1)
	}
}

func defaultString(compare, def string) string {
	if compare != "" {
		return compare
	}
	return def
}
