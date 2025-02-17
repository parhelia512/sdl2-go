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

/* !!! FIXME: several functions in here need Doxygen comments. */

/**
 *  \file SDL_audio.h
 *
 *  Access to the raw audio mixing buffer for the SDL library.
 */

// #ifndef SDL_audio_h_
// #define SDL_audio_h_

// #include <SDL3/SDL_stdinc.h>
// #include <SDL3/SDL_error.h>
// #include <SDL3/SDL_endian.h>
// #include <SDL3/SDL_mutex.h>
// #include <SDL3/SDL_thread.h>
// #include <SDL3/SDL_rwops.h>

// #include <SDL3/SDL_begin_code.h>
// /* Set up for C function definitions, even when using C++ */
// #ifdef __cplusplus
// extern "C" {
// #endif

/**
*  \brief Audio format flags.
*
*  These are what the 16 bits in SDL_AudioFormat currently mean...
*  (Unspecified bits are always zero).
*
*  \verbatim
   ++-----------------------sample is signed if set
   ||
   ||       ++-----------sample is bigendian if set
   ||       ||
   ||       ||          ++---sample is float if set
   ||       ||          ||
   ||       ||          || +---sample bit size---+
   ||       ||          || |                     |
   15 14 13 12 11 10 09 08 07 06 05 04 03 02 01 00
   \endverbatim
*
*  There are macros in SDL 2.0 and later to query these bits.
*/
// typedef Uint16 SDL_AudioFormat;
type SDL_AudioFormat ffcommon.FUint16T

/**
 *  \name Audio flags
 */
/* @{ */

const SDL_AUDIO_MASK_BITSIZE = (0xFF)
const SDL_AUDIO_MASK_DATATYPE = (1 << 8)
const SDL_AUDIO_MASK_ENDIAN = (1 << 12)
const SDL_AUDIO_MASK_SIGNED = (1 << 15)

// #define SDL_AUDIO_BITSIZE(x)         (x & SDL_AUDIO_MASK_BITSIZE)
// #define SDL_AUDIO_ISFLOAT(x)         (x & SDL_AUDIO_MASK_DATATYPE)
// #define SDL_AUDIO_ISBIGENDIAN(x)     (x & SDL_AUDIO_MASK_ENDIAN)
// #define SDL_AUDIO_ISSIGNED(x)        (x & SDL_AUDIO_MASK_SIGNED)
// #define SDL_AUDIO_ISINT(x)           (!SDL_AUDIO_ISFLOAT(x))
// #define SDL_AUDIO_ISLITTLEENDIAN(x)  (!SDL_AUDIO_ISBIGENDIAN(x))
// #define SDL_AUDIO_ISUNSIGNED(x)      (!SDL_AUDIO_ISSIGNED(x))

/**
 *  \name Audio format flags
 *
 *  Defaults to LSB byte order.
 */
/* @{ */
const AUDIO_U8 = 0x0008     /**< Unsigned 8-bit samples */
const AUDIO_S8 = 0x8008     /**< Signed 8-bit samples */
const AUDIO_U16LSB = 0x0010 /**< Unsigned 16-bit samples */
const AUDIO_S16LSB = 0x8010 /**< Signed 16-bit samples */
const AUDIO_U16MSB = 0x1010 /**< As above, but big-endian byte order */
const AUDIO_S16MSB = 0x9010 /**< As above, but big-endian byte order */
const AUDIO_U16 = AUDIO_U16LSB
const AUDIO_S16 = AUDIO_S16LSB

/* @} */

/**
 *  \name int32 support
 */
/* @{ */
const AUDIO_S32LSB = 0x8020 /**< 32-bit integer samples */
const AUDIO_S32MSB = 0x9020 /**< As above, but big-endian byte order */
const AUDIO_S32 = AUDIO_S32LSB

/* @} */

/**
 *  \name float32 support
 */
/* @{ */
const AUDIO_F32LSB = 0x8120 /**< 32-bit floating point samples */
const AUDIO_F32MSB = 0x9120 /**< As above, but big-endian byte order */
const AUDIO_F32 = AUDIO_F32LSB

/* @} */

/**
 *  \name Native audio byte ordering
 */
/* @{ */
// #if SDL_BYTEORDER == SDL_LIL_ENDIAN
const AUDIO_U16SYS = AUDIO_U16LSB
const AUDIO_S16SYS = AUDIO_S16LSB
const AUDIO_S32SYS = AUDIO_S32LSB
const AUDIO_F32SYS = AUDIO_F32LSB

// #else
// #define AUDIO_U16SYS    AUDIO_U16MSB
// #define AUDIO_S16SYS    AUDIO_S16MSB
// #define AUDIO_S32SYS    AUDIO_S32MSB
// #define AUDIO_F32SYS    AUDIO_F32MSB
// #endif
/* @} */

/**
 *  \name Allow change flags
 *
 *  Which audio format changes are allowed when opening a device.
 */
/* @{ */
const SDL_AUDIO_ALLOW_FREQUENCY_CHANGE = 0x00000001
const SDL_AUDIO_ALLOW_FORMAT_CHANGE = 0x00000002
const SDL_AUDIO_ALLOW_CHANNELS_CHANGE = 0x00000004
const SDL_AUDIO_ALLOW_SAMPLES_CHANGE = 0x00000008
const SDL_AUDIO_ALLOW_ANY_CHANGE = (SDL_AUDIO_ALLOW_FREQUENCY_CHANGE | SDL_AUDIO_ALLOW_FORMAT_CHANGE | SDL_AUDIO_ALLOW_CHANNELS_CHANGE | SDL_AUDIO_ALLOW_SAMPLES_CHANGE)

/* @} */

/* @} */ /* Audio flags */

/**
 *  This function is called when the audio device needs more data.
 *
 *  \param userdata An application-specific parameter saved in
 *                  the SDL_AudioSpec structure
 *  \param stream A pointer to the audio data buffer.
 *  \param len    The length of that buffer in bytes.
 *
 *  Once the callback returns, the buffer will no longer be valid.
 *  Stereo samples are stored in a LRLRLR ordering.
 *
 *  You can choose to avoid callbacks and use SDL_QueueAudio() instead, if
 *  you like. Just open your audio device with a NULL callback.
 */
// typedef void (SDLCALL * SDL_AudioCallback) (void *userdata, Uint8 * stream,
//                                             int len);
type SDL_AudioCallback = func(userdata ffcommon.FVoidP, stream *ffcommon.FUint8T, len0 ffcommon.FInt) uintptr

