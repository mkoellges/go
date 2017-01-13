package main

import ("fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)
type etcd struct {
	Hostname string `yaml:"hostname"`
	Port int64 `yaml:"port"`
}

func (c *etcd) getConf() *etcd {

	yamlFile, err := ioutil.ReadFile("test.yml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

func main() {
	var c etcd
	c.getConf()

	fmt.Println(c.Port, c.Hostname)

	fmt.Println("you can now call etcd on", c.Hostname, "and use port", c.Port)
}
