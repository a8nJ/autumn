use std::collections::VecDeque;

fn main() {
   let mut a = VecDeque::new();
   a.push_back(10);
   a.push_back(11);
   println!("{:?}", a);
}