/**
 *  The calculated values in this structure are calculated by SDL_OpenAudioDevice().
 *
 *  For multi-channel audio, the default SDL channel mapping is:
 *  2:  FL FR                       (stereo)
 *  3:  FL FR LFE                   (2.1 surround)
 *  4:  FL FR BL BR                 (quad)
 *  5:  FL FR LFE BL BR             (4.1 surround)
 *  6:  FL FR FC LFE SL SR          (5.1 surround - last two can also be BL BR)
 *  7:  FL FR FC LFE BC SL SR       (6.1 surround)
 *  8:  FL FR FC LFE BL BR SL SR    (7.1 surround)
 */
type SDL_AudioSpec struct {
	Freq     ffcommon.FInt     /**< DSP frequency -- samples per second */
	Format   SDL_AudioFormat   /**< Audio data format */
	Channels ffcommon.FUint8T  /**< Number of channels: 1 mono, 2 stereo */
	Silence  ffcommon.FUint8T  /**< Audio buffer silence value (calculated) */
	Samples  ffcommon.FUint16T /**< Audio buffer size in sample FRAMES (total samples divided by channel count) */
	Padding  ffcommon.FUint16T /**< Necessary for some compile environments */
	Cize     ffcommon.FUint32T /**< Audio buffer size in bytes (calculated) */
	Callback uintptr           //*SDL_AudioCallback // /**< Callback that feeds the audio device (NULL to use SDL_QueueAudio()). */
	Userdata ffcommon.FVoidP   /**< Userdata passed to callback (ignored for NULL callbacks). */
}

/* Function prototypes */

/**
 *  \name Driver discovery functions
 *
 *  These functions return the list of built in audio drivers, in the
 *  order that they are normally initialized by default.
 */
/* @{ */

/**
 * Use this function to get the number of built-in audio drivers.
 *
 * This function returns a hardcoded number. This never returns a negative
 * value; if there are no drivers compiled into this build of SDL, this
 * function returns zero. The presence of a driver in this list does not mean
 * it will function, it just means SDL is capable of interacting with that
 * interface. For example, a build of SDL might have esound support, but if
 * there's no esound server available, SDL's esound driver would fail if used.
 *
 * By default, SDL tries all drivers, in its preferred order, until one is
 * found to be usable.
 *
 * \returns the number of built-in audio drivers.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetAudioDriver
 */
// extern DECLSPEC int SDLCALL SDL_GetNumAudioDrivers(void);
func SDL_GetNumAudioDrivers() (res ffcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetNumAudioDrivers").Call()
	res = ffcommon.FInt(t)
	return
}

/**
 * Use this function to get the name of a built in audio driver.
 *
 * The list of audio drivers is given in the order that they are normally
 * initialized by default; the drivers that seem more reasonable to choose
 * first (as far as the SDL developers believe) are earlier in the list.
 *
 * The names of drivers are all simple, low-ASCII identifiers, like "alsa",
 * "coreaudio" or "xaudio2". These never have Unicode characters, and are not
 * meant to be proper names.
 *
 * \param index the index of the audio driver; the value ranges from 0 to
 *              SDL_GetNumAudioDrivers() - 1
 * \returns the name of the audio driver at the requested index, or NULL if an
 *          invalid index was specified.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetNumAudioDrivers
 */
// extern DECLSPEC const char *SDLCALL SDL_GetAudioDriver(int index);
func SDL_GetAudioDriver() (res ffcommon.FConstCharP) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetAudioDriver").Call()
	res = ffcommon.StringFromPtr(t)
	return
}

/* @} */

/**
 * Get the name of the current audio driver.
 *
 * The returned string points to internal static memory and thus never becomes
 * invalid, even if you quit the audio subsystem and initialize a new driver
 * (although such a case would return a different static string from another
 * call to this function, of course). As such, you should not modify or free
 * the returned string.
 *
 * \returns the name of the current audio driver or NULL if no driver has been
 *          initialized.
 *
 * \since This function is available since SDL 3.0.0.
 */
// extern DECLSPEC const char *SDLCALL SDL_GetCurrentAudioDriver(void);
func SDL_GetCurrentAudioDriver() (res ffcommon.FConstCharP) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetCurrentAudioDriver").Call()
	res = ffcommon.StringFromPtr(t)
	return
}

/**
 *  SDL Audio Device IDs.
 */
// typedef Uint32 SDL_AudioDeviceID;
type SDL_AudioDeviceID ffcommon.FUint32T

/**
 * Get the number of built-in audio devices.
 *
 * This function is only valid after successfully initializing the audio
 * subsystem.
 *
 * Note that audio capture support is not implemented as of SDL 2.0.4, so the
 * `iscapture` parameter is for future expansion and should always be zero for
 * now.
 *
 * This function will return -1 if an explicit list of devices can't be
 * determined. Returning -1 is not an error. For example, if SDL is set up to
 * talk to a remote audio server, it can't list every one available on the
 * Internet, but it will still allow a specific host to be specified in
 * SDL_OpenAudioDevice().
 *
 * In many common cases, when this function returns a value <= 0, it can still
 * successfully open the default device (NULL for first argument of
 * SDL_OpenAudioDevice()).
 *
 * This function may trigger a complete redetect of available hardware. It
 * should not be called for each iteration of a loop, but rather once at the
 * start of a loop:
 *
 * ```c
 * // Don't do this:
 * for (int i = 0; i < SDL_GetNumAudioDevices(0); i++)
 *
 * // do this instead:
 * const int count = SDL_GetNumAudioDevices(0);
 * for (int i = 0; i < count; ++i) { do_something_here(); }
 * ```
 *
 * \param iscapture zero to request playback devices, non-zero to request
 *                  recording devices
 * \returns the number of available devices exposed by the current driver or
 *          -1 if an explicit list of devices can't be determined. A return
 *          value of -1 does not necessarily mean an error condition.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetAudioDeviceName
 * \sa SDL_OpenAudioDevice
 */
