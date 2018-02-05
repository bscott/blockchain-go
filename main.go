package main

import (
    "os"
    "log"
    "github.com/urfave/cli"
    "github.com/hoshsadiq/tescoin/server"
    "github.com/hoshsadiq/tescoin/block"
    "github.com/hoshsadiq/tescoin/helper"
)

func main() {
    app := cli.NewApp()
    app.Name = "tescoin"
    app.Version = "0.0.1"
    app.Usage = "Interact with the Tescoin cryptocurrency"

    app.Commands = []cli.Command{
        {
            Name: "server",
            Flags: []cli.Flag{
                cli.IntFlag{
                    Name:  "port",
                    Value: 8080,
                },
            },
            Action: func(c *cli.Context) error {
                blockchain := block.NewBlockchain()
                addr := helper.GenerateAddress()

                return server.RunServer(blockchain, addr, c.Int("port"))
            },
        },
    }

    err := app.Run(os.Args)
    if err != nil {
        log.Fatal(err.Error())
    }
}
