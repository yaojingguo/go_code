// func castagnoliSSE42(crc uint32, p []byte) uint32
TEXT ·castagnoliSSE42(SB), NOSPLIT, $0
	MOVL crc+0(FP), AX    // CRC value
	MOVQ p+8(FP), SI      // data pointer
	MOVQ p_len+16(FP), CX // len(p)

	NOTL AX

	// If there's less than 8 bytes to process, we do it byte-by-byte.
	CMPQ CX, $8
	JL   cleanup

	// Process individual bytes until the input is 8-byte aligned.
startup:
