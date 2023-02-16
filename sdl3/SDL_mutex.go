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

// #ifndef SDL_mutex_h_
// #define SDL_mutex_h_

// /**
//  *  \file SDL_mutex.h
//  *
//  *  Functions to provide thread synchronization primitives.
//  */

// #include <SDL3/SDL_stdinc.h>
// #include <SDL3/SDL_error.h>

// /******************************************************************************/
// /* Enable thread safety attributes only with clang.
//  * The attributes can be safely erased when compiling with other compilers.
//  */
// #if defined(SDL_THREAD_SAFETY_ANALYSIS) && \
//     defined(__clang__) && (!defined(SWIG))
// #define SDL_THREAD_ANNOTATION_ATTRIBUTE__(x)   __attribute__((x))
// #else
// #define SDL_THREAD_ANNOTATION_ATTRIBUTE__(x)   /* no-op */
// #endif

// #define SDL_CAPABILITY(x) \
//   SDL_THREAD_ANNOTATION_ATTRIBUTE__(capability(x))

// #define SDL_SCOPED_CAPABILITY \
//   SDL_THREAD_ANNOTATION_ATTRIBUTE__(scoped_lockable)

// #define SDL_GUARDED_BY(x) \
//   SDL_THREAD_ANNOTATION_ATTRIBUTE__(guarded_by(x))

// #define SDL_PT_GUARDED_BY(x) \
//   SDL_THREAD_ANNOTATION_ATTRIBUTE__(pt_guarded_by(x))

// #define SDL_ACQUIRED_BEFORE(...) \
//   SDL_THREAD_ANNOTATION_ATTRIBUTE__(acquired_before(__VA_ARGS__))

// #define SDL_ACQUIRED_AFTER(...) \
//   SDL_THREAD_ANNOTATION_ATTRIBUTE__(acquired_after(__VA_ARGS__))

// #define SDL_REQUIRES(...) \
//   SDL_THREAD_ANNOTATION_ATTRIBUTE__(requires_capability(__VA_ARGS__))

// #define SDL_REQUIRES_SHARED(...) \
//   SDL_THREAD_ANNOTATION_ATTRIBUTE__(requires_shared_capability(__VA_ARGS__))

// #define SDL_ACQUIRE(...) \
//   SDL_THREAD_ANNOTATION_ATTRIBUTE__(acquire_capability(__VA_ARGS__))

// #define SDL_ACQUIRE_SHARED(...) \
//   SDL_THREAD_ANNOTATION_ATTRIBUTE__(acquire_shared_capability(__VA_ARGS__))

// #define SDL_RELEASE(...) \
//   SDL_THREAD_ANNOTATION_ATTRIBUTE__(release_capability(__VA_ARGS__))

// #define SDL_RELEASE_SHARED(...) \
//   SDL_THREAD_ANNOTATION_ATTRIBUTE__(release_shared_capability(__VA_ARGS__))

// #define SDL_RELEASE_GENERIC(...) \
//   SDL_THREAD_ANNOTATION_ATTRIBUTE__(release_generic_capability(__VA_ARGS__))

// #define SDL_TRY_ACQUIRE(...) \
//   SDL_THREAD_ANNOTATION_ATTRIBUTE__(try_acquire_capability(__VA_ARGS__))

// #define SDL_TRY_ACQUIRE_SHARED(...) \
//   SDL_THREAD_ANNOTATION_ATTRIBUTE__(try_acquire_shared_capability(__VA_ARGS__))

// #define SDL_EXCLUDES(...) \
//   SDL_THREAD_ANNOTATION_ATTRIBUTE__(locks_excluded(__VA_ARGS__))

// #define SDL_ASSERT_CAPABILITY(x) \
//   SDL_THREAD_ANNOTATION_ATTRIBUTE__(assert_capability(x))

// #define SDL_ASSERT_SHARED_CAPABILITY(x) \
//   SDL_THREAD_ANNOTATION_ATTRIBUTE__(assert_shared_capability(x))

// #define SDL_RETURN_CAPABILITY(x) \
//   SDL_THREAD_ANNOTATION_ATTRIBUTE__(lock_returned(x))

// #define SDL_NO_THREAD_SAFETY_ANALYSIS \
//   SDL_THREAD_ANNOTATION_ATTRIBUTE__(no_thread_safety_analysis)

// /******************************************************************************/

