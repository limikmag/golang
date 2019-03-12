#!/bin/bash

function error() {
    echo "ERROR: $1"
    exit 1
}

function usage() {
    echo -e "\nUsage:\n"
    echo "sudo ./run_from_scratch.sh [-d] [-g] [-s] [-c] [-h]"
    echo "   -d   install docker"
    echo "   -g   install git"
    echo "   -s   compile moby project"
    echo "   -c   clean environment"
    echo -e "   -h   print usage\n"
    echo -e "This script was only tested on CentOS 7.6\n"
    exit 2
}

function install_git() {
    yum install git
}

function remove_git() {
    yum remove git
}

function remove_docker() {
    sudo yum remove docker \
                  docker-ce \
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
    sudo systemctl daemon-reload
    sudo systemctl enable docker
    sudo systemctl start docker
    sudo docker run hello-world
}

# if number of arguments not exceed 2 print usage 
[ $# -lt 1 ] && usage


while getopts ":dgshc" opt; do
  case $opt in
    d) docker=yes;;
    g) git=yes;;
    s) setup=yes;;
    c) clean=yes;;
    h) usage
  esac
done


if [ -n "${git}" ];
then
    install_git
fi

if [ -n "${docker}" ];
then
    install_docker
fi

if [ -n "${setup}" ];
then
    docker info &> /dev/null
    [ $? -ne 0 ] && error "Docker is not installed"
    git status &> /dev/null
    [ $? -ne 0 ] && error "Git is not installed"
    git clone https://github.com/moby/moby
    cd moby
    make
    [ $? -ne 0 ] && error "Something goes wrong while making binaries"
    echo -n "\nBuild succesfully passed!\n"
fi

if [ -n "${clean}" ];
then
    remove_git
    remove_docker
    remove_moby
fi


