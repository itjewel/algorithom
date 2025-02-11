const getMethod = (array, target) => {
  //let results = [];

  for (let key = 0; key < array.length; key++) {
    for (let j = key + 1; j < array.length; j++) {
      // Start from key + 1 to avoid duplicate index
      let res = array[key] + array[j];

      if (res === target) {
        return [key, j]; // Return the first valid pair found
      }
    }
  }
  return []; // Return empty array if no valid pair is found
};

const arr = [1, 7, 2, 3, 4, 5, 6];
const targetNum = 11;
const result = getMethod(arr, targetNum);
console.log(result);
