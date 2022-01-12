const person = {
    name: "john doe",
    age: 25,
    address: {
        city:"Birmingham",
        Country:"United Kindom",
        phone:{
            extension: "1",
            number: "234 54 32"
        },
        code: "505055"
    },
    gender:"male"
}

function flattenObject(obj, prefix, parent = {}) {
    const path = (prefix, key) => (typeof prefix === "string") ? `${prefix}.${key}` : key 

    for (const [ key, value ] of Object.entries(obj)) {
        if (Object.prototype.toString.call(value) === '[object Object]') {
            flattenObject(value, path(prefix, key), parent) 
            continue
        }

        parent[path(prefix, key)] = value
    }

    return parent
}


const flatten = flattenObject(person, "person")
console.log(flatten)