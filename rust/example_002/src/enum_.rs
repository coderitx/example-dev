
// 定义一个枚举
enum IpAddrKind {
    V4,
    V6,
}

#[allow(dead_code)]
pub fn enum_test(){
    let four = IpAddrKind::V4;
    let six = IpAddrKind::V6;
}


enum IpAddr {
    V4(String),
    V6(String),
}

pub fn enum_test1(){
    let home = IpAddr::V4(String::from("127.0.0.1"));

    let loopback = IpAddr::V6(String::from("::1"));
}
