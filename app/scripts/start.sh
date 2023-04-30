#!/bin/sh

#exit immediately if any command exits with a non-zero status.
set -e 

echo "start the app"

# presents all the arguments passed to the script.
exec "$@"
