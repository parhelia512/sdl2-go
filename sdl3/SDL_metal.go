package sdl3

import (
	"unsafe"

	"github.com/moonfdd/sdl2-go/sdlcommon"
)

/*
  Simple DirectMedia Layer
  Copyright (C) 1997-2023 Sam Lantinga <slouken@libsdl.org>

  This software is provided 'as-is', without any express or implied
  warranty.  In no event will the authors be held liable for any damages
  arising from the use of this software.

  Permission is granted to anyone to use this software for any purpose,
  including commercial applications, and to alter it and redistribute it
  freely, subject to the following restrictions:

  1. The origin of this software must not be misrepresented; you must not
     claim that you wrote the original software. If you use this software
     in a product, an acknowledgment in the product documentation would be
     appreciated but is not required.
  2. Altered source versions must be plainly marked as such, and must not be
     misrepresented as being the original software.
  3. This notice may not be removed or altered from any source distribution.
*/

/**
 *  \file SDL_metal.h
 *
 *  Header file for functions to creating Metal layers and views on SDL windows.
 */

// #ifndef SDL_metal_h_
// #define SDL_metal_h_

// #include <SDL3/SDL_video.h>

// #include <SDL3/SDL_begin_code.h>
// /* Set up for C function definitions, even when using C++ */
// #ifdef __cplusplus
// extern "C" {
// #endif

/**
 *  \brief A handle to a CAMetalLayer-backed NSView (macOS) or UIView (iOS/tvOS).
 *
 *  \note This can be cast directly to an NSView or UIView.
 */
// typedef void *SDL_MetalView;
type SDL_MetalView = sdlcommon.FVoidP

/**
 *  \name Metal support functions
 */
/* @{ */

/**
 * Create a CAMetalLayer-backed NSView/UIView and attach it to the specified
 * window.
 *
 * On macOS, this does *not* associate a MTLDevice with the CAMetalLayer on
 * its own. It is up to user code to do that.
 *
 * The returned handle can be casted directly to a NSView or UIView. To access
 * the backing CAMetalLayer, call SDL_Metal_GetLayer().
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_Metal_DestroyView
 * \sa SDL_Metal_GetLayer
 */
// extern DECLSPEC SDL_MetalView SDLCALL SDL_Metal_CreateView(SDL_Window * window);
func (window *SDL_Window) SDL_Metal_CreateView() (res SDL_MetalView) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_Metal_CreateView").Call(
		uintptr(unsafe.Pointer(window)),
	)
	res = t
	return
}

/**
 * Destroy an existing SDL_MetalView object.
 *
 * This should be called before SDL_DestroyWindow, if SDL_Metal_CreateView was
 * called after SDL_CreateWindow.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_Metal_CreateView
 */
// extern DECLSPEC void SDLCALL SDL_Metal_DestroyView(SDL_MetalView view);
func SDL_Metal_DestroyView(view SDL_MetalView) {
	sdlcommon.GetSDL2Dll().NewProc("SDL_Metal_DestroyView").Call(
		view,
	)
}

/**
 * Get a pointer to the backing CAMetalLayer for the given view.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_Metal_CreateView
 */
// extern DECLSPEC void *SDLCALL SDL_Metal_GetLayer(SDL_MetalView view);
func SDL_Metal_GetLayer(view SDL_MetalView) (res sdlcommon.FVoidP) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_Metal_GetLayer").Call(
		view,
	)
	res = t
	return
}

/* @} */ /* Metal support functions */

/* Ends C function definitions when using C++ */
// #ifdef __cplusplus
// }
// #endif
// #include <SDL3/SDL_close_code.h>

// #endif /* SDL_metal_h_ */
