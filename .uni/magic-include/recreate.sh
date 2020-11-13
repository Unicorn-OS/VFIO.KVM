unilib=~/src/uni/ansible
magic="$unilib/magic include/magic.yml"

Example(){
ln -srf $magic ../tasks/core/
ln -srf $magic ../tasks/gui/
ln -srf $magic ../tasks/openstack/
ln -srf $magic ../tasks/migrate/
ln -srf $magic ../tasks/microcloud/
}
