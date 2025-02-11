const getAllPairs = (array, target) => {
  let map = {};
  let result = [];
  let lengthCount = array.length;

  for (let i = 0; i < lengthCount; i++) {
    let complement = target - array[i];

    // 8
    // 2
    // 7
    if (map[complement] !== undefined) {
      result.push([map[complement], i]); // Store index pairs
      // [1,2]
    }

    map[array[i]] = i;
    //{1:0}
    // {7:1}

    // {7:1}
    // console.log(map[array[i]], i, "jewel");
  }

  return result.length ? result : "No pairs found!";
};

0(n);

const arr = [1, 7, 2, 3, 2, 3, 6];
const targetNum = 9;
console.log(getAllPairs(arr, targetNum));
