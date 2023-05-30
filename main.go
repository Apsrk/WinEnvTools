package main

import (
	"fmt"
	"os"
	"strings"

    "os/exec"
    "syscall"
	"golang.org/x/sys/windows/registry"
)

const (
	envKey      = `SYSTEM\CurrentControlSet\Control\Session Manager\Environment`
	envVarSplit = ";"
)

func checkError(err error) {
	if err != nil {
		if isAdministrator() {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		} else {
			fmt.Fprintf(os.Stderr, "Please run with administrator privileges.Error: %s\n", err.Error())
		}
		os.Exit(1)
	}
}

func getEnvVar(key string) (string, error) {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, envKey, registry.QUERY_VALUE)
	if err != nil {
		return "", err
	}
	defer k.Close()

	value, _, err := k.GetStringValue(key)
	if err != nil {
		return "", err
	}

	return value, nil
}

func setEnvVar(key, value string) error {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, envKey, registry.SET_VALUE)
	if err != nil {
		return err
	}
	defer k.Close()

	err = k.SetStringValue(key, value)
	if err != nil {
		return err
	}

	return nil
}


func delEnvVar(key string) error {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, envKey, registry.SET_VALUE)
	if err != nil {
		return err
	}
	defer k.Close()

	err = k.DeleteValue(key)
	if err != nil {
		return err
	}

	return nil
}

func removeSeparator(path string) string {
	return strings.TrimSuffix(strings.TrimSuffix(path, "/"), "\\")
}

func appendEnvVar(key, value string) error {
	curValue, err := getEnvVar(key)
	if err != nil {
		return err
	}
	curValue := strings.Replace(curValue, ";;", ";", -1)
	paths := strings.Split(curValue, envVarSplit)
	for _, path := range paths {
		if strings.EqualFold(path, value) || strings.EqualFold(removeSeparator(path), removeSeparator(value)) {
			return fmt.Errorf("path already exists")
		}
	}

	newValue := strings.Join([]string{curValue, value}, envVarSplit)
	err = setEnvVar(key, newValue)
	if err != nil {
		return err
	}

	return nil
}

func removeEnvVar(key, value string) error {
	curValue, err := getEnvVar(key)
	if err != nil {
		return err
	}

	paths := strings.Split(curValue, envVarSplit)
	newPaths := []string{}
	for _, path := range paths {
		if !strings.EqualFold(path, value) {
			newPaths = append(newPaths, path)
		}
	}

	newValue := strings.Join(newPaths, envVarSplit)
	err = setEnvVar(key, newValue)
	if err != nil {
		return err
	}

	return nil
}

func listEnvVar(key string) error {
	value, err := getEnvVar(key)
	if err != nil {
		return err
	}

	paths := strings.Split(value, envVarSplit)
	for _, path := range paths {
		fmt.Println(path)
	}

	return nil
}

func isAdministrator() bool {
    cmd := exec.Command("net", "session")
    if err := cmd.Run(); err != nil {
        if exitErr, ok := err.(*exec.ExitError); ok {
            if status, ok := exitErr.Sys().(syscall.WaitStatus); ok {
                return status.ExitStatus() == 0
            }
        }
    }
    return false
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: envtool OPTION [KEY] [VALUE]\n")
	fmt.Fprintf(os.Stderr, "Options:\n")
	fmt.Fprintf(os.Stderr, "  get,view		View the value of an environment variable\n")
	fmt.Fprintf(os.Stderr, "  add,append	Add a value to an environment variable\n")
	fmt.Fprintf(os.Stderr, "  rm,remove		Remove a value from an environment variable\n")
	fmt.Fprintf(os.Stderr, "  del,delete	Delete specific environment variable\n")
	fmt.Fprintf(os.Stderr, "  list			List all values of an environment variable\n")
	fmt.Fprintf(os.Stderr, "  set			Set an environment variable\n")
	os.Exit(1)
}

func main() {
	if len(os.Args) < 3 {
		usage()
	}

	option := os.Args[1]
	key := os.Args[2]

	switch option {
	case "view", "get":
		value, err := getEnvVar(key)
		checkError(err)
		fmt.Println(value)
	case "add", "append":
		if len(os.Args) < 4 {
			usage()
		}
		value := os.Args[3]
		err := appendEnvVar(key, value)
		checkError(err)
	case "set":
		if len(os.Args) < 4 {
			usage()
		}
		value := os.Args[3]
		err := setEnvVar(key, value)
		checkError(err)
	case "remove", "rm":
		if len(os.Args) < 4 {
			usage()
		}
		value := os.Args[3]
		err := removeEnvVar(key, value)
		checkError(err)
	case "delete", "del":
		if len(os.Args) < 4 {
			usage()
		}
		err := delEnvVar(key)
		checkError(err)
	case "list":
		err := listEnvVar(key)
		checkError(err)
	default:
		usage()
	}
}
