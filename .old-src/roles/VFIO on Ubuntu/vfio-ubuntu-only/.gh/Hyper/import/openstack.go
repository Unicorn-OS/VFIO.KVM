/*
from: http://libguestfs.org/virt-builder.1.html

Importing into OpenStack
Import the image into Glance (the OpenStack image store) by doing:
*/

 glance image-create --name fedora-27-image --file fedora-27.img \
   --disk-format raw --container-format bare \
   --is-public True
