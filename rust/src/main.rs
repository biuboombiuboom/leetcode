mod dp;
mod easy;
mod everyday;
mod normal;
mod utils;

fn main() {
    // easy::solution::run();

    // everyday::solution::run();

    // let cmd_string: String = String::from(
    //     "tar -zvxf \"/tmp/test abc.tar.gz\" -C /tmp \"/tmp/test abc.tar.gz\" --force-local",
    // );

    // let mut temp_string = String::default();
    // let mut ans: Vec<String> = Vec::new();
    // let mut find_pre = false;
    // for b in cmd_string.as_bytes().iter() {
    //     if b == &b' ' && !find_pre {
    //         ans.push(temp_string.clone());
    //         temp_string.clear();
    //         continue;
    //     }
    //     temp_string.push(*b as char);
    //     if b == &b'"' {
    //         find_pre = !find_pre;
    //     }
    // }
    let mut ttt: Vec<String> = Vec::new();
    let mut ans1: Vec<&str> = Vec::new();
    ttt.push("".to_string());

    let sss1 = ttt.join("");
    let s = sss1.as_str();
    // let mut sss = sss1.as_str();
    // sss = ttt.join("").as_str();

    println!("{}", sss1);

    // let cmd_vec = cmd_string.split("\" ");

    // let mut find_pre = false;
    // for cmd in cmd_vec {
    //     if cmd.starts_with("\"") {
    //         temp_string.push_str(cmd);
    //         find_pre = true;
    //         continue;
    //     }
    //     if find_pre {
    //         temp_string.push_str(" ");
    //         temp_string.push_str(cmd);
    //         if cmd.ends_with("\"") {
    //             find_pre = false;
    //             ans.push(temp_string.clone());
    //             temp_string.clear();
    //         }
    //         continue;
    //     }
    //     ans.push(cmd.to_string());
    // }
    // for part in ans {
    //     println!("{}", part);
    // }
}
