pub fn min_cost_connect_points(points: Vec<Vec<i32>>) -> i32 {
    let mut edges: Vec<Edge> = Vec::new();
    let n = points.len();
    for i in 0..n {
        for j in (i + 1)..n {
            let edge = Edge {
                v: i,
                w: j,
                dis: dis(points.get(i).unwrap(), points.get(j).unwrap()),
            };
            edges.push(edge);
        }
    }
    edges.sort_by(|a, b| a.dis.cmp(&b.dis));
    let mut left = n - 1;
    let mut ans = 0;
    let mut uf = Unionfind::new(n);

    for edge in edges {
        if uf.union(edge.v, edge.w) {
            left -= 1;
            ans += edge.dis;
            if left == 0 {
                break;
            }
        }
    }

    return ans;
}

fn dis(a: &Vec<i32>, b: &Vec<i32>) -> i32 {
    return abs(a[0] - b[0]) + abs(a[1] - b[1]);
}

fn abs(x: i32) -> i32 {
    let mut x = x;
    if x < 0 {
        x = x * (0 - 1);
    }
    x
}

struct Edge {
    v: usize,
    w: usize,
    dis: i32,
}

struct Unionfind {
    parent: Vec<usize>,
}

impl Unionfind {
    fn new(n: usize) -> Self {
        let mut parent = Vec::new();
        for i in 0..n {
            parent.push(i);
        }
        Self { parent }
    }

    fn union(&mut self, x: usize, y: usize) -> bool {
        let fx = self.find(x);
        let fy = self.find(y);
        if fx == fy {
            return false;
        }
        self.parent[fy] = fx;
        return true;
    }

    fn find(&mut self, x: usize) -> usize {
        if self.parent[x] == x {
            self.parent[x] = self.find(self.parent[x]);
        }
        return self.parent[x];
    }
}

pub fn run() {
    let points = vec![vec![3, 12], vec![-2, 5], vec![-4, 1]];
    min_cost_connect_points(points);
}
