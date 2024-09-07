	.file	"switch_eg.c"
	.text
	.globl	switch_eg
	.def	switch_eg;	.scl	2;	.type	32;	.endef
	.seh_proc	switch_eg
switch_eg:
	.seh_endprologue
	subl	$100, %edx
	cmpl	$6, %edx
	ja	.L8
	movl	%edx, %edx
	leaq	.L4(%rip), %r9
	movslq	(%r9,%rdx,4), %rax
	addq	%r9, %rax
	jmp	*%rax
	.section .rdata,"dr"
	.align 4
.L4:
	.long	.L7-.L4
	.long	.L8-.L4
	.long	.L6-.L4
	.long	.L5-.L4
	.long	.L3-.L4
	.long	.L8-.L4
	.long	.L3-.L4
	.text
.L7:
	leal	(%rcx,%rcx,2), %eax
	leal	(%rcx,%rax,4), %ecx
	jmp	.L2
.L6:
	addl	$10, %ecx
.L5:
	addl	$11, %ecx
.L2:
	movl	%ecx, (%r8)
	ret
.L3:
	imull	%ecx, %ecx
	jmp	.L2
.L8:
	movl	$0, %ecx
	jmp	.L2
	.seh_endproc
	.globl	switch_eg_impl
	.def	switch_eg_impl;	.scl	2;	.type	32;	.endef
	.seh_proc	switch_eg_impl
switch_eg_impl:
	.seh_endprologue
	subl	$100, %edx
	cmpl	$6, %edx
	ja	.L10
	leaq	jt.1986(%rip), %rax
	movl	%edx, %edx
	jmp	*(%rax,%rdx,8)
.L11:
	leal	(%rcx,%rcx,2), %eax
	leal	(%rcx,%rax,4), %ecx
	jmp	.L12
.L13:
	addl	$10, %ecx
.L14:
	addl	$11, %ecx
	jmp	.L12
.L15:
	imull	%ecx, %ecx
	jmp	.L12
.L10:
	movl	$0, %ecx
.L12:
	movl	%ecx, (%r8)
	ret
	.seh_endproc
	.section .rdata,"dr"
	.align 32
jt.1986:
	.quad	.L11
	.quad	.L10
	.quad	.L13
	.quad	.L14
	.quad	.L15
	.quad	.L10
	.quad	.L15
	.ident	"GCC: (x86_64-posix-seh-rev0, Built by MinGW-W64 project) 8.1.0"
