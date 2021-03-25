package pkg

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func CheckEnv(args []string) {
	fmt.Println("Printing all of the env variables:")
	for _, val := range os.Environ() {
		fmt.Printf("%s\n", val)
	}
	fmt.Println("Command line argumets:")
	fmt.Println(args)
}

func RunEnvdir(args []string) error {
	path := args[0]

	envVarsMap, err := ReadEnvFiles(path)

	if err != nil {
		return err
	}

	err = RunCmd(args, envVarsMap)
	if err != nil {
		return err
	}
	return nil
}

func ReadEnvFiles(path string) (map[string]string, error) {

	resMap := make(map[string]string)

	files, err := filepath.Glob(path)

	if err != nil {
		return resMap, err
	}

	for _, file_path := range files {
		data, err := ioutil.ReadFile(file_path)
		if err != nil {
			return resMap, err
		}

		key := strings.Split(file_path, `\`)
		ind := len(key)
		resMap[key[ind-1]] = string(data)
	}

	return resMap, nil
}

func RunCmd(args []string, env map[string]string) error {
	currentArgs := args[1:]

	for key, value := range env {
		os.Setenv(key, value)
	}

	c1 := exec.Command(`.\hw-12.exe `, currentArgs...)
	// c1 := exec.Command("ls")

	c1.Stdout = os.Stdout

	if err := c1.Start(); err != nil {
		return err
	}
	if err := c1.Wait(); err != nil {
		return err
	}

	return nil
}
