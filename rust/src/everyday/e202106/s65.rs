struct Solution;
impl Solution {
    pub fn is_number(s: String) -> bool {
        let chars=vec![b'1',b'2',b'3',b'4',b'5',b'6',b'7',b'8',b'9',b'0'];
        let n=s.len();
        let mut has_e=false;
        for i  in 0..n  {
            let b=s.as_bytes()[i];
            if chars.contains(&b){
                continue;
            }
            if b==b'+'||b==b'-'{
                return  Solution::is_number(s[i+1..].to_string());
            }

            if b==b'.'{
                if has_e{
                    return false
                }
                if i==n-1{
                    if n>1{
                        return  true;
                    }else{
                        return false
                    }
                   
                }               
                let next=s.as_bytes()[i+1];
                if chars.contains(&next)||next==b'e'||next==b'E'{
                    continue;
                }
                return false
            }

            if b==b'e'||b==b'E'{
                if has_e{
                    return false;
                }
                has_e=true;                
                if i==n-1||i==0{
                    return false;
                } 
                continue;                
            }
            return  false;
        }
        
        return  true;
    }
}

#[test]
fn is_number_test(){
    assert_eq!(Solution::is_number("2e0".to_string()),true);
}