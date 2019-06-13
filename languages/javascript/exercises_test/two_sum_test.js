import { twoSum } from '../exercises_itemized/two_sum';

test('twoSum test', function(t) {
    t.plan(3);
    t.equal(twoSum([1, 3, 5, -3], 0), [1,3]);
    t.equal(twoSum([1,3,5], 0), []);
    t.equal(twoSum([1, 2, 3, 4, 5,6, -4], 0), [3, 6]);
});