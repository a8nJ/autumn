use std::{env, error};
fn main() -> Result<(), Box<dyn error::Error>> {
   let s1 = env::var("BROWSER")?;
   dbg!(s1);
   Ok(())
}
