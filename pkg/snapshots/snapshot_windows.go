package snapshots

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"os/exec"
)

func WipeSnapshots(l *logrus.Logger) error {
	l.Debug("wiping Windows shadow copies")

	cmd := exec.Command("vssadmin", "delete", "shadows", "/all")

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	if err := cmd.Run(); err != nil {
		l.Debugf("%s", out.String())
		return fmt.Errorf("could not run vssadmin: %+v", err)
	}

	return nil
}
