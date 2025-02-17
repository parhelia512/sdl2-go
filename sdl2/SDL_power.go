package sdl2

import (
	"unsafe"

	"github.com/moonfdd/ffmpeg-go/ffcommon"
	"github.com/moonfdd/sdl2-go/sdlcommon"
)

/**
 *  The basic state for the system's power supply.
 */
type SDL_PowerState = int32

const (
	SDL_POWERSTATE_UNKNOWN    = iota /**< cannot determine power status */
	SDL_POWERSTATE_ON_BATTERY        /**< Not plugged in, running on the battery */
	SDL_POWERSTATE_NO_BATTERY        /**< Plugged in, no battery available */
	SDL_POWERSTATE_CHARGING          /**< Plugged in, charging battery */
	SDL_POWERSTATE_CHARGED           /**< Plugged in, battery charged */
)

/**
 * Get the current power supply details.
 *
 * You should never take a battery status as absolute truth. Batteries
 * (especially failing batteries) are delicate hardware, and the values
 * reported here are best estimates based on what that hardware reports. It's
 * not unffcommon for older batteries to lose stored power much faster than it
 * reports, or completely drain when reporting it has 20 percent left, etc.
 *
 * Battery status can change at any time; if you are concerned with power
 * state, you should call this function frequently, and perhaps ignore changes
 * until they seem to be stable for a few seconds.
 *
 * It's possible a platform can only report battery percentage or time left
 * but not both.
 *
 * \param secs seconds of battery life left, you can pass a NULL here if you
 *             don't care, will return -1 if we can't determine a value, or
 *             we're not running on a battery
 * \param pct percentage of battery life left, between 0 and 100, you can pass
 *            a NULL here if you don't care, will return -1 if we can't
 *            determine a value, or we're not running on a battery
 * \returns an SDL_PowerState enum representing the current battery state.
 */
//extern DECLSPEC SDL_PowerState SDLCALL SDL_GetPowerInfo(int *secs, int *pct);
func SDL_GetPowerInfo(secs, pct *ffcommon.FInt) (res SDL_PowerState) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetPowerInfo").Call(
		uintptr(unsafe.Pointer(secs)),
		uintptr(unsafe.Pointer(pct)),
	)
	if t == 0 {

	}
	res = SDL_PowerState(t)
	return
}
