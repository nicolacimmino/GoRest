#include "textflag.h"

// Implementation of func add(x, y int64) int64, see .go file.
//
// With the TEXT entry we declare the function name and its attributes.
// ·    The middle dot (U+00B7) is the separator between the package name
//      (implicitely the current package here) and the function name.
// SB   the Static Base, a pointer to the beginning of our program.
// $0   the stack size, in bytes, in this case we don't need  a stack.
// 24   the arguments+return value(s) size, in bytes. In this case we have
//      two int64 as args and one int64 as return, so 24 bytes.
TEXT ·add(SB),$0-24
    MOVQ x+0(FP), BX        // Move a Quadword, the first parameter, into register BX
                            // x+0(FP):
                            // x        This is just a label, it's good practice to have it
                            //          as the parameter name in the function declaration
                            // +0(FP)   The value with offset 0 into the memory pointed by
                            //          the Frame Pointer. This is the first argument of
                            //          the function call.

    MOVQ y+8(FP), BP        // As above for the second parameter of the function. This is
                            // at offset +8(FP) as the first argument takes 8 bytes (int64).

    ADDQ BP, BX             // Add the two Quad words in BP and BX.

    MOVQ BX, ret+16(FP)     // Return the value. Also here the label name "ret" is irrelevant
                            // however the convention is to call it "ret". The return value
                            // goes into the memory after the parameters, so +16(FP).

    RET
