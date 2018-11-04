#JVM

## 助记符
首字母 | 数据类型 | 例子
---- | --- | ---
a | reference | aload, astore, areturn
b | byte/boolean | bipush, baload
c | char | caload, castore
d | double | dload, dstore, dadd
f | float | fload, fstore, fadd
i | int | iload, istore, iadd
l | long | load, lstore, ladd
s | short | sipush, sastore

## 11类
1. 常量操作，constant
1. 加载指令，loads
1. 存储指令，store
1. 操作数栈指令，stack
1. 数学指令，math
1. 转换指令，conversions
1. 比较指令，comparision
1. 控制指令，control
1. 引用指令，references
1. 扩展指令，extended
1. 保留指令，reserved
    1. breakpoint, 0XCA, 202
    2. impdep1, OXFE, 254
    3. impdep2, OXFF, 266
