/*
https://pve.proxmox.com/wiki/Shrink_Qcow2_Disk_Files#Shrink_the_Disk_File

Parallel
https://www.redhat.com/archives/libguestfs/2017-November/msg00082.html
*/


compress(){
sudo qemu-img convert -O qcow2 -c win10.qcow2 WhatsApp-Chillah-8-30-2018.qcow2
}

parallel_compress(){
sudo qemu-img convert -W -m -c -O qcow2 win10.qcow2 WhatsApp-Chillah-8-30-2018.qcow2
}
