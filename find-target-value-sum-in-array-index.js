const getMethod = (array, target) => {
  const map = [];
  const countVal = array.length;
  for (let i = 0; i < countVal; i++) {
    for (let j = i + 1; j < countVal; j++) {
      let sumVal = array[i] + array[j];
      if (sumVal === target) {
        map.push([i, j]);
      }
    }
  }
  return map;
};

const arr = [1, 7, 2, 3, 4, 5, 6];
const targetNum = 11;
const result = getMethod(arr, targetNum);
console.log(result);
