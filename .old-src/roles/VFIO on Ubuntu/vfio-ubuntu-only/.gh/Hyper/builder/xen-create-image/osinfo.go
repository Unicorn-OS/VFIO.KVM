/* Uses: libosinfo-bin
https://raymii.org/s/articles/virt-install_introduction_and_copy_paste_distro_install_commands.html
*/

func init(){
  sudo apt -y install libosinfo-bin
}

func query(){
  // Works with virt-build not with xen-create-img!
  osinfo-query os
}

func Works_in_xen(){
  'cosmic'
  'stretch'
  'buster'
}

func Works_in_virt-build(){

}