// extern DECLSPEC int SDLCALL SDL_GetNumAudioDevices(int iscapture);
func SDL_GetNumAudioDevices(iscapture ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetNumAudioDevices").Call(
		uintptr(iscapture),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Get the human-readable name of a specific audio device.
 *
 * This function is only valid after successfully initializing the audio
 * subsystem. The values returned by this function reflect the latest call to
 * SDL_GetNumAudioDevices(); re-call that function to redetect available
 * hardware.
 *
 * The string returned by this function is UTF-8 encoded, read-only, and
 * managed internally. You are not to free it. If you need to keep the string
 * for any length of time, you should make your own copy of it, as it will be
 * invalid next time any of several other SDL functions are called.
 *
 * \param index the index of the audio device; valid values range from 0 to
 *              SDL_GetNumAudioDevices() - 1
 * \param iscapture non-zero to query the list of recording devices, zero to
 *                  query the list of output devices.
 * \returns the name of the audio device at the requested index, or NULL on
 *          error.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetNumAudioDevices
 * \sa SDL_GetDefaultAudioInfo
 */
// extern DECLSPEC const char *SDLCALL SDL_GetAudioDeviceName(int index,
//                                                            int iscapture);
func SDL_GetAudioDeviceName(index, iscapture ffcommon.FInt) (res ffcommon.FConstCharP) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetAudioDeviceName").Call(
		uintptr(index),
		uintptr(iscapture),
	)
	res = ffcommon.StringFromPtr(t)
	return
}

/**
 * Get the preferred audio format of a specific audio device.
 *
 * This function is only valid after a successfully initializing the audio
 * subsystem. The values returned by this function reflect the latest call to
 * SDL_GetNumAudioDevices(); re-call that function to redetect available
 * hardware.
 *
 * `spec` will be filled with the sample rate, sample format, and channel
 * count.
 *
 * \param index the index of the audio device; valid values range from 0 to
 *              SDL_GetNumAudioDevices() - 1
 * \param iscapture non-zero to query the list of recording devices, zero to
 *                  query the list of output devices.
 * \param spec The SDL_AudioSpec to be initialized by this function.
 * \returns 0 on success, nonzero on error
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetNumAudioDevices
 * \sa SDL_GetDefaultAudioInfo
 */
// extern DECLSPEC int SDLCALL SDL_GetAudioDeviceSpec(int index,
//                                                    int iscapture,
//                                                    SDL_AudioSpec *spec);
func SDL_GetAudioDeviceSpec(index, iscapture ffcommon.FInt, spec *SDL_AudioSpec) (res ffcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetAudioDeviceSpec").Call(
		uintptr(index),
		uintptr(iscapture),
		uintptr(unsafe.Pointer(spec)),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Get the name and preferred format of the default audio device.
 *
 * Some (but not all!) platforms have an isolated mechanism to get information
 * about the "default" device. This can actually be a completely different
 * device that's not in the list you get from SDL_GetAudioDeviceSpec(). It can
 * even be a network address! (This is discussed in SDL_OpenAudioDevice().)
 *
 * As a result, this call is not guaranteed to be performant, as it can query
 * the sound server directly every time, unlike the other query functions. You
 * should call this function sparingly!
 *
 * `spec` will be filled with the sample rate, sample format, and channel
 * count, if a default device exists on the system. If `name` is provided,
 * will be filled with either a dynamically-allocated UTF-8 string or NULL.
 *
 * \param name A pointer to be filled with the name of the default device (can
 *             be NULL). Please call SDL_free() when you are done with this
 *             pointer!
 * \param spec The SDL_AudioSpec to be initialized by this function.
 * \param iscapture non-zero to query the default recording device, zero to
 *                  query the default output device.
 * \returns 0 on success, nonzero on error
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetAudioDeviceName
 * \sa SDL_GetAudioDeviceSpec
 * \sa SDL_OpenAudioDevice
 */
// extern DECLSPEC int SDLCALL SDL_GetDefaultAudioInfo(char **name,
//                                                     SDL_AudioSpec *spec,
//                                                     int iscapture);

func SDL_GetDefaultAudioInfo(name **ffcommon.FChar, spec *SDL_AudioSpec, iscapture ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetDefaultAudioInfo").Call(
		uintptr(unsafe.Pointer(name)),
		uintptr(unsafe.Pointer(spec)),
		uintptr(iscapture),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Open a specific audio device.
 *
 * Passing in a `device` name of NULL requests the most reasonable default.
 * The `device` name is a UTF-8 string reported by SDL_GetAudioDeviceName(),
 * but some drivers allow arbitrary and driver-specific strings, such as a
 * hostname/IP address for a remote audio server, or a filename in the
 * diskaudio driver.
 *
 * An opened audio device starts out paused, and should be enabled for playing
 * by calling SDL_PlayAudioDevice(devid) when you are ready for your audio
 * callback function to be called. Since the audio driver may modify the
 * requested size of the audio buffer, you should allocate any local mixing
 * buffers after you open the audio device.
 *
 * The audio callback runs in a separate thread in most cases; you can prevent
 * race conditions between your callback and other threads without fully
 * pausing playback with SDL_LockAudioDevice(). For more information about the
 * callback, see SDL_AudioSpec.
 *
 * Managing the audio spec via 'desired' and 'obtained':
 *
 * When filling in the desired audio spec structure:
 *
 * - `desired->freq` should be the frequency in sample-frames-per-second (Hz).
 * - `desired->format` should be the audio format (`AUDIO_S16SYS`, etc).
 * - `desired->samples` is the desired size of the audio buffer, in _sample
 *   frames_ (with stereo output, two samples--left and right--would make a
 *   single sample frame). This number should be a power of two, and may be
 *   adjusted by the audio driver to a value more suitable for the hardware.
 *   Good values seem to range between 512 and 8096 inclusive, depending on
 *   the application and CPU speed. Smaller values reduce latency, but can
 *   lead to underflow if the application is doing heavy processing and cannot
 *   fill the audio buffer in time. Note that the number of sample frames is
 *   directly related to time by the following formula: `ms =
 *   (sampleframes*1000)/freq`
 * - `desired->size` is the size in _bytes_ of the audio buffer, and is
 *   calculated by SDL_OpenAudioDevice(). You don't initialize this.
 * - `desired->silence` is the value used to set the buffer to silence, and is
 *   calculated by SDL_OpenAudioDevice(). You don't initialize this.
 * - `desired->callback` should be set to a function that will be called when
 *   the audio device is ready for more data. It is passed a pointer to the
 *   audio buffer, and the length in bytes of the audio buffer. This function
 *   usually runs in a separate thread, and so you should protect data
 *   structures that it accesses by calling SDL_LockAudioDevice() and
 *   SDL_UnlockAudioDevice() in your code. Alternately, you may pass a NULL
 *   pointer here, and call SDL_QueueAudio() with some frequency, to queue
 *   more audio samples to be played (or for capture devices, call
 *   SDL_DequeueAudio() with some frequency, to obtain audio samples).
 * - `desired->userdata` is passed as the first parameter to your callback
 *   function. If you passed a NULL callback, this value is ignored.
 *
 * `allowed_changes` can have the following flags OR'd together:
 *
 * - `SDL_AUDIO_ALLOW_FREQUENCY_CHANGE`
 * - `SDL_AUDIO_ALLOW_FORMAT_CHANGE`
 * - `SDL_AUDIO_ALLOW_CHANNELS_CHANGE`
 * - `SDL_AUDIO_ALLOW_SAMPLES_CHANGE`
 * - `SDL_AUDIO_ALLOW_ANY_CHANGE`
 *
 * These flags specify how SDL should behave when a device cannot offer a
 * specific feature. If the application requests a feature that the hardware
 * doesn't offer, SDL will always try to get the closest equivalent.
 *
 * For example, if you ask for float32 audio format, but the sound card only
 * supports int16, SDL will set the hardware to int16. If you had set
 * SDL_AUDIO_ALLOW_FORMAT_CHANGE, SDL will change the format in the `obtained`
 * structure. If that flag was *not* set, SDL will prepare to convert your
 * callback's float32 audio to int16 before feeding it to the hardware and
 * will keep the originally requested format in the `obtained` structure.
 *
 * The resulting audio specs, varying depending on hardware and on what
 * changes were allowed, will then be written back to `obtained`.
 *
 * If your application can only handle one specific data format, pass a zero
 * for `allowed_changes` and let SDL transparently handle any differences.
 *
 * \param device a UTF-8 string reported by SDL_GetAudioDeviceName() or a
 *               driver-specific name as appropriate. NULL requests the most
 *               reasonable default device.
 * \param iscapture non-zero to specify a device should be opened for
 *                  recording, not playback
 * \param desired an SDL_AudioSpec structure representing the desired output
 *                format
 * \param obtained an SDL_AudioSpec structure filled in with the actual output
 *                 format
 * \param allowed_changes 0, or one or more flags OR'd together
 * \returns a valid device ID that is > 0 on success or 0 on failure; call
 *          SDL_GetError() for more information.
 *
 *          For compatibility with SDL 1.2, this will never return 1, since
 *          SDL reserves that ID for the legacy SDL_OpenAudio() function.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_CloseAudioDevice
 * \sa SDL_GetAudioDeviceName
 * \sa SDL_LockAudioDevice
 * \sa SDL_PlayAudioDevice
 * \sa SDL_PauseAudioDevice
 * \sa SDL_UnlockAudioDevice
 */
// extern DECLSPEC SDL_AudioDeviceID SDLCALL SDL_OpenAudioDevice(
//                                                   const char *device,
//                                                   int iscapture,
//                                                   const SDL_AudioSpec *desired,
//                                                   SDL_AudioSpec *obtained,
//                                                   int allowed_changes);

func SDL_OpenAudioDevice(device ffcommon.FConstCharP, iscapture ffcommon.FInt, desired *SDL_AudioSpec, obtained *SDL_AudioSpec, allowed_changes ffcommon.FInt) (res SDL_AudioDeviceID) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_OpenAudioDevice").Call(
		ffcommon.UintPtrFromString(device),
		uintptr(iscapture),
		uintptr(unsafe.Pointer(desired)),
		uintptr(unsafe.Pointer(obtained)),
		uintptr(allowed_changes),
	)
	res = SDL_AudioDeviceID(t)
	return
}

