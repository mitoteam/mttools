package mttools

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"

	_ "embed"
)

//go:embed service_tools.service.template
var systemdServiceUnitFileTemplate string

const SystemdServiceDirPath = "/etc/systemd/system"

type ServiceData struct {
	Name       string //service name
	User       string //user
	Group      string //group
	Executable string //path to executable file
	WorkingDir string //full working directory path
	Autostart  bool   //start service automatically
}

func IsSystemdAvailable() bool {
	return IsDirExists(SystemdServiceDirPath)
}

func (unit *ServiceData) InstallSystemdService() error {
	service_unit_filename := filepath.Join(SystemdServiceDirPath, unit.Name+".service")

	if IsFileExists(service_unit_filename) {
		return fmt.Errorf(
			"File %s already exists. Use 'uninstall' command or remove file manually.\n", service_unit_filename,
		)
	}

	t := template.New("service")
	if _, err := t.Parse(systemdServiceUnitFileTemplate); err != nil {
		return err
	}

	var err error

	unit.Executable, err = os.Executable()

	if err != nil {
		return err
	}

	unit.WorkingDir, err = os.Getwd()

	if err != nil {
		return err
	}

	file, err := os.Create(service_unit_filename)
	if err != nil {
		return err
	}

	if err := t.Execute(file, unit); err != nil {
		return err
	}

	if err := file.Chmod(0644); err != nil {
		return err
	}

	file.Close()

	log.Printf("File %s created.", service_unit_filename)

	log.Println("Reloading systemctl daemon")
	out, err := exec.Command("systemctl", "daemon-reload").Output()
	if err != nil {
		return err
	}

	log.Println(string(out[:]))

	if unit.Autostart {
		log.Printf("Enabling '%s' service autostart.\n", unit.Name)
		out, err := exec.Command("systemctl", "enable", unit.Name).Output()
		if err != nil {
			return err
		}

		log.Println(string(out[:]))
	}

	return nil
}

func (unit *ServiceData) UninstallSystemdService() error {
	service_unit_filename := filepath.Join(SystemdServiceDirPath, unit.Name+".service")

	if !IsFileExists(service_unit_filename) {
		return fmt.Errorf(
			"File %s does not exists. Use 'install' command to create service unit file.\n", service_unit_filename,
		)
	}

	var out []byte
	var err error

	log.Printf("Stopping '%s' service.\n", unit.Name)
	out, err = exec.Command("service", unit.Name, "stop").Output()
	if err != nil {
		return err
	}

	log.Println(string(out[:]))

	log.Printf("Disabling '%s' service autostart.\n", unit.Name)
	out, err = exec.Command("systemctl", "disable", unit.Name).Output()
	if err != nil {
		return err
	}

	log.Println(string(out[:]))

	err = os.Remove(service_unit_filename)
	if err != nil {
		return err
	}

	log.Printf("File %s removed.", service_unit_filename)

	log.Println("Reloading systemctl daemon")
	out, err = exec.Command("systemctl", "daemon-reload").Output()
	if err != nil {
		return err
	}

	log.Println(string(out[:]))

	return nil
}
