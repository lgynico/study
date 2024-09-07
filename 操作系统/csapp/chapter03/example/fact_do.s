	.file	"fact_do.c"
	.text
	.globl	fact_do
	.def	fact_do;	.scl	2;	.type	32;	.endef
	.seh_proc	fact_do
fact_do:
	.seh_endprologue
	movl	$1, %eax
.L2:
	imull	%ecx, %eax
	subl	$1, %ecx
	cmpl	$1, %ecx
	jg	.L2
	ret
	.seh_endproc
	.globl	fact_do_goto
	.def	fact_do_goto;	.scl	2;	.type	32;	.endef
	.seh_proc	fact_do_goto
fact_do_goto:
	.seh_endprologue
	movl	$1, %eax
.L4:
	imull	%ecx, %eax
	subl	$1, %ecx
	cmpl	$1, %ecx
	jg	.L4
	ret
	.seh_endproc
	.ident	"GCC: (x86_64-posix-seh-rev0, Built by MinGW-W64 project) 8.1.0"
