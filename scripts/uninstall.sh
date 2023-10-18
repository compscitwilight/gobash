#!/bin/bash
# This script uninstalls the current Gobash installation.

if [ ! -f /usr/bin/gobash ]
then
    echo "There is not a Gobash installation in your path."
    exit
fi

sudo rm /usr/bin/gobash
echo "Successfully removed Gobash installation."