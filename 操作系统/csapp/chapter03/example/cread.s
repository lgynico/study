	.file	"cread.c"
	.text
	.globl	cread
	.def	cread;	.scl	2;	.type	32;	.endef
	.seh_proc	cread
cread:
	.seh_endprologue
	testq	%rcx, %rcx
	je	.L3
	movl	(%rcx), %eax
.L1:
	ret
.L3:
	movl	$0, %eax
	jmp	.L1
	.seh_endproc
	.ident	"GCC: (x86_64-posix-seh-rev0, Built by MinGW-W64 project) 8.1.0"