/**
 *  \name Audio state
 *
 *  Get the current audio state.
 */
/* @{ */
type SDL_AudioStatus int32

const (
	SDL_AUDIO_STOPPED = iota
	SDL_AUDIO_PLAYING
	SDL_AUDIO_PAUSED
)

/**
 * Use this function to get the current audio state of an audio device.
 *
 * \param dev the ID of an audio device previously opened with
 *            SDL_OpenAudioDevice()
 * \returns the SDL_AudioStatus of the specified audio device.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_PlayAudioDevice
 * \sa SDL_PauseAudioDevice
 */
// extern DECLSPEC SDL_AudioStatus SDLCALL SDL_GetAudioDeviceStatus(SDL_AudioDeviceID dev);
func SDL_GetAudioStatus() (res SDL_AudioStatus) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetAudioStatus").Call()
	res = SDL_AudioStatus(t)
	return
}

/* @} */ /* Audio State */

/**
 * Use this function to play audio on a specified device.
 *
 * Newly-opened audio devices start in the paused state, so you must call this
 * function after opening the specified audio device to start playing sound.
 * This allows you to safely initialize data for your callback function after
 * opening the audio device. Silence will be written to the audio device while
 * paused, and the audio callback is guaranteed to not be called. Pausing one
 * device does not prevent other unpaused devices from running their
 * callbacks.
 *
 * \param dev a device opened by SDL_OpenAudioDevice()
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_LockAudioDevice
 * \sa SDL_PauseAudioDevice
 */
