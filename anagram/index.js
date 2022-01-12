const fs = require("fs")
const path = require("path")

const readFile = () => {
    const data = fs.readFileSync(path.join(__dirname, "dictionary.txt"), { encoding: "utf-8" })
    return data.split('\n')
}

const list = ["dog", "cat", "disk", "bottle", "god", "kids"]

// anagrams function
function anagrams(list) {
    const sorted = []
    for (let i = 0; i < list.length; i++) {
        sorted.push(list[i].split("").sort().join(""))
    }

    const result = {}
    const anagram = []
    const isAnagram = (word) => anagram[word] === true
    const markWordAsAnagram = (word) => anagram[word] = true
    
    for (let i = 0; i < sorted.length; i++) {
        if (isAnagram(list[i])) continue
        if (!Array.isArray(result[list[i]])) result[list[i]] = []

        for (let j = i + 1; j < sorted.length; j++) {
            if (sorted[i].length !== sorted[j].length) continue

            if (sorted[i] === sorted[j]) {
                if (!Array.isArray(result[list[i]])) result[list[i]] = []
                result[list[i]].push(list[j])
                markWordAsAnagram(list[j])
            }
        }
    }

    return result
}

console.log(anagrams(list))
// console.log(anagrams(readFile()))