function Node(data) {
    this.data = data
    this.next = null
    this.prev = null
}

function List(head = null) {
    this.head = head
}


function add(list, node) {
    if (!list.head) {
        list.head = node
        return
    }

    if (contains(list, node.data)) return

    const tail = last(list)
    tail.next = node
    node.prev = tail
}

function last(list) {
    let node = list.head

    while(node && node.next) {
        node = node.next
    }
    return node
}

function find(list, data) {
    let node = list.head
    while(node && node.data !== data) {
        node = node.next
    }

    return node
}

function contains(list, data) {
    let node = list.head
    while(node) {
        if (node.data === data) return true
        node = node.next
    }

    return false
}

function remove(list, data) {
    const node = find(list, data)
    if (!node) return

    if (list.head == node) {
        list.head = node.next
    } else {
        node.prev.next = node.next
        node.next.prev = node.prev
    }
}


function traverse(list) {
    let node = list.head
    let level = 0

    while(node) {
        (level > 0) ? console.log("".padStart(level, "-"), node.data) : console.log(node.data)
        level++
        node = node.next
    }
}

const list = new List()
const n1 = new Node(0)
const n2 = new Node(1)
const n3 = new Node(2)

add(list, n1)
add(list, n2)
add(list, n3)
traverse(list)


// console.log(find(list, 1))
remove(list, 1)
console.log(list)