package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/alecthomas/kong"
	"github.com/evantbyrne/1password-opvault/opvault"
	"golang.org/x/term"
)

const version = "1.1.0"

type exportCommand struct {
	HidePasswords bool   `flag name:"hide-passwords" help:"Hide passwords instead of exporting them."`
	Path          string `arg name:"path" help:"Path of opvault file to export." type:"path"`
	Profile       string `flag name:"profile" help:"Which profile to use." default:"default"`
}

func (cmd *exportCommand) Run() error {
	vault, err := opvault.Open(cmd.Path)
	if err != nil {
		return err
	}
	profiles, err := vault.ProfileNames()
	if err != nil {
		return err
	}
	if len(profiles) == 0 {
		return errors.New("not a valid vault")
	}
	profile, err := vault.Profile(cmd.Profile)
	if err != nil {
		return fmt.Errorf("cannot open profile '%s': %s", cmd.Profile, err)
	}

	// Password prompt.
	password := os.Getenv("ONEPASSWORD_OPVAULT_PASSWORD")
	if len(password) < 1 {
		prompt := "Password: "
		if hint := profile.PasswordHint(); len(hint) > 0 {
			prompt = fmt.Sprintf("Password (hint: %s): ", hint)
		}
		print(prompt)
		passwordBytes, err := term.ReadPassword(int(os.Stdin.Fd()))
		print("\n")
		if err != nil {
			return err
		}
		if len(passwordBytes) < 1 {
			return errors.New("password cannot be empty")
		}
		password = string(passwordBytes)
	}

	if err := profile.Unlock(password); err != nil {
		return errors.New("incorrect password")
	}

	items, err := profile.Items()
	if err != nil {
		return err
	}

	for _, item := range items {
		str, err := item.JsonMarshal(cmd.HidePasswords)
		if err != nil {
			return err
		}
		fmt.Println(string(str))
	}

	return nil
}

type profilesCommand struct {
	Path string `arg name:"path" help:"Path of opvault file to export." type:"path"`
}

func (cmd *profilesCommand) Run() error {
	vault, err := opvault.Open(cmd.Path)
	if err != nil {
		return err
	}
	profiles, err := vault.ProfileNames()
	if err != nil {
		return err
	}
	if len(profiles) == 0 {
		return errors.New("not a valid vault")
	}
	for _, profileName := range profiles {
		fmt.Println(profileName)
	}
	return nil
}

type versionCommand struct{}

func (cmd *versionCommand) Run() error {
	fmt.Println(version)
	return nil
}

var cli struct {
	Export   exportCommand   `cmd help:"Export 1Password opvault."`
	Profiles profilesCommand `cmd help:"Show all profiles in vault."`
	Version  versionCommand  `cmd help:"Show version."`
}

func main() {
	ctx := kong.Parse(&cli)
	err := ctx.Run()
	ctx.FatalIfErrorf(err)
}
