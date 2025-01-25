package shell

import (
	"fmt"
	"os/exec"
)

func ChangeDirectory(dir string) error {
	cmd := exec.Command("powershell", "-NoProfile", "-Command", fmt.Sprintf(`& {Start-Process powershell -ArgumentList '-NoProfile -NoExit -Command Set-Location -Path "%s"'}`, dir))

	fmt.Println("Changing directory to:", dir)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error executing PowerShell script:", err)
		return err
	}
	return nil
}
