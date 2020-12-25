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
#
for dev in {{device0}} {{device1}}
do
 echo "vfio-pci" > /sys/bus/pci/devices/$dev/driver_override
 echo "$dev" > /sys/bus/pci/drivers/vfio-pci/bind
done

exit 0
