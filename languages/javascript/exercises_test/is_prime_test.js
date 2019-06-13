import { isPrime } from '../exercises_itemized/is_prime';


test('isPrime test', function(t) {
    t.plan(5);
    t.equal(isPrime(2), true);
    t.equal(isPrime(3), true);
    t.equal(isPrime(4), false);
    t.equal(isPrime(9), false);
    t.equal(isPrime(227, true));
});