// extern DECLSPEC int SDLCALL SDL_PlayAudioDevice(SDL_AudioDeviceID dev);
func SDL_PlayAudioDevice(dev SDL_AudioDeviceID) (res ffcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_PlayAudioDevice").Call(
		uintptr(dev),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Use this function to pause audio playback on a specified device.
 *
 * This function pauses the audio callback processing for a given device.
 * Silence will be written to the audio device while paused, and the audio
 * callback is guaranteed to not be called. Pausing one device does not
 * prevent other unpaused devices from running their callbacks.
 *
 * If you just need to protect a few variables from race conditions vs your
 * callback, you shouldn't pause the audio device, as it will lead to dropouts
 * in the audio playback. Instead, you should use SDL_LockAudioDevice().
 *
 * \param dev a device opened by SDL_OpenAudioDevice()
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_LockAudioDevice
 * \sa SDL_PlayAudioDevice
 */
// extern DECLSPEC int SDLCALL SDL_PauseAudioDevice(SDL_AudioDeviceID dev);

func SDL_PauseAudioDevice(dev SDL_AudioDeviceID) (res ffcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_PauseAudioDevice").Call(
		uintptr(dev),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Load the audio data of a WAVE file into memory.
 *
 * Loading a WAVE file requires `src`, `spec`, `audio_buf` and `audio_len` to
 * be valid pointers. The entire data portion of the file is then loaded into
 * memory and decoded if necessary.
 *
 * If `freesrc` is non-zero, the data source gets automatically closed and
 * freed before the function returns.
 *
 * Supported formats are RIFF WAVE files with the formats PCM (8, 16, 24, and
 * 32 bits), IEEE Float (32 bits), Microsoft ADPCM and IMA ADPCM (4 bits), and
 * A-law and mu-law (8 bits). Other formats are currently unsupported and
 * cause an error.
 *
 * If this function succeeds, the pointer returned by it is equal to `spec`
 * and the pointer to the audio data allocated by the function is written to
 * `audio_buf` and its length in bytes to `audio_len`. The SDL_AudioSpec
 * members `freq`, `channels`, and `format` are set to the values of the audio
 * data in the buffer. The `samples` member is set to a sane default and all
 * others are set to zero.
 *
 * It's necessary to use SDL_free() to free the audio data returned in
 * `audio_buf` when it is no longer used.
 *
 * Because of the underspecification of the .WAV format, there are many
 * problematic files in the wild that cause issues with strict decoders. To
 * provide compatibility with these files, this decoder is lenient in regards
 * to the truncation of the file, the fact chunk, and the size of the RIFF
 * chunk. The hints `SDL_HINT_WAVE_RIFF_CHUNK_SIZE`,
 * `SDL_HINT_WAVE_TRUNCATION`, and `SDL_HINT_WAVE_FACT_CHUNK` can be used to
 * tune the behavior of the loading process.
 *
 * Any file that is invalid (due to truncation, corruption, or wrong values in
 * the headers), too big, or unsupported causes an error. Additionally, any
 * critical I/O error from the data source will terminate the loading process
 * with an error. The function returns NULL on error and in all cases (with
 * the exception of `src` being NULL), an appropriate error message will be
 * set.
 *
 * It is required that the data source supports seeking.
 *
 * Example:
 *
 * ```c
 * SDL_LoadWAV_RW(SDL_RWFromFile("sample.wav", "rb"), 1, &spec, &buf, &len);
 * ```
 *
 * Note that the SDL_LoadWAV macro does this same thing for you, but in a less
 * messy way:
 *
 * ```c
 * SDL_LoadWAV("sample.wav", &spec, &buf, &len);
 * ```
 *
 * \param src The data source for the WAVE data
 * \param freesrc If non-zero, SDL will _always_ free the data source
 * \param spec An SDL_AudioSpec that will be filled in with the wave file's
 *             format details
 * \param audio_buf A pointer filled with the audio data, allocated by the
 *                  function.
 * \param audio_len A pointer filled with the length of the audio data buffer
 *                  in bytes
 * \returns This function, if successfully called, returns `spec`, which will
 *          be filled with the audio data format of the wave source data.
 *          `audio_buf` will be filled with a pointer to an allocated buffer
 *          containing the audio data, and `audio_len` is filled with the
 *          length of that audio buffer in bytes.
 *
 *          This function returns NULL if the .WAV file cannot be opened, uses
 *          an unknown data format, or is corrupt; call SDL_GetError() for
 *          more information.
 *
 *          When the application is done with the data returned in
 *          `audio_buf`, it should call SDL_free() to dispose of it.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_free
 * \sa SDL_LoadWAV
 */
// extern DECLSPEC SDL_AudioSpec *SDLCALL SDL_LoadWAV_RW(SDL_RWops * src,
//                                                       int freesrc,
//                                                       SDL_AudioSpec * spec,
//                                                       Uint8 ** audio_buf,
//                                                       Uint32 * audio_len);
func (src *SDL_RWops) SDL_LoadWAV_RW(freesrc ffcommon.FInt, spec *SDL_AudioSpec, audio_buf **ffcommon.FUint8T, audio_len *ffcommon.FUint32T) (res *SDL_AudioSpec) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_LoadWAV_RW").Call(
		uintptr(unsafe.Pointer(src)),
		uintptr(freesrc),
		uintptr(unsafe.Pointer(spec)),
		uintptr(unsafe.Pointer(audio_buf)),
		uintptr(unsafe.Pointer(audio_len)),
	)
	if t == 0 {

	}
	res = (*SDL_AudioSpec)(unsafe.Pointer(t))
	return
}

/**
 *  Loads a WAV from a file.
 *  Compatibility convenience function.
 */
// #define SDL_LoadWAV(file, spec, audio_buf, audio_len) \
//     SDL_LoadWAV_RW(SDL_RWFromFile(file, "rb"),1, spec,audio_buf,audio_len)

/* SDL_AudioStream is a new audio conversion interface.
   The benefits vs SDL_AudioCVT:
    - it can handle resampling data in chunks without generating
      artifacts, when it doesn't have the complete buffer available.
    - it can handle incoming data in any variable size.
    - You push data as you have it, and pull it when you need it
*/
/* this is opaque to the outside world. */
// struct SDL_AudioStream;
// typedef struct SDL_AudioStream SDL_AudioStream;
type SDL_AudioStream struct {
	//不透明
}

/**
 * Create a new audio stream.
 *
 * \param src_format The format of the source audio
 * \param src_channels The number of channels of the source audio
 * \param src_rate The sampling rate of the source audio
 * \param dst_format The format of the desired audio output
 * \param dst_channels The number of channels of the desired audio output
 * \param dst_rate The sampling rate of the desired audio output
 * \returns 0 on success, or -1 on error.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_PutAudioStreamData
 * \sa SDL_GetAudioStreamData
 * \sa SDL_GetAudioStreamAvailable
 * \sa SDL_FlushAudioStream
 * \sa SDL_ClearAudioStream
 * \sa SDL_DestroyAudioStream
 */
// extern DECLSPEC SDL_AudioStream *SDLCALL SDL_CreateAudioStream(SDL_AudioFormat src_format,
//                                                             Uint8 src_channels,
//                                                             int src_rate,
//                                                             SDL_AudioFormat dst_format,
//                                                             Uint8 dst_channels,
//                                                             int dst_rate);
func SDL_CreateAudioStream(src_format SDL_AudioFormat,
	src_channels ffcommon.FUint8T,
	src_rate ffcommon.FInt,
	dst_format SDL_AudioFormat,
	dst_channels ffcommon.FUint8T,
	dst_rate ffcommon.FInt) (res *SDL_AudioSpec) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_CreateAudioStream").Call(
		uintptr(src_format),
		uintptr(src_channels),
		uintptr(src_rate),
		uintptr(dst_format),
		uintptr(dst_channels),
		uintptr(dst_rate),
	)
	res = (*SDL_AudioSpec)(unsafe.Pointer(t))
	return
}

