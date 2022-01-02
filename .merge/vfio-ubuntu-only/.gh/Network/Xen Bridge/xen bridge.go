/*  https://help.ubuntu.com/community/Xen#Network_Configuration
*/

func init(){
  sudo stop network-manager
  echo "manual" | sudo tee /etc/init/network-manager.override
  sudo apt-get install bridge-utils

  // My
  sudo systemctl disable network-manager.service
}

func Config(){
  nic=enp0s31f6

echo '''
auto lo enp0s31f6 xenbr0
iface lo inet loopback

iface xenbr0 inet dhcp
  bridge_ports enp0s31f6

iface enp0s31f6 inet manual
''' | sudo tee /etc/network/interfaces

  sudo ifdown $nic && sudo ifup xenbr0 && sudo ifup $nic
}

func Start(){
sudo service networking restart
}
