fn main() {
   #[derive(Debug)]
   struct Day {
      s: Option<String>,
      n: Option<u8>
   }
   let o = Day {
      s: Some(String::from("Sunday")),
      n: None
   };
   println!("{:?}", o);
}
