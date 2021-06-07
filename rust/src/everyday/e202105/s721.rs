use std::collections::HashMap;

fn accounts_merge(accounts: Vec<Vec<String>>) -> Vec<Vec<String>> {
    let mut email_indexs = HashMap::new();
    let mut email_name: HashMap<String, String> = HashMap::new();
    let n = accounts.len();
    for i in 0..n {
        let emails = accounts.get(i).unwrap();
        let an = emails.len();
        let name = emails.get(0).unwrap();
        for j in 1..an {
            let email = &emails[j];
            match email_indexs.get(email) {
                None => {
                    email_indexs.insert(email, email_indexs.len());
                    email_name.insert(email.to_string(), name.to_string());
                }
                _ => {}
            }
        }
    }

    let mut uf = UnionFind::new(email_indexs.len());
    for i in 0..n {
        let emails = accounts.get(i).unwrap();
        let first_email = emails.get(1).unwrap();
        let first_index = email_indexs.get(first_email).unwrap();
        let size = emails.len();
        for j in 2..size {
            let next_email = emails.get(j).unwrap();
            let next_index = email_indexs.get(next_email).unwrap();
            uf.union(*first_index, *next_index);
        }
    }

    let mut index_emails: HashMap<usize, Vec<String>> = HashMap::new();
    for (k, v) in &email_indexs {
        let index = uf.find(*v);
        let account = index_emails.entry(index).or_insert(vec![k.to_string()]);
        if account[account.len() - 1] != k.to_string() {
            account.push(k.to_string());
        }
    }

    let mut ans: Vec<Vec<String>> = Vec::new();

    for (_, mut emails) in index_emails {
        emails.sort();
        emails.insert(0, email_name[&emails[0]].to_string());
        ans.push(emails);
    }
    ans
}

struct UnionFind {
    parent: Vec<usize>,
}

impl UnionFind {
    fn new(n: usize) -> Self {
        let mut parent: Vec<usize> = vec![];
        for i in 0..n {
            parent.push(i);
        }
        Self { parent }
    }

    fn union(&mut self, index1: usize, index2: usize) {
        let index = self.find(index2);
        self.parent[index] = self.find(index1);
    }

    fn find(&mut self, index: usize) -> usize {
        if *self.parent.get(index).unwrap() != index {
            self.parent[index] = self.find(self.parent[index]);
        }
        return self.parent[index];
    }
}

pub fn run() {
    let accounts = accounts_merge(vec![
        vec![String::from("ddd"), String::from("a@mail.com"),String::from("b@mail.com")],
        vec![String::from("sss"), String::from("b@mail.com")],
    ]);
    for account in accounts {
        print!("{}:", account[0]);
        for i in 1..account.len() {
            print!(
                "{},
            ",
                account[i]
            );
        }
        println!("");
    }
}
