#define LL(base, rt)	WORD	$((060<<26)|((base)<<21)|((rt)<<16))
#define LLV(base, rt)	WORD	$((064<<26)|((base)<<21)|((rt)<<16))
#define SC(base, rt)	WORD	$((070<<26)|((base)<<21)|((rt)<<16))
#define SCV(base, rt)	WORD	$((074<<26)|((base)<<21)|((rt)<<16))
#define SYNC WORD $0xf

TEXT ·increment(SB),$0-24
	MOVV	k+0(FP), R2
	SYNC
inc_again:
	LL(2, 1)	// R1 = *R2
	ADDU	$1, R1, R3
	SC(2, 3)	// *R2 = R3
	BEQ	R3, inc_again
	SYNC
	RET
