struct SummaryRanges {
    arr: Vec<i32>,
    intervals: Vec<Vec<i32>>,
}

/**
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl SummaryRanges {
    fn new() -> Self {
        return SummaryRanges {
            arr: Vec::new(),
            intervals: Vec::new(),
        };
    }

    fn add_num(&mut self, val: i32) {
        self.arr.push(val);
        self.arr.sort();
    }

    fn get_intervals(&mut self) -> Vec<Vec<i32>> {
        // for n in &self.arr {
        //     if let Some(&tail_vec) = &self.intervals.last() {
        //         let tail = tail_vec[0];
        //         if n - tail > 1 {
        //             self.intervals.push(vec![*n, *n])
        //         }
        //     } else {
        //         self.intervals.push(vec![*n, *n]);
        //     }
        // }
        return self.intervals.clone();
    }
}
