#!/usr/bin/bash
# Variables sys y arch para compilar a varios sistemas y arquitecturas multiples archivos en go
sys="Linux"
arch="amd64"
extension=""
ejecutables="Linux/amd64/"
compilados="categoric.go"
prefix=".."
compilar"GOOS=$sys GOARCH=$arch go build $prefix"
# Compilo cada uno de los archivos
$compilar/categoric.go
# Muevo ejecutables creados a su resppectiva carpeta
mv $prefix/$compilados -t $ejecutables
