
package provider;

import "os";
import "fmt";
import "log";
import (
	"github.com/hashicorp/terraform/helper/schema"
	"os/exec"
);

func init() {
	start()
	// Terraform is already adding the timestamp for us
	log.SetFlags(log.Lshortfile);
	log.SetPrefix(fmt.Sprintf("pid-%d-", os.Getpid()));
}

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"vagrant_instance": resourceVagrant(),
		},
	}
}


func start() {
	valor:="Vagrant.configure(\"2\") do |config|\n"
	valor += "config.vm.provider \"virtualbox\" do |vb|\n";
	valor += "end\n";

	var commands string = "echo '" + valor + "' > $HOME/.terraform/Vagrantfile";

	_, err := exec.Command("sh", "-c", commands).CombinedOutput();
	if (err != nil) {
		os.Stderr.WriteString(err.Error());
	}

}