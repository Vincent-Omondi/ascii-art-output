#!/bin/bash

# Add all changes
git add .

# Commit changes
git commit -m "add: commit scripts"

# Push to Gitea
git push origin master

# Push to GitHub
git push https://github.com/Vincent-Omondi/ascii-art-output.git master
