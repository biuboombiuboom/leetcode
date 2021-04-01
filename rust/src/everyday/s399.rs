fn calc_equation(
    equations: Vec<Vec<String>>,
    values: Vec<f64>,
    queries: Vec<Vec<String>>,
) -> Vec<f64> {
    let mut ans: Vec<f64> = Vec::new();
    let mut list_equation: Vec<Vec<String>> = Vec::new();
    let n= equations.len();

    for i in 0..n {
        let equation = &equations[i];
        let n1=list_equation.len();
       for j in 0..n1 {
           
       }
        
    }

    for query in queries {
        let qa = &query[0];
        let qb = &query[1];
        
    }
    ans
}

pub fn run() {
    let equations = vec![
        vec![String::from("a"), String::from("b")],
        vec![String::from("b"), String::from("c")],
    ];
    let values = vec![1.2, 1.3];
    let queries = vec![
        vec![String::from("a"), String::from("c")],
        vec![String::from("b"), String::from("a")],
        vec![String::from("a"), String::from("e")],
        vec![String::from("a"), String::from("a")],
        vec![String::from("x"), String::from("x")],
    ];
    let rs = calc_equation(equations, values, queries);
    for item in rs {
        println!("{}", item);
    }
}

#[test]
fn calc_equation_test() {}