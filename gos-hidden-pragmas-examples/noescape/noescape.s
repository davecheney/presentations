#include "textflag.h"

TEXT ·length(SB), NOSPLIT, $0-24
	MOVQ len+8(FP), SI
	MOVB SI, ret+16(FP)
	RET

