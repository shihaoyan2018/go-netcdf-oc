// netcdf_group
package netcdf_group

// #cgo LDFLAGS: -lnetcdf
// #include<stdlib.h>
// #include<netcdf.h>
import "C"
import (
	_ "fmt"
	"unsafe"

	"github.com/fhs/go-netcdf/netcdf"
)

type Dataset C.int
type Group struct {
	ds Dataset
	id C.int
}
type Var struct {
	gr Dataset
	id C.int
}

// inquire dataset'groups ids;
func (ds Dataset) Groupn(gpn []int32, gps []int32) {
	C.nc_inq_grps(C.int(ds), (*C.int)(unsafe.Pointer(&gpn[0])), (*C.int)(unsafe.Pointer(&gps[0])))
}
func (ds Dataset) Groupbin(ids int32, gps []int32) (gp Group) {
	id := C.int(gps[ids])
	gp = Group{ds, id}
	return gp
}
func (gp Group) NVar(vrn []int32, vrs []int32) {

	C.nc_inq_varids(C.int(gp.id), (*C.int)(unsafe.Pointer(&vrn[0])), (*C.int)(unsafe.Pointer(&vrs[0])))
}
func (gp Group) Varn(ids int32, vrs []int32) (v netcdf.Var) {
	id := C.int(vrs[ids])
	vv := Var{Dataset(gp.id), id}
	v = *(*netcdf.Var)(unsafe.Pointer(&vv))
	return v
}
