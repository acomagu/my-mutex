#define LL(base, rt)	WORD	$((060<<26)|((base)<<21)|((rt)<<16))
#define LLV(base, rt)	WORD	$((064<<26)|((base)<<21)|((rt)<<16))
#define SC(base, rt)	WORD	$((070<<26)|((base)<<21)|((rt)<<16))
#define SCV(base, rt)	WORD	$((074<<26)|((base)<<21)|((rt)<<16))
#define SYNC WORD $0xf

TEXT ·_lock(SB),$0-24
	MOVV	flag+0(FP), R2
	MOVV	$1, R4
	SYNC
lock_again:
	MOVV	R4, R3
	LL(2, 1)	// R1 = *R2
	SC(2, 3)	// *R2 = R3
	BEQ	R3, lock_again
	BEQ	R4, R1, lock_again
	SYNC
	RET
