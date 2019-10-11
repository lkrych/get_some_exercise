class Node():
    def __init__(self, val, next = None):
        self.val = val
        self.next = next

class LinkedList():
    def __init__(self, data):
        if type data is List: #stub out until you know proper syntax
            self.root = self.create_ll_arr(data)
        else
            self.root = self.create_ll_node(data)
        self.data = data
    
    def create_ll_from_arr(self, arr):
        root = Node(arr[0])
        current_node = root
        for _ , val in enumerate(arr[1:]):
            new_node = Node(val)
            current_node.next = new_node
            new_node = current_node
        return root