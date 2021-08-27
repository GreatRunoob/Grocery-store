#!/bin/bash

scriptfile="./monitor.py"

if [ -x "/usr/bin/python3" ]
then
	if [ -e $scriptfile ]
	then
		python3 $scriptfile
	else
		echo "The target scrip file '$scriptfile' is not exist!"
	fi
else
	echo "The python3 runtime enviroment is not exist!"
fi

