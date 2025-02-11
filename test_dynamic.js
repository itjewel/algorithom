// dynamic
let arr = [7, 6, 8, 9, 3, 5, 4, 6, 2, 9, 5, 1];
let target = 11;

/* Need to find target value sum from array give me index */
let map = {};
let res = [];
let getSumValue = (array, tar) => {
  let total = array.length;
  for (i = 0; i < total; i++) {
    let complement = tar - arr[i];
    if (map[complement] !== undefined) {
      res.push([map[complement], i]);
    }
    map[array[i]] = i;
    //console.log(complement);
  }
  return res ? res : "not found";
};

let result = getSumValue(arr, target);
console.log(result);
