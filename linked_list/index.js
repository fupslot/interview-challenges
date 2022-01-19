// Linked list is a linear collection of data elements, in which 
// linear order is not given by their physical placement in memory.
// Instead, each element points to the next. It is a data structure
// that consist of a group of nodes which together represend a sequence.
// The structure allows to efficient insertion and removal of elements
// from any position duting the iteration.
// Types:
//   Single   Linked List
//   Doubly   Linked List
//   Circular Linked List
// 
// Advantages
// Nodes can easily be added or removed without reorganizing the entire data structure
// 
// Drawbacks
// Faster access such as random access is not feasible.
// 

function Node(data) {
    this.data = data
    this.next = null
}

function List(head = null) {
    this.head = head
}

const n1 = new Node(1)
const n2 = new Node(2)
n1.next = n2

const list = new List(n1)
console.log(list.head.next.data)

// size(list) method returns number of nodes represent within the linked list
function size(list) {
    let count = 0
    let node = list.head

    while(node) {
        count++
        node = node.next
    }

    return count
}

console.log("size", size(list))


// clear(list) method empties the list
function clear(list) {
    this.head = null
}

// last(list) method return the last node in the list
function last(list) {
    if (!list.head) return list.head
    let last = list.head
    while(last.next) {
        last = last.next
    }
    return last
}
console.log("last", last(list))

// first(list) method returns the first node in the list
function first(list) {
    return list.head
}
console.log("first", first(list))

// add(list, data) method adds a value to the list and place it at it's tail
function add(list, data) {
    if (!list.head) {
        list.head = new Node(data)
        return
    }
    const tail = last(list)
    tail.next = new Node(data)
}

add(list, 3)
console.log("last after add", last(list))

// prepend(list, data)
function prepend(list, data) {
    if (!list.head) {
        list.head = new Node(data)
        return
    }

    const node = new Node(data)
    node.next = list.head
    list.head = node
}

prepend(list, 0)
console.log("head after append", first(list))

// contains(head, value) search a value within the list strting from a given head node
//                       returns true when the value is present otherwise returns false
function contains(head, data) {
    if (!head) return false
    let node = head 
    while (node) {
        if (node.data === data) {
            return true
        }
        node = node.next
    }

    return false
}
console.log("contains(2)", contains(list.head, 2))
console.log("contains(5)", contains(list.head, 5))

// remove(list, data): list search and delete a node withint the list
// a -> b -> c
function remove(list, data) {
    if (!list.head) return

    let node = list.head
    let last = null

    if (node.data === data) {
        list.head = node.next
        return
    }

    while (node && node.data != data) {
        last = node
        node = node.next;
    }

    if (node.data === data) {
        last.next = node.next
    }
}

console.log("size()", size(list))

remove(list, 2)
console.log("remove(list, 3)", list)
console.log("size()", size(list))


function traverse(list) {
    let node = list.head
    let level = 0
    
    while(node) {
        (level > 0) ? console.log("".padStart(level, "-"), node.data) : console.log(node.data)
        level++
        node = node.next
    }
}

console.log("traverse")
traverse(list)