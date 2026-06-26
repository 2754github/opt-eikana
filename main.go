package main

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Carbon -framework CoreGraphics -framework CoreFoundation

#include <Carbon/Carbon.h>
#include <CoreGraphics/CoreGraphics.h>

#define MY_kVK_JIS_Eisu    0x66
#define MY_kVK_JIS_Kana    0x68
#define MY_kVK_Option      0x3A
#define MY_kVK_RightOption 0x3D

static CFMachPortRef tap   = NULL;
static int otherKeyPressed = 0;
static int leftOptionDown  = 0;
static int rightOptionDown = 0;

static void sendKey(CGKeyCode keyCode) {
    CGEventRef down = CGEventCreateKeyboardEvent(NULL, keyCode, true);
    CGEventRef up   = CGEventCreateKeyboardEvent(NULL, keyCode, false);
    CGEventPost(kCGSessionEventTap, down);
    CGEventPost(kCGSessionEventTap, up);
    CFRelease(down);
    CFRelease(up);
}

static CGEventRef eventCallback(CGEventTapProxy proxy, CGEventType type, CGEventRef event, void *refcon) {
    if (type == kCGEventTapDisabledByTimeout || type == kCGEventTapDisabledByUserInput) {
        CGEventTapEnable(tap, true);
        return event;
    }

    CGKeyCode keyCode = (CGKeyCode)CGEventGetIntegerValueField(event, kCGKeyboardEventKeycode);

    if (type == kCGEventKeyDown) {
        if (leftOptionDown || rightOptionDown) {
            otherKeyPressed = 1;
        }
    } else if (type == kCGEventFlagsChanged) {
        CGEventFlags flags = CGEventGetFlags(event);

        if (keyCode == MY_kVK_Option && !(flags & kCGEventFlagMaskAlternate)) {
            if (leftOptionDown && !otherKeyPressed) {
                sendKey(MY_kVK_JIS_Eisu);
            }
            leftOptionDown = 0;
            if (!rightOptionDown) otherKeyPressed = 0;
            return NULL;
        } else if (keyCode == MY_kVK_RightOption && !(flags & kCGEventFlagMaskAlternate)) {
            if (rightOptionDown && !otherKeyPressed) {
                sendKey(MY_kVK_JIS_Kana);
            }
            rightOptionDown = 0;
            if (!leftOptionDown) otherKeyPressed = 0;
            return NULL;
        } else if (keyCode == MY_kVK_Option) {
            leftOptionDown = 1;
            return NULL;
        } else if (keyCode == MY_kVK_RightOption) {
            rightOptionDown = 1;
            return NULL;
        }
    }

    return event;
}

static void startEventTap() {
    CGEventMask mask = CGEventMaskBit(kCGEventKeyDown)
                     | CGEventMaskBit(kCGEventFlagsChanged);

    tap = CGEventTapCreate(
        kCGSessionEventTap,
        kCGHeadInsertEventTap,
        kCGEventTapOptionDefault,
        mask,
        eventCallback,
        NULL
    );

    if (!tap) {
        fprintf(stderr, "Failed to create CGEventTap. Please check your system permissions.\n");
        exit(1);
    }

    CFRunLoopSourceRef src = CFMachPortCreateRunLoopSource(kCFAllocatorDefault, tap, 0);
    CFRunLoopAddSource(CFRunLoopGetCurrent(), src, kCFRunLoopCommonModes);
    CGEventTapEnable(tap, true);

    CFRunLoopRun();
}
*/
import (
	"C"
)

func main() {
	C.startEventTap()
}
