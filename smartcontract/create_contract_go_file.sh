#!/usr/bin/env bash

$(which solc) --abi contracts/Lillybox.sol --include-path node_modules --base-path . -o build --overwrite
$(which abigen) --abi build/Lillybox.abi --pkg main --type Lillybox --out Lillybox.go
