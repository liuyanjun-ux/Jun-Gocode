#!/usr/bin/env bash
# dep ensure

# echo "Compiling functions to bin/ ..."

# rm -rf bin/

# cd handlers/
# for f in *; do
#   if GOOS=linux go build -ldflags="-s -w" -o "../bin/${f}" ../handlers/${f}/main.go; then
#     echo "✓ Compiled ${f}"
#   else
#     echo "✕ Failed to compile ${f}!"
#     exit 1
#   fi
# done

# echo "Done."