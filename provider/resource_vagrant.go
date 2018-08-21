
package provider;

import "os";
import "fmt";
//import "bufio";
//import "errors";
import "os/exec";
import (
	"github.com/hashicorp/terraform/helper/schema"
);

func resourceVagrant() *schema.Resource {
	return &schema.Resource{
		Exists: resourceVagrantExists,
		Create: resourceVagrantCreate,
		Read:   resourceVagrantRead,
		Update: resourceVagrantUpdate,
		Delete: resourceVagrantDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"network": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Required:    true,
			},
			"box": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
			},
			"path": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
			},
			"provision": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func print(valor string) {
	fmt.Println(valor);
	//log.Printf(valor);

	var commands string = "echo '" + valor + "'  >> $HOME/.terraform/Vagrantfile";

	output, err := exec.Command("sh", "-c", commands).CombinedOutput();
	if (err != nil) {
	  	os.Stderr.WriteString(err.Error());
	}
	fmt.Println(string(output));

}

func resourceVagrantExists(d *schema.ResourceData, meta interface{}) (bool, error) {


	print("[resourceVagrantExists]---------------------");

	return true, nil;
}

func resourceVagrantCreate(d *schema.ResourceData, meta interface{}) error {
	//box := d.Get("box").(string)
	//print("[resourceVagrantCreate]---------------------");

	name:=d.Get("name").(string);
	network:=d.Get("network").(string);
	box:=d.Get("box").(string);
	path:=d.Get("path").(string);
	provision:=d.Get("provision").(string);
	ip:=d.Get("ip").(string);
	config:= "\tconfig.vm.define \""+name+"\" do |"+name+"|\n\t\t"

	config+=name+".vm.hostname = \""+name+"\"\n\t\t"
	config+=name+".vm.box = \""+box+"\"\n\t\t"
	config+=name+".vm.network \""+network+"\", ip: \""+ip+"\"\n\t\t"
	config+=name+".vm.provision \""+provision+"\", path: \""+path+"\"\n"
	config+="\tend"
	print(config)

	//fmt.Fprint(file,ip)
	//fmt.Fprint(file,reflect.TypeOf(ip))
    //print(ip.Next().Value.(string))
	//network:= d.Get("private_network").(*schema.Set).List();
	//
	//for _, value :=range network{
	//	for _, value2 :=range value.(string) {
	//		print(value2)
	//	}
	//}
	//print (address.Get("ip").(string))

    return nil;
}

func resourceVagrantRead(d *schema.ResourceData, meta interface{}) error {
	print("[resourceVagrantRead]---------------------");

	return nil;
}

func resourceVagrantUpdate(d *schema.ResourceData, meta interface{}) error {
	print("[resourceVagrantUpdate]---------------------");

	// Enable partial state mode
	d.Partial(true);
	 
	if d.HasChange("address") {
		// Try updating the address
		//if err := updateAddress(d, m); err != nil {
		//	return err
		//}
		d.SetPartial("address");
	}
	// If we were to return here, before disabling partial mode below,
	// then only the "address" field would be saved.

	// We succeeded, disable partial mode. This causes Terraform to save
	// save all fields again.
	d.Partial(false);
	 
	return resourceVagrantRead(d, meta);
}

func resourceVagrantDelete(d *schema.ResourceData, meta interface{}) error {
	print("[resourceVagrantDelete]---------------------");

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("");

	return nil;
}
