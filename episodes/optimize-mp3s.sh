#!/bin/bash

find -type f -iname "*.mp3" -exec sh -c 'ffmpeg -i "$0" -codec:a libmp3lame -b:a 64k "${0%.*}2.mp3"' {} \;

find . -name "content2.mp3" -exec bash -c 'mv "$1" "${1/content2/content}"' _ {} \;


