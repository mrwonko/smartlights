package benchmark

import (
	"bytes"
	"encoding/binary"
	"os"
	"reflect"
	"syscall"
	"unsafe"
)

const (
	memLength       = 4096
	bcm2835BaseAddr = 0x20000000
	gpioOffset      = 0x200000
)

var (
	gpioBaseAddr = getBaseAddr() + gpioOffset
)

type gpio struct {
	mem      []uint32
	mem8     []uint8
	closeErr error
}

func newGPIO() (*gpio, error) {
	file, err := os.OpenFile("/dev/gpiomem", os.O_RDWR|os.O_SYNC, 0)
	if err != nil {
		return nil, err
	}
	// FD can be closed after memory mapping
	defer file.Close()

	var res gpio
	res.mem, res.mem8, err = memMap(file.Fd(), gpioBaseAddr)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// WritePin sets a given pin High or Low
// by setting the clear or set registers respectively
func (g *gpio) WritePin(p uint8, high bool) {

	if high {
		// Clear register, 10 / 11 depending on bank
		clearReg := p/32 + 10
		g.mem[clearReg] = 1 << (p & 31)
	} else {
		// Set register, 7 / 8 depending on bank
		setReg := p/32 + 7
		g.mem[setReg] = 1 << (p & 31)
	}
}

// WritePins sets a given set of pins High or Low
// by setting the clear or set registers respectively
func (g *gpio) WritePins(mask uint64, high bool) {
	loMask := uint32(mask & 0xFFFFFFFF)
	hiMask := uint32((mask >> 32) & 0xFFFFFFFF)
	if high {
		g.mem[10] = loMask
		g.mem[11] = hiMask
	} else {
		g.mem[7] = uint32(mask & 0xFFFFFFFF)
		g.mem[8] = uint32((mask >> 32) & 0xFFFFFFFF)
	}
}

func (g *gpio) Close() error {
	if g.mem == nil {
		return g.closeErr
	}
	g.closeErr = syscall.Munmap(g.mem8)
	g.mem = nil
	g.mem8 = nil
	return g.closeErr
}

func memMap(fd uintptr, base int64) (mem []uint32, mem8 []byte, err error) {
	mem8, err = syscall.Mmap(
		int(fd),
		base,
		memLength,
		syscall.PROT_READ|syscall.PROT_WRITE,
		syscall.MAP_SHARED,
	)
	if err != nil {
		return
	}
	// Convert mapped byte memory to unsafe []uint32 pointer, adjust length as needed
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&mem8))
	header.Len /= (32 / 8) // (32 bit = 4 bytes)
	header.Cap /= (32 / 8)
	mem = *(*[]uint32)(unsafe.Pointer(&header))
	return
}

// getBaseAddr reads /proc/device-tree/soc/ranges and determines the base address.
// Uses the default Raspberry Pi 1 base address if this fails.
func getBaseAddr() int64 {
	ranges, err := os.Open("/proc/device-tree/soc/ranges")
	defer ranges.Close()
	if err != nil {
		return bcm2835BaseAddr
	}
	b := make([]byte, 4)
	n, err := ranges.ReadAt(b, 4)
	if n != 4 || err != nil {
		return bcm2835BaseAddr
	}
	buf := bytes.NewReader(b)
	var out uint32
	err = binary.Read(buf, binary.BigEndian, &out)
	if err != nil {
		return bcm2835BaseAddr
	}
	return int64(out)
}
