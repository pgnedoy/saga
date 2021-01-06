#!/bin/bash
echo GITHUB_DEPLOY_KEY
if [ "$GITHUB_DEPLOY_KEY" != "" ]; then
    echo "Adding deploy key and config"

    # Ensure we have an ssh folder
    if [ ! -d ~/.ssh ]; then
      mkdir -p ~/.ssh
      chmod 700 ~/.ssh
    fi

    # Load the private key into a file.
    echo "$GITHUB_DEPLOY_KEY" > ~/.ssh/deploy_key

    # Change the permissions on the file to
    # be read-only for this user.
    chmod 400 ~/.ssh/deploy_key

    # Setup the ssh config file.
    echo -e "Host github.com\n" \
      " IdentityFile ~/.ssh/deploy_key\n" \
      " IdentitiesOnly yes\n" \
      " UserKnownHostsFile=/dev/null\n" \
      " StrictHostKeyChecking no" \
      > ~/.ssh/config
fi
