const arr = [1, 7, 2, 3, 2, 3, 6];
const targetNum = 9;

const getDynamic = (array, targer) => {
  const mapCalled = new Map();
  const result = [];
  const count = array.length;
  for (let i = 0; i < count; i++) {
    let acc = targer - array[i];
    if (mapCalled.has(acc)) {
      result.push([mapCalled.get(acc), i]);
    }
    mapCalled.set(array[i], i);
  }
  return result;
};
console.log(getDynamic(arr, targetNum));
