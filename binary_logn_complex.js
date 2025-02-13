const arr = [1, 2, 5, 7, 8, 9, 10];
const findIndexOf = 2;

const getResult = (arr, findInd) => {
  let left = 0;
  let right = arr.length - 1;
  while (left <= right) {
    let mid = Math.ceil((left + right) / 2);
    if (arr[mid] === findInd) {
      return mid;
    } else if (arr[mid] < findInd) {
      left = mid + 1;
    } else {
      right = mid - 1;
    }
  }
  return -1;
};

const result = getResult(arr, findIndexOf);
console.log(result);
