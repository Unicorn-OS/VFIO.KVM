https://docs.ansible.com/ansible/latest/reference_appendices/special_variables.html

# Role
{{ role_path }}/files/checkGroups

# Useful for setting /home/$USER variables!

`home: '/home/{{ ansible_user }}'`
