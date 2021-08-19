package cmd

import (
	"fmt"

	"github.com/konoui/go-alfred"
)

func printUpdateResults(err error) (_ error) {
	if err != nil {
		fmt.Fprintf(outStream, "update failed due to %s", err)
	} else {
		fmt.Fprintf(outStream, "update succeeded")
	}
	return
}

func (cfg *config) updateDB() error {
	if cfg.confirm {
		// update explicitly
		awf.Logger().Infoln("updating tldr database...")
		err := cfg.tldrClient.Update()
		return printUpdateResults(err)
	}

	awf.Append(
		alfred.NewItem().
			Title("Please Enter if update tldr database").
			Arg(fmt.Sprintf("--%s --%s", longUpdateFlag, confirmFlag)),
	).
		Variable(nextActionKey, nextActionShell).
		Output()

	return nil
}
