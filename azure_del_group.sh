#!/bin/bash

azLogin=k.lazar@dtpoland.com
groupName=klazar_rg
line="======================================="

echo $line
echo "            AZURE LOGIN"
echo $line
echo "Please enter Azure password:"
read -s azPassword
az login -u $azLogin -p $azPassword
echo

echo $line
echo "  REMOVING REGISTRY GROUP: $groupName"
echo $line
az group delete --name $groupName --yes
echo

echo $line
echo -e "          THE END\a"
echo $line
echo