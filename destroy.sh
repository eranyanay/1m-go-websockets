#!/bin/bash
## This script removes *all* containers, even those that wasn't spawned using connect script
## Use with caution

docker rm -vf $(docker ps -q)