#!/bin/bash
find . -type f -o -type d | wc -l | sed -e 's/[ \t]*//'
