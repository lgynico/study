	.file	"diff.c"
	.text
	.globl	absdiff
	.def	absdiff;	.scl	2;	.type	32;	.endef
	.seh_proc	absdiff
absdiff:
	.seh_endprologue
	cmpl	%edx, %ecx
	jge	.L2
	movl	%edx, %eax
	subl	%ecx, %eax
.L1:
	ret
.L2:
	movl	%ecx, %eax
	subl	%edx, %eax
	jmp	.L1
	.seh_endproc
	.globl	cmovdiff
	.def	cmovdiff;	.scl	2;	.type	32;	.endef
	.seh_proc	cmovdiff
cmovdiff:
	.seh_endprologue
	movl	%edx, %r8d
	subl	%ecx, %r8d
	movl	%ecx, %eax
	subl	%edx, %eax
	cmpl	%ecx, %edx
	jle	.L4
	movl	%r8d, %eax
.L4:
	ret
	.seh_endproc
	.ident	"GCC: (x86_64-posix-seh-rev0, Built by MinGW-W64 project) 8.1.0"
