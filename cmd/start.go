package main

import (
	"fmt"
	"io/ioutil"
	"path"

	"github.com/codegangsta/cli"

	"github.com/coreos/coreinit/job"
)

func newStartUnitCommand() cli.Command {
	return cli.Command{
		Name: "start",
		Flags: []cli.Flag{
			cli.IntFlag{"count", 1, "Run N instances of these units."},
		},
		Usage:  "Start (activate) one or more units",
		Action: startUnitAction,
	}
}

func startUnitAction(c *cli.Context) {
	r := getRegistry(c)

	payloads := make([]job.JobPayload, len(c.Args()))

	for i, v := range c.Args() {
		out, err := ioutil.ReadFile(v)
		if err != nil {
			fmt.Printf("%s: No such file or directory\n", v)
			continue
		}

		name := path.Base(v)
		payload := job.JobPayload{name, string(out)}
		if err != nil {
			fmt.Print(err)
		} else {
			payloads[i] = payload
		}
	}

	//TODO: Handle error response from NewJobRequest
	req, _ := job.NewJobRequest(payloads, nil)
	req.Count = c.Int("count")

	r.AddRequest(req)
}
