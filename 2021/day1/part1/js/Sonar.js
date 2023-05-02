const fs = require("fs")



fs.readFile("../../input.txt", { encoding: 'utf8' }, function callback(err, data) {

    const arr = data.trim().split("\n")

    let prev = Number(arr[0])
    let count = 0

    for (let i of arr) {
        let n = Number(i)

        if (prev < n) {
            console.log(n + " (increased)")
            prev = n
            count++
        } else if (prev > n) {
            console.log(n + " (decreased)")
            prev = n
        } else {
            console.log(n + " (N/A - no previous measurement)")
        }

    }

    console.log(count)
})
