package sdl2

import (
	"unsafe"

	"github.com/moonfdd/ffmpeg-go/ffcommon"
	"github.com/moonfdd/sdl2-go/sdlcommon"
)

/* ! Definitions for test case structures */
const TEST_ENABLED = 1
const TEST_DISABLED = 0

/* ! Definition of all the possible test return values of the test case method */
const TEST_ABORTED = -1
const TEST_STARTED = 0
const TEST_COMPLETED = 1
const TEST_SKIPPED = 2

/* ! Definition of all the possible test results for the harness */
const TEST_RESULT_PASSED = 0
const TEST_RESULT_FAILED = 1
const TEST_RESULT_NO_ASSERT = 2
const TEST_RESULT_SKIPPED = 3
const TEST_RESULT_SETUP_FAILURE = 4

/* !< Function pointer to a test case setup function (run before every test) */
//typedef void (*SDLTest_TestCaseSetUpFp)(void *arg);
type SDLTest_TestCaseSetUpFp = func(arg ffcommon.FVoidP)

/* !< Function pointer to a test case function */
//typedef int (*SDLTest_TestCaseFp)(void *arg);
type SDLTest_TestCaseFp = func(arg ffcommon.FVoidP) ffcommon.FInt

/* !< Function pointer to a test case teardown function (run after every test) */
//typedef void  (*SDLTest_TestCaseTearDownFp)(void *arg);
type SDLTest_TestCaseTearDownFp = func(arg ffcommon.FVoidP)

/**
 * Holds information about a single test case.
 */
type SDLTest_TestCaseReference struct {

	/* !< Func2Stress */
	TestCase SDLTest_TestCaseFp
	/* !< Short name (or function name) "Func2Stress" */
	Name ffcommon.FBuf
	/* !< Long name or full description "This test pushes func2() to the limit." */
	Description ffcommon.FBuf
	/* !< Set to TEST_ENABLED or TEST_DISABLED (test won't be run) */
	Enabled ffcommon.FInt
}

/**
 * Holds information about a test suite (multiple test cases).
 */
type SDLTest_TestSuiteReference struct {

	/* !< "PlatformSuite" */
	name ffcommon.FBuf
	/* !< The function that is run before each test. NULL skips. */
	testSetUp SDLTest_TestCaseSetUpFp
	/* !< The test cases that are run as part of the suite. Last item should be NULL. */
	testCases **SDLTest_TestCaseReference
	/* !< The function that is run after each test. NULL skips. */
	testTearDown SDLTest_TestCaseTearDownFp
}

/**
 * \brief Generates a random run seed string for the harness. The generated seed will contain alphanumeric characters (0-9A-Z).
 *
 * Note: The returned string needs to be deallocated by the caller.
 *
 * \param length The length of the seed string to generate
 *
 * \returns the generated seed string
 */
//char *SDLTest_GenerateRunSeed(const int length);
func SDLTest_GenerateRunSeed(length ffcommon.FInt) (res ffcommon.FCharP) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDLTest_GenerateRunSeed").Call(
		uintptr(length),
	)
	if t == 0 {

	}
	res = ffcommon.StringFromPtr(t)
	return
}

/**
 * \brief Execute a test suite using the given run seed and execution key.
 *
 * \param testSuites Suites containing the test case.
 * \param userRunSeed Custom run seed provided by user, or NULL to autogenerate one.
 * \param userExecKey Custom execution key provided by user, or 0 to autogenerate one.
 * \param filter Filter specification. NULL disables. Case sensitive.
 * \param testIterations Number of iterations to run each test case.
 *
 * \returns the test run result: 0 when all tests passed, 1 if any tests failed.
 */
//int SDLTest_RunSuites(SDLTest_TestSuiteReference *testSuites[], const char *userRunSeed, Uint64 userExecKey, const char *filter, int testIterations);
func SDLTest_RunSuites(testSuites *[]SDLTest_TestSuiteReference, userRunSeed ffcommon.FConstCharP, userExecKey ffcommon.FUint64T, filter ffcommon.FConstCharP, testIterations ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDLTest_RunSuites").Call(
		uintptr(unsafe.Pointer(testSuites)),
		uintptr(unsafe.Pointer(ffcommon.BytePtrFromString(userRunSeed))),
		uintptr(userExecKey),
		uintptr(unsafe.Pointer(ffcommon.BytePtrFromString(filter))),
		uintptr(testIterations),
	)
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}
