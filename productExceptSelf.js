var productExceptSelf = function(nums) {
    let n = nums.length;
    let prefixProducts = new Array(n).fill(1)
    let suffixProducts = new Array(n).fill(1)
    
    let prefix = 1;
    for(let i = 0; i<n; i++){
        prefixProducts[i] = prefix;
        prefix *= nums[i]
    }


    let suffix = 1;
    for(let i = n-1; i>=0; i--){
        suffixProducts[i] = suffix;
        suffix *= nums[i]
    }

    let results = [];
    for(let i=0; i<n; i++){
        results[i] = prefixProducts[i] * suffixProducts[i]
    }
    return results;
    
};

const numberd = [1,2,3,4];
console.log(productExceptSelf(numberd))