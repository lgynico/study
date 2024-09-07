	.file	"fact_while.c"
	.text
	.globl	fact_while
	.def	fact_while;	.scl	2;	.type	32;	.endef
	.seh_proc	fact_while
fact_while:
	.seh_endprologue
	movl	$1, %eax
.L2:
	cmpl	$1, %ecx
	jle	.L4
	imull	%ecx, %eax
	subl	$1, %ecx
	jmp	.L2
.L4:
	ret
	.seh_endproc
	.globl	fact_while_jm_goto
	.def	fact_while_jm_goto;	.scl	2;	.type	32;	.endef
	.seh_proc	fact_while_jm_goto
fact_while_jm_goto:
	.seh_endprologue
	movl	$1, %eax
.L6:
	cmpl	$1, %ecx
	jle	.L5
	imull	%ecx, %eax
	subl	$1, %ecx
	jmp	.L6
.L5:
	ret
	.seh_endproc
	.globl	fact_while_gd_goto
	.def	fact_while_gd_goto;	.scl	2;	.type	32;	.endef
	.seh_proc	fact_while_gd_goto
fact_while_gd_goto:
	.seh_endprologue
	cmpl	$1, %ecx
	jle	.L11
	movl	$1, %eax
.L10:
	imull	%ecx, %eax
	subl	$1, %ecx
	cmpl	$1, %ecx
	jne	.L10
.L8:
	ret
.L11:
	movl	$1, %eax
.L9:
	jmp	.L8
	.seh_endproc
	.ident	"GCC: (x86_64-posix-seh-rev0, Built by MinGW-W64 project) 8.1.0"
