#!/bin/bash

BINDIR="./bin"

# Build binaries
gox

# Create bindir if it doesn't exist
test -d ${BINDIR} || mkdir -p ${BINDIR}

# Move binaries to subdirectory
mv GoKurac_*_* ${BINDIR}
