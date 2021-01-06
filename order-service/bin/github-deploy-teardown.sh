#!/bin/bash

if [ "$GITHUB_DEPLOY_KEY" != "" ]; then
    echo "Removing deploy key and config"

    # Now that npm has finished running, we shouldn't need the ssh key/config anymore.
    # Remove the files that we created.
    rm -f ~/.ssh/config
    rm -f ~/.ssh/deploy_key

    # Clear that sensitive key data from the environment
    export GITHUB_DEPLOY_KEY=0
fi
