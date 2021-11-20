package main

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/configs"
	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/dtos"
	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/handlers"
)

type adminCommand struct {
	email   string
	promote bool
	demote  bool
}

func init() {
	RegisterCommand("admin", func(flags *flag.FlagSet) Command {
		cmd := new(adminCommand)
		flags.StringVar(&cmd.email, "email", "", "Email of user to set administrator status. User must already be registered into gallery-api.")
		flags.BoolVar(&cmd.promote, "promote", false, "Promote user to administrator.")
		flags.BoolVar(&cmd.demote, "demote", false, "Demote user from administrator.")
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
	if c.email == "" {
		return fmt.Errorf("email is required")
	}

	if c.promote && c.demote || (!c.promote && !c.demote) {
		return fmt.Errorf("choose either to promote or demote user")
	}

	configs.InitializeConfig()
	if !configs.AppConfig.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	if err := handlers.InitializeHandler(); err != nil {
		return err
	}

	if _, err := handlers.Handler.UserGetOneByEmail(c.email); err != nil {
		return err
	}

	if err := handlers.Handler.UserUpdate(dtos.User{
		Email:   c.email,
		IsAdmin: c.promote,
	}); err != nil {
		return err
	}

	if c.promote {
		fmt.Printf("User %s is now administrator.\n", c.email)
	} else {
		fmt.Printf("User %s is no longer administrator.\n", c.email)
	}

	return nil
}
