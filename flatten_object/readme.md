**Task**

Write a function that would iterate over a given object and flattens its structure. 

Note: Do not use any 3rd-party libraries

```js
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
```

For example `flattenObject(person)` should print in the console

```
{
  'name': 'john doe',
  'age': 25,
  'address.city': 'Birmingham',
  'address.Country': 'United Kindom',
  ....
}
```
