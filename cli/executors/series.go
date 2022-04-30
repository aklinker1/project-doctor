package executors

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/aklinker1/project-doctor/cli"
	"github.com/aklinker1/project-doctor/cli/config"
)

type SeriesExecutor struct {
	UI cli.UI
}

func (s *SeriesExecutor) Validate(checks []cli.Check) []error {
	slowMo := os.Getenv("SLOW_MO") == "true"
	errs := []error{}

	for _, check := range checks {
		err := s.check(check, slowMo)
		if err != nil {
			fmt.Println("    Error:", cli.ErrorMessage(err))
			errs = append(errs, err)
		}
	}

	return errs
}

func (s *SeriesExecutor) check(check cli.Check, slowMo bool) error {
	status := check.DisplayName

	stop := s.UI.Spinner(status)
	if slowMo {
		time.Sleep(2 * time.Second)
	}
	err := check.Verify(s.UI)
	stop(err)

	if errors.Is(err, config.NotInPathError) {
		fmt.Println("    Not installed")
		err = check.Fix(s.UI)
	}
	if errors.Is(err, config.WrongVersionError) {
		fmt.Printf("    Installed version: %s\n", config.AsWrongVersionError(err).InstalledVersion)
	}

	return err
}
