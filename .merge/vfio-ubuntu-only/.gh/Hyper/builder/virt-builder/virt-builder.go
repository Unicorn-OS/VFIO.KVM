/*
 * http://libguestfs.org/virt-builder.1.html
 */

/*
Password
http://libguestfs.org/virt-builder.1.html#users-and-passwords
*/

pw = '--root-password password:u'

init(){
sudo apt install libguestfs-tools
}

ubuntu(){
  virt-builder ubuntu-18.04 --format qcow2 --size 33G --hostname ubuLTS
}

centOS(){
  virt-builder ubuntu-18.04 --format qcow2 --size 33G --hostname ubuLTS
}
