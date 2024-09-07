	.file	"scale.c"
	.text
	.globl	scale
	.def	scale;	.scl	2;	.type	32;	.endef
	.seh_proc	scale
scale:
	.seh_endprologue
	leal	(%rcx,%rdx,4), %eax
	leal	(%r8,%r8,2), %ecx
	leal	0(,%rcx,4), %edx
	addl	%edx, %eax
	ret
	.seh_endproc
	.ident	"GCC: (x86_64-win32-seh-rev0, Built by MinGW-W64 project) 8.1.0"
