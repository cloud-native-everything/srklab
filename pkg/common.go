package srklab

import (
        "sync"
        "gopkg.in/yaml.v2"
        "io/ioutil"
        "log"
	"fmt"

)

type stdOut struct{}


var wg = sync.WaitGroup{}

type Kind_Data struct {
    Cluster []Cluster_Data `yaml:"clusters"`
    Links []Link_Data `yaml:"links"`
    Network string `yaml:"network"`
    Prefix string `yaml:"prefix"`
    ClabTopology string `yaml:"clabTopology"`
}

type Cluster_Data struct {
    //Info of every K8s Kind Cluster
    Name string    `yaml:"name"`
    Config  string    `yaml:"config"`
    Kubeconfig  string    `yaml:"kubeconfig"`
    Image string `yaml:"image"` 
    ImagesToLoad []ImagesToLoad_Data `yaml:"imagesToLoad"` 
}

type Link_Data struct {
    K8sNode string    `yaml:"k8sNode"`
    ClabNode  string    `yaml:"clabNode"`
    K8sIpv4  string    `yaml:"k8sIpv4"`
}

type ImagesToLoad_Data struct {
    Image string    `yaml:"image"`
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

func (s *stdOut) Write(p []byte) (n int, err error) {
        fmt.Print(string(p))
        return len(p), nil
}
