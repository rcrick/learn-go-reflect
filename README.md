# benchmark
`benchmark`将对象的创建也放到了基准测试中,`benchmark2`没有  
## benchmark

    | Func                              |    ops/s |    ns/op    |
    | --------------------------------- | -------: | :---------: |
    | Benchmark_Direct_Set-16           | 45673189 | 25.48 ns/op |
    | Benchmark_Reflect_Set-16          | 27042921 | 40.57 ns/op |
    | Benchmark_Reflect_FiledByIndex-16 | 20692100 | 52.14 ns/op |
    | Benchmark_Reflect_FiledByName-16  |  6617115 | 190.1 ns/op |





## benchmark2
    | Func                              |      ops/s |    ns/op     |
    | --------------------------------- | ---------: | :----------: |
    | Benchmark_Direct_Set-16           | 1000000000 | 0.2491 ns/op |
    | Benchmark_Reflect_Set-16          |  838860561 | 1.415 ns/op  |
    | Benchmark_Reflect_FiledByIndex-16 |   50546955 | 23.33 ns/op  |
    | Benchmark_Reflect_FiledByName-16  |    4385136 | 264.9 ns/op  |


## struct/interface call method
    | Func                                |      ops/s |    ns/op     |
    | ----------------------------------- | ---------: | :----------: |
    | Benchmark_StructPointerMetodCall-16 | 1000000000 | 0.2387 ns/op |
    | Benchmark_StructInterfaceCall-16    | 1000000000 | 1.208 ns/op  |

    | Func                    | ops/s |      ns/op      |
    | ----------------------- | ----: | :-------------: |
    | Benchmark_SortIface-16  |     6 | 166848069 ns/op |
    | Benchmark_SortStruct-16 |    19 | 59353636 ns/op  |