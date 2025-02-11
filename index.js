function increasingTriplet(nums) {
    if (nums.length < 3) {
        return false;
    }

    let min = Number.MAX_SAFE_INTEGER;
    let secondMin = Number.MAX_SAFE_INTEGER;

    for (let num of nums) {
        if (num <= min) {
            min = num;
        } else if (num <= secondMin) {
            secondMin = num;
        } else {
            return true;
        }
    }

    return false;
}

// Example usage:
let arr1 = [1, 2, 3, 4, 5];
console.log(increasingTriplet(arr1)); // Output: true (1 < 2 < 3)

let arr2 = [5, 4, 3, 2, 1];
console.log(increasingTriplet(arr2)); // Output: false (No increasing triplet)