/**
 * Add data to be converted/resampled to the stream.
 *
 * \param stream The stream the audio data is being added to
 * \param buf A pointer to the audio data to add
 * \param len The number of bytes to write to the stream
 * \returns 0 on success, or -1 on error.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_CreateAudioStream
 * \sa SDL_GetAudioStreamData
 * \sa SDL_GetAudioStreamAvailable
 * \sa SDL_FlushAudioStream
 * \sa SDL_ClearAudioStream
 * \sa SDL_DestroyAudioStream
 */
// extern DECLSPEC int SDLCALL SDL_PutAudioStreamData(SDL_AudioStream *stream, const void *buf, int len);
func (stream *SDL_AudioStream) SDL_PutAudioStreamData(buf ffcommon.FConstVoidP, len0 ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_PutAudioStreamData").Call(
		uintptr(unsafe.Pointer(stream)),
		buf,
		uintptr(len0),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Get converted/resampled data from the stream
 *
 * \param stream The stream the audio is being requested from
 * \param buf A buffer to fill with audio data
 * \param len The maximum number of bytes to fill
 * \returns the number of bytes read from the stream, or -1 on error
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_CreateAudioStream
 * \sa SDL_PutAudioStreamData
 * \sa SDL_GetAudioStreamAvailable
 * \sa SDL_FlushAudioStream
 * \sa SDL_ClearAudioStream
 * \sa SDL_DestroyAudioStream
 */
// extern DECLSPEC int SDLCALL SDL_GetAudioStreamData(SDL_AudioStream *stream, void *buf, int len);
func (stream *SDL_AudioStream) SDL_GetAudioStreamData(buf ffcommon.FConstVoidP, len0 ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetAudioStreamData").Call(
		uintptr(unsafe.Pointer(stream)),
		buf,
		uintptr(len0),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Get the number of converted/resampled bytes available.
 *
 * The stream may be buffering data behind the scenes until it has enough to
 * resample correctly, so this number might be lower than what you expect, or
 * even be zero. Add more data or flush the stream if you need the data now.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_CreateAudioStream
 * \sa SDL_PutAudioStreamData
 * \sa SDL_GetAudioStreamData
 * \sa SDL_FlushAudioStream
 * \sa SDL_ClearAudioStream
 * \sa SDL_DestroyAudioStream
 */
// extern DECLSPEC int SDLCALL SDL_GetAudioStreamAvailable(SDL_AudioStream *stream);
func (stream *SDL_AudioStream) SDL_GetAudioStreamAvailable() (res ffcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetAudioStreamAvailable").Call(
		uintptr(unsafe.Pointer(stream)),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Tell the stream that you're done sending data, and anything being buffered
 * should be converted/resampled and made available immediately.
 *
 * It is legal to add more data to a stream after flushing, but there will be
 * audio gaps in the output. Generally this is intended to signal the end of
 * input, so the complete output becomes available.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_CreateAudioStream
 * \sa SDL_PutAudioStreamData
 * \sa SDL_GetAudioStreamData
 * \sa SDL_GetAudioStreamAvailable
 * \sa SDL_ClearAudioStream
 * \sa SDL_DestroyAudioStream
 */
// extern DECLSPEC int SDLCALL SDL_FlushAudioStream(SDL_AudioStream *stream);
func (stream *SDL_AudioStream) SDL_FlushAudioStream() (res ffcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_FlushAudioStream").Call(
		uintptr(unsafe.Pointer(stream)),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Clear any pending data in the stream without converting it
 *
 * \param   stream The audio stream to clear
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_CreateAudioStream
 * \sa SDL_PutAudioStreamData
 * \sa SDL_GetAudioStreamData
 * \sa SDL_GetAudioStreamAvailable
 * \sa SDL_FlushAudioStream
 * \sa SDL_DestroyAudioStream
 */
// extern DECLSPEC int SDLCALL SDL_ClearAudioStream(SDL_AudioStream *stream);
func (stream *SDL_AudioStream) SDL_ClearAudioStream() (res ffcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_ClearAudioStream").Call(
		uintptr(unsafe.Pointer(stream)),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Free an audio stream
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_CreateAudioStream
 * \sa SDL_PutAudioStreamData
 * \sa SDL_GetAudioStreamData
 * \sa SDL_GetAudioStreamAvailable
 * \sa SDL_FlushAudioStream
 * \sa SDL_ClearAudioStream
 */
// extern DECLSPEC void SDLCALL SDL_DestroyAudioStream(SDL_AudioStream *stream);
func (stream *SDL_AudioStream) SDL_DestroyAudioStream() {
	sdlcommon.GetSDL2Dll().NewProc("SDL_DestroyAudioStream").Call(
		uintptr(unsafe.Pointer(stream)),
	)
}

const SDL_MIX_MAXVOLUME = 128

/**
 * Mix audio data in a specified format.
 *
 * This takes an audio buffer `src` of `len` bytes of `format` data and mixes
 * it into `dst`, performing addition, volume adjustment, and overflow
 * clipping. The buffer pointed to by `dst` must also be `len` bytes of
 * `format` data.
 *
 * This is provided for convenience -- you can mix your own audio data.
 *
 * Do not use this function for mixing together more than two streams of
 * sample data. The output from repeated application of this function may be
 * distorted by clipping, because there is no accumulator with greater range
 * than the input (not to mention this being an inefficient way of doing it).
 *
 * It is a common misconception that this function is required to write audio
 * data to an output stream in an audio callback. While you can do that,
 * SDL_MixAudioFormat() is really only needed when you're mixing a single
 * audio stream with a volume adjustment.
 *
 * \param dst the destination for the mixed audio
 * \param src the source audio buffer to be mixed
 * \param format the SDL_AudioFormat structure representing the desired audio
 *               format
 * \param len the length of the audio buffer in bytes
 * \param volume ranges from 0 - 128, and should be set to SDL_MIX_MAXVOLUME
 *               for full audio volume
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 */
// extern DECLSPEC int SDLCALL SDL_MixAudioFormat(Uint8 * dst,
//                                                 const Uint8 * src,
//                                                 SDL_AudioFormat format,
//                                                 Uint32 len, int volume);
func SDL_MixAudioFormat(dst *ffcommon.FUint8T, src *ffcommon.FUint8T, format SDL_AudioFormat, len0 ffcommon.FUint32T, volume ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_MixAudioFormat").Call(
		uintptr(unsafe.Pointer(dst)),
		uintptr(unsafe.Pointer(src)),
		uintptr(format),
		uintptr(len0),
		uintptr(volume),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Queue more audio on non-callback devices.
 *
 * If you are looking to retrieve queued audio from a non-callback capture
 * device, you want SDL_DequeueAudio() instead. SDL_QueueAudio() will return
 * -1 to signify an error if you use it with capture devices.
 *
 * SDL offers two ways to feed audio to the device: you can either supply a
 * callback that SDL triggers with some frequency to obtain more audio (pull
 * method), or you can supply no callback, and then SDL will expect you to
 * supply data at regular intervals (push method) with this function.
 *
 * There are no limits on the amount of data you can queue, short of
 * exhaustion of address space. Queued data will drain to the device as
 * necessary without further intervention from you. If the device needs audio
 * but there is not enough queued, it will play silence to make up the
 * difference. This means you will have skips in your audio playback if you
 * aren't routinely queueing sufficient data.
 *
 * This function copies the supplied data, so you are safe to free it when the
 * function returns. This function is thread-safe, but queueing to the same
 * device from two threads at once does not promise which buffer will be
 * queued first.
 *
 * You may not queue audio on a device that is using an application-supplied
 * callback; doing so returns an error. You have to use the audio callback or
 * queue audio with this function, but not both.
 *
 * You should not call SDL_LockAudio() on the device before queueing; SDL
 * handles locking internally for this function.
 *
 * Note that SDL does not support planar audio. You will need to resample from
 * planar audio formats into a non-planar one (see SDL_AudioFormat) before
 * queuing audio.
 *
 * \param dev the device ID to which we will queue audio
 * \param data the data to queue to the device for later playback
 * \param len the number of bytes (not samples!) to which `data` points
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_ClearQueuedAudio
 * \sa SDL_GetQueuedAudioSize
 */
// extern DECLSPEC int SDLCALL SDL_QueueAudio(SDL_AudioDeviceID dev, const void *data, Uint32 len);
func SDL_QueueAudio(dev SDL_AudioDeviceID, data ffcommon.FConstVoidP, len0 ffcommon.FUint32T) (res ffcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_QueueAudio").Call(
		uintptr(dev),
		data,
		uintptr(len0),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Dequeue more audio on non-callback devices.
 *
 * If you are looking to queue audio for output on a non-callback playback
 * device, you want SDL_QueueAudio() instead. SDL_DequeueAudio() will always
 * return 0 if you use it with playback devices.
 *
 * SDL offers two ways to retrieve audio from a capture device: you can either
 * supply a callback that SDL triggers with some frequency as the device
 * records more audio data, (push method), or you can supply no callback, and
 * then SDL will expect you to retrieve data at regular intervals (pull
 * method) with this function.
 *
 * There are no limits on the amount of data you can queue, short of
 * exhaustion of address space. Data from the device will keep queuing as
 * necessary without further intervention from you. This means you will
 * eventually run out of memory if you aren't routinely dequeueing data.
 *
 * Capture devices will not queue data when paused; if you are expecting to
 * not need captured audio for some length of time, use SDL_PauseAudioDevice()
 * to stop the capture device from queueing more data. This can be useful
 * during, say, level loading times. When unpaused, capture devices will start
 * queueing data from that point, having flushed any capturable data available
 * while paused.
 *
 * This function is thread-safe, but dequeueing from the same device from two
 * threads at once does not promise which thread will dequeue data first.
 *
 * You may not dequeue audio from a device that is using an
 * application-supplied callback; doing so returns an error. You have to use
 * the audio callback, or dequeue audio with this function, but not both.
 *
 * You should not call SDL_LockAudio() on the device before dequeueing; SDL
 * handles locking internally for this function.
 *
 * \param dev the device ID from which we will dequeue audio
 * \param data a pointer into where audio data should be copied
 * \param len the number of bytes (not samples!) to which (data) points
 * \returns the number of bytes dequeued, which could be less than requested;
 *          call SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_ClearQueuedAudio
 * \sa SDL_GetQueuedAudioSize
 */
// extern DECLSPEC Uint32 SDLCALL SDL_DequeueAudio(SDL_AudioDeviceID dev, void *data, Uint32 len);
func SDL_DequeueAudio(dev SDL_AudioDeviceID, data ffcommon.FConstVoidP, len0 ffcommon.FUint32T) (res ffcommon.FUint32T) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_DequeueAudio").Call(
		uintptr(dev),
		data,
		uintptr(len0),
	)
	res = ffcommon.FUint32T(t)
	return
}

/**
 * Get the number of bytes of still-queued audio.
 *
 * For playback devices: this is the number of bytes that have been queued for
 * playback with SDL_QueueAudio(), but have not yet been sent to the hardware.
 *
 * Once we've sent it to the hardware, this function can not decide the exact
 * byte boundary of what has been played. It's possible that we just gave the
 * hardware several kilobytes right before you called this function, but it
 * hasn't played any of it yet, or maybe half of it, etc.
 *
 * For capture devices, this is the number of bytes that have been captured by
 * the device and are waiting for you to dequeue. This number may grow at any
 * time, so this only informs of the lower-bound of available data.
 *
 * You may not queue or dequeue audio on a device that is using an
 * application-supplied callback; calling this function on such a device
 * always returns 0. You have to use the audio callback or queue audio, but
 * not both.
 *
 * You should not call SDL_LockAudio() on the device before querying; SDL
 * handles locking internally for this function.
 *
 * \param dev the device ID of which we will query queued audio size
 * \returns the number of bytes (not samples!) of queued audio.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_ClearQueuedAudio
 * \sa SDL_QueueAudio
 * \sa SDL_DequeueAudio
 */
// extern DECLSPEC Uint32 SDLCALL SDL_GetQueuedAudioSize(SDL_AudioDeviceID dev);
func SDL_GetQueuedAudioSize(dev SDL_AudioDeviceID) (res ffcommon.FUint32T) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetQueuedAudioSize").Call(
		uintptr(dev),
	)
	res = ffcommon.FUint32T(t)
	return
}

/**
 * Drop any queued audio data waiting to be sent to the hardware.
 *
 * Immediately after this call, SDL_GetQueuedAudioSize() will return 0. For
 * output devices, the hardware will start playing silence if more audio isn't
 * queued. For capture devices, the hardware will start filling the empty
 * queue with new data if the capture device isn't paused.
 *
 * This will not prevent playback of queued audio that's already been sent to
 * the hardware, as we can not undo that, so expect there to be some fraction
 * of a second of audio that might still be heard. This can be useful if you
 * want to, say, drop any pending music or any unprocessed microphone input
 * during a level change in your game.
 *
 * You may not queue or dequeue audio on a device that is using an
 * application-supplied callback; calling this function on such a device
 * always returns 0. You have to use the audio callback or queue audio, but
 * not both.
 *
 * You should not call SDL_LockAudio() on the device before clearing the
 * queue; SDL handles locking internally for this function.
 *
 * This function always succeeds and thus returns void.
 *
 * \param dev the device ID of which to clear the audio queue
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetQueuedAudioSize
 * \sa SDL_QueueAudio
 * \sa SDL_DequeueAudio
 */
// extern DECLSPEC int SDLCALL SDL_ClearQueuedAudio(SDL_AudioDeviceID dev);
func SDL_ClearQueuedAudio(dev SDL_AudioDeviceID) (res ffcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_ClearQueuedAudio").Call(
		uintptr(dev),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 *  \name Audio lock functions
 *
 *  The lock manipulated by these functions protects the callback function.
 *  During a SDL_LockAudio()/SDL_UnlockAudio() pair, you can be guaranteed that
 *  the callback function is not running.  Do not call these from the callback
 *  function or you will cause deadlock.
 */
/* @{ */

/**
 * Use this function to lock out the audio callback function for a specified
 * device.
 *
 * The lock manipulated by these functions protects the audio callback
 * function specified in SDL_OpenAudioDevice(). During a
 * SDL_LockAudioDevice()/SDL_UnlockAudioDevice() pair, you can be guaranteed
 * that the callback function for that device is not running, even if the
 * device is not paused. While a device is locked, any other unpaused,
 * unlocked devices may still run their callbacks.
 *
 * Calling this function from inside your audio callback is unnecessary. SDL
 * obtains this lock before calling your function, and releases it when the
 * function returns.
 *
 * You should not hold the lock longer than absolutely necessary. If you hold
 * it too long, you'll experience dropouts in your audio playback. Ideally,
 * your application locks the device, sets a few variables and unlocks again.
 * Do not do heavy work while holding the lock for a device.
 *
 * It is safe to lock the audio device multiple times, as long as you unlock
 * it an equivalent number of times. The callback will not run until the
 * device has been unlocked completely in this way. If your application fails
 * to unlock the device appropriately, your callback will never run, you might
 * hear repeating bursts of audio, and SDL_CloseAudioDevice() will probably
 * deadlock.
 *
 * Internally, the audio device lock is a mutex; if you lock from two threads
 * at once, not only will you block the audio callback, you'll block the other
 * thread.
 *
 * \param dev the ID of the device to be locked
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_UnlockAudioDevice
 */
// extern DECLSPEC int SDLCALL SDL_LockAudioDevice(SDL_AudioDeviceID dev);
func SDL_LockAudioDevice(dev SDL_AudioDeviceID) (res ffcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_LockAudioDevice").Call(
		uintptr(dev),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Use this function to unlock the audio callback function for a specified
 * device.
 *
 * This function should be paired with a previous SDL_LockAudioDevice() call.
 *
 * \param dev the ID of the device to be unlocked
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_LockAudioDevice
 */
// extern DECLSPEC void SDLCALL SDL_UnlockAudioDevice(SDL_AudioDeviceID dev);
func SDL_UnlockAudioDevice(dev SDL_AudioDeviceID) {
	sdlcommon.GetSDL2Dll().NewProc("SDL_UnlockAudioDevice").Call(
		uintptr(dev),
	)
}

/* @} */ /* Audio lock functions */

/**
 * Use this function to shut down audio processing and close the audio device.
 *
 * The application should close open audio devices once they are no longer
 * needed. Calling this function will wait until the device's audio callback
 * is not running, release the audio hardware and then clean up internal
 * state. No further audio will play from this device once this function
 * returns.
 *
 * This function may block briefly while pending audio data is played by the
 * hardware, so that applications don't drop the last buffer of data they
 * supplied.
 *
 * The device ID is invalid as soon as the device is closed, and is eligible
 * for reuse in a new SDL_OpenAudioDevice() call immediately.
 *
 * \param dev an audio device previously opened with SDL_OpenAudioDevice()
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_OpenAudioDevice
 */
// extern DECLSPEC void SDLCALL SDL_CloseAudioDevice(SDL_AudioDeviceID dev);
func SDL_CloseAudioDevice(dev SDL_AudioDeviceID) {
	sdlcommon.GetSDL2Dll().NewProc("SDL_CloseAudioDevice").Call(
		uintptr(dev),
	)
	return
}

/**
 * Convert some audio data of one format to another format.
 *
 * \param src_format The format of the source audio
 * \param src_channels The number of channels of the source audio
 * \param src_rate The sampling rate of the source audio
 * \param src_data The audio data to be converted
 * \param src_len The len of src_data
 * \param dst_format The format of the desired audio output
 * \param dst_channels The number of channels of the desired audio output
 * \param dst_rate The sampling rate of the desired audio output
 * \param dst_data Will be filled with a pointer to converted audio data,
 *                 which should be freed with SDL_free().
 * \param dst_len Will be filled with the len of dst_data
 * \returns 0 on success or a negative error code on failure. On error,
 *          *dst_data will be NULL and so not allocated.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_CreateAudioStream
 */
// extern DECLSPEC int SDLCALL SDL_ConvertAudioSamples(SDL_AudioFormat src_format,
//                                                     Uint8 src_channels,
//                                                     int src_rate,
//                                                     const Uint8 *src_data,
//                                                     int src_len,
//                                                     SDL_AudioFormat dst_format,
//                                                     Uint8 dst_channels,
//                                                     int dst_rate,
//                                                     Uint8 **dst_data,
//                                                     int *dst_len);
func SDL_ConvertAudioSamples(src_format SDL_AudioFormat,
	src_channels ffcommon.FUint8T,
	src_rate ffcommon.FInt,
	src_data *ffcommon.FUint8T,
	src_len ffcommon.FInt,
	dst_format SDL_AudioFormat,
	dst_channels ffcommon.FUint8T,
	dst_rate ffcommon.FInt,
	dst_data **ffcommon.FUint8T,
	dst_len *ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_ConvertAudioSamples").Call(
		uintptr(src_format),
		uintptr(src_channels),
		uintptr(src_rate),
		uintptr(unsafe.Pointer(src_data)),
		uintptr(src_len),
		uintptr(dst_format),
		uintptr(dst_channels),
		uintptr(dst_rate),
		uintptr(unsafe.Pointer(dst_data)),
		uintptr(unsafe.Pointer(dst_len)),
	)
	res = ffcommon.FInt(t)
	return
}

// /* Ends C function definitions when using C++ */
// #ifdef __cplusplus
// }
// #endif
// #include <SDL3/SDL_close_code.h>

// #endif /* SDL_audio_h_ */
