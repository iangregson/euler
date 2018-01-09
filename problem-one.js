const keepMultiplesOf3Or5 = v => v % 3 === 0 || v % 5 === 0
const sum = (acc, next) => acc += next
const numbers = new Array(...Array(1000).keys())
  .filter(keepMultiplesOf3Or5)

console.log(numbers.reduce(sum, 0))
