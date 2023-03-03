package srklab

import (
	"fmt"
	"os/exec"
	"strings"
)

func KNetScript(myipv4 string, myif string, mygw string, mynode string) {
	cmd := exec.Command("docker", "exec", "-i", mynode, "bash")
	cmd.Stdin = strings.NewReader(fmt.Sprintf(`#!/bin/bash	
	ip link set %s down
	ip link set %s up mtu 9500
	ip addr add %s dev %s
	ip link set %s up
	ip route del default
	ip route add default via %s
	`, myif, myif, myipv4, myif, myif, mygw))
	fmt.Printf("INFO: Executing bash scripts to change settings for node %s and interface %s\n", mynode, myif)
	output, err := cmd.CombinedOutput()
	fmt.Println("-->", string(output))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(output))
}
