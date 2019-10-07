class Node():
    def __init__(self, val, next = None):
        self.val = val
        self.next = next

class LinkedList():
    def __init__(self, data):
        self.root = self.create_ll(data)
    
    def create_ll(self, arr):
        root = Node(arr[0])
        current_node = root
        for i in range(arr[1:]):
            new_node = Node(arr[i])
            current_node.next = new_node
            new_node = current_node

        return root