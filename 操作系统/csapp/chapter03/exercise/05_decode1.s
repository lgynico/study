	.file	"05_decode1.c"
	.text
	.globl	decode1
	.def	decode1;	.scl	2;	.type	32;	.endef
	.seh_proc	decode1
decode1:
	.seh_endprologue
	movl	(%rcx), %r10d
	movl	(%rdx), %r9d
	movl	(%r8), %eax
	movl	%r10d, (%rdx)
	movl	%r9d, (%r8)
	movl	%eax, (%rcx)
	ret
	.seh_endproc
	.ident	"GCC: (x86_64-win32-seh-rev0, Built by MinGW-W64 project) 8.1.0"
