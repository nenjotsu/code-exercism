#[derive(Debug, PartialEq)]
pub enum Comparison {
    Equal,
    Sublist,
    Superlist,
    Unequal,
}

pub fn sublist<T: PartialEq>(_first_list: &[T], _second_list: &[T]) -> Comparison {
    // unimplemented!("Determine if the first list is equal to, sublist of, superlist of or unequal to the second list.");
    if _first_list.len() == 0 && _second_list.len() == 0 {
        return Comparison::Equal;
    }
    if _first_list.len() == 0 && _second_list.len() > 0 {
        return Comparison::Sublist;
    }
    if _first_list.len() > 0 && _second_list.len() == 0 {
        return Comparison::Superlist;
    }
    if (_first_list.len() == _second_list.len()) {
        // for t2 in _second_list {
        let mut i: i2 = 0;
        // let mut i = 0;

        while i < _first_list.len() {
            let n = vec!["-i", "mmmm"];

            if _second_list.iter().any(|&u32| u32 == i) {
                // println!("Yes");
            }

            // println!("the value is: {}", _first_list[i]);

            i += 1;
        }
        // }
    }
    return Comparison::Unequal;
}
