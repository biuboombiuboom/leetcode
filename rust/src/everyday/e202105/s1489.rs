fn find_critical_and_pseudo_critical_edges(n: i32, edges: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
    let edge_count = edges.len();
    let mut sorted_edges: Vec<Vec<i32>> = Vec::new();
    for i in 0..edge_count {
        let edge = edges.get(i).unwrap();
        let new_edge = vec![edge[0], edge[1], edge[2], i as i32];
        sorted_edges.push(new_edge);
    }
    sorted_edges.sort_by(|e1, e2| e1[2].cmp(&e2[2]));

    let mut uf = Unionfind::new(n);
    let min_mst = uf.find_mst(&sorted_edges, -1);
    let mut key_edges: Vec<i32> = Vec::new();
    let mut pseudo_key_edges: Vec<i32> = Vec::new();

    for i in 0..edge_count {
        let edge = &sorted_edges[i];
        let mut uf = Unionfind::new(n);
        let mst = uf.find_mst(&sorted_edges, i as i32);
        if mst > min_mst {
            key_edges.push(edge[3]);
            continue;
        }
        let mut uf = Unionfind::new(n);
        uf.union(edge[0] as usize, edge[1] as usize);
        let mst = uf.find_mst(&sorted_edges, i as i32);
        if edge[2] as i64 + mst == min_mst {
            pseudo_key_edges.push(edge[3]);
        }
    }

    vec![key_edges, pseudo_key_edges]
}

struct Unionfind {
    parent: Vec<usize>,
    rank: Vec<usize>,
    node_count: i32,
}

impl Unionfind {
    fn new(n: i32) -> Self {
        let mut parent = Vec::new();
        let mut rank = Vec::new();
        for i in 0..(n as usize) {
            parent.push(i);
            rank.push(1);
        }
        return Self {
            parent: parent,
            rank: rank,
            node_count: n,
        };
    }

    fn union(&mut self, x: usize, y: usize) -> bool {
        let (mut fx, mut fy) = (self.find(x), self.find(y));
        if fx == fy {
            return false;
        }
        if fx < fy {
            fx = swap(fx, fy).1;
            fy = swap(fx, fy).0;
        }
        self.rank[fx] += self.rank[fy];
        self.parent[fy] = fx;
        self.node_count -= 1;
        return true;
    }

    fn find(&mut self, x: usize) -> usize {
        if self.parent[x] != x {
            self.parent[x] = self.find(self.parent[x]);
        }
        return self.parent[x];
    }

    fn find_mst(&mut self, edges: &Vec<Vec<i32>>, skip_id: i32) -> i64 {
        let mut mst_wight = 0;
        for i in 0..edges.len() {
            if skip_id != i as i32 && self.union(edges[i][0] as usize, edges[i][1] as usize) {
                mst_wight += edges[i][2];
            }
        }
        if self.node_count > 1 {
            return 1 << 31;
        }
        mst_wight as i64
    }
}

pub fn run() {
    let edges = vec![
        vec![0, 1, 1],
        vec![1, 2, 1],
        vec![0, 2, 1],
        vec![2, 3, 4],
        vec![3, 4, 2],
        vec![3, 5, 2],
        vec![4, 5, 2],
    ];
    let edges = find_critical_and_pseudo_critical_edges(6, edges);
    for edge in edges {
        for elem in edge {
            println!("{}", elem)
        }
    }
}

fn swap(mut x: usize, mut y: usize) -> (usize, usize) {
    x = x + y;
    y = x - y;
    x = x - y;
    (x, y)
}

#[test]
fn test_swap() {
    let swap_tuple = swap(1, 2);
    assert_eq!((swap_tuple.0, swap_tuple.1), (2, 1))
}

#[test]
fn test_find_critical_and_pseudo_critical_edges() {
    let edges = vec![
        vec![0, 1, 1],
        vec![1, 2, 1],
        vec![0, 2, 1],
        vec![2, 3, 4],
        vec![3, 4, 2],
        vec![3, 5, 2],
        vec![4, 5, 2],
    ];
    let edges = find_critical_and_pseudo_critical_edges(6, edges);
    assert_eq!(edges[1].len(), 6)
}

#[test]
fn test_find_mst() {
    let mut edges = vec![
        vec![0, 1, 1, 0],
        vec![1, 2, 1, 1],
        vec![0, 2, 1, 2],
        vec![2, 3, 4, 3],
        vec![3, 4, 2, 4],
        vec![3, 5, 2, 5],
        vec![4, 5, 2, 6],
    ];
    let mut uf = Unionfind::new(6);
    edges.sort_by(|e1, e2| e1[2].cmp(&e2[2]));
    let mst = uf.find_mst(&edges, -1);
    assert_eq!(mst, 10);
}
