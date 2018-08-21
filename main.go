
package main;

import "github.com/hashicorp/terraform/plugin";
import "github.com/hashicorp/terraform/terraform";
import (
	"./provider"
	"fmt"
	"os"
	"os/exec"
	"log"
);

func main() {

	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return provider.Provider();
		},
	});
	defer end("end")
}

func end(valor string) {
	fmt.Println(valor);
	log.Printf(valor);

	var commands string = "echo '" + valor + "'>> $HOME/.terraform/Vagrantfile";

	output, err := exec.Command("sh", "-c", commands).CombinedOutput();
	if (err != nil) {
		os.Stderr.WriteString(err.Error());
	}
	_, err1 := exec.Command("sh", "-c", "cd $HOME/.terraform && vagrant up").CombinedOutput();
	if (err1 != nil) {
		os.Stderr.WriteString(err.Error());
	}
	fmt.Println(string(output));
}
