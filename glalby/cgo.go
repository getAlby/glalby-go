package glalby

/*
#cgo LDFLAGS: -lglalby_bindings
#cgo linux,amd64 LDFLAGS: -Wl,-rpath,${SRCDIR}/x86_64-unknown-linux-gnu -L${SRCDIR}/x86_64-unknown-linux-gnu
#cgo darwin LDFLAGS: -Wl,-rpath,${SRCDIR}/universal-macos -L${SRCDIR}/universal-macos
#cgo windows,amd64 LDFLAGS: -Wl,-rpath,${SRCDIR}/x86_64-pc-windows-msvc -L${SRCDIR}/x86_64-pc-windows-msvc
#cgo linux,arm.6 LDFLAGS: -Wl,-rpath,${SRCDIR}/arm-unknown-linux-gnueabihf -L${SRCDIR}/arm-unknown-linux-gnueabihf
*/
import "C"
