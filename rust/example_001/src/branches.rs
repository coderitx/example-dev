
#[allow(dead_code)]
pub fn branches(num: i32){
    if num < 50{
        println!("{num} 小于50");
    }else{
        println!("{num} 大于 50");
    }

    // let number = 3;
    // number 不是bool类型会出错
    // if number {
    //     println!("number was three");
    // }
}

// 多分支判断
#[allow(dead_code)]
pub fn branch_by_else_if(num: i32){
    if num % 5 == 0 {
        println!("{num} 可以被5整除");
    }else if num % 4 == 0{
        println!("{num} 可以被4整除");
    }else if num % 3 == 0 {
        println!("{num} 可以被3整除");
    }   
}


// 在let中使用if-else
#[allow(dead_code)]
pub fn branch_by_let_if(){
    let ok: bool = false;
    let number: i32 = if ok {999} else {-999};
    println!("number value is :{number}")
}

// loop 循环
#[allow(dead_code)]
pub fn branch_by_loop(){
    let mut num = 0;
    let result = loop {
        num += 1;
        if num == 10{
            break num * 2;
        }
    };
    println!("result value is {result} num value is: {num}")
}


// loop 标签
#[allow(dead_code)]
pub fn branch_by_loop_lable(){
    let mut count = 0;
    'coutinue_lable: loop{
        println!("count = {count}");
        let mut remaining = 10;
        loop {
            println!("remaining = {remaining}");
            if remaining == 5{
                break;
            }
            if count == 5{
                break 'coutinue_lable;
            }
            remaining -= 1;
        }
        count += 1;
    }
    println!("End count = {count}");
}


// for 循环
#[allow(dead_code)]
pub fn branch_by_for(){
    let _arr = [1,2,4,5,6];
    for val in _arr{
        println!("value = {val}");
    }

    // for number in -10..-5
    for number in (-10..-5).rev(){
        println!("numebr = {number}");
    }
}

// while 循环
#[allow(dead_code)]
pub fn branch_by_while(){
    let mut count = 10;
    while count > 0 {
        println!("count value is : {count}");
        count = count - 1;
    }
    println!("while end.");
}