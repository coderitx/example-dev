
// 定义结构体
struct User{
    active: bool,
    username: String,
    email: String,
    sign_in_count: u64,
}

#[allow(dead_code)]
pub fn struct_test(){
    let mut u = User{
        active: true,
        username: String::from("admin"),
        email: String::from("admin@admin.com"),
        sign_in_count: 999,
    };
    u.email = String::from("xx@xx.com");
}

#[allow(dead_code)]
fn build_user(user: String,email: String) -> User{
    User{
        active: false,
        username: user,
        email: email,
        sign_in_count: 1,
    }
}

// 元组结构体
#[allow(dead_code)]
struct Color(i32, i32, i32);
struct Point(i32, i32, i32);

#[allow(dead_code)]
fn tuple_struct(){
    let mut color = Color(10,20,30);
    let mut point = Point(11,22,33); 
}


// 类单元结构体
#[allow(dead_code)]
struct AlwaysEqual;


// 结构体的使用
#[allow(dead_code)]
struct Rectangle{
    width: u32,
    height: u32,
}

impl Rectangle{
    fn area(&self) -> u32{
        self.width * self.height
    }
}

pub fn print_area(){
    let rect = Rectangle{
        width: 20,
        height: 30,
    };
    println!("area = {}",rect.area());
}