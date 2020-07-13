	.section	__TEXT,__text,regular,pure_instructions
	.macosx_version_min 10, 14
	.globl	_tailcaller             ## -- Begin function tailcaller
	.p2align	4, 0x90
_tailcaller:                            ## @tailcaller
	.cfi_startproc
## %bb.0:
                                        ## kill: def $esi killed $esi def $rsi
                                        ## kill: def $edi killed $edi def $rdi
	leal	(%rdi,%rsi), %ecx
                                        ## kill: def $esi killed $esi killed $rsi
	movl	%edi, %edx
	jmp	_tailcallee             ## TAILCALL
	.cfi_endproc
                                        ## -- End function

.subsections_via_symbols
