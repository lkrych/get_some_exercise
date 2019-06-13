import { thirdGreatest } from '../exercises_itemized/third_greatest';


test('thirdGreatest test', function(t) {
    t.plan(3);
    t.equal(thirdGreatest([1,2,3,4,5]), 3);
    t.equal(thirdGreatest([1,2,3,4,5,6,7,8,9,10]), 8);
    t.equal(thirdGreatest([1,2,3,4,5,6,7,8,9,10,11,12,13,14,15]), 13);
});