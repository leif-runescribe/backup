use std::io;

fn main() {
    loop{
        println!("select");
        println!("1. Fahrenheit to Celsius");
        println!("2. Celsius to Fahrenheit");

        let mut inp = String::new();
        io::stdin().
        read_line(&mut inp)
        .expect("error, retry!");
        let inp = inp.trim();
        let inp = match inp{
            "1" => 1,
            "2" => 2,
            _ => {
                println!("Please input 1 or 2!");
                continue;
            }
        };
        
        println!("Now, enter the temperature u wish to convert");
        let mut temp = String::new();

        io::stdin().
        read_line(&mut temp).
        expect("error, retry");
        let temp: f64 = match temp.trim().parse(){
            Ok(num) => num,
            Err(_) => {
                println!("invalid, you are not valid, unvalid, disvalid");
                continue;
            }
        };
        let res;
        if inp ==1 {
            res = (temp-32.)/1.8;
        }
        else {
            res = temp * 1.8 + 32.;
        }
        println!("converted output: {}", res);
    }
}