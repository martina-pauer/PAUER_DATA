#!/usr/bin/bash
# Variables sys y arch para compilar a varios sistemas y arquitecturas multiples archivos en go
sys="linux"
arch="386"
prefix=".."
extension=""
compilados="categoric$extension"
ejecutables="Linux/32-bits/"
compilar="GOOS=$sys GOARCH=$arch go build"
# Compilo cada uno de los archivos
$compilar $prefix/categoric.go
# Muevo ejecutables creados a su resppectiva carpeta
mv $compilados -t $ejecutables
