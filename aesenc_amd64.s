
// func aesenc(k, s*byte)
TEXT ·aesenc(SB),$0
    MOVQ k+0(FP), AX
    MOVQ s+8(FP), BX
    MOVUPS 0(AX), X1
    MOVUPS 0(BX), X0
    AESENC X1, X0
    MOVUPS X0, 0(BX)
    RET

