package sdl

import "github.com/moonfdd/sdl2-go/common"

/**
 * Dynamically load a shared object.
 *
 * \param sofile a system-dependent name of the object file
 * \returns an opaque pointer to the object handle or NULL if there was an
 *          error; call SDL_GetError() for more information.
 *
 * \sa SDL_LoadFunction
 * \sa SDL_UnloadObject
 */
//extern DECLSPEC void *SDLCALL SDL_LoadObject(const char *sofile);
func SDL_LoadObject(sofile common.FConstCharP) (res common.FConstCharP) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_LoadObject").Call(
		common.UintPtrFromString(sofile),
	)
	if t == 0 {

	}
	res = common.StringFromPtr(t)
	return
}

/**
 * Look up the address of the named function in a shared object.
 *
 * This function pointer is no longer valid after calling SDL_UnloadObject().
 *
 * This function can only look up C function names. Other languages may have
 * name mangling and intrinsic language support that varies from compiler to
 * compiler.
 *
 * Make sure you declare your function pointers with the same calling
 * convention as the actual library function. Your code will crash
 * mysteriously if you do not do this.
 *
 * If the requested function doesn't exist, NULL is returned.
 *
 * \param handle a valid shared object handle returned by SDL_LoadObject()
 * \param name the name of the function to look up
 * \returns a pointer to the function or NULL if there was an error; call
 *          SDL_GetError() for more information.
 *
 * \sa SDL_LoadObject
 * \sa SDL_UnloadObject
 */
//extern DECLSPEC void *SDLCALL SDL_LoadFunction(void *handle,
//const char *name);
func SDL_LoadFunction(handle common.FVoidP, name common.FConstCharP) (res common.FConstCharP) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_LoadFunction").Call(
		handle,
		common.UintPtrFromString(name),
	)
	if t == 0 {

	}
	res = common.StringFromPtr(t)
	return
}

/**
 * Unload a shared object from memory.
 *
 * \param handle a valid shared object handle returned by SDL_LoadObject()
 *
 * \sa SDL_LoadFunction
 * \sa SDL_LoadObject
 */
//extern DECLSPEC void SDLCALL SDL_UnloadObject(void *handle);
func SDL_UnloadObject(p common.FVoidP) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_UnloadObject").Call(
		p,
	)
	if t == 0 {

	}
	return
}
