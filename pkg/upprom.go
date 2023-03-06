package srklab

import (
	"fmt"
	"os/exec"
	"strings"
)

func MyNodes(mycluster string) {

	cmd := exec.Command("/root/go/bin/kind", "get", "nodes", "--name", mycluster)
	out, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	files := strings.Split(string(out), "\n")
	filesSlice := make([]string, 0, len(files))
	for _, file := range files {
		if file != "" {
			getIP(file)
			filesSlice = append(filesSlice, file)
		}
	}
	fmt.Println(filesSlice)
}

func getIP(container string) {
	cmd := exec.Command("docker", "inspect", "-f", "'{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}'", container)
	out, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	ipadd := strings.ReplaceAll(string(out), "'", "")
	fmt.Println("------->", ipadd)

}
