//go:build windows
// +build windows

package vulkan

import "unsafe"

func Win32SurfaceCreateInfoKHRHelper(hwnd, hInstance unsafe.Pointer, instance Instance, surface *Surface) Result {
	cinstance, _ := *(*C.VkInstance)(unsafe.Pointer(&instance)), cgoAllocsUnknown
	createInfo := C.VkWin32SurfaceCreateInfoKHR{}
	createInfo.sType = C.VkStructureType(StructureTypeWin32SurfaceCreateInfo)
	createInfo.hwnd = C.HWND(hwnd)
	createInfo.hinstance = C.HINSTANCE(hInstance)
	cSurface, _ := (*C.VkSurfaceKHR)(unsafe.Pointer(surface)), cgoAllocsUnknown
	__ret := C.callVkCreateWin32SurfaceKHR(cinstance, &createInfo, nil, cSurface)
	return (Result)(__ret)
}
