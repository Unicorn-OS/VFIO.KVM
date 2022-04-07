Works(){

name='deb-kube'
size=110
os='stretch'
store=~/.uni/disk

mkdir -p $store

sudo xen-create-image --hostname=$name \
--size=${size}Gb \
--noswap \
--memory=2G \
--maxmem=6G \
--vcpus=8 \
--dir=$store \
--dhcp \
--pygrub \
--image=sparse \
--output=$store \
--password=u \
--dist=$os
}

Option(){
    // for easy img conver
    --noswap \

    // boot
    --nopygrub \
    OR
    --pygrub \

    --password=u \
}

get_pass(){
  sudo cat /var/log/xen-tools/dev.log
}

/* Breaks!
--fs=btrfs \
*/
