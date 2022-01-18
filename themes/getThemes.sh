#!/bin/bash
# Downloads every file from sources.txt
xargs -I {} wget {} <sources.txt
