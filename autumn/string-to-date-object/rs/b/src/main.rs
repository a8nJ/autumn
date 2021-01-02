use chrono::{
   TimeZone,
   Utc
};

fn main() -> Result<(), chrono::ParseError> {
   let s = "2020-12-31 23:59:59";
   let o  = Utc.datetime_from_str(s, "%Y-%m-%d %H:%M:%S")?;
   println!("{:?}", o);
   Ok(())
}
