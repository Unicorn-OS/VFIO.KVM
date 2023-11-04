docs.ansible.com/ansible/2.3/playbooks_conditionals.html
serverfault.com/questions/875247/whats-the-difference-between-include-tasks-and-import-tasks

example:
```
- include_tasks: '{{ ansible_distribution }}.yml'
```
