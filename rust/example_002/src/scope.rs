
#[allow(dead_code)]
pub fn str_scope(){
    let mut s = String::from("hello");
    s.push_str(", world!");
    println!("{s}");

    {
        let _s = String::from("scope...");
        println!("作用域失效");
    }
}

#[allow(dead_code)]
pub fn str_copy(){
    let s1 = String::from("Rust");

    // 浅拷贝
    // s1 被 移动 到了 s2 中
    // s1 失效（ -- move occurs because `s1` has type `String`, which does not implement the `Copy` trait）
    let s2 = s1;
    println!("s2 = {s2}");

    // 深拷贝： clone()
    let s3 = s2.clone();
    println!("s3 = {s3}");

    // 栈空间的数据不会发生移动
    // 原因是像整型这样的在编译时已知大小的类型被整个存储在栈上，所以拷贝其实际的值是快速的
    let number = 99;
    let number1 = number;
    println!("number = {number} number1 = {number1}");
}

// 比较str类型和int类型copy过程
#[allow(dead_code)]
pub fn str_int_copy_diff() {
    let s = String::from("hello");  // s 进入作用域

    takes_ownership(s);             // s 的值移动到函数里 ...
    // ... 所以到这里不再有效 再次调用会编译失败
    // println!(s);                                    
    let x = 5;                      // x 进入作用域

    makes_copy(x);                  // x 应该移动函数里，
                                    // 但 i32 是 Copy 的，
                                    // 所以在后面可继续使用 x
} // 这里，x 先移出了作用域，然后是 s。但因为 s 的值已被移走，
  // 没有特殊之处


// 值的移动，作用域的转移
#[allow(dead_code)]
pub fn value_move(){
    let s1 = gives_ownership();         // gives_ownership 将返回值
                                        // 转移给 s1
    println!("{s1}");
    let s2 = String::from("hello");     // s2 进入作用域

    let s3 = takes_and_gives_back(s2);  // s2 被移动到
                                        // takes_and_gives_back 中，
                                        // 它也将返回值移给 s3
    println!("{s3}");
} // 这里，s3 移出作用域并被丢弃。s2 也移出作用域，但已被移走，
  // 所以什么也不会发生。s1 离开作用域并被丢弃


fn takes_ownership(some_string: String){
    // some_string 进入作用域
    println!("{}", some_string);
} // 这里，some_string 移出作用域并调用 `drop` 方法。
  // 占用的内存被释放


fn makes_copy(some_integer: i32) { 
    // some_integer 进入作用域
    println!("{}", some_integer);
} // 这里，some_integer 移出作用域。没有特殊之处


fn gives_ownership() -> String {             // gives_ownership 会将
    // 返回值移动给
    // 调用它的函数
    let some_string = String::from("yours"); // some_string 进入作用域。

    some_string                              // 返回 some_string 
        // 并移出给调用的函数
        // 
}

// takes_and_gives_back 将传入字符串并返回该值
fn takes_and_gives_back(a_string: String) -> String { // a_string 进入作用域
    a_string  // 返回 a_string 并移出给调用的函数
}