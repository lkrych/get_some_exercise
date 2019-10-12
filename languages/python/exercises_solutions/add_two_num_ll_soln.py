from helper import Node

def add_two_nums_ll(l1: Node, l2: Node) -> Node:
        root_node = Node(0)
        current_node = root_node
        carry = 0
        while l1 or l2:
            sum = carry
            if l1:
                sum += l1.val
                l1 = l1.next
            if l2:
                sum += l2.val
                l2 = l2.next
                
            val = sum % 10
            carry = sum // 10
            
            current_node.val = val
            
            if l1 or l2:
                current_node.next = Node(None)
                current_node = current_node.next  
        
        if carry > 0:
            current_node.next = Node(1)
            
        return root_node
