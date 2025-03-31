#!/usr/bin/bash
# Variables sys y arch para compilar a varios sistemas y arquitecturas multiples archivos en go
sys="windows"
arch="386"
prefix=".."
extension=".exe"
compilados="categoric.go"
ejecutables="windows/32-bits/"
# Compilo cada uno de los archivos
GOOS=$sys GOARCH=$arch go build $prefix/categoric.go
# Muevo ejecutables creados a su resppectiva carpeta
mv $compilados -t $ejecutables
