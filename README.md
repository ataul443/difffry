# Diffry

## Overview
It is an basic implementation of Meyers diff O(ND) variation algorithm.
Give it two strings and it will tell you what changes need to be done in
one string to get the other string.

Suppose you gave it two strings `pika` and `copapika`. It will tell you 
in how many add, replace, delete operations it will transform the `pika`
`copika`.

## Why I build it ?
I wanted to learn about diff algorithms. In our daily life we always use them
while comparing changes in version control like git. I was curious how these
algorithms works and hence I tried implementing it.

## Usage
```shell script
git clone https://github.com/ataul443/diffry
cd diffry
go build .
./diffry
```

### Note

It is work in progress and needs refinement in presentation logic.