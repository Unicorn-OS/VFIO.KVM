func resize(){
qemu-img resize $img 200G
virt-customize -a $img --run-command 'growpart /dev/sda 1'
virt-customize -a $img --run-command 'resize2fs /dev/sda1'
} 
