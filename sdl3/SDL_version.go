package sdl3

import (
	"unsafe"

	"github.com/moonfdd/ffmpeg-go/ffcommon"
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
 *  \file SDL_version.h
 *
 *  This header defines the current SDL version.
 */

// #ifndef SDL_version_h_
// #define SDL_version_h_

// #include <SDL3/SDL_stdinc.h>

// #include <SDL3/SDL_begin_code.h>
// /* Set up for C function definitions, even when using C++ */
// #ifdef __cplusplus
// extern "C" {
// #endif

/**
 * Information about the version of SDL in use.
 *
 * Represents the library's version as three levels: major revision
 * (increments with massive changes, additions, and enhancements),
 * minor revision (increments with backwards-compatible changes to the
 * major revision), and patchlevel (increments with fixes to the minor
 * revision).
 *
 * \sa SDL_VERSION
 * \sa SDL_GetVersion
 */
type SDL_version struct {
	Major ffcommon.FUint8T /**< major version */
	Minor ffcommon.FUint8T /**< minor version */
	Patch ffcommon.FUint8T /**< update version */
}

/* Printable format: "%d.%d.%d", MAJOR, MINOR, PATCHLEVEL
 */
const SDL_MAJOR_VERSION = 3
const SDL_MINOR_VERSION = 0
const SDL_PATCHLEVEL = 0

/**
 * Macro to determine SDL version program was compiled against.
 *
 * This macro fills in a SDL_version structure with the version of the
 * library you compiled against. This is determined by what header the
 * compiler uses. Note that if you dynamically linked the library, you might
 * have a slightly newer or older version at runtime. That version can be
 * determined with SDL_GetVersion(), which, unlike SDL_VERSION(),
 * is not a macro.
 *
 * \param x A pointer to a SDL_version struct to initialize.
 *
 * \sa SDL_version
 * \sa SDL_GetVersion
 */
// #define SDL_VERSION(x)                          \
// {                                   \
//     (x)->major = SDL_MAJOR_VERSION;                 \
//     (x)->minor = SDL_MINOR_VERSION;                 \
//     (x)->patch = SDL_PATCHLEVEL;                    \
// }

// /**
//  *  This macro turns the version numbers into a numeric value:
//  *  \verbatim
//     (1,2,3) -> (0x1000203)
//     \endverbatim
//  */
// #define SDL_VERSIONNUM(X, Y, Z) \
//     ((X) << 24 | (Y) << 8 | (Z) << 0)

// /**
//  *  This is the version number macro for the current SDL version.
//  */
// #define SDL_COMPILEDVERSION \
//     SDL_VERSIONNUM(SDL_MAJOR_VERSION, SDL_MINOR_VERSION, SDL_PATCHLEVEL)

// /**
//  *  This macro will evaluate to true if compiled with SDL at least X.Y.Z.
//  */
// #define SDL_VERSION_ATLEAST(X, Y, Z) \
//     (SDL_COMPILEDVERSION >= SDL_VERSIONNUM(X, Y, Z))

/**
 * Get the version of SDL that is linked against your program.
 *
 * If you are linking to SDL dynamically, then it is possible that the current
 * version will be different than the version you compiled against. This
 * function returns the current version, while SDL_VERSION() is a macro that
 * tells you what version you compiled with.
 *
 * This function may be called safely at any time, even before SDL_Init().
 *
 * \param ver the SDL_version structure that contains the version information
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetRevision
 */
// extern DECLSPEC int SDLCALL SDL_GetVersion(SDL_version * ver);
func (ver *SDL_version) SDL_GetVersion() (res ffcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetVersion").Call(
		uintptr(unsafe.Pointer(ver)),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Get the code revision of SDL that is linked against your program.
 *
 * This value is the revision of the code you are linked with and may be
 * different from the code you are compiling with, which is found in the
 * constant SDL_REVISION.
 *
 * The revision is arbitrary string (a hash value) uniquely identifying the
 * exact revision of the SDL library in use, and is only useful in comparing
 * against other revisions. It is NOT an incrementing number.
 *
 * If SDL wasn't built from a git repository with the appropriate tools, this
 * will return an empty string.
 *
 * Prior to SDL 2.0.16, before development moved to GitHub, this returned a
 * hash for a Mercurial repository.
 *
 * You shouldn't use this function for anything but logging it for debugging
 * purposes. The string is not intended to be reliable in any way.
 *
 * \returns an arbitrary string, uniquely identifying the exact revision of
 *          the SDL library in use.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetVersion
 */
// extern DECLSPEC const char *SDLCALL SDL_GetRevision(void);
func SDL_GetRevision() (res ffcommon.FConstCharP) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetRevision").Call()
	res = ffcommon.StringFromPtr(t)
	return
}

/* Ends C function definitions when using C++ */
// #ifdef __cplusplus
// }
// #endif
// #include <SDL3/SDL_close_code.h>

// #endif /* SDL_version_h_ */