// #include <SDL3/SDL_begin_code.h>
// /* Set up for C function definitions, even when using C++ */
// #ifdef __cplusplus
// extern "C" {
// #endif

/**
 *  Synchronization functions which can time out return this value
 *  if they time out.
 */
const SDL_MUTEX_TIMEDOUT = 1

/**
 *  This is the timeout value which corresponds to never time out.
 */
const SDL_MUTEX_MAXWAIT = -1

/**
 *  \name Mutex functions
 */
/* @{ */

/* The SDL mutex structure, defined in SDL_sysmutex.c */
// struct SDL_mutex;
// typedef struct SDL_mutex SDL_mutex;
type SDL_mutex struct {
	// 不可见
}

/**
 * Create a new mutex.
 *
 * All newly-created mutexes begin in the _unlocked_ state.
 *
 * Calls to SDL_LockMutex() will not return while the mutex is locked by
 * another thread. See SDL_TryLockMutex() to attempt to lock without blocking.
 *
 * SDL mutexes are reentrant.
 *
 * \returns the initialized and unlocked mutex or NULL on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_DestroyMutex
 * \sa SDL_LockMutex
 * \sa SDL_TryLockMutex
 * \sa SDL_UnlockMutex
 */
// extern DECLSPEC SDL_mutex *SDLCALL SDL_CreateMutex(void);
func SDL_CreateMutex() (res *SDL_mutex) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_CreateMutex").Call()
	res = (*SDL_mutex)(unsafe.Pointer(t))
	return
}

/**
 * Lock the mutex.
 *
 * This will block until the mutex is available, which is to say it is in the
 * unlocked state and the OS has chosen the caller as the next thread to lock
 * it. Of all threads waiting to lock the mutex, only one may do so at a time.
 *
 * It is legal for the owning thread to lock an already-locked mutex. It must
 * unlock it the same number of times before it is actually made available for
 * other threads in the system (this is known as a "recursive mutex").
 *
 * \param mutex the mutex to lock
 * \return 0, or -1 on error.
 *
 * \since This function is available since SDL 3.0.0.
 */
