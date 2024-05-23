#!/bin/bash
find . -type f -name "*.sh" -print | cut -d '/' -f 2- | rev | cut -d '.' -f 2- | cut -d '/' -f -1 | rev | sort -nr
