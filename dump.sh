dd if=/proc/21423/mem of=./pokestack.dd skip=824633720832 bs=1 count=4194304
hexdump ./pokestack.dd >./pokestack.hex
