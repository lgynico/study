	.file	"arith.c"
	.text
	.globl	arith
	.def	arith;	.scl	2;	.type	32;	.endef
	.seh_proc	arith
arith:
	.seh_endprologue
	xorl	%edx, %ecx
	leal	(%r8,%r8,2), %edx
	movl	%edx, %eax
	sall	$4, %eax
	andl	$252645135, %ecx
	subl	%ecx, %eax
	ret
	.seh_endproc
	.ident	"GCC: (x86_64-win32-seh-rev0, Built by MinGW-W64 project) 8.1.0"
