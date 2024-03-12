# ansible-clone

This Ansible playbook ensures that `git` is installed on your machine and then clones a specified GitHub repository to a defined directory.

## Prerequisites

- Ansible installed on the machine.

## How to use

1. Clone this repository:

       git clone https://github.com/ibilalkayy/ansible-clone.git
       cd ansible-clone

2. Modify the playbook `file.yml`.

3. Run the playbook:

       bash ansible-playbook file.yml

## What it does

- Ensures the `git` package is installed on your machine.
- Clones the specified GitHub repository to the provided directory path.

## Notes

- This playbook is set to run on the `localhost` and will not attempt an SSH connection as the connection is set to `local`.
- The playbook uses the `apt` package manager, which means it's designed for Debian-based distributions such as Ubuntu.

- [Here is the video explanation of it.](https://www.youtube.com/watch?v=NhRfR2H4bWs)