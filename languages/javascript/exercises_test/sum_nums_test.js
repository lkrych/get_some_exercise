import { sumNums } from './exercises';


test('sumNums test', function(t) {
    t.plan(5);
    t.equal(sumNums(3), 6);
    t.equal(sumNums(6), 21);
    t.equal(sumNums(7), 28);
    t.equal(sumNums(9), 45);
    t.equal(sumNums(15, 120));
});