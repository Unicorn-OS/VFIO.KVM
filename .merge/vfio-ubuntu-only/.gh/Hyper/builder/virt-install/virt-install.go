/* Works for OpenStack!
*/

cd.example = 'kubuntu-18.10-desktop-amd64.iso'
img.example = 'kub.qcow2'

virt-install --virt-type xen --hvm --name bionic --ram 1024 \
  --cdrom=$cd \
  --disk $img/bionic.qcow2,format=qcow2 \
  --network network=default \
  --graphics vnc,listen=0.0.0.0 --noautoconsole \
  --os-type=linux
