# !/bin/bash
# This script is the quick start installation script used to
# build and install Gobash and set it up to be used in the 
# environment.

exec_install_path=/usr/bin
user=$(whoami)

# Go build
if [ ! -f /usr/bin/go ]
then
    echo "Failed to build : Missing 'go' dependency"
    exit
fi

if [ -f /usr/bin/gobash ]
then
    echo "There seems to be an already existing Gobash installation in /usr/bin/gobash."
    sudo rm $exec_install_path/gobash
fi

go build .
sudo mv ./gobash $exec_install_path

# Configuration build
config_path=/home/$user/.config/gobash
if [ ! -f $config_path ]
then
    echo "Creating configuration file..."
    touch $config_path
    echo "Created at ${config_path}."
else
    echo "Configuration file detected."
fi