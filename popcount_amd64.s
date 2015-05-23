// func Fast32(i uint32) uint8
TEXT ·Fast32(SB),$0
    MOVL i+0(FP), BP        // i
    MOVL BP, BX             // i
    SHRL $1, BX             // i >> 1
    ANDL $0x055555555, BX   // (i >> 1) & 0x55555555
    SUBL BX, BP             // w = i - ((i >> 1) & 0x55555555)
    MOVL BP, AX             // w
    SHRL $2, BP             // w >> 2
    ANDL $0x033333333, AX   // w & 0x33333333
    ANDL $0x033333333, BP   // (w >> 2) & 0x33333333
    ADDL BP, AX             // x = (w & 0x33333333) + ((w >> 2) & 0x33333333)
    MOVL AX, BX             // x
    SHRL $4, BX             // x >> 4
    ADDL AX, BX             // x + (x >> 4)
    ANDL $0x00F0F0F0F, BX   // y = (x + (x >> 4) & 0x0F0F0F0F)
    IMULL $0x001010101, BX  // y * 0x01010101
    SHRL $24, BX            // population count = (y * 0x01010101) >> 24
    MOVL BX, toReturn+8(FP) // Store result.
    RET                     // return

// func Fast64(i uint32) uint8
TEXT ·Fast64(SB),$0
    MOVQ i+0(FP), BP        // i
    MOVQ BP, BX             // i
    SHRQ $1, BX             // i >> 1
    MOVQ $0x5555555555555555, R11
    ANDQ R11, BX            // (i >> 1) & 0x55555555
    SUBQ BX, BP             // w = i - ((i >> 1) & 0x55555555)
    MOVQ BP, AX             // w
    SHRQ $2, BP             // w >> 2
    MOVQ $0x3333333333333333, R11
    ANDQ R11, AX            // w & 0x33333333
    ANDQ R11, BP            // (w >> 2) & 0x33333333
    ADDQ BP, AX             // x = (w & 0x33333333) + ((w >> 2) & 0x33333333)
    MOVQ AX, BX             // x
    SHRQ $4, BX             // x >> 4
    ADDQ AX, BX             // x + (x >> 4)
    MOVQ $0x0F0F0F0F0F0F0F0F, R11
    ANDQ R11, BX            // y = (x + (x >> 4) & 0x0F0F0F0F)
    MOVQ $0x0101010101010101, R11
    IMULQ R11, BX           // y * 0x01010101
    SHRQ $56, BX            // population count = (y * 0x01010101) >> 24
    MOVQ BX, toReturn+8(FP) // Store result.
    RET                     // return

// func PopCnt32(i uint32) uint8
// Using POPCNT instruction, nehalem and later
// See: 
// http://www.felixcloutier.com/x86/POPCNT.html
// http://wiki.osdev.org/X86-64_Instruction_Encoding#Registers  
TEXT ·PopCnt32(SB),$0
    MOVL i+0(FP), BP        // i
    BYTE $0xF3; BYTE $0x0F; BYTE $0xB8; BYTE $0xdd  // POPCNT BX, BP
    // BX<-BP = 11011101 = 0xdd
    MOVL BX, toReturn+8(FP) // Store result.
    RET                     // return

// func PopCnt64(i uint64) uint8
TEXT ·PopCnt64(SB),$0
    MOVQ i+0(FP), BP        // i
    BYTE $0xF3; BYTE $0x48; BYTE $0x0F; BYTE $0xB8; BYTE $0xdd  // POPCNT RBX, RBP
    // 01001000 = 0x48
    // BX<-BP = 11011101 = 0xdd
    MOVQ BX, toReturn+8(FP) // Store result.
    RET                     // return
