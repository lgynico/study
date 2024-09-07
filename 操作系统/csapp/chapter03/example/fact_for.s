	.file	"fact_for.c"
	.text
	.globl	fact_for
	.def	fact_for;	.scl	2;	.type	32;	.endef
	.seh_proc	fact_for
fact_for:
	.seh_endprologue
	movl	$1, %eax
	movl	$2, %edx
.L2:
	cmpl	%ecx, %edx
	jg	.L4
	imull	%edx, %eax
	addl	$1, %edx
	jmp	.L2
.L4:
	ret
	.seh_endproc
	.globl	fact_for_while
	.def	fact_for_while;	.scl	2;	.type	32;	.endef
	.seh_proc	fact_for_while
fact_for_while:
	.seh_endprologue
	movl	$1, %eax
	movl	$2, %edx
.L6:
	cmpl	%ecx, %edx
	jg	.L8
	imull	%edx, %eax
	addl	$1, %edx
	jmp	.L6
.L8:
	ret
	.seh_endproc
	.globl	fact_for_jm_goto
	.def	fact_for_jm_goto;	.scl	2;	.type	32;	.endef
	.seh_proc	fact_for_jm_goto
fact_for_jm_goto:
	.seh_endprologue
	movl	$1, %eax
	movl	$2, %edx
.L10:
	cmpl	%ecx, %edx
	jg	.L9
	imull	%edx, %eax
	addl	$1, %edx
	jmp	.L10
.L9:
	ret
	.seh_endproc
	.ident	"GCC: (x86_64-posix-seh-rev0, Built by MinGW-W64 project) 8.1.0"
