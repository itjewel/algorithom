
const  gcdOfStrings = function(str1, str2) {
    const param1 = str1.length;
    const param2 = str2.length;
    if((str1 + str2) !== (str2 + str1)){
        return '';
    }
   
    const gcd = (a,b)=>(b=== 0? a : gcd(b, a%b));
    const gcdLenth = gcd(param1,param2);
    return str1.substring(0, gcdLenth)
};
const str1 ="ABCABC"
const str2 = "ABC"
console.log(gcdOfStrings(str1,str2))