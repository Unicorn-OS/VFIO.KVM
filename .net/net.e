[lan]
192.168.1.113

[lan:vars]
ansible_become_password=u
ansible_user=me


[net:children]
local
lan
