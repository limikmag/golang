#!/bin/bash

function error() {
    echo "ERROR: $1"
    exit 1
}

function usage() {
    echo "Usage:"
    echo ""
    echo "./run_from_scratch.sh [-d] [-g] [-s]"
    echo "-d   install docker"
    echo "-g   install git"
    echo "-s   compile moby project"
    echo "-c   clean environment"
    echo "-h   print usage"
    echo ""
    echo "This script was only tested on CentOS 7.6"
}

function install_git() {
    yum install git
}

function remove_git() {
    yum remove git
}

function remove_docker() {
    sudo yum remove docker \
                  docker-client \
                  docker-client-latest \
                  docker-common \
                  docker-latest \
                  docker-latest-logrotate \
                  docker-logrotate \
                  docker-engine

    rm -rf /var/lib/docker
}

function remove_moby() {
    rm -rf moby
}

function install_docker() {
    remove_docker
    sudo yum install -y yum-utils \
    device-mapper-persistent-data \
    lvm2

    sudo yum-config-manager \
    --add-repo \
    https://download.docker.com/linux/centos/docker-ce.repo
    sudo yum install docker-ce docker-ce-cli containerd.io
    sudo systemctl start docker
    sudo docker run hello-world
}

while getopts ":dgshc" opt; do
  case $opt in
    d)docker=yes;;
    g)git=yes;;
    s)setup=yes;;
    c)clean=yes;;
    h) usage
    \?) echo "Invalid option: -$OPTARG" >&2;;
  esac
done


if [ -n ${git} ];
then
    install_git
fi

if [ -n $docker ];
then
    install_docker
fi

if [ -n ${setup} ];
then
    git clone https://github.com/moby/moby
    cd moby
    make
fi

if [ -n ${clean} ];
then
    remove_git
    remove_docker
    remove_moby
fi


