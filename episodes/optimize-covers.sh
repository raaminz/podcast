#!/bin/bash

find . -name "cover.jpg" -exec bash -c 'jpegoptim --size=300k "$1"' _ {} \;
