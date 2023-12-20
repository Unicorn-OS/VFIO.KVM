/* Xen
https://rubenerd.com/xen-with-qcow2/
https://www.virtuatopia.com/index.php/Using_QEMU_Disk_Images_for_Xen_DomainU_Systems

https://en.wikibooks.org/wiki/QEMU/Images#Creating_an_image
*/


size = 110
img = "ubu"
store = "~/.uni/disk"


init(){
sudo apt install qemu-utils -y
mkdir -p {store}
}

func create(){
func qcow(){
qemu-img create -f qcow2 {img}.qcow2 {size}G
}

func raw(){
qemu-img create -f raw ~/uni/try.raw 20G
}
}

config(){
// For raw:
disk = ['file:/home/me/img/disk.raw,xvda,w']

// For qcow2, the syntax is:
disk = ['tap:qcow2:/path/images/disk.qcow2,xvda,w']
}
