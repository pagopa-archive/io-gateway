#!/bin/bash

USERNAME=iodev

cp -R /tmp/.ssh/* /home/$USERNAME/.ssh
chmod 700 /home/$USERNAME/.ssh
find /home/$USERNAME/.ssh -name "id_rsa*" -not -name *.pub -exec chmod 600 {} \;
find /home/$USERNAME/.ssh -name "id_rsa*" -name *.pub -exec chmod 644 {} \;
