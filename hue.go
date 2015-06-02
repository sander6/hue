// package hue is a command line client for interacting with Philips Hue lighting
package main

import (
	"github.com/codegangsta/cli"

	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
)

var (
	MaxSupportedLights int = 50
)

const (
	contentType = "application/json"
)

func main() {
	var bridge *Bridge

	app := cli.NewApp()
	app.Name = "hue"
	app.Usage = "control Philips Hue lighting"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "host",
			Usage:  "hostname of the Hue bridge",
			Value:  "192.168.0.1",
			EnvVar: "HUE_BRIDGE",
		},
		cli.StringFlag{
			Name:   "user",
			Usage:  "API username registered with the Hue bridge",
			Value:  "",
			EnvVar: "HUE_USER",
		},
		cli.IntFlag{
			Name:  "transition-time, t",
			Usage: "effect transition time in milliseconds",
			Value: 0,
		},
		cli.BoolFlag{
			Name:  "verbose",
			Usage: "log debug output to stderr",
		},
	}
	app.Before = func(ctx *cli.Context) error {
		if ctx.Bool("verbose") {
			log.SetOutput(os.Stderr)
		} else {
			log.SetOutput(ioutil.Discard)
		}
		bridge = NewBridge(ctx.String("host"), ctx.String("user"))
		return nil
	}
	app.Commands = []cli.Command{
		{
			Name:        "on",
			Description: "turn on all lights",
			Action: func(_ *cli.Context) {
				if err := bridge.AllLights().On(); err != nil {
					fmt.Fprintf(os.Stderr, err.Error())
					os.Exit(1)
				}
			},
		},
		{
			Name:        "off",
			Description: "turn off all lights",
			Action: func(_ *cli.Context) {
				if err := bridge.AllLights().Off(); err != nil {
					fmt.Fprintf(os.Stderr, err.Error())
					os.Exit(1)
				}
			},
		},
		{
			Name:        "lights",
			Description: "interact with individual lights",
			Subcommands: []cli.Command{
				{
					Name:        "list",
					Description: "list all lights and their current state",
					Action: func(_ *cli.Context) {
						lights, err := bridge.AllLights().Find()
						if err != nil {
							fmt.Fprintln(os.Stderr, err.Error())
							os.Exit(1)
						}
						sort.Sort(BulbsSortedByID(lights))
						for _, light := range lights {
							LightPrinter{light}.WriteTo(os.Stdout)
						}
					},
				},
				{
					Name:        "on",
					Description: "turn the given lights on",
					Action: func(c *cli.Context) {
						for _, id := range c.Args() {
							if err := bridge.Light(id).On(); err != nil {
								fmt.Fprintf(os.Stderr, err.Error())
								os.Exit(1)
							}
						}
					},
				},
				{
					Name:        "off",
					Description: "turn the given lights off",
					Action: func(c *cli.Context) {
						for _, id := range c.Args() {
							if err := bridge.Light(id).Off(); err != nil {
								fmt.Fprintf(os.Stderr, err.Error())
								os.Exit(1)
							}
						}
					},
				},
				{
					Name:        "alert",
					Description: "blink the given lights once",
					Action: func(c *cli.Context) {
						for _, id := range c.Args() {
							if err := bridge.Light(id).Alert(); err != nil {
								fmt.Fprintln(os.Stderr, err.Error())
								os.Exit(1)
							}
						}
					},
				},
			},
		},
		{
			Name:        "groups",
			Description: "interact with groups of lights",
			Subcommands: []cli.Command{
				{
					Name:        "list",
					Description: "list all groups and their members",
					Action: func(_ *cli.Context) {
						groups, err := bridge.AllGroups().Find()
						if err != nil {
							fmt.Fprintln(os.Stderr, err.Error())
							os.Exit(1)
						}
						sort.Sort(GroupsSortedByID(groups))
						for _, group := range groups {
							GroupPrinter{group}.WriteTo(os.Stdout)
						}
					},
				},
				{
					Name:        "on",
					Description: "turn all lights in the given groups on",
					Action: func(c *cli.Context) {
						for _, group := range c.Args() {
							if err := bridge.Group(group).On(); err != nil {
								fmt.Fprintf(os.Stderr, err.Error())
								os.Exit(1)
							}
						}
					},
				},
				{
					Name:        "off",
					Description: "turn all lights in the given groups off",
					Action: func(c *cli.Context) {
						for _, group := range c.Args() {
							if err := bridge.Group(group).Off(); err != nil {
								fmt.Fprintf(os.Stderr, err.Error())
								os.Exit(1)
							}
						}
					},
				},
			},
		},
		{
			Name:        "scenes",
			Description: "interact with scenes",
			Subcommands: []cli.Command{
				{
					Name:        "list",
					Description: "list all lights and their current state",
				},
				{
					Name:        "apply",
					Description: "apply the given scene",
				},
			},
		},
	}

	app.Run(os.Args)
}