// extern DECLSPEC int SDLCALL SDL_LockMutex(SDL_mutex * mutex) SDL_ACQUIRE(mutex);
// #define SDL_mutexP(m)   SDL_LockMutex(m)
func (mutex *SDL_mutex) SDL_LockMutex() (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_CreateMutex").Call(
		uintptr(unsafe.Pointer(mutex)),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Try to lock a mutex without blocking.
 *
 * This works just like SDL_LockMutex(), but if the mutex is not available,
 * this function returns `SDL_MUTEX_TIMEOUT` immediately.
 *
 * This technique is useful if you need exclusive access to a resource but
 * don't want to wait for it, and will return to it to try again later.
 *
 * \param mutex the mutex to try to lock
 * \returns 0, `SDL_MUTEX_TIMEDOUT`, or -1 on error; call SDL_GetError() for
 *          more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_CreateMutex
 * \sa SDL_DestroyMutex
 * \sa SDL_LockMutex
 * \sa SDL_UnlockMutex
 */
// extern DECLSPEC int SDLCALL SDL_TryLockMutex(SDL_mutex * mutex) SDL_TRY_ACQUIRE(0, mutex);
func (mutex *SDL_mutex) SDL_TryLockMutex() (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_TryLockMutex").Call(
		uintptr(unsafe.Pointer(mutex)),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Unlock the mutex.
 *
 * It is legal for the owning thread to lock an already-locked mutex. It must
 * unlock it the same number of times before it is actually made available for
 * other threads in the system (this is known as a "recursive mutex").
 *
 * It is an error to unlock a mutex that has not been locked by the current
 * thread, and doing so results in undefined behavior.
 *
 * It is also an error to unlock a mutex that isn't locked at all.
 *
 * \param mutex the mutex to unlock.
 * \returns 0, or -1 on error.
 *
 * \since This function is available since SDL 3.0.0.
 */
// extern DECLSPEC int SDLCALL SDL_UnlockMutex(SDL_mutex * mutex) SDL_RELEASE(mutex);
// #define SDL_mutexV(m)   SDL_UnlockMutex(m)
func (mutex *SDL_mutex) SDL_UnlockMutex() (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_UnlockMutex").Call(
		uintptr(unsafe.Pointer(mutex)),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Destroy a mutex created with SDL_CreateMutex().
 *
 * This function must be called on any mutex that is no longer needed. Failure
 * to destroy a mutex will result in a system memory or resource leak. While
 * it is safe to destroy a mutex that is _unlocked_, it is not safe to attempt
 * to destroy a locked mutex, and may result in undefined behavior depending
 * on the platform.
 *
 * \param mutex the mutex to destroy
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_CreateMutex
 * \sa SDL_LockMutex
 * \sa SDL_TryLockMutex
 * \sa SDL_UnlockMutex
 */
// extern DECLSPEC void SDLCALL SDL_DestroyMutex(SDL_mutex * mutex);
func (mutex *SDL_mutex) SDL_DestroyMutex() {
	sdlcommon.GetSDL2Dll().NewProc("SDL_DestroyMutex").Call(
		uintptr(unsafe.Pointer(mutex)),
	)
}

/* @} */ /* Mutex functions */

/**
 *  \name Semaphore functions
 */
/* @{ */

/* The SDL semaphore structure, defined in SDL_syssem.c */
// struct SDL_semaphore;
// typedef struct SDL_semaphore SDL_sem;
type SDL_semaphore struct {
}
type SDL_sem = SDL_semaphore

/**
 * Create a semaphore.
 *
 * This function creates a new semaphore and initializes it with the value
 * `initial_value`. Each wait operation on the semaphore will atomically
 * decrement the semaphore value and potentially block if the semaphore value
 * is 0. Each post operation will atomically increment the semaphore value and
 * wake waiting threads and allow them to retry the wait operation.
 *
 * \param initial_value the starting value of the semaphore
 * \returns a new semaphore or NULL on failure; call SDL_GetError() for more
 *          information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_DestroySemaphore
 * \sa SDL_SemPost
 * \sa SDL_SemTryWait
 * \sa SDL_SemValue
 * \sa SDL_SemWait
 * \sa SDL_SemWaitTimeout
 */
// extern DECLSPEC SDL_sem *SDLCALL SDL_CreateSemaphore(Uint32 initial_value);
func SDL_CreateSemaphore(initial_value sdlcommon.FUint32T) (res *SDL_sem) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_CreateSemaphore").Call(
		uintptr(initial_value),
	)
	res = (*SDL_sem)(unsafe.Pointer(t))
	return
}

/**
 * Destroy a semaphore.
 *
 * It is not safe to destroy a semaphore if there are threads currently
 * waiting on it.
 *
 * \param sem the semaphore to destroy
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_CreateSemaphore
 * \sa SDL_SemPost
 * \sa SDL_SemTryWait
 * \sa SDL_SemValue
 * \sa SDL_SemWait
 * \sa SDL_SemWaitTimeout
 */
// extern DECLSPEC void SDLCALL SDL_DestroySemaphore(SDL_sem *sem);
func (sem *SDL_sem) SDL_DestroySemaphore() {
	sdlcommon.GetSDL2Dll().NewProc("SDL_DestroySemaphore").Call(
		uintptr(unsafe.Pointer(sem)),
	)
}

/**
 * Wait until a semaphore has a positive value and then decrements it.
 *
 * This function suspends the calling thread until either the semaphore
 * pointed to by `sem` has a positive value or the call is interrupted by a
 * signal or error. If the call is successful it will atomically decrement the
 * semaphore value.
 *
 * This function is the equivalent of calling SDL_SemWaitTimeout() with a time
 * length of `SDL_MUTEX_MAXWAIT`.
 *
 * \param sem the semaphore wait on
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_CreateSemaphore
 * \sa SDL_DestroySemaphore
 * \sa SDL_SemPost
 * \sa SDL_SemTryWait
 * \sa SDL_SemValue
 * \sa SDL_SemWait
 * \sa SDL_SemWaitTimeout
 */
// extern DECLSPEC int SDLCALL SDL_SemWait(SDL_sem *sem);
func (sem *SDL_sem) SDL_SemWait() (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SemWait").Call(
		uintptr(unsafe.Pointer(sem)),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * See if a semaphore has a positive value and decrement it if it does.
 *
 * This function checks to see if the semaphore pointed to by `sem` has a
 * positive value and atomically decrements the semaphore value if it does. If
 * the semaphore doesn't have a positive value, the function immediately
 * returns SDL_MUTEX_TIMEDOUT.
 *
 * \param sem the semaphore to wait on
 * \returns 0 if the wait succeeds, `SDL_MUTEX_TIMEDOUT` if the wait would
 *          block, or a negative error code on failure; call SDL_GetError()
 *          for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_CreateSemaphore
 * \sa SDL_DestroySemaphore
 * \sa SDL_SemPost
 * \sa SDL_SemValue
 * \sa SDL_SemWait
 * \sa SDL_SemWaitTimeout
 */
// extern DECLSPEC int SDLCALL SDL_SemTryWait(SDL_sem *sem);
func (sem *SDL_sem) SDL_SemTryWait() (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SemTryWait").Call(
		uintptr(unsafe.Pointer(sem)),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Wait until a semaphore has a positive value and then decrements it.
 *
 * This function suspends the calling thread until either the semaphore
 * pointed to by `sem` has a positive value, the call is interrupted by a
 * signal or error, or the specified time has elapsed. If the call is
 * successful it will atomically decrement the semaphore value.
 *
 * \param sem the semaphore to wait on
 * \param timeoutMS the length of the timeout, in milliseconds
 * \returns 0 if the wait succeeds, `SDL_MUTEX_TIMEDOUT` if the wait does not
 *          succeed in the allotted time, or a negative error code on failure;
 *          call SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_CreateSemaphore
 * \sa SDL_DestroySemaphore
 * \sa SDL_SemPost
 * \sa SDL_SemTryWait
 * \sa SDL_SemValue
 * \sa SDL_SemWait
 */
// extern DECLSPEC int SDLCALL SDL_SemWaitTimeout(SDL_sem *sem, Sint32 timeoutMS);
func (sem *SDL_sem) SDL_SemWaitTimeout(ms sdlcommon.FUint32T) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SemWaitTimeout").Call(
		uintptr(unsafe.Pointer(sem)),
		uintptr(ms),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Atomically increment a semaphore's value and wake waiting threads.
 *
 * \param sem the semaphore to increment
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_CreateSemaphore
 * \sa SDL_DestroySemaphore
 * \sa SDL_SemTryWait
 * \sa SDL_SemValue
 * \sa SDL_SemWait
 * \sa SDL_SemWaitTimeout
 */
// extern DECLSPEC int SDLCALL SDL_SemPost(SDL_sem *sem);
func (sem *SDL_sem) SDL_SemPost() (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SemPost").Call(
		uintptr(unsafe.Pointer(sem)),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Get the current value of a semaphore.
 *
 * \param sem the semaphore to query
 * \returns the current value of the semaphore.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_CreateSemaphore
 */
// extern DECLSPEC Uint32 SDLCALL SDL_SemValue(SDL_sem *sem);
func (sem *SDL_sem) SDL_SemValue() (res sdlcommon.FUint32T) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SemValue").Call(
		uintptr(unsafe.Pointer(sem)),
	)
	res = sdlcommon.FUint32T(t)
	return
}

/* @} */ /* Semaphore functions */

/**
 *  \name Condition variable functions
 */
/* @{ */

/* The SDL condition variable structure, defined in SDL_syscond.c */
// struct SDL_cond;
// typedef struct SDL_cond SDL_cond;
type SDL_cond struct {
}

/**
 * Create a condition variable.
 *
 * \returns a new condition variable or NULL on failure; call SDL_GetError()
 *          for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_CondBroadcast
 * \sa SDL_CondSignal
 * \sa SDL_CondWait
 * \sa SDL_CondWaitTimeout
 * \sa SDL_DestroyCond
 */
// extern DECLSPEC SDL_cond *SDLCALL SDL_CreateCond(void);
func SDL_CreateCond() (res *SDL_cond) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_CreateCond").Call()
	res = (*SDL_cond)(unsafe.Pointer(t))
	return
}

/**
 * Destroy a condition variable.
 *
 * \param cond the condition variable to destroy
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_CondBroadcast
 * \sa SDL_CondSignal
 * \sa SDL_CondWait
 * \sa SDL_CondWaitTimeout
 * \sa SDL_CreateCond
 */
// extern DECLSPEC void SDLCALL SDL_DestroyCond(SDL_cond *cond);
func (cond *SDL_cond) SDL_DestroyCond() {
	sdlcommon.GetSDL2Dll().NewProc("SDL_DestroyCond").Call(
		uintptr(unsafe.Pointer(cond)),
	)
}

/**
 * Restart one of the threads that are waiting on the condition variable.
 *
 * \param cond the condition variable to signal
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_CondBroadcast
 * \sa SDL_CondWait
 * \sa SDL_CondWaitTimeout
 * \sa SDL_CreateCond
 * \sa SDL_DestroyCond
 */
// extern DECLSPEC int SDLCALL SDL_CondSignal(SDL_cond *cond);
func (cond *SDL_cond) SDL_CondSignal() (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_CondSignal").Call(
		uintptr(unsafe.Pointer(cond)),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Restart all threads that are waiting on the condition variable.
 *
 * \param cond the condition variable to signal
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_CondSignal
 * \sa SDL_CondWait
 * \sa SDL_CondWaitTimeout
 * \sa SDL_CreateCond
 * \sa SDL_DestroyCond
 */
// extern DECLSPEC int SDLCALL SDL_CondBroadcast(SDL_cond *cond);
func (cond *SDL_cond) SDL_CondBroadcast() (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_CondBroadcast").Call(
		uintptr(unsafe.Pointer(cond)),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Wait until a condition variable is signaled.
 *
 * This function unlocks the specified `mutex` and waits for another thread to
 * call SDL_CondSignal() or SDL_CondBroadcast() on the condition variable
 * `cond`. Once the condition variable is signaled, the mutex is re-locked and
 * the function returns.
 *
 * The mutex must be locked before calling this function. Locking the mutex
 * recursively (more than once) is not supported and leads to undefined
 * behavior.
 *
 * This function is the equivalent of calling SDL_CondWaitTimeout() with a
 * time length of `SDL_MUTEX_MAXWAIT`.
 *
 * \param cond the condition variable to wait on
 * \param mutex the mutex used to coordinate thread access
 * \returns 0 when it is signaled or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_CondBroadcast
 * \sa SDL_CondSignal
 * \sa SDL_CondWaitTimeout
 * \sa SDL_CreateCond
 * \sa SDL_DestroyCond
 */
// extern DECLSPEC int SDLCALL SDL_CondWait(SDL_cond *cond, SDL_mutex *mutex);
func (cond *SDL_cond) SDL_CondWait(mutex *SDL_mutex) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_CondWait").Call(
		uintptr(unsafe.Pointer(cond)),
		uintptr(unsafe.Pointer(mutex)),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Wait until a condition variable is signaled or a certain time has passed.
 *
 * This function unlocks the specified `mutex` and waits for another thread to
 * call SDL_CondSignal() or SDL_CondBroadcast() on the condition variable
 * `cond`, or for the specified time to elapse. Once the condition variable is
 * signaled or the time elapsed, the mutex is re-locked and the function
 * returns.
 *
 * The mutex must be locked before calling this function. Locking the mutex
 * recursively (more than once) is not supported and leads to undefined
 * behavior.
 *
 * \param cond the condition variable to wait on
 * \param mutex the mutex used to coordinate thread access
 * \param timeoutMS the maximum time to wait, in milliseconds, or
 *                  `SDL_MUTEX_MAXWAIT` to wait indefinitely
 * \returns 0 if the condition variable is signaled, `SDL_MUTEX_TIMEDOUT` if
 *          the condition is not signaled in the allotted time, or a negative
 *          error code on failure; call SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_CondBroadcast
 * \sa SDL_CondSignal
 * \sa SDL_CondWait
 * \sa SDL_CreateCond
 * \sa SDL_DestroyCond
 */
// extern DECLSPEC int SDLCALL SDL_CondWaitTimeout(SDL_cond *cond,
//                                                 SDL_mutex *mutex, Sint32 timeoutMS);
func (cond *SDL_cond) SDL_CondWaitTimeout(mutex *SDL_mutex, ms sdlcommon.FUint32T) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_CondWaitTimeout").Call(
		uintptr(unsafe.Pointer(cond)),
		uintptr(unsafe.Pointer(mutex)),
		uintptr(ms),
	)
	res = sdlcommon.FInt(t)
	return
}

/* @} */ /* Condition variable functions */

/* Ends C function definitions when using C++ */
// #ifdef __cplusplus
// }
// #endif
// #include <SDL3/SDL_close_code.h>

// #endif /* SDL_mutex_h_ */
