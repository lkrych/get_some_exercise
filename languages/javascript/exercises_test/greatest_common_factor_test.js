import { greatestCommonFactor } from '../exercises_itemized/greatest_common_factor';


test('greatestCommonFactor test', function(t) {
    t.plan(3);
    t.equal(greatestCommonFactor(3, 9), 3);
    t.equal(greatestCommonFactor(16, 24), 8);
    t.equal(greatestCommonFactor(7, 38), 1);
});