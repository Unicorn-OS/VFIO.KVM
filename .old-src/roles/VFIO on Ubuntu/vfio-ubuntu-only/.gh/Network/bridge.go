func init(){
sudo apt-get install bridge-utils ifupdown -y

sudo stop network-manager
echo "manual" | sudo tee /etc/init/network-manager.override
}

func create_XenBridge():
/* From
https://help.ubuntu.com/community/Xen#Network_Configuration
*/
nic="enp2s0"

echo "auto lo $nic xenbr0
iface lo inet loopback

iface xenbr0 inet dhcp
  bridge_ports $nic

iface $nic inet manual" | sudo tee /etc/network/interfaces

sudo ifdown $nic && sudo ifup xenbr0 && sudo ifup $nic
}


func Delete(i){
// https://unix.stackexchange.com/questions/31763/bring-down-and-delete-bridge-interface-thats-up 

ip link set br100 down
brctl delbr br100
}
