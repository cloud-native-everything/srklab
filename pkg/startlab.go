package srklab

import (
        "os"
        "fmt"
        "os/exec"

        "sigs.k8s.io/kind/pkg/cmd"
        "github.com/cloud-native-everything/srklab/pkg/kind"
        "github.com/cloud-native-everything/srklab/pkg/dockerapi"
)

func Start(configFile string) {


        kindVars := getVars(configFile)
	DockerNet.DockerNetworkCreate(kindVars.Network, kindVars.Prefix)
	wg.Add(1)
	go clab(kindVars.ClabTopology)
	runKind(*kindVars)
        wg.Wait()
}


// Main is the kind main(), it will invoke Run(), if an error is returned
// it will then call os.Exit

func runKind(karray Kind_Data) {
        
	KindExec := func(Args *[]string) {
		if err := kindApp.Run(cmd.NewLogger(), cmd.StandardIOStreams(), *Args); err != nil {
                    os.Exit(1)
                }
		wg.Done()
	}

	fmt.Println("--->", karray.Cluster)
	fmt.Println(len(karray.Cluster))
        for i := 0; i < len(karray.Cluster); i++ {
		aux := []string{"create" , "cluster" , "--name" , karray.Cluster[i].Name}
		aux = append( aux, "--image", karray.Cluster[i].Image)
		aux = append( aux, "--config", karray.Cluster[i].Config)
		aux = append( aux, "--kubeconfig", karray.Cluster[i].Kubeconfig)
		fmt.Println("--->", aux)
		wg.Add(1)
		go KindExec(&aux)

	}
}


func clab(mytopofile string){

    // Create a new command to run "ls" with no arguments
    cmd := exec.Command("/usr/bin/containerlab", "deploy",  "--topo" , mytopofile )
//	cmd := exec.Command("ls")
	cmd.Stdout = &stdOut{}
    // Run the command and wait for it to finish
    err := cmd.Run()
    if err != nil {
        panic(err)
    }
    wg.Done()
}


