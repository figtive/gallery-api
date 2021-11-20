package main

import (
	"flag"
	"fmt"
)

type adminCommand struct {
	email string
	admin bool
}

func init() {
	RegisterCommand("admin", func(flags *flag.FlagSet) Command {
		cmd := new(adminCommand)
		flags.StringVar(&cmd.email, "email", "", "Email of user to set administrator status. User must already registered into gallery-api.")
		flags.BoolVar(&cmd.admin, "admin", false, "User administrator status.")
		return cmd
	})
}

func (c *adminCommand) Description() string {
	return "Set user administrator status."
}

func (c *adminCommand) Usage() string {
	return "Usage: gallery-cli admin [options]"
}

func (c *adminCommand) Run(args []string) error {
	fmt.Fprintf(Stdout, "run admin: %s, %v\n", c.email, c.admin)
	return nil
}
