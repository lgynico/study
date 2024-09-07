// 寄存器：
//      %rax = 0x100
//      %rcx = 0x1
//      %rdx = 0x3
// 地址：
//      0x100 = 0xFF
//      0x104 = 0xAB
//      0x108 = 0x13
//      0x10C = 0x11
// 求：
//      %rax
//      0x104
//      $0x108
//      (%rax)
//      4(%rax)
//      9(%rax, %rdx)
//      260(%rcx, %rdx)
//      0xFC(, %rcx, 4)
//      (%rax, %rdx, 4)