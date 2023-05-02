const { log } = require("console")
const fs = require("fs")



fs.readFile("../../input.txt", { encoding: 'utf8' }, function cb(err, data) {

    const arr = data.trim().split("\n").map((n) => Number(n))

    const res = []

    let end = 3
    let sum = 0
    let count = 0

    for (let i = 0; i < arr.length; i++) {

        for (let index = i; index < i + end; index++) {
            sum += arr[index]
        }
        if (sum > res[i - 1]) {
            count++
        }
        res.push(sum)
        sum = 0
    }
    log(String(res))

    log(arr.length)
    log(count)



})


