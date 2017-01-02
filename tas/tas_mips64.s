#define LL(base, rt)	WORD	$((060<<26)|((base)<<21)|((rt)<<16))
#define LLV(base, rt)	WORD	$((064<<26)|((base)<<21)|((rt)<<16))
#define SC(base, rt)	WORD	$((070<<26)|((base)<<21)|((rt)<<16))
#define SCV(base, rt)	WORD	$((074<<26)|((base)<<21)|((rt)<<16))
#define SYNC WORD $0xf

TEXT ·testAndSet(SB),$0-24
	MOVV	lock+0(FP), R2
	MOVW	v+8(FP), R4
	SYNC
tas_again:
	MOVV	R4, R3
	LL(2, 1)	// R1 = *R2
	SC(2, 3)	// *R2 = R3
	BEQ	R3, tas_again
	MOVW	R1, ret+16(FP)
	SYNC
	RET
