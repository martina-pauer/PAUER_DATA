#!/usr/bin/bash
# Variables sys y arch para compilar a varios sistemas y arquitecturas multiples archivos en go
<<<<<<< HEAD
sys="linux"
arch="386"
prefix=".."
extension=""
compilados="categoric$extension"
ejecutables="Linux/32-bits/"
=======
sys="Linux"
arch="amd64"
extension=""
ejecutables="Linux/amd64/"
compilados="categoric$extension"
prefix=".."
compilar"GOOS=$sys GOARCH=$arch go build $prefix"
>>>>>>> 550a5efd122ea4da200e789bb2615704d3d5c5d8
# Compilo cada uno de los archivos
$compilar/categoric.go
# Muevo ejecutables creados a su resppectiva carpeta
mv $prefix/$compilados -t $ejecutables
