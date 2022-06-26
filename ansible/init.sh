inventory=".this/"

copy_template(){
  cp -r .this/inventory.example .this/inventory
}

set_user(){
  sed -i "s/ansible_user=me/ansible_user=$USER/" .this/inventory/test/hosts
}

link_template(){
  ln -srf .this/ansible.cfg .
  ln -srf .this/uni .
}

my_optional_feature(){
  ln -srf roles/.my/feature/vars.example roles/.my/feature/vars
}

checkLocation(){
  SCRIPTPATH="$( cd -- "$(dirname "$0")" >/dev/null 2>&1 ; pwd -P )"
  cd $SCRIPTPATH

  # Debug
  # dirname $SCRIPTPATH
}


main(){
  checkLocation
  copy_template
  set_user
  link_template
  my_optional_feature
}

main
