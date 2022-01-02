instance=win10

func convertConfFmt(){
/* https://libvirt.org/drvxen.html#imex */

  toNative(){
    // virsh domxml-to-native
    virsh -c xen:///system domxml-to-native xen-xl --domain $instance >$instance.cfg
  }

  toLibvirt(){
    // https://libvirt.org/drvxen.html#xmlimport
    virsh -c xen:///system domxml-from-native xen-xl $instance.cfg
  }
)
