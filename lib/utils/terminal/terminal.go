package terminal

import (
	"os"
	"syscall"
	"unsafe"
)

func GetDimensions() (int, int) {
	out, err := os.OpenFile("/dev/tty", syscall.O_WRONLY, 0)
	if err != nil {
		return 0, 0
	}
	defer out.Close()

	var sz winsize
	_, _, _ = syscall.Syscall(syscall.SYS_IOCTL,
		out.Fd(), uintptr(syscall.TIOCGWINSZ), uintptr(unsafe.Pointer(&sz)))

	return int(sz.cols), int(sz.rows)
}

type winsize struct {
	rows    uint16
	cols    uint16
	xpixels uint16
	ypixels uint16
}
