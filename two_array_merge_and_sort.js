// two array_merge_and_sort

const arrayOne = [1, 2, 3, 4, 5, 6];
const arrayTwo = [7, 2, 3, 4, 9, 6];

/*
const mergeArray = (arr1, arr2) => {
  const merge = [...arr1, ...arr2];
  return merge;
};
*/

const mergeArray = (arr1, arr2) => {
  const merge = arr1.reduce(
    (acc, next) => {
      acc.push(next);
      return acc;
    },
    [...arr2]
  );
  return merge;
};

const result = mergeArray(arrayOne, arrayTwo);
console.log(result);
