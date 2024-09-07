	.file	"diff_se.c"
	.text
	.globl	absdiff_se
	.def	absdiff_se;	.scl	2;	.type	32;	.endef
	.seh_proc	absdiff_se
absdiff_se:
	.seh_endprologue
	cmpl	%edx, %ecx
	jge	.L2
	addl	$1, lt_cnt(%rip)
	movl	%edx, %eax
	subl	%ecx, %eax
.L1:
	ret
.L2:
	addl	$1, ge_cnt(%rip)
	movl	%ecx, %eax
	subl	%edx, %eax
	jmp	.L1
	.seh_endproc
	.globl	gotodiff_se
	.def	gotodiff_se;	.scl	2;	.type	32;	.endef
	.seh_proc	gotodiff_se
gotodiff_se:
	.seh_endprologue
	cmpl	%edx, %ecx
	jge	.L7
	addl	$1, lt_cnt(%rip)
	movl	%edx, %eax
	subl	%ecx, %eax
.L4:
	ret
.L7:
	addl	$1, ge_cnt(%rip)
	movl	%ecx, %eax
	subl	%edx, %eax
	jmp	.L4
	.seh_endproc
	.globl	ge_cnt
	.bss
	.align 4
ge_cnt:
	.space 4
	.globl	lt_cnt
	.align 4
lt_cnt:
	.space 4
	.ident	"GCC: (x86_64-posix-seh-rev0, Built by MinGW-W64 project) 8.1.0"
