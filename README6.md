##  在使用cgo时，编译中gcc和cgo的编译缓存问题

###  在使用cgo时，编译中gcc和cgo的编译生命周期不同

gcc和cgo编译缓存不同，cgo的缓存允许使用go clean清除，但是无法清除gcc缓存。gcc缓存的清除考虑改变gcc编译的选项来清除。
