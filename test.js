let arr = [7, 6, 8, 9, 3, 5, 4, 6, 2, 9, 5, 1];
let target = 11;

/* Need to find target value sum from array give me index */
let map = [];
let getSumValue = (array, tar) => {
  let total = array.length;
  for (i = 0; i < total; i++) {
    for (j = i + 1; j < total; j++) {
      let sum = arr[i] + arr[j];
      if (sum === tar) {
        map.push([i, j]);
      }
    }
    //
  }
  return map;
};

let result = getSumValue(arr, target);
console.log(result);
