#!/bin/sh

PREREQ=""

prereqs()
{
   echo "$PREREQ"
}

case $1 in
prereqs)
   prereqs
   exit 0
   ;;
esac

# Need to make this use ansible Lists instead!
# RTX 2070 Super, Elgato HD60, & 4K60 Pro
for dev in 0000:01:00.0 0000:01:00.1 0000:01:00.2 0000:01:00.3 0000:02:00.0 0000:04:00.0
do
 echo "vfio-pci" > /sys/bus/pci/devices/$dev/driver_override
 echo "$dev" > /sys/bus/pci/drivers/vfio-pci/bind
done

exit 0
