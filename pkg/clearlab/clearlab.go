package clearlab

import (
	"fmt"
	"os"
	"os/exec"
	"sync"

	"io/ioutil"
	"log"

	kindApp "github.com/cloud-native-everything/srklab/pkg/kind"
	"github.com/cloud-native-everything/srklab/pkg/dockerapi"
	"gopkg.in/yaml.v2"
	"sigs.k8s.io/kind/pkg/cmd"
)

type stdOut struct{}

var wg = sync.WaitGroup{}

type Kind_Data struct {
	Cluster []Cluster_Data `yaml:"clusters"`
}

type Cluster_Data struct {
	//
	Name       string `yaml:"name"`
	Config     string `yaml:"config"`
	Kubeconfig string `yaml:"kubeconfig"`
	Image      string `yaml:"image"`
}

func Main(configFile string) {

	kindVars := getVars(configFile)
	wg.Add(1)
	go delclab()
	delKind(*kindVars)
	wg.Wait()
	DockerNet.DockerNetworkDelete("kind")

}

// Main is the kind main(), it will invoke Run(), if an error is returned
// it will then call os.Exit

func delKind(karray Kind_Data) {

	KindExec := func(Args *[]string) {
		if err := kindApp.Run(cmd.NewLogger(), cmd.StandardIOStreams(), *Args); err != nil {
			os.Exit(1)
		}
		wg.Done()
	}

	fmt.Println("--->", karray.Cluster)
	fmt.Println(len(karray.Cluster))

	aux := []string{"delete", "clusters", "--all"}
	fmt.Println("--->", aux)
	wg.Add(1)
	go KindExec(&aux)

}

func checkerr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getVars(configFile string) (kdata *Kind_Data) {
	gtt_config, err := ioutil.ReadFile(configFile)
	checkerr(err)
	err = yaml.Unmarshal(gtt_config, &kdata)
	return kdata

}

func delclab() {

	// Create a new command to run "ls" with no arguments
	cmd := exec.Command("/usr/bin/containerlab", "destroy", "--topo", "/root/srklab/topo.yml")
	cmd.Stdout = &stdOut{}
	// Run the command and wait for it to finish
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
	wg.Done()
}

func (s *stdOut) Write(p []byte) (n int, err error) {
	fmt.Print(string(p))
	return len(p), nil
}
