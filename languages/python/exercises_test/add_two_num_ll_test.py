from helper import LinkedList, create_arr_from_ll

class TestAddTwoNumsLL(unittest.TestCase):
    def test_5(self):
        """
        Test that add_two_nums can handle small integers
        """
        ll1 = LinkedList([5])
        ll2 = LinkedList([6])
        ans = LinkedList([1, 1])
        result = add_two_nums_ll(ll1.root, ll2.root)
        self.assertEqual(create_arr_from_ll(result), ans.data)

    def test_6(self):
        """
        Test that add_two_nums_ll can handle small integers
        """
        ll1 = LinkedList([3])
        ll2 = LinkedList([8])
        ans = LinkedList([1, 1])
        result = add_two_nums_ll(ll1.root, ll2.root)
        self.assertEqual(create_arr_from_ll(result), ans.data)

    def test_8(self):
        """
        Test that add_two_nums_ll can handle small integers
        """
        ll1 = LinkedList([8])
        ll2 = LinkedList([2, 4, 5])
        ans = LinkedList([2, 5, 3])
        result = add_two_nums_ll(ll1.root, ll2.root)
        self.assertEqual(create_arr_from_ll(result), ans.data)
    def test_65536(self):
        """
        Test that add_two_nums_ll can handle larger integers
        """
        ll1 = LinkedList([6, 5, 5, 3, 6])
        ll2 = LinkedList([4, 3])
        ans = LinkedList([6, 5, 5, 7, 9])
        result = add_two_nums_ll(ll1.root, ll2.root)
        self.assertEqual(create_arr_from_ll(result), ans.data)
 