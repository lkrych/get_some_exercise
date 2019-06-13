import { factorial } from '../exercises_itemized/factorial';


test('factorial test', function(t) {
    t.plan(5);
    t.equal(factorial(3), 6);
    t.equal(factorial(6), 720);
    t.equal(factorial(7), 5040);
    t.equal(factorial(9), 362880);
    t.equal(factorial(15, 1307674368000));
});