//go:build linux || darwin

package vulkan

/*
#cgo CFLAGS: -DVK_USE_PLATFORM_XLIB_KHR

#include "vk_wrapper.h"
#include "vk_bridge.h"
#include "vulkan/vulkan_xlib.h"
*/
import "C"
import "unsafe"

func XlibSurfaceCreateInfoKHRHelper(window, display unsafe.Pointer, instance Instance, surface *Surface) Result {
	cinstance, _ := *(*C.VkInstance)(unsafe.Pointer(&instance)), cgoAllocsUnknown
	createInfo := C.VkXlibSurfaceCreateInfoKHR{}
	createInfo.sType = C.VkStructureType(StructureTypeXlibSurfaceCreateInfo)
	println(uintptr(display), uintptr(window))
	createInfo.dpy = (*C.Display)(display)
	createInfo.window = *(*C.Window)(window)
	cSurface, _ := (*C.VkSurfaceKHR)(unsafe.Pointer(surface)), cgoAllocsUnknown
	__ret := C.callVkCreateXlibSurfaceKHR(cinstance, &createInfo, nil, cSurface)
	return (Result)(__ret)
}
